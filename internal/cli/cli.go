package cli

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	cli "github.com/urfave/cli/v2"

	"github.com/nevalang/neva/internal/builder"
	"github.com/nevalang/neva/internal/compiler"
	"github.com/nevalang/neva/internal/compiler/sourcecode"
	"github.com/nevalang/neva/internal/compiler/sourcecode/graphviz"
	"github.com/nevalang/neva/internal/interpreter"
	"github.com/nevalang/neva/pkg"
)

func NewApp( //nolint:funlen
	workdir string,
	parser compiler.Parser,
	bldr builder.Builder,
	goc compiler.Compiler,
	nativec compiler.Compiler,
	wasmc compiler.Compiler,
	jsonc compiler.Compiler,
	dotc compiler.Compiler,
) *cli.App {
	var (
		target string
		main   string
		debug  bool
	)

	return &cli.App{
		Name:  "neva",
		Usage: "Flow-based programming language",
		Commands: []*cli.Command{
			{
				Name:  "version",
				Usage: "Get current Nevalang version",
				Action: func(_ *cli.Context) error {
					fmt.Println(pkg.Version)
					return nil
				},
			},
			{
				Name:  "new",
				Usage: "Create new Nevalang project",
				Args:  true,
				Action: func(cCtx *cli.Context) error {
					if path := cCtx.Args().First(); path != "" {
						if err := os.Mkdir(path, 0755); err != nil {
							return err
						}
						return createNevaMod(path)
					}
					return createNevaMod(workdir)
				},
			},
			{
				Name:      "get",
				Usage:     "Add dependency to current module",
				Args:      true,
				ArgsUsage: "Provide path to the module",
				Action: func(cCtx *cli.Context) error {
					installedPath, err := bldr.Get(
						workdir,
						cCtx.Args().Get(0),
						cCtx.Args().Get(1),
					)
					if err != nil {
						return err
					}
					fmt.Printf(
						"%s installed to %s\n", cCtx.Args().Get(0),
						installedPath,
					)
					return nil
				},
			},
			{
				Name:      "run",
				Usage:     "Run neva program from source code in interpreter mode",
				Args:      true,
				ArgsUsage: "Provide path to the executable package",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "debug",
						Usage:       "Show message events in stdout",
						Destination: &debug,
					},
				},
				Action: func(cCtx *cli.Context) error {
					dirFromArg, err := getMainPkgFromArgs(cCtx)
					if err != nil {
						return err
					}
					intr := interpreter.New(bldr, goc, debug)
					if err := intr.Interpret(
						context.Background(),
						workdir,
						dirFromArg,
					); err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "build",
				Usage: "Build executable binary from neva program source code",
				Args:  true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "target",
						Required:    false,
						Usage:       "Emit Go or WASM instead of machine code",
						Destination: &target,
						Action: func(ctx *cli.Context, s string) error {
							switch s {
							case "go", "wasm", "native", "json", "dot":
							default:
								return fmt.Errorf("Unknown target %s", s)
							}
							return nil
						},
					},
				},
				ArgsUsage: "Provide path to the executable package",
				Action: func(cCtx *cli.Context) error {
					dirFromArg, err := getMainPkgFromArgs(cCtx)
					if err != nil {
						return err
					}
					switch target {
					case "go":
						return goc.Compile(
							workdir, dirFromArg, workdir,
						)
					case "wasm":
						return wasmc.Compile(
							workdir, dirFromArg, workdir,
						)
					case "json":
						return jsonc.Compile(
							workdir, dirFromArg, workdir,
						)
					case "dot":
						return dotc.Compile(
							workdir, dirFromArg, workdir,
						)
					default:
						return nativec.Compile(
							workdir, dirFromArg, workdir,
						)
					}
				},
			},
			{
				Name:      "parse",
				Usage:     "Parse a neva source file and output a Graphviz visualization",
				Args:      true,
				ArgsUsage: "Provide path to the executable package",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "main",
						Required:    false,
						Usage:       "Name of Main component to output",
						Destination: &main,
						DefaultText: "Main",
					},
				},
				Action: func(cCtx *cli.Context) error {
					inputFile, err := getInputFileFromArgs(cCtx)
					if err != nil {
						return err
					}
					f, err := os.Open(inputFile)
					if err != nil {
						return err
					}
					data, err := io.ReadAll(f)
					if err != nil {
						return err
					}
					f.Close()

					loc := sourcecode.Location{
						FileName: inputFile,
					}
					sf, err := parser.ParseFile(loc, data)
					if err != (*compiler.Error)(nil) {
						return err
					}

					var cb graphviz.ClusterBuilder
					for name, e := range sf.Entities {
						if err := cb.AddEntity(nil, name, e); err != nil {
							return err
						}
					}
					return cb.Build(os.Stdout)
				},
			},
		},
	}
}

func getMainPkgFromArgs(cCtx *cli.Context) (string, error) {
	firstArg := cCtx.Args().First()
	dirFromArg := strings.TrimSuffix(firstArg, "/main.neva")
	if filepath.Ext(dirFromArg) != "" {
		return "", errors.New(
			"Use path to directory with executable package, relative to module root",
		)
	}
	return dirFromArg, nil
}

func getInputFileFromArgs(cCtx *cli.Context) (string, error) {
	inputFile := cCtx.Args().First()
	if filepath.Ext(inputFile) != ".neva" {
		return "", errors.New(
			`command takes a single neva input file, relative to module root (ex. "./pkg/main.neva")`,
		)
	}
	return inputFile, nil
}

func createNevaMod(path string) error {
	// Create neva.yml file
	nevaYmlContent := fmt.Sprintf("neva: %s", pkg.Version)
	if err := os.WriteFile(
		filepath.Join(path, "neva.yml"),
		[]byte(nevaYmlContent),
		0644,
	); err != nil {
		return err
	}

	// Create src sub-directory
	srcPath := filepath.Join(path, "src")
	if err := os.Mkdir(srcPath, 0755); err != nil {
		return err
	}

	// Create main.neva file
	mainNevaContent := `component Main(start any) (stop any) {
	nodes {

	}
	net {
		:start -> :stop
	}
}
`

	if err := os.WriteFile(
		filepath.Join(srcPath, "main.neva"),
		[]byte(mainNevaContent),
		0644,
	); err != nil {
		return err
	}

	return nil
}

package main

import (
	"fmt"
	"os"

	"github.com/nevalang/neva/internal/builder"
	"github.com/nevalang/neva/internal/cli"
	"github.com/nevalang/neva/internal/compiler"
	"github.com/nevalang/neva/internal/compiler/analyzer"
	"github.com/nevalang/neva/internal/compiler/backend/dot"
	"github.com/nevalang/neva/internal/compiler/backend/golang"
	"github.com/nevalang/neva/internal/compiler/backend/golang/native"
	"github.com/nevalang/neva/internal/compiler/backend/golang/wasm"
	"github.com/nevalang/neva/internal/compiler/backend/json"
	"github.com/nevalang/neva/internal/compiler/desugarer"
	"github.com/nevalang/neva/internal/compiler/irgen"
	"github.com/nevalang/neva/internal/compiler/parser"
	"github.com/nevalang/neva/internal/compiler/sourcecode/typesystem"
	"github.com/nevalang/neva/pkg"
)

func main() { //nolint:funlen
	// current working directory (hopefully with neva.yaml)
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// type-system
	terminator := typesystem.Terminator{}
	checker := typesystem.MustNewSubtypeChecker(terminator)
	resolver := typesystem.MustNewResolver(typesystem.Validator{}, checker, terminator)

	// parser
	prsr := parser.New(false)

	// pkg manager
	pkgMngr := builder.MustNew(prsr)

	// compiler frontend
	desugarer := desugarer.New()
	analyzer := analyzer.MustNew(pkg.Version, resolver)
	irgen := irgen.New()

	// golang backend
	golangBackend := golang.NewBackend()

	// this one can emit go code
	goCompiler := compiler.New(
		pkgMngr,
		prsr,
		desugarer,
		analyzer,
		irgen,
		golang.NewBackend(),
	)

	// this one can emit native code
	nativeCompiler := compiler.New(
		pkgMngr,
		prsr,
		desugarer,
		analyzer,
		irgen,
		native.NewBackend(
			golangBackend,
		),
	)

	// this one can emit wasm code
	wasmCompiler := compiler.New(
		pkgMngr,
		prsr,
		desugarer,
		analyzer,
		irgen,
		wasm.NewBackend(
			golangBackend,
		),
	)

	jsonCompiler := compiler.New(
		pkgMngr,
		prsr,
		desugarer,
		analyzer,
		irgen,
		json.NewBackend(),
	)

	dotCompiler := compiler.New(
		pkgMngr,
		prsr,
		desugarer,
		analyzer,
		irgen,
		dot.NewBackend(),
	)

	// command-line app that can compile and interpret neva code
	app := cli.NewApp(
		wd,
		prsr,
		pkgMngr,
		goCompiler,
		nativeCompiler,
		wasmCompiler,
		jsonCompiler,
		dotCompiler,
	)

	// run CLI app
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		return
	}
}

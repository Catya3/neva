package golang

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"fmt"
	"text/template"

	"github.com/emil14/neva/internal/compiler/ir"
)

//go:embed tmpl/main.go.tmpl
var efs embed.FS

type Backend struct{}

var ErrExecTmpl = errors.New("execute template")

func (b Backend) GenerateTarget(ctx context.Context, prog ir.Program) ([]byte, error) {
	tmpl, err := template.ParseFS(efs, "tmpl/main.go.tmpl")
	if err != nil {
		return nil, err
	}

	tmpl.Funcs(template.FuncMap{
		"getMsg": b.getMsg,
		"getIO":  b.getPorts(prog.Ports),
	})

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, prog); err != nil {
		return nil, errors.Join(ErrExecTmpl, err)
	}

	return buf.Bytes(), nil
}

var ErrUnknownMsgType = errors.New("unknown msg type")

func (b Backend) getMsg(msg ir.Msg) (string, error) {
	switch msg.Type {
	case ir.IntMsg:
		return fmt.Sprintf("runtime.NewIntMsg(%d)", msg.Int), nil
	}
	return "", fmt.Errorf("%w: %v", ErrUnknownMsgType, msg.Type)
}

func (b Backend) getPorts(ports map[ir.PortAddr]uint8) func(path, port string) (string, error) {
	return func(path, port string) (string, error) {
		return "", nil
	}
}

package core

import (
	"github.com/emil14/stream/internal/types"
)

type Component interface {
	Interface() Interface
}

type Interface struct {
	In  InportsInterface
	Out OutportsInterface
}

func (want Interface) Compare(got Interface) error {
	if err := PortsInterface(want.In).Compare(PortsInterface(got.In)); err != nil {
		return err
	}

	return PortsInterface(want.Out).Compare(PortsInterface(got.Out))
}

type InportsInterface PortsInterface

type OutportsInterface PortsInterface

type PortsInterface map[string]PortType

func (want PortsInterface) Compare(got PortsInterface) error {
	if len(want) != len(got) {
		return errPortsLen(len(want), len(got))
	}

	for name, typ := range want {
		_, ok := got[name]
		if !ok {
			return errPortNotFound(name, typ)
		}

		if err := typ.Compare(got[name]); err != nil {
			return errPortInvalid(name, err)
		}
	}

	return nil
}

func (io PortsInterface) Arr() map[string]PortType {
	m := map[string]PortType{}

	for name, typ := range io {
		if typ.Arr {
			m[name] = typ
		}
	}

	return m
}

type PortType struct {
	Type types.Type
	Arr  bool
}

func (p1 PortType) Compare(p2 PortType) error {
	if p1.Arr != p2.Arr || p1.Type != p2.Type {
		return errPortTypes(p1, p2)
	}

	return nil
}

func (pt PortType) String() (s string) {
	if pt.Arr {
		s += "array"
	}

	return s + "port of type " + pt.Type.String()
}
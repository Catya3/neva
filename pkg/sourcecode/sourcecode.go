// This package defines source code entities - abstractions that end-user (a programmer) operates on.
// For convenience these structures have json tags. This is not clean architecture but it's very handy for LSP.
package sourcecode

import (
	"fmt"

	ts "github.com/nevalang/neva/pkg/typesystem"
)

type Build struct {
	EntryModRef ModuleRef
	Modules     map[ModuleRef]Module
}

type Module struct {
	Manifest ModuleManifest     `json:"manifest,omitempty"`
	Packages map[string]Package `json:"packages,omitempty"`
}

type ModuleManifest struct {
	WantCompilerVersion string               `json:"compiler,omitempty" yaml:"compiler,omitempty"`
	Deps                map[string]ModuleRef `json:"deps,omitempty"     yaml:"deps,omitempty"`
}

type ModuleRef struct {
	Path    string `json:"path,omitempty"`
	Version string `json:"version,omitempty"`
}

var ErrEntityNotFound = fmt.Errorf("entity not found")

type Package map[string]File

// Just like program's Entity
func (p Package) Entity(entityName string) (entity Entity, filename string, ok bool) {
	for fileName, file := range p {
		entity, ok := file.Entities[entityName]
		if ok {
			return entity, fileName, true
		}
	}
	return Entity{}, "", false
}

// Entities takes function that can return error. If function returns error iteration stops.
func (p Package) Entities(f func(entity Entity, entityName string, fileName string) error) error {
	for fileName, file := range p {
		for entityName, entity := range file.Entities {
			if err := f(entity, entityName, fileName); err != nil {
				return err
			}
		}
	}
	return nil
}

type File struct {
	Imports  map[string]Import `json:"imports,omitempty"`
	Entities map[string]Entity `json:"entities,omitempty"`
}

type Import struct {
	ModuleName string `json:"moduleName,omitempty"`
	PkgName    string `json:"pkgName,omitempty"`
}

type Entity struct {
	IsPublic  bool       `json:"exported,omitempty"`
	Kind      EntityKind `json:"kind,omitempty"`
	Const     Const      `json:"const,omitempty"`
	Type      ts.Def     `json:"type,omitempty"`
	Interface Interface  `json:"interface,omitempty"`
	Component Component  `json:"component,omitempty"`
}

func (e Entity) Meta() *Meta {
	m := Meta{}
	switch e.Kind {
	case ConstEntity:
		m = e.Const.Meta
	case TypeEntity:
		m = e.Type.Meta.(Meta) //nolint
	case InterfaceEntity:
		m = e.Interface.Meta
	case ComponentEntity:
		m = e.Component.Meta
	}
	return &m
}

type EntityKind string // It's handy to transmit strings enum instead of digital

const (
	ComponentEntity EntityKind = "component_entity"
	ConstEntity     EntityKind = "const_entity"
	TypeEntity      EntityKind = "type_entity"
	InterfaceEntity EntityKind = "interface_entity"
)

type Component struct {
	Directives map[Directive][]string `json:"directives,omitempty"`
	Interface  `json:"interface,omitempty"`
	Nodes      map[string]Node `json:"nodes,omitempty"`
	Net        []Connection    `json:"net,omitempty"`
	Meta       Meta            `json:"meta,omitempty"`
}

type Directive string

type Interface struct {
	TypeParams TypeParams `json:"typeParams,omitempty"`
	IO         IO         `json:"io,omitempty,"`
	Meta       Meta       `json:"meta,omitempty"`
}

type TypeParams struct {
	Params []ts.Param `json:"params,omitempty"`
	Meta   Meta       `json:"meta,omitempty"`
}

type Node struct {
	Directives map[Directive][]string `json:"directives,omitempty"`
	EntityRef  EntityRef              `json:"entityRef,omitempty"`
	TypeArgs   []ts.Expr              `json:"typeArgs,omitempty"`
	Deps       map[string]Node        `json:"componentDi,omitempty"`
	Meta       Meta                   `json:"meta,omitempty"`
}

type EntityRef struct {
	Pkg  string `json:"pkg,omitempty"`
	Name string `json:"name,omitempty"`
	Meta Meta   `json:"meta,omitempty"`
}

func (e EntityRef) String() string {
	if e.Pkg == "" {
		return e.Name
	}
	return fmt.Sprintf("%s.%s", e.Pkg, e.Name)
}

type Const struct {
	Ref   *EntityRef `json:"ref,omitempty"`
	Value *Msg       `json:"value,omitempty"`
	Meta  Meta       `json:"meta,omitempty"`
}

type Msg struct {
	TypeExpr ts.Expr          `json:"typeExpr,omitempty"`
	Bool     *bool            `json:"bool,omitempty"`
	Int      *int             `json:"int,omitempty"`
	Float    *float64         `json:"float,omitempty"`
	Str      *string          `json:"str,omitempty"`
	List     []Const          `json:"vec,omitempty"`
	Map      map[string]Const `json:"map,omitempty"`
	Meta     Meta             `json:"meta,omitempty"`
}

type IO struct {
	In  map[string]Port `json:"in,omitempty"`
	Out map[string]Port `json:"out,omitempty"`
}

type Port struct {
	TypeExpr ts.Expr `json:"typeExpr,omitempty"`
	IsArray  bool    `json:"isArray,omitempty"`
	Meta     Meta    `json:"meta,omitempty"`
}

type Connection struct {
	SenderSide    SenderConnectionSide     `json:"senderSide,omitempty"`
	ReceiverSides []ReceiverConnectionSide `json:"receiverSide,omitempty"`
	Meta          Meta                     `json:"meta,omitempty"`
}

type ReceiverConnectionSide struct {
	PortAddr  PortAddr `json:"portAddr,omitempty"`
	Selectors []string `json:"selectors,omitempty"`
	Meta      Meta     `json:"meta,omit"`
}

// SenderConnectionSide unlike ReceiverConnectionSide could refer to constant.
type SenderConnectionSide struct {
	PortAddr  *PortAddr  `json:"portAddr,omitempty"`
	ConstRef  *EntityRef `json:"constRef,omitempty"` // Only sugared form
	Selectors []string   `json:"selectors,omitempty"`
	Meta      Meta       `json:"meta,omitempty"`
}

type PortAddr struct {
	Node string `json:"node,omitempty"`
	Port string `json:"port,omitempty"`
	Idx  *uint8 `json:"idx,omitempty"`
	Meta Meta   `json:"meta,omitempty"`
}

func (p PortAddr) String() string {
	if p.Node == "" {
		return fmt.Sprintf("%v[%v]", p.Port, p.Idx)
	}
	return fmt.Sprintf("%v.%v[%v]", p.Node, p.Port, p.Idx)
}

// Meta keeps info about original text related to the structured object
type Meta struct {
	Text  string   `json:"text,omitempty"`
	Start Position `json:"start,omitempty"`
	Stop  Position `json:"stop,omitempty"`
}

type Position struct {
	Line   int `json:"line,omitempty"`
	Column int `json:"column,omitempty"`
}
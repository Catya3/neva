package runtime

import "fmt"

// Msg represents set of methods only one of which should return real value.
type Msg interface {
	Str() string
	Int() int
	Bool() bool
	Struct() map[string]Msg
}

type emptyMsg struct{} // emptyMsg exists to allow normal messages define only reasonable methods.

func (msg emptyMsg) Str() (_ string)            { return }
func (msg emptyMsg) Int() (_ int)               { return }
func (msg emptyMsg) Bool() (_ bool)             { return }
func (msg emptyMsg) Struct() (_ map[string]Msg) { return }

type IntMsg struct {
	emptyMsg
	v int
}

func (msg IntMsg) Int() int {
	return msg.v
}

func (msg IntMsg) String() string {
	return fmt.Sprintf("'%d'", msg.v)
}

func NewIntMsg(n int) IntMsg {
	return IntMsg{
		emptyMsg: emptyMsg{},
		v:        n,
	}
}

type StrMsg struct {
	emptyMsg
	v string
}

func (msg StrMsg) Str() string {
	return msg.v
}

func NewStrMsg(s string) StrMsg {
	return StrMsg{
		emptyMsg: emptyMsg{},
		v:        s,
	}
}

type BoolMsg struct {
	emptyMsg
	v bool
}

func (msg BoolMsg) Bool() bool {
	return msg.v
}

func (msg BoolMsg) String() string {
	return fmt.Sprintf("%v", msg.v)
}

func NewBoolMsg(b bool) BoolMsg {
	return BoolMsg{
		emptyMsg: emptyMsg{},
		v:        b,
	}
}

type StructMsg struct {
	emptyMsg
	v map[string]Msg
}

func (msg StructMsg) Struct() map[string]Msg {
	return msg.v
}

func NewMsgStruct(v map[string]Msg) StructMsg {
	return StructMsg{
		emptyMsg: emptyMsg{},
		v:        v,
	}
}

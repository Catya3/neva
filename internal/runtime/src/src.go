package src

type (
	Program struct {
		Effects       Effects
		StartPortAddr PortAddr
		Ports         map[PortAddr]uint8 // value is buf size
		Net           []Connection
	}

	PortAddr struct {
		Path string
		Port string
		Idx  uint8
	}

	Connection struct {
		SenderSide    ConnectionSide
		ReceiverSides []ConnectionSide
	}

	ConnectionSide struct {
		PortAddr   PortAddr
		DictUnpack []string // if not nil connector should unpack dict message
	}

	Effects struct {
		Void     []PortAddr       // read from unused outport to avoid lock
		Giver    map[PortAddr]Msg // send message to port
		Operator []OpRef          // do computation, not map because component can have many instances
	}

	OpRef struct {
		Pkg, Name string
		PortAddrs OpPortAddrs
	}

	Msg struct {
		Type  MsgType
		Bool  bool
		Int   int64
		Float float64
		Str   string
		List  []Msg
		Dict  map[string]Msg
	}

	MsgType uint8

	FuncRef struct {
		Pkg, Name string
	}

	OpPortAddrs struct {
		In, Out []PortAddr
	}
)

const (
	BoolMsg MsgType = iota + 1
	IntMsg
	FloatMsg
	StrMsg
	DictMsg
	List
)

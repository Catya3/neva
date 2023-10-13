// Code generated from ./neva.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type nevaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var NevaLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func nevalexerLexerInit() {
	staticData := &NevaLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'use'", "'{'", "'}'", "'@/'", "'/'", "'types'", "'<'", "'>'", "','",
		"'['", "']'", "'|'", "'interfaces'", "'('", "')'", "'const'", "'true'",
		"'false'", "'nil'", "':'", "'components'", "'nodes'", "'.'", "'net'",
		"'->'", "'in'", "'out'", "", "'pub'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "COMMENT", "PUB_KW", "IDENTIFIER",
		"INT", "FLOAT", "STRING", "NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "COMMENT", "PUB_KW", "IDENTIFIER", "LETTER", "INT",
		"FLOAT", "STRING", "NEWLINE", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 35, 236, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 180, 8, 27,
		10, 27, 12, 27, 183, 9, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1,
		29, 5, 29, 192, 8, 29, 10, 29, 12, 29, 195, 9, 29, 1, 30, 1, 30, 1, 31,
		4, 31, 200, 8, 31, 11, 31, 12, 31, 201, 1, 32, 5, 32, 205, 8, 32, 10, 32,
		12, 32, 208, 9, 32, 1, 32, 1, 32, 4, 32, 212, 8, 32, 11, 32, 12, 32, 213,
		1, 33, 1, 33, 5, 33, 218, 8, 33, 10, 33, 12, 33, 221, 9, 33, 1, 33, 1,
		33, 1, 34, 3, 34, 226, 8, 34, 1, 34, 1, 34, 1, 35, 4, 35, 231, 8, 35, 11,
		35, 12, 35, 232, 1, 35, 1, 35, 1, 219, 0, 36, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 0, 63, 31,
		65, 32, 67, 33, 69, 34, 71, 35, 1, 0, 4, 2, 0, 10, 10, 13, 13, 3, 0, 65,
		90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0, 9, 9, 32, 32, 243, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0,
		0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1,
		0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 1, 73,
		1, 0, 0, 0, 3, 77, 1, 0, 0, 0, 5, 79, 1, 0, 0, 0, 7, 81, 1, 0, 0, 0, 9,
		84, 1, 0, 0, 0, 11, 86, 1, 0, 0, 0, 13, 92, 1, 0, 0, 0, 15, 94, 1, 0, 0,
		0, 17, 96, 1, 0, 0, 0, 19, 98, 1, 0, 0, 0, 21, 100, 1, 0, 0, 0, 23, 102,
		1, 0, 0, 0, 25, 104, 1, 0, 0, 0, 27, 115, 1, 0, 0, 0, 29, 117, 1, 0, 0,
		0, 31, 119, 1, 0, 0, 0, 33, 125, 1, 0, 0, 0, 35, 130, 1, 0, 0, 0, 37, 136,
		1, 0, 0, 0, 39, 140, 1, 0, 0, 0, 41, 142, 1, 0, 0, 0, 43, 153, 1, 0, 0,
		0, 45, 159, 1, 0, 0, 0, 47, 161, 1, 0, 0, 0, 49, 165, 1, 0, 0, 0, 51, 168,
		1, 0, 0, 0, 53, 171, 1, 0, 0, 0, 55, 175, 1, 0, 0, 0, 57, 184, 1, 0, 0,
		0, 59, 188, 1, 0, 0, 0, 61, 196, 1, 0, 0, 0, 63, 199, 1, 0, 0, 0, 65, 206,
		1, 0, 0, 0, 67, 215, 1, 0, 0, 0, 69, 225, 1, 0, 0, 0, 71, 230, 1, 0, 0,
		0, 73, 74, 5, 117, 0, 0, 74, 75, 5, 115, 0, 0, 75, 76, 5, 101, 0, 0, 76,
		2, 1, 0, 0, 0, 77, 78, 5, 123, 0, 0, 78, 4, 1, 0, 0, 0, 79, 80, 5, 125,
		0, 0, 80, 6, 1, 0, 0, 0, 81, 82, 5, 64, 0, 0, 82, 83, 5, 47, 0, 0, 83,
		8, 1, 0, 0, 0, 84, 85, 5, 47, 0, 0, 85, 10, 1, 0, 0, 0, 86, 87, 5, 116,
		0, 0, 87, 88, 5, 121, 0, 0, 88, 89, 5, 112, 0, 0, 89, 90, 5, 101, 0, 0,
		90, 91, 5, 115, 0, 0, 91, 12, 1, 0, 0, 0, 92, 93, 5, 60, 0, 0, 93, 14,
		1, 0, 0, 0, 94, 95, 5, 62, 0, 0, 95, 16, 1, 0, 0, 0, 96, 97, 5, 44, 0,
		0, 97, 18, 1, 0, 0, 0, 98, 99, 5, 91, 0, 0, 99, 20, 1, 0, 0, 0, 100, 101,
		5, 93, 0, 0, 101, 22, 1, 0, 0, 0, 102, 103, 5, 124, 0, 0, 103, 24, 1, 0,
		0, 0, 104, 105, 5, 105, 0, 0, 105, 106, 5, 110, 0, 0, 106, 107, 5, 116,
		0, 0, 107, 108, 5, 101, 0, 0, 108, 109, 5, 114, 0, 0, 109, 110, 5, 102,
		0, 0, 110, 111, 5, 97, 0, 0, 111, 112, 5, 99, 0, 0, 112, 113, 5, 101, 0,
		0, 113, 114, 5, 115, 0, 0, 114, 26, 1, 0, 0, 0, 115, 116, 5, 40, 0, 0,
		116, 28, 1, 0, 0, 0, 117, 118, 5, 41, 0, 0, 118, 30, 1, 0, 0, 0, 119, 120,
		5, 99, 0, 0, 120, 121, 5, 111, 0, 0, 121, 122, 5, 110, 0, 0, 122, 123,
		5, 115, 0, 0, 123, 124, 5, 116, 0, 0, 124, 32, 1, 0, 0, 0, 125, 126, 5,
		116, 0, 0, 126, 127, 5, 114, 0, 0, 127, 128, 5, 117, 0, 0, 128, 129, 5,
		101, 0, 0, 129, 34, 1, 0, 0, 0, 130, 131, 5, 102, 0, 0, 131, 132, 5, 97,
		0, 0, 132, 133, 5, 108, 0, 0, 133, 134, 5, 115, 0, 0, 134, 135, 5, 101,
		0, 0, 135, 36, 1, 0, 0, 0, 136, 137, 5, 110, 0, 0, 137, 138, 5, 105, 0,
		0, 138, 139, 5, 108, 0, 0, 139, 38, 1, 0, 0, 0, 140, 141, 5, 58, 0, 0,
		141, 40, 1, 0, 0, 0, 142, 143, 5, 99, 0, 0, 143, 144, 5, 111, 0, 0, 144,
		145, 5, 109, 0, 0, 145, 146, 5, 112, 0, 0, 146, 147, 5, 111, 0, 0, 147,
		148, 5, 110, 0, 0, 148, 149, 5, 101, 0, 0, 149, 150, 5, 110, 0, 0, 150,
		151, 5, 116, 0, 0, 151, 152, 5, 115, 0, 0, 152, 42, 1, 0, 0, 0, 153, 154,
		5, 110, 0, 0, 154, 155, 5, 111, 0, 0, 155, 156, 5, 100, 0, 0, 156, 157,
		5, 101, 0, 0, 157, 158, 5, 115, 0, 0, 158, 44, 1, 0, 0, 0, 159, 160, 5,
		46, 0, 0, 160, 46, 1, 0, 0, 0, 161, 162, 5, 110, 0, 0, 162, 163, 5, 101,
		0, 0, 163, 164, 5, 116, 0, 0, 164, 48, 1, 0, 0, 0, 165, 166, 5, 45, 0,
		0, 166, 167, 5, 62, 0, 0, 167, 50, 1, 0, 0, 0, 168, 169, 5, 105, 0, 0,
		169, 170, 5, 110, 0, 0, 170, 52, 1, 0, 0, 0, 171, 172, 5, 111, 0, 0, 172,
		173, 5, 117, 0, 0, 173, 174, 5, 116, 0, 0, 174, 54, 1, 0, 0, 0, 175, 176,
		5, 47, 0, 0, 176, 177, 5, 47, 0, 0, 177, 181, 1, 0, 0, 0, 178, 180, 8,
		0, 0, 0, 179, 178, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181, 179, 1, 0, 0,
		0, 181, 182, 1, 0, 0, 0, 182, 56, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184,
		185, 5, 112, 0, 0, 185, 186, 5, 117, 0, 0, 186, 187, 5, 98, 0, 0, 187,
		58, 1, 0, 0, 0, 188, 193, 3, 61, 30, 0, 189, 192, 3, 61, 30, 0, 190, 192,
		3, 63, 31, 0, 191, 189, 1, 0, 0, 0, 191, 190, 1, 0, 0, 0, 192, 195, 1,
		0, 0, 0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 60, 1, 0, 0,
		0, 195, 193, 1, 0, 0, 0, 196, 197, 7, 1, 0, 0, 197, 62, 1, 0, 0, 0, 198,
		200, 7, 2, 0, 0, 199, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 199,
		1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 64, 1, 0, 0, 0, 203, 205, 7, 2,
		0, 0, 204, 203, 1, 0, 0, 0, 205, 208, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0,
		206, 207, 1, 0, 0, 0, 207, 209, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 209,
		211, 5, 46, 0, 0, 210, 212, 7, 2, 0, 0, 211, 210, 1, 0, 0, 0, 212, 213,
		1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 66, 1, 0,
		0, 0, 215, 219, 5, 39, 0, 0, 216, 218, 9, 0, 0, 0, 217, 216, 1, 0, 0, 0,
		218, 221, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 220,
		222, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 222, 223, 5, 39, 0, 0, 223, 68,
		1, 0, 0, 0, 224, 226, 5, 13, 0, 0, 225, 224, 1, 0, 0, 0, 225, 226, 1, 0,
		0, 0, 226, 227, 1, 0, 0, 0, 227, 228, 5, 10, 0, 0, 228, 70, 1, 0, 0, 0,
		229, 231, 7, 3, 0, 0, 230, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232,
		230, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 235,
		6, 35, 0, 0, 235, 72, 1, 0, 0, 0, 10, 0, 181, 191, 193, 201, 206, 213,
		219, 225, 232, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// nevaLexerInit initializes any static state used to implement nevaLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewnevaLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func NevaLexerInit() {
	staticData := &NevaLexerLexerStaticData
	staticData.once.Do(nevalexerLexerInit)
}

// NewnevaLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewnevaLexer(input antlr.CharStream) *nevaLexer {
	NevaLexerInit()
	l := new(nevaLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &NevaLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "neva.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// nevaLexer tokens.
const (
	nevaLexerT__0       = 1
	nevaLexerT__1       = 2
	nevaLexerT__2       = 3
	nevaLexerT__3       = 4
	nevaLexerT__4       = 5
	nevaLexerT__5       = 6
	nevaLexerT__6       = 7
	nevaLexerT__7       = 8
	nevaLexerT__8       = 9
	nevaLexerT__9       = 10
	nevaLexerT__10      = 11
	nevaLexerT__11      = 12
	nevaLexerT__12      = 13
	nevaLexerT__13      = 14
	nevaLexerT__14      = 15
	nevaLexerT__15      = 16
	nevaLexerT__16      = 17
	nevaLexerT__17      = 18
	nevaLexerT__18      = 19
	nevaLexerT__19      = 20
	nevaLexerT__20      = 21
	nevaLexerT__21      = 22
	nevaLexerT__22      = 23
	nevaLexerT__23      = 24
	nevaLexerT__24      = 25
	nevaLexerT__25      = 26
	nevaLexerT__26      = 27
	nevaLexerCOMMENT    = 28
	nevaLexerPUB_KW     = 29
	nevaLexerIDENTIFIER = 30
	nevaLexerINT        = 31
	nevaLexerFLOAT      = 32
	nevaLexerSTRING     = 33
	nevaLexerNEWLINE    = 34
	nevaLexerWS         = 35
)

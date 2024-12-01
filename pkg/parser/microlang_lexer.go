// Code generated from MicroLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

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

type MicroLangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var MicroLangLexerLexerStaticData struct {
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

func microlanglexerLexerInit() {
	staticData := &MicroLangLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "';'", "'while'", "'{'", "'}'", "'if'", "'*'", "'/'", "'+'", "'-'",
		"'('", "')'", "'='", "'>'", "'<'", "'=='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "MULT", "DIV", "PLUS", "MINUS", "LPAREN", "RPAREN",
		"ASSIGN", "GREAT", "LESS", "EQUAL", "ID", "INT", "WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "MULT", "DIV", "PLUS", "MINUS",
		"LPAREN", "RPAREN", "ASSIGN", "GREAT", "LESS", "EQUAL", "ID", "INT",
		"WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 20, 121, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 5, 15, 80, 8, 15, 10, 15, 12, 15, 83, 9, 15, 1, 16, 4, 16,
		86, 8, 16, 11, 16, 12, 16, 87, 1, 17, 4, 17, 91, 8, 17, 11, 17, 12, 17,
		92, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 101, 8, 18, 10, 18,
		12, 18, 104, 9, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		19, 1, 19, 5, 19, 115, 8, 19, 10, 19, 12, 19, 118, 9, 19, 1, 19, 1, 19,
		1, 102, 0, 20, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18,
		37, 19, 39, 20, 1, 0, 5, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65,
		90, 95, 95, 97, 122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10,
		10, 13, 13, 125, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37,
		1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 1, 41, 1, 0, 0, 0, 3, 43, 1, 0, 0, 0, 5,
		49, 1, 0, 0, 0, 7, 51, 1, 0, 0, 0, 9, 53, 1, 0, 0, 0, 11, 56, 1, 0, 0,
		0, 13, 58, 1, 0, 0, 0, 15, 60, 1, 0, 0, 0, 17, 62, 1, 0, 0, 0, 19, 64,
		1, 0, 0, 0, 21, 66, 1, 0, 0, 0, 23, 68, 1, 0, 0, 0, 25, 70, 1, 0, 0, 0,
		27, 72, 1, 0, 0, 0, 29, 74, 1, 0, 0, 0, 31, 77, 1, 0, 0, 0, 33, 85, 1,
		0, 0, 0, 35, 90, 1, 0, 0, 0, 37, 96, 1, 0, 0, 0, 39, 110, 1, 0, 0, 0, 41,
		42, 5, 59, 0, 0, 42, 2, 1, 0, 0, 0, 43, 44, 5, 119, 0, 0, 44, 45, 5, 104,
		0, 0, 45, 46, 5, 105, 0, 0, 46, 47, 5, 108, 0, 0, 47, 48, 5, 101, 0, 0,
		48, 4, 1, 0, 0, 0, 49, 50, 5, 123, 0, 0, 50, 6, 1, 0, 0, 0, 51, 52, 5,
		125, 0, 0, 52, 8, 1, 0, 0, 0, 53, 54, 5, 105, 0, 0, 54, 55, 5, 102, 0,
		0, 55, 10, 1, 0, 0, 0, 56, 57, 5, 42, 0, 0, 57, 12, 1, 0, 0, 0, 58, 59,
		5, 47, 0, 0, 59, 14, 1, 0, 0, 0, 60, 61, 5, 43, 0, 0, 61, 16, 1, 0, 0,
		0, 62, 63, 5, 45, 0, 0, 63, 18, 1, 0, 0, 0, 64, 65, 5, 40, 0, 0, 65, 20,
		1, 0, 0, 0, 66, 67, 5, 41, 0, 0, 67, 22, 1, 0, 0, 0, 68, 69, 5, 61, 0,
		0, 69, 24, 1, 0, 0, 0, 70, 71, 5, 62, 0, 0, 71, 26, 1, 0, 0, 0, 72, 73,
		5, 60, 0, 0, 73, 28, 1, 0, 0, 0, 74, 75, 5, 61, 0, 0, 75, 76, 5, 61, 0,
		0, 76, 30, 1, 0, 0, 0, 77, 81, 7, 0, 0, 0, 78, 80, 7, 1, 0, 0, 79, 78,
		1, 0, 0, 0, 80, 83, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0,
		82, 32, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 84, 86, 7, 2, 0, 0, 85, 84, 1,
		0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88,
		34, 1, 0, 0, 0, 89, 91, 7, 3, 0, 0, 90, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0,
		0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95,
		6, 17, 0, 0, 95, 36, 1, 0, 0, 0, 96, 97, 5, 47, 0, 0, 97, 98, 5, 42, 0,
		0, 98, 102, 1, 0, 0, 0, 99, 101, 9, 0, 0, 0, 100, 99, 1, 0, 0, 0, 101,
		104, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 103, 105,
		1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 105, 106, 5, 42, 0, 0, 106, 107, 5, 47,
		0, 0, 107, 108, 1, 0, 0, 0, 108, 109, 6, 18, 0, 0, 109, 38, 1, 0, 0, 0,
		110, 111, 5, 47, 0, 0, 111, 112, 5, 47, 0, 0, 112, 116, 1, 0, 0, 0, 113,
		115, 8, 4, 0, 0, 114, 113, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114,
		1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 119, 1, 0, 0, 0, 118, 116, 1, 0,
		0, 0, 119, 120, 6, 19, 0, 0, 120, 40, 1, 0, 0, 0, 6, 0, 81, 87, 92, 102,
		116, 1, 6, 0, 0,
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

// MicroLangLexerInit initializes any static state used to implement MicroLangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMicroLangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MicroLangLexerInit() {
	staticData := &MicroLangLexerLexerStaticData
	staticData.once.Do(microlanglexerLexerInit)
}

// NewMicroLangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMicroLangLexer(input antlr.CharStream) *MicroLangLexer {
	MicroLangLexerInit()
	l := new(MicroLangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &MicroLangLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "MicroLang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MicroLangLexer tokens.
const (
	MicroLangLexerT__0         = 1
	MicroLangLexerT__1         = 2
	MicroLangLexerT__2         = 3
	MicroLangLexerT__3         = 4
	MicroLangLexerT__4         = 5
	MicroLangLexerMULT         = 6
	MicroLangLexerDIV          = 7
	MicroLangLexerPLUS         = 8
	MicroLangLexerMINUS        = 9
	MicroLangLexerLPAREN       = 10
	MicroLangLexerRPAREN       = 11
	MicroLangLexerASSIGN       = 12
	MicroLangLexerGREAT        = 13
	MicroLangLexerLESS         = 14
	MicroLangLexerEQUAL        = 15
	MicroLangLexerID           = 16
	MicroLangLexerINT          = 17
	MicroLangLexerWS           = 18
	MicroLangLexerCOMMENT      = 19
	MicroLangLexerLINE_COMMENT = 20
)

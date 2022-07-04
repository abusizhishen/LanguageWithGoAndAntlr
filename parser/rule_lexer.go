// Code generated from Rule.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type RuleLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var rulelexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rulelexerLexerInit() {
	staticData := &rulelexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'('", "')'", "'='", "';'", "'{'", "'}'", "'return'", "':'", "','",
		"'['", "']'", "", "'now()'", "'&&'", "'||'", "'!'", "'if'", "'elsif'",
		"'else'", "'true'", "'false'", "", "'>'", "'>='", "'=='", "'<='", "'<'",
		"'+'", "'-'", "'*'", "'/'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "NUM", "NOW", "AND",
		"OR", "NOT", "IF", "ELSIF", "ELSE", "TRUE", "FALSE", "IDENTIFY", "GT",
		"GTE", "EQ", "LTE", "LT", "ADD", "SUB", "MUL", "DIV", "WS", "Str",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "NUM", "NOW", "AND", "OR", "NOT", "IF", "ELSIF", "ELSE",
		"TRUE", "FALSE", "IDENTIFY", "GT", "GTE", "EQ", "LTE", "LT", "ADD",
		"SUB", "MUL", "DIV", "WS", "Str",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 33, 179, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 4, 11, 96, 8, 11, 11,
		11, 12, 11, 97, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		21, 4, 21, 140, 8, 21, 11, 21, 12, 21, 141, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1,
		27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 4, 31, 166, 8, 31,
		11, 31, 12, 31, 167, 1, 31, 1, 31, 1, 32, 1, 32, 4, 32, 174, 8, 32, 11,
		32, 12, 32, 175, 1, 32, 1, 32, 0, 0, 33, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29,
		15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47,
		24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65,
		33, 1, 0, 4, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10,
		13, 13, 32, 32, 6, 0, 32, 32, 39, 39, 48, 57, 65, 90, 95, 95, 97, 122,
		182, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0,
		0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0,
		0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0,
		0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1,
		0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 1, 67, 1, 0, 0, 0, 3, 69,
		1, 0, 0, 0, 5, 71, 1, 0, 0, 0, 7, 73, 1, 0, 0, 0, 9, 75, 1, 0, 0, 0, 11,
		77, 1, 0, 0, 0, 13, 79, 1, 0, 0, 0, 15, 86, 1, 0, 0, 0, 17, 88, 1, 0, 0,
		0, 19, 90, 1, 0, 0, 0, 21, 92, 1, 0, 0, 0, 23, 95, 1, 0, 0, 0, 25, 99,
		1, 0, 0, 0, 27, 105, 1, 0, 0, 0, 29, 108, 1, 0, 0, 0, 31, 111, 1, 0, 0,
		0, 33, 113, 1, 0, 0, 0, 35, 116, 1, 0, 0, 0, 37, 122, 1, 0, 0, 0, 39, 127,
		1, 0, 0, 0, 41, 132, 1, 0, 0, 0, 43, 139, 1, 0, 0, 0, 45, 143, 1, 0, 0,
		0, 47, 145, 1, 0, 0, 0, 49, 148, 1, 0, 0, 0, 51, 151, 1, 0, 0, 0, 53, 154,
		1, 0, 0, 0, 55, 156, 1, 0, 0, 0, 57, 158, 1, 0, 0, 0, 59, 160, 1, 0, 0,
		0, 61, 162, 1, 0, 0, 0, 63, 165, 1, 0, 0, 0, 65, 171, 1, 0, 0, 0, 67, 68,
		5, 40, 0, 0, 68, 2, 1, 0, 0, 0, 69, 70, 5, 41, 0, 0, 70, 4, 1, 0, 0, 0,
		71, 72, 5, 61, 0, 0, 72, 6, 1, 0, 0, 0, 73, 74, 5, 59, 0, 0, 74, 8, 1,
		0, 0, 0, 75, 76, 5, 123, 0, 0, 76, 10, 1, 0, 0, 0, 77, 78, 5, 125, 0, 0,
		78, 12, 1, 0, 0, 0, 79, 80, 5, 114, 0, 0, 80, 81, 5, 101, 0, 0, 81, 82,
		5, 116, 0, 0, 82, 83, 5, 117, 0, 0, 83, 84, 5, 114, 0, 0, 84, 85, 5, 110,
		0, 0, 85, 14, 1, 0, 0, 0, 86, 87, 5, 58, 0, 0, 87, 16, 1, 0, 0, 0, 88,
		89, 5, 44, 0, 0, 89, 18, 1, 0, 0, 0, 90, 91, 5, 91, 0, 0, 91, 20, 1, 0,
		0, 0, 92, 93, 5, 93, 0, 0, 93, 22, 1, 0, 0, 0, 94, 96, 7, 0, 0, 0, 95,
		94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0,
		0, 98, 24, 1, 0, 0, 0, 99, 100, 5, 110, 0, 0, 100, 101, 5, 111, 0, 0, 101,
		102, 5, 119, 0, 0, 102, 103, 5, 40, 0, 0, 103, 104, 5, 41, 0, 0, 104, 26,
		1, 0, 0, 0, 105, 106, 5, 38, 0, 0, 106, 107, 5, 38, 0, 0, 107, 28, 1, 0,
		0, 0, 108, 109, 5, 124, 0, 0, 109, 110, 5, 124, 0, 0, 110, 30, 1, 0, 0,
		0, 111, 112, 5, 33, 0, 0, 112, 32, 1, 0, 0, 0, 113, 114, 5, 105, 0, 0,
		114, 115, 5, 102, 0, 0, 115, 34, 1, 0, 0, 0, 116, 117, 5, 101, 0, 0, 117,
		118, 5, 108, 0, 0, 118, 119, 5, 115, 0, 0, 119, 120, 5, 105, 0, 0, 120,
		121, 5, 102, 0, 0, 121, 36, 1, 0, 0, 0, 122, 123, 5, 101, 0, 0, 123, 124,
		5, 108, 0, 0, 124, 125, 5, 115, 0, 0, 125, 126, 5, 101, 0, 0, 126, 38,
		1, 0, 0, 0, 127, 128, 5, 116, 0, 0, 128, 129, 5, 114, 0, 0, 129, 130, 5,
		117, 0, 0, 130, 131, 5, 101, 0, 0, 131, 40, 1, 0, 0, 0, 132, 133, 5, 102,
		0, 0, 133, 134, 5, 97, 0, 0, 134, 135, 5, 108, 0, 0, 135, 136, 5, 115,
		0, 0, 136, 137, 5, 101, 0, 0, 137, 42, 1, 0, 0, 0, 138, 140, 7, 1, 0, 0,
		139, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141,
		142, 1, 0, 0, 0, 142, 44, 1, 0, 0, 0, 143, 144, 5, 62, 0, 0, 144, 46, 1,
		0, 0, 0, 145, 146, 5, 62, 0, 0, 146, 147, 5, 61, 0, 0, 147, 48, 1, 0, 0,
		0, 148, 149, 5, 61, 0, 0, 149, 150, 5, 61, 0, 0, 150, 50, 1, 0, 0, 0, 151,
		152, 5, 60, 0, 0, 152, 153, 5, 61, 0, 0, 153, 52, 1, 0, 0, 0, 154, 155,
		5, 60, 0, 0, 155, 54, 1, 0, 0, 0, 156, 157, 5, 43, 0, 0, 157, 56, 1, 0,
		0, 0, 158, 159, 5, 45, 0, 0, 159, 58, 1, 0, 0, 0, 160, 161, 5, 42, 0, 0,
		161, 60, 1, 0, 0, 0, 162, 163, 5, 47, 0, 0, 163, 62, 1, 0, 0, 0, 164, 166,
		7, 2, 0, 0, 165, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 165, 1, 0,
		0, 0, 167, 168, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 170, 6, 31, 0, 0,
		170, 64, 1, 0, 0, 0, 171, 173, 5, 34, 0, 0, 172, 174, 7, 3, 0, 0, 173,
		172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 175, 176,
		1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178, 5, 34, 0, 0, 178, 66, 1, 0,
		0, 0, 5, 0, 97, 141, 167, 175, 1, 6, 0, 0,
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

// RuleLexerInit initializes any static state used to implement RuleLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewRuleLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func RuleLexerInit() {
	staticData := &rulelexerLexerStaticData
	staticData.once.Do(rulelexerLexerInit)
}

// NewRuleLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewRuleLexer(input antlr.CharStream) *RuleLexer {
	RuleLexerInit()
	l := new(RuleLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &rulelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Rule.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RuleLexer tokens.
const (
	RuleLexerT__0     = 1
	RuleLexerT__1     = 2
	RuleLexerT__2     = 3
	RuleLexerT__3     = 4
	RuleLexerT__4     = 5
	RuleLexerT__5     = 6
	RuleLexerT__6     = 7
	RuleLexerT__7     = 8
	RuleLexerT__8     = 9
	RuleLexerT__9     = 10
	RuleLexerT__10    = 11
	RuleLexerNUM      = 12
	RuleLexerNOW      = 13
	RuleLexerAND      = 14
	RuleLexerOR       = 15
	RuleLexerNOT      = 16
	RuleLexerIF       = 17
	RuleLexerELSIF    = 18
	RuleLexerELSE     = 19
	RuleLexerTRUE     = 20
	RuleLexerFALSE    = 21
	RuleLexerIDENTIFY = 22
	RuleLexerGT       = 23
	RuleLexerGTE      = 24
	RuleLexerEQ       = 25
	RuleLexerLTE      = 26
	RuleLexerLT       = 27
	RuleLexerADD      = 28
	RuleLexerSUB      = 29
	RuleLexerMUL      = 30
	RuleLexerDIV      = 31
	RuleLexerWS       = 32
	RuleLexerStr      = 33
)

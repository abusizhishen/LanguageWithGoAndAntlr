// Code generated from Rule.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type RuleParser struct {
	*antlr.BaseParser
}

var ruleParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ruleParserInit() {
	staticData := &ruleParserStaticData
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
		"boolOperate", "calculate", "compare", "logical", "compareStatement",
		"num", "boolValue", "identify", "stringValue", "calculateValue", "calculateStatement",
		"boolStatement", "valueType", "setValueStatement", "ifStatement", "elseIfStatement",
		"elseStatement", "returnStatement", "mapKey", "pair", "mapValue", "getMapOrArrayValue",
		"funCall", "array", "statement", "init",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 33, 262, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 68, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 74, 8, 4, 10,
		4, 12, 4, 77, 9, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		9, 1, 9, 3, 9, 89, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10,
		97, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 105, 8, 10,
		10, 10, 12, 10, 108, 9, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 3, 11, 118, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 124, 8,
		11, 10, 11, 12, 11, 127, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 3, 12, 137, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 143,
		8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 149, 8, 14, 10, 14, 12, 14, 152,
		9, 14, 1, 14, 1, 14, 5, 14, 156, 8, 14, 10, 14, 12, 14, 159, 9, 14, 1,
		14, 3, 14, 162, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 168, 8, 15, 10,
		15, 12, 15, 171, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 5, 16, 178,
		8, 16, 10, 16, 12, 16, 181, 9, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 3, 17, 189, 8, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20,
		1, 20, 1, 20, 1, 20, 5, 20, 201, 8, 20, 10, 20, 12, 20, 204, 9, 20, 3,
		20, 206, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 4, 21,
		215, 8, 21, 11, 21, 12, 21, 216, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5,
		22, 224, 8, 22, 10, 22, 12, 22, 227, 9, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 23, 5, 23, 235, 8, 23, 10, 23, 12, 23, 238, 9, 23, 3, 23, 240,
		8, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3,
		24, 251, 8, 24, 1, 25, 5, 25, 254, 8, 25, 10, 25, 12, 25, 257, 9, 25, 1,
		25, 3, 25, 260, 8, 25, 1, 25, 0, 3, 8, 20, 22, 26, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
		50, 0, 7, 1, 0, 14, 16, 1, 0, 28, 31, 1, 0, 23, 27, 1, 0, 20, 21, 1, 0,
		30, 31, 1, 0, 28, 29, 3, 0, 12, 12, 22, 22, 33, 33, 275, 0, 52, 1, 0, 0,
		0, 2, 54, 1, 0, 0, 0, 4, 56, 1, 0, 0, 0, 6, 58, 1, 0, 0, 0, 8, 67, 1, 0,
		0, 0, 10, 78, 1, 0, 0, 0, 12, 80, 1, 0, 0, 0, 14, 82, 1, 0, 0, 0, 16, 84,
		1, 0, 0, 0, 18, 88, 1, 0, 0, 0, 20, 96, 1, 0, 0, 0, 22, 117, 1, 0, 0, 0,
		24, 136, 1, 0, 0, 0, 26, 138, 1, 0, 0, 0, 28, 144, 1, 0, 0, 0, 30, 163,
		1, 0, 0, 0, 32, 174, 1, 0, 0, 0, 34, 188, 1, 0, 0, 0, 36, 190, 1, 0, 0,
		0, 38, 192, 1, 0, 0, 0, 40, 196, 1, 0, 0, 0, 42, 209, 1, 0, 0, 0, 44, 218,
		1, 0, 0, 0, 46, 230, 1, 0, 0, 0, 48, 250, 1, 0, 0, 0, 50, 255, 1, 0, 0,
		0, 52, 53, 7, 0, 0, 0, 53, 1, 1, 0, 0, 0, 54, 55, 7, 1, 0, 0, 55, 3, 1,
		0, 0, 0, 56, 57, 7, 2, 0, 0, 57, 5, 1, 0, 0, 0, 58, 59, 7, 0, 0, 0, 59,
		7, 1, 0, 0, 0, 60, 61, 6, 4, -1, 0, 61, 68, 3, 18, 9, 0, 62, 68, 3, 20,
		10, 0, 63, 64, 5, 1, 0, 0, 64, 65, 3, 8, 4, 0, 65, 66, 5, 2, 0, 0, 66,
		68, 1, 0, 0, 0, 67, 60, 1, 0, 0, 0, 67, 62, 1, 0, 0, 0, 67, 63, 1, 0, 0,
		0, 68, 75, 1, 0, 0, 0, 69, 70, 10, 4, 0, 0, 70, 71, 3, 4, 2, 0, 71, 72,
		3, 8, 4, 5, 72, 74, 1, 0, 0, 0, 73, 69, 1, 0, 0, 0, 74, 77, 1, 0, 0, 0,
		75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 9, 1, 0, 0, 0, 77, 75, 1, 0,
		0, 0, 78, 79, 5, 12, 0, 0, 79, 11, 1, 0, 0, 0, 80, 81, 7, 3, 0, 0, 81,
		13, 1, 0, 0, 0, 82, 83, 5, 22, 0, 0, 83, 15, 1, 0, 0, 0, 84, 85, 5, 33,
		0, 0, 85, 17, 1, 0, 0, 0, 86, 89, 3, 14, 7, 0, 87, 89, 3, 10, 5, 0, 88,
		86, 1, 0, 0, 0, 88, 87, 1, 0, 0, 0, 89, 19, 1, 0, 0, 0, 90, 91, 6, 10,
		-1, 0, 91, 97, 3, 18, 9, 0, 92, 93, 5, 1, 0, 0, 93, 94, 3, 20, 10, 0, 94,
		95, 5, 2, 0, 0, 95, 97, 1, 0, 0, 0, 96, 90, 1, 0, 0, 0, 96, 92, 1, 0, 0,
		0, 97, 106, 1, 0, 0, 0, 98, 99, 10, 4, 0, 0, 99, 100, 7, 4, 0, 0, 100,
		105, 3, 20, 10, 5, 101, 102, 10, 3, 0, 0, 102, 103, 7, 5, 0, 0, 103, 105,
		3, 20, 10, 4, 104, 98, 1, 0, 0, 0, 104, 101, 1, 0, 0, 0, 105, 108, 1, 0,
		0, 0, 106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 21, 1, 0, 0, 0,
		108, 106, 1, 0, 0, 0, 109, 110, 6, 11, -1, 0, 110, 118, 7, 3, 0, 0, 111,
		118, 5, 22, 0, 0, 112, 118, 3, 8, 4, 0, 113, 114, 5, 1, 0, 0, 114, 115,
		3, 22, 11, 0, 115, 116, 5, 2, 0, 0, 116, 118, 1, 0, 0, 0, 117, 109, 1,
		0, 0, 0, 117, 111, 1, 0, 0, 0, 117, 112, 1, 0, 0, 0, 117, 113, 1, 0, 0,
		0, 118, 125, 1, 0, 0, 0, 119, 120, 10, 5, 0, 0, 120, 121, 3, 0, 0, 0, 121,
		122, 3, 22, 11, 6, 122, 124, 1, 0, 0, 0, 123, 119, 1, 0, 0, 0, 124, 127,
		1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 23, 1, 0,
		0, 0, 127, 125, 1, 0, 0, 0, 128, 137, 3, 12, 6, 0, 129, 137, 3, 14, 7,
		0, 130, 137, 3, 10, 5, 0, 131, 137, 3, 20, 10, 0, 132, 137, 3, 40, 20,
		0, 133, 137, 3, 42, 21, 0, 134, 137, 3, 16, 8, 0, 135, 137, 3, 46, 23,
		0, 136, 128, 1, 0, 0, 0, 136, 129, 1, 0, 0, 0, 136, 130, 1, 0, 0, 0, 136,
		131, 1, 0, 0, 0, 136, 132, 1, 0, 0, 0, 136, 133, 1, 0, 0, 0, 136, 134,
		1, 0, 0, 0, 136, 135, 1, 0, 0, 0, 137, 25, 1, 0, 0, 0, 138, 139, 5, 22,
		0, 0, 139, 140, 5, 3, 0, 0, 140, 142, 3, 24, 12, 0, 141, 143, 5, 4, 0,
		0, 142, 141, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 27, 1, 0, 0, 0, 144,
		145, 5, 17, 0, 0, 145, 146, 3, 22, 11, 0, 146, 150, 5, 5, 0, 0, 147, 149,
		3, 48, 24, 0, 148, 147, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1,
		0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 153, 1, 0, 0, 0, 152, 150, 1, 0, 0,
		0, 153, 157, 5, 6, 0, 0, 154, 156, 3, 30, 15, 0, 155, 154, 1, 0, 0, 0,
		156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158,
		161, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 162, 3, 32, 16, 0, 161, 160,
		1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 29, 1, 0, 0, 0, 163, 164, 5, 18,
		0, 0, 164, 165, 3, 22, 11, 0, 165, 169, 5, 5, 0, 0, 166, 168, 3, 48, 24,
		0, 167, 166, 1, 0, 0, 0, 168, 171, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169,
		170, 1, 0, 0, 0, 170, 172, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 172, 173,
		5, 6, 0, 0, 173, 31, 1, 0, 0, 0, 174, 175, 5, 19, 0, 0, 175, 179, 5, 5,
		0, 0, 176, 178, 3, 48, 24, 0, 177, 176, 1, 0, 0, 0, 178, 181, 1, 0, 0,
		0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 182, 1, 0, 0, 0, 181,
		179, 1, 0, 0, 0, 182, 183, 5, 6, 0, 0, 183, 33, 1, 0, 0, 0, 184, 185, 5,
		7, 0, 0, 185, 189, 3, 24, 12, 0, 186, 189, 3, 20, 10, 0, 187, 189, 3, 22,
		11, 0, 188, 184, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 188, 187, 1, 0, 0, 0,
		189, 35, 1, 0, 0, 0, 190, 191, 7, 6, 0, 0, 191, 37, 1, 0, 0, 0, 192, 193,
		3, 36, 18, 0, 193, 194, 5, 8, 0, 0, 194, 195, 3, 24, 12, 0, 195, 39, 1,
		0, 0, 0, 196, 205, 5, 5, 0, 0, 197, 202, 3, 38, 19, 0, 198, 199, 5, 9,
		0, 0, 199, 201, 3, 38, 19, 0, 200, 198, 1, 0, 0, 0, 201, 204, 1, 0, 0,
		0, 202, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 206, 1, 0, 0, 0, 204,
		202, 1, 0, 0, 0, 205, 197, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 207,
		1, 0, 0, 0, 207, 208, 5, 6, 0, 0, 208, 41, 1, 0, 0, 0, 209, 214, 3, 14,
		7, 0, 210, 211, 5, 10, 0, 0, 211, 212, 3, 36, 18, 0, 212, 213, 5, 11, 0,
		0, 213, 215, 1, 0, 0, 0, 214, 210, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216,
		214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 43, 1, 0, 0, 0, 218, 219, 3,
		14, 7, 0, 219, 220, 5, 1, 0, 0, 220, 225, 3, 24, 12, 0, 221, 222, 5, 9,
		0, 0, 222, 224, 3, 24, 12, 0, 223, 221, 1, 0, 0, 0, 224, 227, 1, 0, 0,
		0, 225, 223, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 228, 1, 0, 0, 0, 227,
		225, 1, 0, 0, 0, 228, 229, 5, 2, 0, 0, 229, 45, 1, 0, 0, 0, 230, 239, 5,
		10, 0, 0, 231, 236, 3, 24, 12, 0, 232, 233, 5, 9, 0, 0, 233, 235, 3, 24,
		12, 0, 234, 232, 1, 0, 0, 0, 235, 238, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0,
		236, 237, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 239,
		231, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 242,
		5, 11, 0, 0, 242, 47, 1, 0, 0, 0, 243, 251, 3, 44, 22, 0, 244, 251, 3,
		20, 10, 0, 245, 251, 3, 22, 11, 0, 246, 251, 3, 8, 4, 0, 247, 251, 3, 28,
		14, 0, 248, 251, 3, 26, 13, 0, 249, 251, 3, 34, 17, 0, 250, 243, 1, 0,
		0, 0, 250, 244, 1, 0, 0, 0, 250, 245, 1, 0, 0, 0, 250, 246, 1, 0, 0, 0,
		250, 247, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 250, 249, 1, 0, 0, 0, 251,
		49, 1, 0, 0, 0, 252, 254, 3, 48, 24, 0, 253, 252, 1, 0, 0, 0, 254, 257,
		1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 259, 1, 0,
		0, 0, 257, 255, 1, 0, 0, 0, 258, 260, 5, 0, 0, 1, 259, 258, 1, 0, 0, 0,
		259, 260, 1, 0, 0, 0, 260, 51, 1, 0, 0, 0, 25, 67, 75, 88, 96, 104, 106,
		117, 125, 136, 142, 150, 157, 161, 169, 179, 188, 202, 205, 216, 225, 236,
		239, 250, 255, 259,
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

// RuleParserInit initializes any static state used to implement RuleParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRuleParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RuleParserInit() {
	staticData := &ruleParserStaticData
	staticData.once.Do(ruleParserInit)
}

// NewRuleParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRuleParser(input antlr.TokenStream) *RuleParser {
	RuleParserInit()
	this := new(RuleParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ruleParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Rule.g4"

	return this
}

// RuleParser tokens.
const (
	RuleParserEOF      = antlr.TokenEOF
	RuleParserT__0     = 1
	RuleParserT__1     = 2
	RuleParserT__2     = 3
	RuleParserT__3     = 4
	RuleParserT__4     = 5
	RuleParserT__5     = 6
	RuleParserT__6     = 7
	RuleParserT__7     = 8
	RuleParserT__8     = 9
	RuleParserT__9     = 10
	RuleParserT__10    = 11
	RuleParserNUM      = 12
	RuleParserNOW      = 13
	RuleParserAND      = 14
	RuleParserOR       = 15
	RuleParserNOT      = 16
	RuleParserIF       = 17
	RuleParserELSIF    = 18
	RuleParserELSE     = 19
	RuleParserTRUE     = 20
	RuleParserFALSE    = 21
	RuleParserIDENTIFY = 22
	RuleParserGT       = 23
	RuleParserGTE      = 24
	RuleParserEQ       = 25
	RuleParserLTE      = 26
	RuleParserLT       = 27
	RuleParserADD      = 28
	RuleParserSUB      = 29
	RuleParserMUL      = 30
	RuleParserDIV      = 31
	RuleParserWS       = 32
	RuleParserStr      = 33
)

// RuleParser rules.
const (
	RuleParserRULE_boolOperate        = 0
	RuleParserRULE_calculate          = 1
	RuleParserRULE_compare            = 2
	RuleParserRULE_logical            = 3
	RuleParserRULE_compareStatement   = 4
	RuleParserRULE_num                = 5
	RuleParserRULE_boolValue          = 6
	RuleParserRULE_identify           = 7
	RuleParserRULE_stringValue        = 8
	RuleParserRULE_calculateValue     = 9
	RuleParserRULE_calculateStatement = 10
	RuleParserRULE_boolStatement      = 11
	RuleParserRULE_valueType          = 12
	RuleParserRULE_setValueStatement  = 13
	RuleParserRULE_ifStatement        = 14
	RuleParserRULE_elseIfStatement    = 15
	RuleParserRULE_elseStatement      = 16
	RuleParserRULE_returnStatement    = 17
	RuleParserRULE_mapKey             = 18
	RuleParserRULE_pair               = 19
	RuleParserRULE_mapValue           = 20
	RuleParserRULE_getMapOrArrayValue = 21
	RuleParserRULE_funCall            = 22
	RuleParserRULE_array              = 23
	RuleParserRULE_statement          = 24
	RuleParserRULE_init               = 25
)

// IBoolOperateContext is an interface to support dynamic dispatch.
type IBoolOperateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolOperateContext differentiates from other interfaces.
	IsBoolOperateContext()
}

type BoolOperateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolOperateContext() *BoolOperateContext {
	var p = new(BoolOperateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_boolOperate
	return p
}

func (*BoolOperateContext) IsBoolOperateContext() {}

func NewBoolOperateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolOperateContext {
	var p = new(BoolOperateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_boolOperate

	return p
}

func (s *BoolOperateContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolOperateContext) AND() antlr.TerminalNode {
	return s.GetToken(RuleParserAND, 0)
}

func (s *BoolOperateContext) OR() antlr.TerminalNode {
	return s.GetToken(RuleParserOR, 0)
}

func (s *BoolOperateContext) NOT() antlr.TerminalNode {
	return s.GetToken(RuleParserNOT, 0)
}

func (s *BoolOperateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolOperateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolOperateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterBoolOperate(s)
	}
}

func (s *BoolOperateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitBoolOperate(s)
	}
}

func (p *RuleParser) BoolOperate() (localctx IBoolOperateContext) {
	this := p
	_ = this

	localctx = NewBoolOperateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RuleParserRULE_boolOperate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserAND)|(1<<RuleParserOR)|(1<<RuleParserNOT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICalculateContext is an interface to support dynamic dispatch.
type ICalculateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCalculateContext differentiates from other interfaces.
	IsCalculateContext()
}

type CalculateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalculateContext() *CalculateContext {
	var p = new(CalculateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_calculate
	return p
}

func (*CalculateContext) IsCalculateContext() {}

func NewCalculateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalculateContext {
	var p = new(CalculateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_calculate

	return p
}

func (s *CalculateContext) GetParser() antlr.Parser { return s.parser }

func (s *CalculateContext) ADD() antlr.TerminalNode {
	return s.GetToken(RuleParserADD, 0)
}

func (s *CalculateContext) SUB() antlr.TerminalNode {
	return s.GetToken(RuleParserSUB, 0)
}

func (s *CalculateContext) MUL() antlr.TerminalNode {
	return s.GetToken(RuleParserMUL, 0)
}

func (s *CalculateContext) DIV() antlr.TerminalNode {
	return s.GetToken(RuleParserDIV, 0)
}

func (s *CalculateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalculateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalculateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCalculate(s)
	}
}

func (s *CalculateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCalculate(s)
	}
}

func (p *RuleParser) Calculate() (localctx ICalculateContext) {
	this := p
	_ = this

	localctx = NewCalculateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RuleParserRULE_calculate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserADD)|(1<<RuleParserSUB)|(1<<RuleParserMUL)|(1<<RuleParserDIV))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICompareContext is an interface to support dynamic dispatch.
type ICompareContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompareContext differentiates from other interfaces.
	IsCompareContext()
}

type CompareContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareContext() *CompareContext {
	var p = new(CompareContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_compare
	return p
}

func (*CompareContext) IsCompareContext() {}

func NewCompareContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareContext {
	var p = new(CompareContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_compare

	return p
}

func (s *CompareContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareContext) GT() antlr.TerminalNode {
	return s.GetToken(RuleParserGT, 0)
}

func (s *CompareContext) GTE() antlr.TerminalNode {
	return s.GetToken(RuleParserGTE, 0)
}

func (s *CompareContext) EQ() antlr.TerminalNode {
	return s.GetToken(RuleParserEQ, 0)
}

func (s *CompareContext) LTE() antlr.TerminalNode {
	return s.GetToken(RuleParserLTE, 0)
}

func (s *CompareContext) LT() antlr.TerminalNode {
	return s.GetToken(RuleParserLT, 0)
}

func (s *CompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCompare(s)
	}
}

func (s *CompareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCompare(s)
	}
}

func (p *RuleParser) Compare() (localctx ICompareContext) {
	this := p
	_ = this

	localctx = NewCompareContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RuleParserRULE_compare)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserGT)|(1<<RuleParserGTE)|(1<<RuleParserEQ)|(1<<RuleParserLTE)|(1<<RuleParserLT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILogicalContext is an interface to support dynamic dispatch.
type ILogicalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicalContext differentiates from other interfaces.
	IsLogicalContext()
}

type LogicalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalContext() *LogicalContext {
	var p = new(LogicalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_logical
	return p
}

func (*LogicalContext) IsLogicalContext() {}

func NewLogicalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalContext {
	var p = new(LogicalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_logical

	return p
}

func (s *LogicalContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalContext) AND() antlr.TerminalNode {
	return s.GetToken(RuleParserAND, 0)
}

func (s *LogicalContext) OR() antlr.TerminalNode {
	return s.GetToken(RuleParserOR, 0)
}

func (s *LogicalContext) NOT() antlr.TerminalNode {
	return s.GetToken(RuleParserNOT, 0)
}

func (s *LogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterLogical(s)
	}
}

func (s *LogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitLogical(s)
	}
}

func (p *RuleParser) Logical() (localctx ILogicalContext) {
	this := p
	_ = this

	localctx = NewLogicalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RuleParserRULE_logical)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserAND)|(1<<RuleParserOR)|(1<<RuleParserNOT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICompareStatementContext is an interface to support dynamic dispatch.
type ICompareStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompareStatementContext differentiates from other interfaces.
	IsCompareStatementContext()
}

type CompareStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareStatementContext() *CompareStatementContext {
	var p = new(CompareStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_compareStatement
	return p
}

func (*CompareStatementContext) IsCompareStatementContext() {}

func NewCompareStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareStatementContext {
	var p = new(CompareStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_compareStatement

	return p
}

func (s *CompareStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareStatementContext) CopyFrom(ctx *CompareStatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CompareStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type COMPAREXContext struct {
	*CompareStatementContext
}

func NewCOMPAREXContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *COMPAREXContext {
	var p = new(COMPAREXContext)

	p.CompareStatementContext = NewEmptyCompareStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CompareStatementContext))

	return p
}

func (s *COMPAREXContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *COMPAREXContext) CompareStatement() ICompareStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompareStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompareStatementContext)
}

func (s *COMPAREXContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCOMPAREX(s)
	}
}

func (s *COMPAREXContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCOMPAREX(s)
	}
}

type COMPAREContext struct {
	*CompareStatementContext
	op ICompareContext
}

func NewCOMPAREContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *COMPAREContext {
	var p = new(COMPAREContext)

	p.CompareStatementContext = NewEmptyCompareStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CompareStatementContext))

	return p
}

func (s *COMPAREContext) GetOp() ICompareContext { return s.op }

func (s *COMPAREContext) SetOp(v ICompareContext) { s.op = v }

func (s *COMPAREContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *COMPAREContext) AllCompareStatement() []ICompareStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICompareStatementContext); ok {
			len++
		}
	}

	tst := make([]ICompareStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICompareStatementContext); ok {
			tst[i] = t.(ICompareStatementContext)
			i++
		}
	}

	return tst
}

func (s *COMPAREContext) CompareStatement(i int) ICompareStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompareStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompareStatementContext)
}

func (s *COMPAREContext) Compare() ICompareContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompareContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompareContext)
}

func (s *COMPAREContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCOMPARE(s)
	}
}

func (s *COMPAREContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCOMPARE(s)
	}
}

type CalcuContext struct {
	*CompareStatementContext
}

func NewCalcuContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CalcuContext {
	var p = new(CalcuContext)

	p.CompareStatementContext = NewEmptyCompareStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CompareStatementContext))

	return p
}

func (s *CalcuContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalcuContext) CalculateStatement() ICalculateStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalculateStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
}

func (s *CalcuContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCalcu(s)
	}
}

func (s *CalcuContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCalcu(s)
	}
}

type ITEMCOMPContext struct {
	*CompareStatementContext
}

func NewITEMCOMPContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ITEMCOMPContext {
	var p = new(ITEMCOMPContext)

	p.CompareStatementContext = NewEmptyCompareStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CompareStatementContext))

	return p
}

func (s *ITEMCOMPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ITEMCOMPContext) CalculateValue() ICalculateValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalculateValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalculateValueContext)
}

func (s *ITEMCOMPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterITEMCOMP(s)
	}
}

func (s *ITEMCOMPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitITEMCOMP(s)
	}
}

func (p *RuleParser) CompareStatement() (localctx ICompareStatementContext) {
	return p.compareStatement(0)
}

func (p *RuleParser) compareStatement(_p int) (localctx ICompareStatementContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewCompareStatementContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICompareStatementContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, RuleParserRULE_compareStatement, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewITEMCOMPContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(61)
			p.CalculateValue()
		}

	case 2:
		localctx = NewCalcuContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(62)
			p.calculateStatement(0)
		}

	case 3:
		localctx = NewCOMPAREXContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(63)
			p.Match(RuleParserT__0)
		}
		{
			p.SetState(64)
			p.compareStatement(0)
		}
		{
			p.SetState(65)
			p.Match(RuleParserT__1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCOMPAREContext(p, NewCompareStatementContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_compareStatement)
			p.SetState(69)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
			}
			{
				p.SetState(70)

				var _x = p.Compare()

				localctx.(*COMPAREContext).op = _x
			}
			{
				p.SetState(71)
				p.compareStatement(5)
			}

		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// INumContext is an interface to support dynamic dispatch.
type INumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumContext differentiates from other interfaces.
	IsNumContext()
}

type NumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumContext() *NumContext {
	var p = new(NumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_num
	return p
}

func (*NumContext) IsNumContext() {}

func NewNumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumContext {
	var p = new(NumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_num

	return p
}

func (s *NumContext) GetParser() antlr.Parser { return s.parser }

func (s *NumContext) CopyFrom(ctx *NumContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NUMContext struct {
	*NumContext
}

func NewNUMContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NUMContext {
	var p = new(NUMContext)

	p.NumContext = NewEmptyNumContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumContext))

	return p
}

func (s *NUMContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NUMContext) NUM() antlr.TerminalNode {
	return s.GetToken(RuleParserNUM, 0)
}

func (s *NUMContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterNUM(s)
	}
}

func (s *NUMContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitNUM(s)
	}
}

func (p *RuleParser) Num() (localctx INumContext) {
	this := p
	_ = this

	localctx = NewNumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RuleParserRULE_num)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewNUMContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(RuleParserNUM)
	}

	return localctx
}

// IBoolValueContext is an interface to support dynamic dispatch.
type IBoolValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolValueContext differentiates from other interfaces.
	IsBoolValueContext()
}

type BoolValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolValueContext() *BoolValueContext {
	var p = new(BoolValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_boolValue
	return p
}

func (*BoolValueContext) IsBoolValueContext() {}

func NewBoolValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolValueContext {
	var p = new(BoolValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_boolValue

	return p
}

func (s *BoolValueContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(RuleParserTRUE, 0)
}

func (s *BoolValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(RuleParserFALSE, 0)
}

func (s *BoolValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterBoolValue(s)
	}
}

func (s *BoolValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitBoolValue(s)
	}
}

func (p *RuleParser) BoolValue() (localctx IBoolValueContext) {
	this := p
	_ = this

	localctx = NewBoolValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RuleParserRULE_boolValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RuleParserTRUE || _la == RuleParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIdentifyContext is an interface to support dynamic dispatch.
type IIdentifyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifyContext differentiates from other interfaces.
	IsIdentifyContext()
}

type IdentifyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifyContext() *IdentifyContext {
	var p = new(IdentifyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_identify
	return p
}

func (*IdentifyContext) IsIdentifyContext() {}

func NewIdentifyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifyContext {
	var p = new(IdentifyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_identify

	return p
}

func (s *IdentifyContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifyContext) CopyFrom(ctx *IdentifyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *IdentifyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IDENTIFYContext struct {
	*IdentifyContext
}

func NewIDENTIFYContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IDENTIFYContext {
	var p = new(IDENTIFYContext)

	p.IdentifyContext = NewEmptyIdentifyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IdentifyContext))

	return p
}

func (s *IDENTIFYContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IDENTIFYContext) IDENTIFY() antlr.TerminalNode {
	return s.GetToken(RuleParserIDENTIFY, 0)
}

func (s *IDENTIFYContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterIDENTIFY(s)
	}
}

func (s *IDENTIFYContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitIDENTIFY(s)
	}
}

func (p *RuleParser) Identify() (localctx IIdentifyContext) {
	this := p
	_ = this

	localctx = NewIdentifyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RuleParserRULE_identify)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewIDENTIFYContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(RuleParserIDENTIFY)
	}

	return localctx
}

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringValueContext differentiates from other interfaces.
	IsStringValueContext()
}

type StringValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringValueContext() *StringValueContext {
	var p = new(StringValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_stringValue
	return p
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_stringValue

	return p
}

func (s *StringValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValueContext) CopyFrom(ctx *StringValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type STRINGContext struct {
	*StringValueContext
}

func NewSTRINGContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *STRINGContext {
	var p = new(STRINGContext)

	p.StringValueContext = NewEmptyStringValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StringValueContext))

	return p
}

func (s *STRINGContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *STRINGContext) Str() antlr.TerminalNode {
	return s.GetToken(RuleParserStr, 0)
}

func (s *STRINGContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterSTRING(s)
	}
}

func (s *STRINGContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitSTRING(s)
	}
}

func (p *RuleParser) StringValue() (localctx IStringValueContext) {
	this := p
	_ = this

	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RuleParserRULE_stringValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewSTRINGContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(RuleParserStr)
	}

	return localctx
}

// ICalculateValueContext is an interface to support dynamic dispatch.
type ICalculateValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCalculateValueContext differentiates from other interfaces.
	IsCalculateValueContext()
}

type CalculateValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalculateValueContext() *CalculateValueContext {
	var p = new(CalculateValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_calculateValue
	return p
}

func (*CalculateValueContext) IsCalculateValueContext() {}

func NewCalculateValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalculateValueContext {
	var p = new(CalculateValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_calculateValue

	return p
}

func (s *CalculateValueContext) GetParser() antlr.Parser { return s.parser }

func (s *CalculateValueContext) Identify() IIdentifyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifyContext)
}

func (s *CalculateValueContext) Num() INumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumContext)
}

func (s *CalculateValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalculateValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalculateValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCalculateValue(s)
	}
}

func (s *CalculateValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCalculateValue(s)
	}
}

func (p *RuleParser) CalculateValue() (localctx ICalculateValueContext) {
	this := p
	_ = this

	localctx = NewCalculateValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RuleParserRULE_calculateValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(88)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuleParserIDENTIFY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Identify()
		}

	case RuleParserNUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
			p.Num()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICalculateStatementContext is an interface to support dynamic dispatch.
type ICalculateStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCalculateStatementContext differentiates from other interfaces.
	IsCalculateStatementContext()
}

type CalculateStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalculateStatementContext() *CalculateStatementContext {
	var p = new(CalculateStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_calculateStatement
	return p
}

func (*CalculateStatementContext) IsCalculateStatementContext() {}

func NewCalculateStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalculateStatementContext {
	var p = new(CalculateStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_calculateStatement

	return p
}

func (s *CalculateStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CalculateStatementContext) CopyFrom(ctx *CalculateStatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CalculateStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalculateStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ITEMCALCUContext struct {
	*CalculateStatementContext
}

func NewITEMCALCUContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ITEMCALCUContext {
	var p = new(ITEMCALCUContext)

	p.CalculateStatementContext = NewEmptyCalculateStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CalculateStatementContext))

	return p
}

func (s *ITEMCALCUContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ITEMCALCUContext) CalculateValue() ICalculateValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalculateValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalculateValueContext)
}

func (s *ITEMCALCUContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterITEMCALCU(s)
	}
}

func (s *ITEMCALCUContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitITEMCALCU(s)
	}
}

type ADDSUBContext struct {
	*CalculateStatementContext
	op antlr.Token
}

func NewADDSUBContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ADDSUBContext {
	var p = new(ADDSUBContext)

	p.CalculateStatementContext = NewEmptyCalculateStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CalculateStatementContext))

	return p
}

func (s *ADDSUBContext) GetOp() antlr.Token { return s.op }

func (s *ADDSUBContext) SetOp(v antlr.Token) { s.op = v }

func (s *ADDSUBContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ADDSUBContext) AllCalculateStatement() []ICalculateStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICalculateStatementContext); ok {
			len++
		}
	}

	tst := make([]ICalculateStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICalculateStatementContext); ok {
			tst[i] = t.(ICalculateStatementContext)
			i++
		}
	}

	return tst
}

func (s *ADDSUBContext) CalculateStatement(i int) ICalculateStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalculateStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
}

func (s *ADDSUBContext) ADD() antlr.TerminalNode {
	return s.GetToken(RuleParserADD, 0)
}

func (s *ADDSUBContext) SUB() antlr.TerminalNode {
	return s.GetToken(RuleParserSUB, 0)
}

func (s *ADDSUBContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterADDSUB(s)
	}
}

func (s *ADDSUBContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitADDSUB(s)
	}
}

type MULDIVContext struct {
	*CalculateStatementContext
	op antlr.Token
}

func NewMULDIVContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MULDIVContext {
	var p = new(MULDIVContext)

	p.CalculateStatementContext = NewEmptyCalculateStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CalculateStatementContext))

	return p
}

func (s *MULDIVContext) GetOp() antlr.Token { return s.op }

func (s *MULDIVContext) SetOp(v antlr.Token) { s.op = v }

func (s *MULDIVContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MULDIVContext) AllCalculateStatement() []ICalculateStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICalculateStatementContext); ok {
			len++
		}
	}

	tst := make([]ICalculateStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICalculateStatementContext); ok {
			tst[i] = t.(ICalculateStatementContext)
			i++
		}
	}

	return tst
}

func (s *MULDIVContext) CalculateStatement(i int) ICalculateStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalculateStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
}

func (s *MULDIVContext) MUL() antlr.TerminalNode {
	return s.GetToken(RuleParserMUL, 0)
}

func (s *MULDIVContext) DIV() antlr.TerminalNode {
	return s.GetToken(RuleParserDIV, 0)
}

func (s *MULDIVContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterMULDIV(s)
	}
}

func (s *MULDIVContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitMULDIV(s)
	}
}

type CALCULATEXContext struct {
	*CalculateStatementContext
}

func NewCALCULATEXContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CALCULATEXContext {
	var p = new(CALCULATEXContext)

	p.CalculateStatementContext = NewEmptyCalculateStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CalculateStatementContext))

	return p
}

func (s *CALCULATEXContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CALCULATEXContext) CalculateStatement() ICalculateStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalculateStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
}

func (s *CALCULATEXContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCALCULATEX(s)
	}
}

func (s *CALCULATEXContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCALCULATEX(s)
	}
}

func (p *RuleParser) CalculateStatement() (localctx ICalculateStatementContext) {
	return p.calculateStatement(0)
}

func (p *RuleParser) calculateStatement(_p int) (localctx ICalculateStatementContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewCalculateStatementContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICalculateStatementContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, RuleParserRULE_calculateStatement, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(96)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuleParserNUM, RuleParserIDENTIFY:
		localctx = NewITEMCALCUContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(91)
			p.CalculateValue()
		}

	case RuleParserT__0:
		localctx = NewCALCULATEXContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(92)
			p.Match(RuleParserT__0)
		}
		{
			p.SetState(93)
			p.calculateStatement(0)
		}
		{
			p.SetState(94)
			p.Match(RuleParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMULDIVContext(p, NewCalculateStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_calculateStatement)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(99)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MULDIVContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RuleParserMUL || _la == RuleParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MULDIVContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(100)
					p.calculateStatement(5)
				}

			case 2:
				localctx = NewADDSUBContext(p, NewCalculateStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_calculateStatement)
				p.SetState(101)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(102)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ADDSUBContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RuleParserADD || _la == RuleParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ADDSUBContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(103)
					p.calculateStatement(4)
				}

			}

		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IBoolStatementContext is an interface to support dynamic dispatch.
type IBoolStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolStatementContext differentiates from other interfaces.
	IsBoolStatementContext()
}

type BoolStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolStatementContext() *BoolStatementContext {
	var p = new(BoolStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_boolStatement
	return p
}

func (*BoolStatementContext) IsBoolStatementContext() {}

func NewBoolStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolStatementContext {
	var p = new(BoolStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_boolStatement

	return p
}

func (s *BoolStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolStatementContext) CopyFrom(ctx *BoolStatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BoolStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BOOLContext struct {
	*BoolStatementContext
}

func NewBOOLContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BOOLContext {
	var p = new(BOOLContext)

	p.BoolStatementContext = NewEmptyBoolStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolStatementContext))

	return p
}

func (s *BOOLContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BOOLContext) TRUE() antlr.TerminalNode {
	return s.GetToken(RuleParserTRUE, 0)
}

func (s *BOOLContext) FALSE() antlr.TerminalNode {
	return s.GetToken(RuleParserFALSE, 0)
}

func (s *BOOLContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterBOOL(s)
	}
}

func (s *BOOLContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitBOOL(s)
	}
}

type BOOLOPContext struct {
	*BoolStatementContext
	op IBoolOperateContext
}

func NewBOOLOPContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BOOLOPContext {
	var p = new(BOOLOPContext)

	p.BoolStatementContext = NewEmptyBoolStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolStatementContext))

	return p
}

func (s *BOOLOPContext) GetOp() IBoolOperateContext { return s.op }

func (s *BOOLOPContext) SetOp(v IBoolOperateContext) { s.op = v }

func (s *BOOLOPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BOOLOPContext) AllBoolStatement() []IBoolStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolStatementContext); ok {
			len++
		}
	}

	tst := make([]IBoolStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolStatementContext); ok {
			tst[i] = t.(IBoolStatementContext)
			i++
		}
	}

	return tst
}

func (s *BOOLOPContext) BoolStatement(i int) IBoolStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolStatementContext)
}

func (s *BOOLOPContext) BoolOperate() IBoolOperateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolOperateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolOperateContext)
}

func (s *BOOLOPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterBOOLOP(s)
	}
}

func (s *BOOLOPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitBOOLOP(s)
	}
}

type IDENBOOLContext struct {
	*BoolStatementContext
}

func NewIDENBOOLContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IDENBOOLContext {
	var p = new(IDENBOOLContext)

	p.BoolStatementContext = NewEmptyBoolStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolStatementContext))

	return p
}

func (s *IDENBOOLContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IDENBOOLContext) IDENTIFY() antlr.TerminalNode {
	return s.GetToken(RuleParserIDENTIFY, 0)
}

func (s *IDENBOOLContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterIDENBOOL(s)
	}
}

func (s *IDENBOOLContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitIDENBOOL(s)
	}
}

type COMPAREBOOLContext struct {
	*BoolStatementContext
}

func NewCOMPAREBOOLContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *COMPAREBOOLContext {
	var p = new(COMPAREBOOLContext)

	p.BoolStatementContext = NewEmptyBoolStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolStatementContext))

	return p
}

func (s *COMPAREBOOLContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *COMPAREBOOLContext) CompareStatement() ICompareStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompareStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompareStatementContext)
}

func (s *COMPAREBOOLContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCOMPAREBOOL(s)
	}
}

func (s *COMPAREBOOLContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCOMPAREBOOL(s)
	}
}

type BOOLOPXContext struct {
	*BoolStatementContext
}

func NewBOOLOPXContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BOOLOPXContext {
	var p = new(BOOLOPXContext)

	p.BoolStatementContext = NewEmptyBoolStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolStatementContext))

	return p
}

func (s *BOOLOPXContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BOOLOPXContext) BoolStatement() IBoolStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolStatementContext)
}

func (s *BOOLOPXContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterBOOLOPX(s)
	}
}

func (s *BOOLOPXContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitBOOLOPX(s)
	}
}

func (p *RuleParser) BoolStatement() (localctx IBoolStatementContext) {
	return p.boolStatement(0)
}

func (p *RuleParser) boolStatement(_p int) (localctx IBoolStatementContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBoolStatementContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBoolStatementContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, RuleParserRULE_boolStatement, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBOOLContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(110)
			_la = p.GetTokenStream().LA(1)

			if !(_la == RuleParserTRUE || _la == RuleParserFALSE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		localctx = NewIDENBOOLContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(111)
			p.Match(RuleParserIDENTIFY)
		}

	case 3:
		localctx = NewCOMPAREBOOLContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(112)
			p.compareStatement(0)
		}

	case 4:
		localctx = NewBOOLOPXContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(113)
			p.Match(RuleParserT__0)
		}
		{
			p.SetState(114)
			p.boolStatement(0)
		}
		{
			p.SetState(115)
			p.Match(RuleParserT__1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBOOLOPContext(p, NewBoolStatementContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_boolStatement)
			p.SetState(119)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
			}
			{
				p.SetState(120)

				var _x = p.BoolOperate()

				localctx.(*BOOLOPContext).op = _x
			}
			{
				p.SetState(121)
				p.boolStatement(6)
			}

		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IValueTypeContext is an interface to support dynamic dispatch.
type IValueTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueTypeContext differentiates from other interfaces.
	IsValueTypeContext()
}

type ValueTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueTypeContext() *ValueTypeContext {
	var p = new(ValueTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_valueType
	return p
}

func (*ValueTypeContext) IsValueTypeContext() {}

func NewValueTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueTypeContext {
	var p = new(ValueTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_valueType

	return p
}

func (s *ValueTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueTypeContext) BoolValue() IBoolValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolValueContext)
}

func (s *ValueTypeContext) Identify() IIdentifyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifyContext)
}

func (s *ValueTypeContext) Num() INumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumContext)
}

func (s *ValueTypeContext) CalculateStatement() ICalculateStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalculateStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
}

func (s *ValueTypeContext) MapValue() IMapValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapValueContext)
}

func (s *ValueTypeContext) GetMapOrArrayValue() IGetMapOrArrayValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGetMapOrArrayValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGetMapOrArrayValueContext)
}

func (s *ValueTypeContext) StringValue() IStringValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *ValueTypeContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ValueTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterValueType(s)
	}
}

func (s *ValueTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitValueType(s)
	}
}

func (p *RuleParser) ValueType() (localctx IValueTypeContext) {
	this := p
	_ = this

	localctx = NewValueTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RuleParserRULE_valueType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.BoolValue()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Identify()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.Num()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(131)
			p.calculateStatement(0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(132)
			p.MapValue()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(133)
			p.GetMapOrArrayValue()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(134)
			p.StringValue()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(135)
			p.Array()
		}

	}

	return localctx
}

// ISetValueStatementContext is an interface to support dynamic dispatch.
type ISetValueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetValueStatementContext differentiates from other interfaces.
	IsSetValueStatementContext()
}

type SetValueStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetValueStatementContext() *SetValueStatementContext {
	var p = new(SetValueStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_setValueStatement
	return p
}

func (*SetValueStatementContext) IsSetValueStatementContext() {}

func NewSetValueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetValueStatementContext {
	var p = new(SetValueStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_setValueStatement

	return p
}

func (s *SetValueStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SetValueStatementContext) IDENTIFY() antlr.TerminalNode {
	return s.GetToken(RuleParserIDENTIFY, 0)
}

func (s *SetValueStatementContext) ValueType() IValueTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *SetValueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetValueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetValueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterSetValueStatement(s)
	}
}

func (s *SetValueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitSetValueStatement(s)
	}
}

func (p *RuleParser) SetValueStatement() (localctx ISetValueStatementContext) {
	this := p
	_ = this

	localctx = NewSetValueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RuleParserRULE_setValueStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(RuleParserIDENTIFY)
	}
	{
		p.SetState(139)
		p.Match(RuleParserT__2)
	}
	{
		p.SetState(140)
		p.ValueType()
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RuleParserT__3 {
		{
			p.SetState(141)
			p.Match(RuleParserT__3)
		}

	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(RuleParserIF, 0)
}

func (s *IfStatementContext) BoolStatement() IBoolStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolStatementContext)
}

func (s *IfStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStatementContext) AllElseIfStatement() []IElseIfStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseIfStatementContext); ok {
			len++
		}
	}

	tst := make([]IElseIfStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseIfStatementContext); ok {
			tst[i] = t.(IElseIfStatementContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) ElseIfStatement(i int) IElseIfStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseIfStatementContext)
}

func (s *IfStatementContext) ElseStatement() IElseStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *RuleParser) IfStatement() (localctx IIfStatementContext) {
	this := p
	_ = this

	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RuleParserRULE_ifStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(RuleParserIF)
	}
	{
		p.SetState(145)
		p.boolStatement(0)
	}
	{
		p.SetState(146)
		p.Match(RuleParserT__4)
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserT__0)|(1<<RuleParserT__6)|(1<<RuleParserNUM)|(1<<RuleParserIF)|(1<<RuleParserTRUE)|(1<<RuleParserFALSE)|(1<<RuleParserIDENTIFY))) != 0 {
		{
			p.SetState(147)
			p.Statement()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(153)
		p.Match(RuleParserT__5)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RuleParserELSIF {
		{
			p.SetState(154)
			p.ElseIfStatement()
		}

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RuleParserELSE {
		{
			p.SetState(160)
			p.ElseStatement()
		}

	}

	return localctx
}

// IElseIfStatementContext is an interface to support dynamic dispatch.
type IElseIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseIfStatementContext differentiates from other interfaces.
	IsElseIfStatementContext()
}

type ElseIfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfStatementContext() *ElseIfStatementContext {
	var p = new(ElseIfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_elseIfStatement
	return p
}

func (*ElseIfStatementContext) IsElseIfStatementContext() {}

func NewElseIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfStatementContext {
	var p = new(ElseIfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_elseIfStatement

	return p
}

func (s *ElseIfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfStatementContext) ELSIF() antlr.TerminalNode {
	return s.GetToken(RuleParserELSIF, 0)
}

func (s *ElseIfStatementContext) BoolStatement() IBoolStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolStatementContext)
}

func (s *ElseIfStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ElseIfStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ElseIfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterElseIfStatement(s)
	}
}

func (s *ElseIfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitElseIfStatement(s)
	}
}

func (p *RuleParser) ElseIfStatement() (localctx IElseIfStatementContext) {
	this := p
	_ = this

	localctx = NewElseIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RuleParserRULE_elseIfStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(RuleParserELSIF)
	}
	{
		p.SetState(164)
		p.boolStatement(0)
	}
	{
		p.SetState(165)
		p.Match(RuleParserT__4)
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserT__0)|(1<<RuleParserT__6)|(1<<RuleParserNUM)|(1<<RuleParserIF)|(1<<RuleParserTRUE)|(1<<RuleParserFALSE)|(1<<RuleParserIDENTIFY))) != 0 {
		{
			p.SetState(166)
			p.Statement()
		}

		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(172)
		p.Match(RuleParserT__5)
	}

	return localctx
}

// IElseStatementContext is an interface to support dynamic dispatch.
type IElseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseStatementContext differentiates from other interfaces.
	IsElseStatementContext()
}

type ElseStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStatementContext() *ElseStatementContext {
	var p = new(ElseStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_elseStatement
	return p
}

func (*ElseStatementContext) IsElseStatementContext() {}

func NewElseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStatementContext {
	var p = new(ElseStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_elseStatement

	return p
}

func (s *ElseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(RuleParserELSE, 0)
}

func (s *ElseStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ElseStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ElseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterElseStatement(s)
	}
}

func (s *ElseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitElseStatement(s)
	}
}

func (p *RuleParser) ElseStatement() (localctx IElseStatementContext) {
	this := p
	_ = this

	localctx = NewElseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RuleParserRULE_elseStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(RuleParserELSE)
	}
	{
		p.SetState(175)
		p.Match(RuleParserT__4)
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserT__0)|(1<<RuleParserT__6)|(1<<RuleParserNUM)|(1<<RuleParserIF)|(1<<RuleParserTRUE)|(1<<RuleParserFALSE)|(1<<RuleParserIDENTIFY))) != 0 {
		{
			p.SetState(176)
			p.Statement()
		}

		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(182)
		p.Match(RuleParserT__5)
	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value rule contexts.
	GetValue() IValueTypeContext

	// SetValue sets the value rule contexts.
	SetValue(IValueTypeContext)

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  IValueTypeContext
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) GetValue() IValueTypeContext { return s.value }

func (s *ReturnStatementContext) SetValue(v IValueTypeContext) { s.value = v }

func (s *ReturnStatementContext) ValueType() IValueTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *ReturnStatementContext) CalculateStatement() ICalculateStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalculateStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
}

func (s *ReturnStatementContext) BoolStatement() IBoolStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolStatementContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (p *RuleParser) ReturnStatement() (localctx IReturnStatementContext) {
	this := p
	_ = this

	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RuleParserRULE_returnStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(184)
			p.Match(RuleParserT__6)
		}
		{
			p.SetState(185)

			var _x = p.ValueType()

			localctx.(*ReturnStatementContext).value = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.calculateStatement(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(187)
			p.boolStatement(0)
		}

	}

	return localctx
}

// IMapKeyContext is an interface to support dynamic dispatch.
type IMapKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapKeyContext differentiates from other interfaces.
	IsMapKeyContext()
}

type MapKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapKeyContext() *MapKeyContext {
	var p = new(MapKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_mapKey
	return p
}

func (*MapKeyContext) IsMapKeyContext() {}

func NewMapKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapKeyContext {
	var p = new(MapKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_mapKey

	return p
}

func (s *MapKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *MapKeyContext) IDENTIFY() antlr.TerminalNode {
	return s.GetToken(RuleParserIDENTIFY, 0)
}

func (s *MapKeyContext) Str() antlr.TerminalNode {
	return s.GetToken(RuleParserStr, 0)
}

func (s *MapKeyContext) NUM() antlr.TerminalNode {
	return s.GetToken(RuleParserNUM, 0)
}

func (s *MapKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterMapKey(s)
	}
}

func (s *MapKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitMapKey(s)
	}
}

func (p *RuleParser) MapKey() (localctx IMapKeyContext) {
	this := p
	_ = this

	localctx = NewMapKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, RuleParserRULE_mapKey)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-12)&-(0x1f+1)) == 0 && ((1<<uint((_la-12)))&((1<<(RuleParserNUM-12))|(1<<(RuleParserIDENTIFY-12))|(1<<(RuleParserStr-12)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) MapKey() IMapKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapKeyContext)
}

func (s *PairContext) ValueType() IValueTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *RuleParser) Pair() (localctx IPairContext) {
	this := p
	_ = this

	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, RuleParserRULE_pair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.MapKey()
	}
	{
		p.SetState(193)
		p.Match(RuleParserT__7)
	}
	{
		p.SetState(194)
		p.ValueType()
	}

	return localctx
}

// IMapValueContext is an interface to support dynamic dispatch.
type IMapValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapValueContext differentiates from other interfaces.
	IsMapValueContext()
}

type MapValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapValueContext() *MapValueContext {
	var p = new(MapValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_mapValue
	return p
}

func (*MapValueContext) IsMapValueContext() {}

func NewMapValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapValueContext {
	var p = new(MapValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_mapValue

	return p
}

func (s *MapValueContext) GetParser() antlr.Parser { return s.parser }

func (s *MapValueContext) AllPair() []IPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPairContext); ok {
			len++
		}
	}

	tst := make([]IPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPairContext); ok {
			tst[i] = t.(IPairContext)
			i++
		}
	}

	return tst
}

func (s *MapValueContext) Pair(i int) IPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *MapValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterMapValue(s)
	}
}

func (s *MapValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitMapValue(s)
	}
}

func (p *RuleParser) MapValue() (localctx IMapValueContext) {
	this := p
	_ = this

	localctx = NewMapValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, RuleParserRULE_mapValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(RuleParserT__4)
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-12)&-(0x1f+1)) == 0 && ((1<<uint((_la-12)))&((1<<(RuleParserNUM-12))|(1<<(RuleParserIDENTIFY-12))|(1<<(RuleParserStr-12)))) != 0 {
		{
			p.SetState(197)
			p.Pair()
		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RuleParserT__8 {
			{
				p.SetState(198)
				p.Match(RuleParserT__8)
			}
			{
				p.SetState(199)
				p.Pair()
			}

			p.SetState(204)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(207)
		p.Match(RuleParserT__5)
	}

	return localctx
}

// IGetMapOrArrayValueContext is an interface to support dynamic dispatch.
type IGetMapOrArrayValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGetMapOrArrayValueContext differentiates from other interfaces.
	IsGetMapOrArrayValueContext()
}

type GetMapOrArrayValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetMapOrArrayValueContext() *GetMapOrArrayValueContext {
	var p = new(GetMapOrArrayValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_getMapOrArrayValue
	return p
}

func (*GetMapOrArrayValueContext) IsGetMapOrArrayValueContext() {}

func NewGetMapOrArrayValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetMapOrArrayValueContext {
	var p = new(GetMapOrArrayValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_getMapOrArrayValue

	return p
}

func (s *GetMapOrArrayValueContext) GetParser() antlr.Parser { return s.parser }

func (s *GetMapOrArrayValueContext) Identify() IIdentifyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifyContext)
}

func (s *GetMapOrArrayValueContext) AllMapKey() []IMapKeyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMapKeyContext); ok {
			len++
		}
	}

	tst := make([]IMapKeyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMapKeyContext); ok {
			tst[i] = t.(IMapKeyContext)
			i++
		}
	}

	return tst
}

func (s *GetMapOrArrayValueContext) MapKey(i int) IMapKeyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapKeyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapKeyContext)
}

func (s *GetMapOrArrayValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetMapOrArrayValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetMapOrArrayValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterGetMapOrArrayValue(s)
	}
}

func (s *GetMapOrArrayValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitGetMapOrArrayValue(s)
	}
}

func (p *RuleParser) GetMapOrArrayValue() (localctx IGetMapOrArrayValueContext) {
	this := p
	_ = this

	localctx = NewGetMapOrArrayValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, RuleParserRULE_getMapOrArrayValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Identify()
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == RuleParserT__9 {
		{
			p.SetState(210)
			p.Match(RuleParserT__9)
		}
		{
			p.SetState(211)
			p.MapKey()
		}
		{
			p.SetState(212)
			p.Match(RuleParserT__10)
		}

		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunCallContext is an interface to support dynamic dispatch.
type IFunCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunCallContext differentiates from other interfaces.
	IsFunCallContext()
}

type FunCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunCallContext() *FunCallContext {
	var p = new(FunCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_funCall
	return p
}

func (*FunCallContext) IsFunCallContext() {}

func NewFunCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunCallContext {
	var p = new(FunCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_funCall

	return p
}

func (s *FunCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunCallContext) Identify() IIdentifyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifyContext)
}

func (s *FunCallContext) AllValueType() []IValueTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueTypeContext); ok {
			len++
		}
	}

	tst := make([]IValueTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueTypeContext); ok {
			tst[i] = t.(IValueTypeContext)
			i++
		}
	}

	return tst
}

func (s *FunCallContext) ValueType(i int) IValueTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *FunCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterFunCall(s)
	}
}

func (s *FunCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitFunCall(s)
	}
}

func (p *RuleParser) FunCall() (localctx IFunCallContext) {
	this := p
	_ = this

	localctx = NewFunCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, RuleParserRULE_funCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Identify()
	}
	{
		p.SetState(219)
		p.Match(RuleParserT__0)
	}
	{
		p.SetState(220)
		p.ValueType()
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RuleParserT__8 {
		{
			p.SetState(221)
			p.Match(RuleParserT__8)
		}
		{
			p.SetState(222)
			p.ValueType()
		}

		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(228)
		p.Match(RuleParserT__1)
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) AllValueType() []IValueTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueTypeContext); ok {
			len++
		}
	}

	tst := make([]IValueTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueTypeContext); ok {
			tst[i] = t.(IValueTypeContext)
			i++
		}
	}

	return tst
}

func (s *ArrayContext) ValueType(i int) IValueTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitArray(s)
	}
}

func (p *RuleParser) Array() (localctx IArrayContext) {
	this := p
	_ = this

	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, RuleParserRULE_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(RuleParserT__9)
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserT__0)|(1<<RuleParserT__4)|(1<<RuleParserT__9)|(1<<RuleParserNUM)|(1<<RuleParserTRUE)|(1<<RuleParserFALSE)|(1<<RuleParserIDENTIFY))) != 0) || _la == RuleParserStr {
		{
			p.SetState(231)
			p.ValueType()
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RuleParserT__8 {
			{
				p.SetState(232)
				p.Match(RuleParserT__8)
			}
			{
				p.SetState(233)
				p.ValueType()
			}

			p.SetState(238)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(241)
		p.Match(RuleParserT__10)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) FunCall() IFunCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunCallContext)
}

func (s *StatementContext) CalculateStatement() ICalculateStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalculateStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
}

func (s *StatementContext) BoolStatement() IBoolStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolStatementContext)
}

func (s *StatementContext) CompareStatement() ICompareStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompareStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompareStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) SetValueStatement() ISetValueStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetValueStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetValueStatementContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *RuleParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, RuleParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(243)
			p.FunCall()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(244)
			p.calculateStatement(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(245)
			p.boolStatement(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(246)
			p.compareStatement(0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(247)
			p.IfStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(248)
			p.SetValueStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(249)
			p.ReturnStatement()
		}

	}

	return localctx
}

// IInitContext is an interface to support dynamic dispatch.
type IInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitContext differentiates from other interfaces.
	IsInitContext()
}

type InitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitContext() *InitContext {
	var p = new(InitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_init
	return p
}

func (*InitContext) IsInitContext() {}

func NewInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitContext {
	var p = new(InitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_init

	return p
}

func (s *InitContext) GetParser() antlr.Parser { return s.parser }

func (s *InitContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *InitContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *InitContext) EOF() antlr.TerminalNode {
	return s.GetToken(RuleParserEOF, 0)
}

func (s *InitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterInit(s)
	}
}

func (s *InitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitInit(s)
	}
}

func (p *RuleParser) Init() (localctx IInitContext) {
	this := p
	_ = this

	localctx = NewInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, RuleParserRULE_init)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserT__0)|(1<<RuleParserT__6)|(1<<RuleParserNUM)|(1<<RuleParserIF)|(1<<RuleParserTRUE)|(1<<RuleParserFALSE)|(1<<RuleParserIDENTIFY))) != 0 {
		{
			p.SetState(252)
			p.Statement()
		}

		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(258)
			p.Match(RuleParserEOF)
		}

	}

	return localctx
}

func (p *RuleParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *CompareStatementContext = nil
		if localctx != nil {
			t = localctx.(*CompareStatementContext)
		}
		return p.CompareStatement_Sempred(t, predIndex)

	case 10:
		var t *CalculateStatementContext = nil
		if localctx != nil {
			t = localctx.(*CalculateStatementContext)
		}
		return p.CalculateStatement_Sempred(t, predIndex)

	case 11:
		var t *BoolStatementContext = nil
		if localctx != nil {
			t = localctx.(*BoolStatementContext)
		}
		return p.BoolStatement_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RuleParser) CompareStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RuleParser) CalculateStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RuleParser) BoolStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

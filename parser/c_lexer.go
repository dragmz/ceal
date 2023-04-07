// Code generated from ../C.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var clexerLexerStaticData struct {
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

func clexerLexerInit() {
	staticData := &clexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'#include'", "'('", "')'", "'{'", "'}'", "'struct'", "';'", "'const'",
		"','", "'='", "'.'", "'if'", "'return'", "'-'", "'else if'", "'else'",
		"'*'", "'/'", "'+'", "'=='", "'!='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "INT", "STRING", "ID", "WS", "SINGLE_COMMENT", "MULTILINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "INT", "STRING", "ID", "WS", "SINGLE_COMMENT",
		"MULTILINE_COMMENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 27, 187, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1,
		21, 4, 21, 132, 8, 21, 11, 21, 12, 21, 133, 1, 22, 1, 22, 1, 22, 1, 22,
		5, 22, 140, 8, 22, 10, 22, 12, 22, 143, 9, 22, 1, 22, 1, 22, 1, 23, 1,
		23, 5, 23, 149, 8, 23, 10, 23, 12, 23, 152, 9, 23, 1, 24, 4, 24, 155, 8,
		24, 11, 24, 12, 24, 156, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25,
		165, 8, 25, 10, 25, 12, 25, 168, 9, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 5, 26, 178, 8, 26, 10, 26, 12, 26, 181, 9, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 2, 166, 179, 0, 27, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22,
		45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 1, 0, 5, 1, 0, 48, 57, 2, 0, 34,
		34, 92, 92, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 193, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0,
		0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 1, 55, 1, 0, 0, 0, 3, 64, 1, 0,
		0, 0, 5, 66, 1, 0, 0, 0, 7, 68, 1, 0, 0, 0, 9, 70, 1, 0, 0, 0, 11, 72,
		1, 0, 0, 0, 13, 79, 1, 0, 0, 0, 15, 81, 1, 0, 0, 0, 17, 87, 1, 0, 0, 0,
		19, 89, 1, 0, 0, 0, 21, 91, 1, 0, 0, 0, 23, 93, 1, 0, 0, 0, 25, 96, 1,
		0, 0, 0, 27, 103, 1, 0, 0, 0, 29, 105, 1, 0, 0, 0, 31, 113, 1, 0, 0, 0,
		33, 118, 1, 0, 0, 0, 35, 120, 1, 0, 0, 0, 37, 122, 1, 0, 0, 0, 39, 124,
		1, 0, 0, 0, 41, 127, 1, 0, 0, 0, 43, 131, 1, 0, 0, 0, 45, 135, 1, 0, 0,
		0, 47, 146, 1, 0, 0, 0, 49, 154, 1, 0, 0, 0, 51, 160, 1, 0, 0, 0, 53, 173,
		1, 0, 0, 0, 55, 56, 5, 35, 0, 0, 56, 57, 5, 105, 0, 0, 57, 58, 5, 110,
		0, 0, 58, 59, 5, 99, 0, 0, 59, 60, 5, 108, 0, 0, 60, 61, 5, 117, 0, 0,
		61, 62, 5, 100, 0, 0, 62, 63, 5, 101, 0, 0, 63, 2, 1, 0, 0, 0, 64, 65,
		5, 40, 0, 0, 65, 4, 1, 0, 0, 0, 66, 67, 5, 41, 0, 0, 67, 6, 1, 0, 0, 0,
		68, 69, 5, 123, 0, 0, 69, 8, 1, 0, 0, 0, 70, 71, 5, 125, 0, 0, 71, 10,
		1, 0, 0, 0, 72, 73, 5, 115, 0, 0, 73, 74, 5, 116, 0, 0, 74, 75, 5, 114,
		0, 0, 75, 76, 5, 117, 0, 0, 76, 77, 5, 99, 0, 0, 77, 78, 5, 116, 0, 0,
		78, 12, 1, 0, 0, 0, 79, 80, 5, 59, 0, 0, 80, 14, 1, 0, 0, 0, 81, 82, 5,
		99, 0, 0, 82, 83, 5, 111, 0, 0, 83, 84, 5, 110, 0, 0, 84, 85, 5, 115, 0,
		0, 85, 86, 5, 116, 0, 0, 86, 16, 1, 0, 0, 0, 87, 88, 5, 44, 0, 0, 88, 18,
		1, 0, 0, 0, 89, 90, 5, 61, 0, 0, 90, 20, 1, 0, 0, 0, 91, 92, 5, 46, 0,
		0, 92, 22, 1, 0, 0, 0, 93, 94, 5, 105, 0, 0, 94, 95, 5, 102, 0, 0, 95,
		24, 1, 0, 0, 0, 96, 97, 5, 114, 0, 0, 97, 98, 5, 101, 0, 0, 98, 99, 5,
		116, 0, 0, 99, 100, 5, 117, 0, 0, 100, 101, 5, 114, 0, 0, 101, 102, 5,
		110, 0, 0, 102, 26, 1, 0, 0, 0, 103, 104, 5, 45, 0, 0, 104, 28, 1, 0, 0,
		0, 105, 106, 5, 101, 0, 0, 106, 107, 5, 108, 0, 0, 107, 108, 5, 115, 0,
		0, 108, 109, 5, 101, 0, 0, 109, 110, 5, 32, 0, 0, 110, 111, 5, 105, 0,
		0, 111, 112, 5, 102, 0, 0, 112, 30, 1, 0, 0, 0, 113, 114, 5, 101, 0, 0,
		114, 115, 5, 108, 0, 0, 115, 116, 5, 115, 0, 0, 116, 117, 5, 101, 0, 0,
		117, 32, 1, 0, 0, 0, 118, 119, 5, 42, 0, 0, 119, 34, 1, 0, 0, 0, 120, 121,
		5, 47, 0, 0, 121, 36, 1, 0, 0, 0, 122, 123, 5, 43, 0, 0, 123, 38, 1, 0,
		0, 0, 124, 125, 5, 61, 0, 0, 125, 126, 5, 61, 0, 0, 126, 40, 1, 0, 0, 0,
		127, 128, 5, 33, 0, 0, 128, 129, 5, 61, 0, 0, 129, 42, 1, 0, 0, 0, 130,
		132, 7, 0, 0, 0, 131, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 131,
		1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 44, 1, 0, 0, 0, 135, 141, 5, 34,
		0, 0, 136, 140, 8, 1, 0, 0, 137, 138, 5, 92, 0, 0, 138, 140, 9, 0, 0, 0,
		139, 136, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141,
		139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 144, 1, 0, 0, 0, 143, 141,
		1, 0, 0, 0, 144, 145, 5, 34, 0, 0, 145, 46, 1, 0, 0, 0, 146, 150, 7, 2,
		0, 0, 147, 149, 7, 3, 0, 0, 148, 147, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0,
		150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 48, 1, 0, 0, 0, 152, 150,
		1, 0, 0, 0, 153, 155, 7, 4, 0, 0, 154, 153, 1, 0, 0, 0, 155, 156, 1, 0,
		0, 0, 156, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0,
		158, 159, 6, 24, 0, 0, 159, 50, 1, 0, 0, 0, 160, 161, 5, 47, 0, 0, 161,
		162, 5, 47, 0, 0, 162, 166, 1, 0, 0, 0, 163, 165, 9, 0, 0, 0, 164, 163,
		1, 0, 0, 0, 165, 168, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 166, 164, 1, 0,
		0, 0, 167, 169, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 169, 170, 5, 10, 0, 0,
		170, 171, 1, 0, 0, 0, 171, 172, 6, 25, 0, 0, 172, 52, 1, 0, 0, 0, 173,
		174, 5, 47, 0, 0, 174, 175, 5, 42, 0, 0, 175, 179, 1, 0, 0, 0, 176, 178,
		9, 0, 0, 0, 177, 176, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 180, 1, 0,
		0, 0, 179, 177, 1, 0, 0, 0, 180, 182, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0,
		182, 183, 5, 42, 0, 0, 183, 184, 5, 47, 0, 0, 184, 185, 1, 0, 0, 0, 185,
		186, 6, 26, 0, 0, 186, 54, 1, 0, 0, 0, 8, 0, 133, 139, 141, 150, 156, 166,
		179, 1, 6, 0, 0,
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

// CLexerInit initializes any static state used to implement CLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CLexerInit() {
	staticData := &clexerLexerStaticData
	staticData.once.Do(clexerLexerInit)
}

// NewCLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCLexer(input antlr.CharStream) *CLexer {
	CLexerInit()
	l := new(CLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &clexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "C.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CLexer tokens.
const (
	CLexerT__0              = 1
	CLexerT__1              = 2
	CLexerT__2              = 3
	CLexerT__3              = 4
	CLexerT__4              = 5
	CLexerT__5              = 6
	CLexerT__6              = 7
	CLexerT__7              = 8
	CLexerT__8              = 9
	CLexerT__9              = 10
	CLexerT__10             = 11
	CLexerT__11             = 12
	CLexerT__12             = 13
	CLexerT__13             = 14
	CLexerT__14             = 15
	CLexerT__15             = 16
	CLexerT__16             = 17
	CLexerT__17             = 18
	CLexerT__18             = 19
	CLexerT__19             = 20
	CLexerT__20             = 21
	CLexerINT               = 22
	CLexerSTRING            = 23
	CLexerID                = 24
	CLexerWS                = 25
	CLexerSINGLE_COMMENT    = 26
	CLexerMULTILINE_COMMENT = 27
)

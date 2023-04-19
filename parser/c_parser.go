// Code generated from C.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // C
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CParser struct {
	*antlr.BaseParser
}

var cParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cParserInit() {
	staticData := &cParserStaticData
	staticData.literalNames = []string{
		"", "';'", "'#include'", "'('", "')'", "'{'", "'}'", "'struct'", "','",
		"'asm'", "'if'", "'return'", "'goto'", "':'", "'for'", "'while'", "'do'",
		"'switch'", "'break'", "'continue'", "'-'", "'!'", "'&'", "'^'", "'|'",
		"'&&'", "'||'", "'?'", "'.'", "'['", "']'", "'='", "'const'", "'+='",
		"'-='", "'case'", "'default'", "'else if'", "'else'", "'*'", "'/'",
		"'%'", "'+'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'++'",
		"'--'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"INT", "STRING", "ID", "WS", "SINGLE_COMMENT", "MULTILINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"program", "include", "function", "struct", "field", "type", "params",
		"param", "stmt", "expr", "value_expr", "subscript_expr", "comment",
		"constant", "assign_expr", "const", "asdexpr", "asd", "case", "default",
		"forInit", "forCondition", "forIter", "global", "declaration", "definition",
		"elseif", "else", "args", "call_expr", "muldiv", "addsub", "eqneq",
		"incdec",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 56, 459, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 5, 0, 78, 8, 0, 10, 0, 12, 0, 81, 9, 0, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 95, 8, 2, 10,
		2, 12, 2, 98, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 106, 8, 3,
		10, 3, 12, 3, 109, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 3, 5,
		118, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 125, 8, 6, 10, 6, 12, 6,
		128, 9, 6, 3, 6, 130, 8, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 150,
		8, 8, 10, 8, 12, 8, 153, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 166, 8, 8, 10, 8, 12, 8, 169, 9, 8, 1, 8,
		1, 8, 3, 8, 173, 8, 8, 1, 8, 5, 8, 176, 8, 8, 10, 8, 12, 8, 179, 9, 8,
		1, 8, 3, 8, 182, 8, 8, 1, 8, 1, 8, 3, 8, 186, 8, 8, 1, 8, 1, 8, 1, 8, 5,
		8, 191, 8, 8, 10, 8, 12, 8, 194, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 232, 8, 8, 10, 8, 12, 8, 235,
		9, 8, 1, 8, 3, 8, 238, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		3, 8, 247, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 266, 8, 9, 10, 9,
		12, 9, 269, 9, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 276, 8, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 311, 8, 9, 10, 9,
		12, 9, 314, 9, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 5, 14, 330, 8, 14, 10, 14, 12,
		14, 333, 9, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16,
		5, 16, 343, 8, 16, 10, 16, 12, 16, 346, 9, 16, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 357, 8, 18, 10, 18, 12, 18,
		360, 9, 18, 1, 19, 1, 19, 1, 19, 5, 19, 365, 8, 19, 10, 19, 12, 19, 368,
		9, 19, 1, 20, 1, 20, 5, 20, 372, 8, 20, 10, 20, 12, 20, 375, 9, 20, 3,
		20, 377, 8, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 5, 22, 384, 8, 22, 10,
		22, 12, 22, 387, 9, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 5, 26, 408, 8, 26, 10, 26, 12, 26, 411, 9, 26, 1, 26, 1, 26,
		3, 26, 415, 8, 26, 1, 27, 1, 27, 1, 27, 5, 27, 420, 8, 27, 10, 27, 12,
		27, 423, 9, 27, 1, 27, 1, 27, 3, 27, 427, 8, 27, 1, 28, 1, 28, 1, 28, 5,
		28, 432, 8, 28, 10, 28, 12, 28, 435, 9, 28, 3, 28, 437, 8, 28, 1, 29, 1,
		29, 1, 29, 5, 29, 442, 8, 29, 10, 29, 12, 29, 445, 9, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1,
		33, 0, 1, 18, 34, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
		66, 0, 7, 1, 0, 55, 56, 1, 0, 51, 52, 1, 0, 33, 34, 1, 0, 39, 41, 2, 0,
		20, 20, 42, 42, 1, 0, 43, 48, 1, 0, 49, 50, 494, 0, 79, 1, 0, 0, 0, 2,
		84, 1, 0, 0, 0, 4, 87, 1, 0, 0, 0, 6, 101, 1, 0, 0, 0, 8, 112, 1, 0, 0,
		0, 10, 117, 1, 0, 0, 0, 12, 129, 1, 0, 0, 0, 14, 131, 1, 0, 0, 0, 16, 246,
		1, 0, 0, 0, 18, 275, 1, 0, 0, 0, 20, 315, 1, 0, 0, 0, 22, 317, 1, 0, 0,
		0, 24, 322, 1, 0, 0, 0, 26, 324, 1, 0, 0, 0, 28, 326, 1, 0, 0, 0, 30, 337,
		1, 0, 0, 0, 32, 339, 1, 0, 0, 0, 34, 350, 1, 0, 0, 0, 36, 352, 1, 0, 0,
		0, 38, 361, 1, 0, 0, 0, 40, 376, 1, 0, 0, 0, 42, 378, 1, 0, 0, 0, 44, 380,
		1, 0, 0, 0, 46, 388, 1, 0, 0, 0, 48, 393, 1, 0, 0, 0, 50, 396, 1, 0, 0,
		0, 52, 401, 1, 0, 0, 0, 54, 416, 1, 0, 0, 0, 56, 436, 1, 0, 0, 0, 58, 438,
		1, 0, 0, 0, 60, 450, 1, 0, 0, 0, 62, 452, 1, 0, 0, 0, 64, 454, 1, 0, 0,
		0, 66, 456, 1, 0, 0, 0, 68, 78, 3, 2, 1, 0, 69, 78, 3, 24, 12, 0, 70, 71,
		3, 46, 23, 0, 71, 72, 5, 1, 0, 0, 72, 78, 1, 0, 0, 0, 73, 78, 3, 4, 2,
		0, 74, 75, 3, 6, 3, 0, 75, 76, 5, 1, 0, 0, 76, 78, 1, 0, 0, 0, 77, 68,
		1, 0, 0, 0, 77, 69, 1, 0, 0, 0, 77, 70, 1, 0, 0, 0, 77, 73, 1, 0, 0, 0,
		77, 74, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1,
		0, 0, 0, 80, 82, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 82, 83, 5, 0, 0, 1, 83,
		1, 1, 0, 0, 0, 84, 85, 5, 2, 0, 0, 85, 86, 5, 52, 0, 0, 86, 3, 1, 0, 0,
		0, 87, 88, 3, 10, 5, 0, 88, 89, 5, 53, 0, 0, 89, 90, 5, 3, 0, 0, 90, 91,
		3, 12, 6, 0, 91, 92, 5, 4, 0, 0, 92, 96, 5, 5, 0, 0, 93, 95, 3, 16, 8,
		0, 94, 93, 1, 0, 0, 0, 95, 98, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97,
		1, 0, 0, 0, 97, 99, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 99, 100, 5, 6, 0, 0,
		100, 5, 1, 0, 0, 0, 101, 102, 5, 7, 0, 0, 102, 103, 5, 53, 0, 0, 103, 107,
		5, 5, 0, 0, 104, 106, 3, 8, 4, 0, 105, 104, 1, 0, 0, 0, 106, 109, 1, 0,
		0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 110, 1, 0, 0, 0,
		109, 107, 1, 0, 0, 0, 110, 111, 5, 6, 0, 0, 111, 7, 1, 0, 0, 0, 112, 113,
		3, 10, 5, 0, 113, 114, 5, 53, 0, 0, 114, 115, 5, 1, 0, 0, 115, 9, 1, 0,
		0, 0, 116, 118, 3, 30, 15, 0, 117, 116, 1, 0, 0, 0, 117, 118, 1, 0, 0,
		0, 118, 119, 1, 0, 0, 0, 119, 120, 5, 53, 0, 0, 120, 11, 1, 0, 0, 0, 121,
		126, 3, 14, 7, 0, 122, 123, 5, 8, 0, 0, 123, 125, 3, 14, 7, 0, 124, 122,
		1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0,
		0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129, 121, 1, 0, 0, 0,
		129, 130, 1, 0, 0, 0, 130, 13, 1, 0, 0, 0, 131, 132, 3, 10, 5, 0, 132,
		133, 5, 53, 0, 0, 133, 15, 1, 0, 0, 0, 134, 135, 3, 48, 24, 0, 135, 136,
		5, 1, 0, 0, 136, 247, 1, 0, 0, 0, 137, 138, 3, 50, 25, 0, 138, 139, 5,
		1, 0, 0, 139, 247, 1, 0, 0, 0, 140, 141, 3, 28, 14, 0, 141, 142, 5, 1,
		0, 0, 142, 247, 1, 0, 0, 0, 143, 144, 3, 32, 16, 0, 144, 145, 5, 1, 0,
		0, 145, 247, 1, 0, 0, 0, 146, 147, 5, 9, 0, 0, 147, 151, 5, 3, 0, 0, 148,
		150, 5, 52, 0, 0, 149, 148, 1, 0, 0, 0, 150, 153, 1, 0, 0, 0, 151, 149,
		1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 154, 1, 0, 0, 0, 153, 151, 1, 0,
		0, 0, 154, 155, 5, 4, 0, 0, 155, 247, 5, 1, 0, 0, 156, 157, 3, 58, 29,
		0, 157, 158, 5, 1, 0, 0, 158, 247, 1, 0, 0, 0, 159, 160, 5, 10, 0, 0, 160,
		161, 5, 3, 0, 0, 161, 162, 3, 18, 9, 0, 162, 172, 5, 4, 0, 0, 163, 167,
		5, 5, 0, 0, 164, 166, 3, 16, 8, 0, 165, 164, 1, 0, 0, 0, 166, 169, 1, 0,
		0, 0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 170, 1, 0, 0, 0,
		169, 167, 1, 0, 0, 0, 170, 173, 5, 6, 0, 0, 171, 173, 3, 16, 8, 0, 172,
		163, 1, 0, 0, 0, 172, 171, 1, 0, 0, 0, 173, 177, 1, 0, 0, 0, 174, 176,
		3, 52, 26, 0, 175, 174, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 175, 1,
		0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 177, 1, 0, 0,
		0, 180, 182, 3, 54, 27, 0, 181, 180, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0,
		182, 247, 1, 0, 0, 0, 183, 185, 5, 11, 0, 0, 184, 186, 3, 18, 9, 0, 185,
		184, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 247,
		5, 1, 0, 0, 188, 192, 5, 5, 0, 0, 189, 191, 3, 16, 8, 0, 190, 189, 1, 0,
		0, 0, 191, 194, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0,
		193, 195, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 195, 247, 5, 6, 0, 0, 196,
		197, 5, 12, 0, 0, 197, 198, 5, 53, 0, 0, 198, 247, 5, 1, 0, 0, 199, 200,
		5, 53, 0, 0, 200, 247, 5, 13, 0, 0, 201, 202, 5, 14, 0, 0, 202, 203, 5,
		3, 0, 0, 203, 204, 3, 40, 20, 0, 204, 205, 5, 1, 0, 0, 205, 206, 3, 42,
		21, 0, 206, 207, 5, 1, 0, 0, 207, 208, 3, 44, 22, 0, 208, 209, 5, 4, 0,
		0, 209, 210, 3, 16, 8, 0, 210, 247, 1, 0, 0, 0, 211, 212, 5, 15, 0, 0,
		212, 213, 5, 3, 0, 0, 213, 214, 3, 18, 9, 0, 214, 215, 5, 4, 0, 0, 215,
		216, 3, 16, 8, 0, 216, 247, 1, 0, 0, 0, 217, 218, 5, 16, 0, 0, 218, 219,
		3, 16, 8, 0, 219, 220, 5, 15, 0, 0, 220, 221, 5, 3, 0, 0, 221, 222, 3,
		18, 9, 0, 222, 223, 5, 4, 0, 0, 223, 224, 5, 1, 0, 0, 224, 247, 1, 0, 0,
		0, 225, 226, 5, 17, 0, 0, 226, 227, 5, 3, 0, 0, 227, 228, 3, 18, 9, 0,
		228, 229, 5, 4, 0, 0, 229, 233, 5, 5, 0, 0, 230, 232, 3, 36, 18, 0, 231,
		230, 1, 0, 0, 0, 232, 235, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 233, 234,
		1, 0, 0, 0, 234, 237, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 236, 238, 3, 38,
		19, 0, 237, 236, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0,
		239, 240, 5, 6, 0, 0, 240, 247, 1, 0, 0, 0, 241, 242, 5, 18, 0, 0, 242,
		247, 5, 1, 0, 0, 243, 244, 5, 19, 0, 0, 244, 247, 5, 1, 0, 0, 245, 247,
		3, 24, 12, 0, 246, 134, 1, 0, 0, 0, 246, 137, 1, 0, 0, 0, 246, 140, 1,
		0, 0, 0, 246, 143, 1, 0, 0, 0, 246, 146, 1, 0, 0, 0, 246, 156, 1, 0, 0,
		0, 246, 159, 1, 0, 0, 0, 246, 183, 1, 0, 0, 0, 246, 188, 1, 0, 0, 0, 246,
		196, 1, 0, 0, 0, 246, 199, 1, 0, 0, 0, 246, 201, 1, 0, 0, 0, 246, 211,
		1, 0, 0, 0, 246, 217, 1, 0, 0, 0, 246, 225, 1, 0, 0, 0, 246, 241, 1, 0,
		0, 0, 246, 243, 1, 0, 0, 0, 246, 245, 1, 0, 0, 0, 247, 17, 1, 0, 0, 0,
		248, 249, 6, 9, -1, 0, 249, 250, 5, 53, 0, 0, 250, 276, 3, 66, 33, 0, 251,
		252, 3, 66, 33, 0, 252, 253, 5, 53, 0, 0, 253, 276, 1, 0, 0, 0, 254, 255,
		5, 20, 0, 0, 255, 276, 3, 18, 9, 18, 256, 257, 5, 21, 0, 0, 257, 276, 3,
		18, 9, 17, 258, 276, 3, 28, 14, 0, 259, 276, 3, 32, 16, 0, 260, 276, 3,
		26, 13, 0, 261, 276, 3, 22, 11, 0, 262, 267, 3, 20, 10, 0, 263, 264, 5,
		28, 0, 0, 264, 266, 3, 20, 10, 0, 265, 263, 1, 0, 0, 0, 266, 269, 1, 0,
		0, 0, 267, 265, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 276, 1, 0, 0, 0,
		269, 267, 1, 0, 0, 0, 270, 276, 3, 58, 29, 0, 271, 272, 5, 3, 0, 0, 272,
		273, 3, 18, 9, 0, 273, 274, 5, 4, 0, 0, 274, 276, 1, 0, 0, 0, 275, 248,
		1, 0, 0, 0, 275, 251, 1, 0, 0, 0, 275, 254, 1, 0, 0, 0, 275, 256, 1, 0,
		0, 0, 275, 258, 1, 0, 0, 0, 275, 259, 1, 0, 0, 0, 275, 260, 1, 0, 0, 0,
		275, 261, 1, 0, 0, 0, 275, 262, 1, 0, 0, 0, 275, 270, 1, 0, 0, 0, 275,
		271, 1, 0, 0, 0, 276, 312, 1, 0, 0, 0, 277, 278, 10, 16, 0, 0, 278, 279,
		3, 60, 30, 0, 279, 280, 3, 18, 9, 17, 280, 311, 1, 0, 0, 0, 281, 282, 10,
		15, 0, 0, 282, 283, 3, 62, 31, 0, 283, 284, 3, 18, 9, 16, 284, 311, 1,
		0, 0, 0, 285, 286, 10, 14, 0, 0, 286, 287, 3, 64, 32, 0, 287, 288, 3, 18,
		9, 15, 288, 311, 1, 0, 0, 0, 289, 290, 10, 13, 0, 0, 290, 291, 5, 22, 0,
		0, 291, 311, 3, 18, 9, 14, 292, 293, 10, 12, 0, 0, 293, 294, 5, 23, 0,
		0, 294, 311, 3, 18, 9, 13, 295, 296, 10, 11, 0, 0, 296, 297, 5, 24, 0,
		0, 297, 311, 3, 18, 9, 12, 298, 299, 10, 10, 0, 0, 299, 300, 5, 25, 0,
		0, 300, 311, 3, 18, 9, 11, 301, 302, 10, 9, 0, 0, 302, 303, 5, 26, 0, 0,
		303, 311, 3, 18, 9, 10, 304, 305, 10, 6, 0, 0, 305, 306, 5, 27, 0, 0, 306,
		307, 3, 18, 9, 0, 307, 308, 5, 13, 0, 0, 308, 309, 3, 18, 9, 7, 309, 311,
		1, 0, 0, 0, 310, 277, 1, 0, 0, 0, 310, 281, 1, 0, 0, 0, 310, 285, 1, 0,
		0, 0, 310, 289, 1, 0, 0, 0, 310, 292, 1, 0, 0, 0, 310, 295, 1, 0, 0, 0,
		310, 298, 1, 0, 0, 0, 310, 301, 1, 0, 0, 0, 310, 304, 1, 0, 0, 0, 311,
		314, 1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 19, 1,
		0, 0, 0, 314, 312, 1, 0, 0, 0, 315, 316, 5, 53, 0, 0, 316, 21, 1, 0, 0,
		0, 317, 318, 5, 53, 0, 0, 318, 319, 5, 29, 0, 0, 319, 320, 3, 18, 9, 0,
		320, 321, 5, 30, 0, 0, 321, 23, 1, 0, 0, 0, 322, 323, 7, 0, 0, 0, 323,
		25, 1, 0, 0, 0, 324, 325, 7, 1, 0, 0, 325, 27, 1, 0, 0, 0, 326, 331, 5,
		53, 0, 0, 327, 328, 5, 28, 0, 0, 328, 330, 5, 53, 0, 0, 329, 327, 1, 0,
		0, 0, 330, 333, 1, 0, 0, 0, 331, 329, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0,
		332, 334, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 334, 335, 5, 31, 0, 0, 335,
		336, 3, 18, 9, 0, 336, 29, 1, 0, 0, 0, 337, 338, 5, 32, 0, 0, 338, 31,
		1, 0, 0, 0, 339, 344, 5, 53, 0, 0, 340, 341, 5, 28, 0, 0, 341, 343, 5,
		53, 0, 0, 342, 340, 1, 0, 0, 0, 343, 346, 1, 0, 0, 0, 344, 342, 1, 0, 0,
		0, 344, 345, 1, 0, 0, 0, 345, 347, 1, 0, 0, 0, 346, 344, 1, 0, 0, 0, 347,
		348, 3, 34, 17, 0, 348, 349, 3, 18, 9, 0, 349, 33, 1, 0, 0, 0, 350, 351,
		7, 2, 0, 0, 351, 35, 1, 0, 0, 0, 352, 353, 5, 35, 0, 0, 353, 354, 3, 18,
		9, 0, 354, 358, 5, 13, 0, 0, 355, 357, 3, 16, 8, 0, 356, 355, 1, 0, 0,
		0, 357, 360, 1, 0, 0, 0, 358, 356, 1, 0, 0, 0, 358, 359, 1, 0, 0, 0, 359,
		37, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 361, 362, 5, 36, 0, 0, 362, 366,
		5, 13, 0, 0, 363, 365, 3, 16, 8, 0, 364, 363, 1, 0, 0, 0, 365, 368, 1,
		0, 0, 0, 366, 364, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 39, 1, 0, 0,
		0, 368, 366, 1, 0, 0, 0, 369, 377, 3, 50, 25, 0, 370, 372, 3, 18, 9, 0,
		371, 370, 1, 0, 0, 0, 372, 375, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 373,
		374, 1, 0, 0, 0, 374, 377, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 376, 369,
		1, 0, 0, 0, 376, 373, 1, 0, 0, 0, 377, 41, 1, 0, 0, 0, 378, 379, 3, 18,
		9, 0, 379, 43, 1, 0, 0, 0, 380, 385, 3, 18, 9, 0, 381, 382, 5, 8, 0, 0,
		382, 384, 3, 18, 9, 0, 383, 381, 1, 0, 0, 0, 384, 387, 1, 0, 0, 0, 385,
		383, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 45, 1, 0, 0, 0, 387, 385, 1,
		0, 0, 0, 388, 389, 3, 10, 5, 0, 389, 390, 5, 53, 0, 0, 390, 391, 5, 31,
		0, 0, 391, 392, 3, 26, 13, 0, 392, 47, 1, 0, 0, 0, 393, 394, 3, 10, 5,
		0, 394, 395, 5, 53, 0, 0, 395, 49, 1, 0, 0, 0, 396, 397, 3, 10, 5, 0, 397,
		398, 5, 53, 0, 0, 398, 399, 5, 31, 0, 0, 399, 400, 3, 18, 9, 0, 400, 51,
		1, 0, 0, 0, 401, 402, 5, 37, 0, 0, 402, 403, 5, 3, 0, 0, 403, 404, 3, 18,
		9, 0, 404, 414, 5, 4, 0, 0, 405, 409, 5, 5, 0, 0, 406, 408, 3, 16, 8, 0,
		407, 406, 1, 0, 0, 0, 408, 411, 1, 0, 0, 0, 409, 407, 1, 0, 0, 0, 409,
		410, 1, 0, 0, 0, 410, 412, 1, 0, 0, 0, 411, 409, 1, 0, 0, 0, 412, 415,
		5, 6, 0, 0, 413, 415, 3, 16, 8, 0, 414, 405, 1, 0, 0, 0, 414, 413, 1, 0,
		0, 0, 415, 53, 1, 0, 0, 0, 416, 426, 5, 38, 0, 0, 417, 421, 5, 5, 0, 0,
		418, 420, 3, 16, 8, 0, 419, 418, 1, 0, 0, 0, 420, 423, 1, 0, 0, 0, 421,
		419, 1, 0, 0, 0, 421, 422, 1, 0, 0, 0, 422, 424, 1, 0, 0, 0, 423, 421,
		1, 0, 0, 0, 424, 427, 5, 6, 0, 0, 425, 427, 3, 16, 8, 0, 426, 417, 1, 0,
		0, 0, 426, 425, 1, 0, 0, 0, 427, 55, 1, 0, 0, 0, 428, 433, 3, 18, 9, 0,
		429, 430, 5, 8, 0, 0, 430, 432, 3, 18, 9, 0, 431, 429, 1, 0, 0, 0, 432,
		435, 1, 0, 0, 0, 433, 431, 1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 434, 437,
		1, 0, 0, 0, 435, 433, 1, 0, 0, 0, 436, 428, 1, 0, 0, 0, 436, 437, 1, 0,
		0, 0, 437, 57, 1, 0, 0, 0, 438, 443, 5, 53, 0, 0, 439, 440, 5, 28, 0, 0,
		440, 442, 5, 53, 0, 0, 441, 439, 1, 0, 0, 0, 442, 445, 1, 0, 0, 0, 443,
		441, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444, 446, 1, 0, 0, 0, 445, 443,
		1, 0, 0, 0, 446, 447, 5, 3, 0, 0, 447, 448, 3, 56, 28, 0, 448, 449, 5,
		4, 0, 0, 449, 59, 1, 0, 0, 0, 450, 451, 7, 3, 0, 0, 451, 61, 1, 0, 0, 0,
		452, 453, 7, 4, 0, 0, 453, 63, 1, 0, 0, 0, 454, 455, 7, 5, 0, 0, 455, 65,
		1, 0, 0, 0, 456, 457, 7, 6, 0, 0, 457, 67, 1, 0, 0, 0, 35, 77, 79, 96,
		107, 117, 126, 129, 151, 167, 172, 177, 181, 185, 192, 233, 237, 246, 267,
		275, 310, 312, 331, 344, 358, 366, 373, 376, 385, 409, 414, 421, 426, 433,
		436, 443,
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

// CParserInit initializes any static state used to implement CParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CParserInit() {
	staticData := &cParserStaticData
	staticData.once.Do(cParserInit)
}

// NewCParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCParser(input antlr.TokenStream) *CParser {
	CParserInit()
	this := new(CParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "C.g4"

	return this
}

// CParser tokens.
const (
	CParserEOF               = antlr.TokenEOF
	CParserT__0              = 1
	CParserT__1              = 2
	CParserT__2              = 3
	CParserT__3              = 4
	CParserT__4              = 5
	CParserT__5              = 6
	CParserT__6              = 7
	CParserT__7              = 8
	CParserT__8              = 9
	CParserT__9              = 10
	CParserT__10             = 11
	CParserT__11             = 12
	CParserT__12             = 13
	CParserT__13             = 14
	CParserT__14             = 15
	CParserT__15             = 16
	CParserT__16             = 17
	CParserT__17             = 18
	CParserT__18             = 19
	CParserT__19             = 20
	CParserT__20             = 21
	CParserT__21             = 22
	CParserT__22             = 23
	CParserT__23             = 24
	CParserT__24             = 25
	CParserT__25             = 26
	CParserT__26             = 27
	CParserT__27             = 28
	CParserT__28             = 29
	CParserT__29             = 30
	CParserT__30             = 31
	CParserT__31             = 32
	CParserT__32             = 33
	CParserT__33             = 34
	CParserT__34             = 35
	CParserT__35             = 36
	CParserT__36             = 37
	CParserT__37             = 38
	CParserT__38             = 39
	CParserT__39             = 40
	CParserT__40             = 41
	CParserT__41             = 42
	CParserT__42             = 43
	CParserT__43             = 44
	CParserT__44             = 45
	CParserT__45             = 46
	CParserT__46             = 47
	CParserT__47             = 48
	CParserT__48             = 49
	CParserT__49             = 50
	CParserINT               = 51
	CParserSTRING            = 52
	CParserID                = 53
	CParserWS                = 54
	CParserSINGLE_COMMENT    = 55
	CParserMULTILINE_COMMENT = 56
)

// CParser rules.
const (
	CParserRULE_program        = 0
	CParserRULE_include        = 1
	CParserRULE_function       = 2
	CParserRULE_struct         = 3
	CParserRULE_field          = 4
	CParserRULE_type           = 5
	CParserRULE_params         = 6
	CParserRULE_param          = 7
	CParserRULE_stmt           = 8
	CParserRULE_expr           = 9
	CParserRULE_value_expr     = 10
	CParserRULE_subscript_expr = 11
	CParserRULE_comment        = 12
	CParserRULE_constant       = 13
	CParserRULE_assign_expr    = 14
	CParserRULE_const          = 15
	CParserRULE_asdexpr        = 16
	CParserRULE_asd            = 17
	CParserRULE_case           = 18
	CParserRULE_default        = 19
	CParserRULE_forInit        = 20
	CParserRULE_forCondition   = 21
	CParserRULE_forIter        = 22
	CParserRULE_global         = 23
	CParserRULE_declaration    = 24
	CParserRULE_definition     = 25
	CParserRULE_elseif         = 26
	CParserRULE_else           = 27
	CParserRULE_args           = 28
	CParserRULE_call_expr      = 29
	CParserRULE_muldiv         = 30
	CParserRULE_addsub         = 31
	CParserRULE_eqneq          = 32
	CParserRULE_incdec         = 33
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllInclude() []IIncludeContext
	Include(i int) IIncludeContext
	AllComment() []ICommentContext
	Comment(i int) ICommentContext
	AllGlobal() []IGlobalContext
	Global(i int) IGlobalContext
	AllFunction() []IFunctionContext
	Function(i int) IFunctionContext
	AllStruct_() []IStructContext
	Struct_(i int) IStructContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(CParserEOF, 0)
}

func (s *ProgramContext) AllInclude() []IIncludeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIncludeContext); ok {
			len++
		}
	}

	tst := make([]IIncludeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIncludeContext); ok {
			tst[i] = t.(IIncludeContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Include(i int) IIncludeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncludeContext); ok {
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

	return t.(IIncludeContext)
}

func (s *ProgramContext) AllComment() []ICommentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommentContext); ok {
			len++
		}
	}

	tst := make([]ICommentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommentContext); ok {
			tst[i] = t.(ICommentContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Comment(i int) ICommentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentContext); ok {
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

	return t.(ICommentContext)
}

func (s *ProgramContext) AllGlobal() []IGlobalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGlobalContext); ok {
			len++
		}
	}

	tst := make([]IGlobalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGlobalContext); ok {
			tst[i] = t.(IGlobalContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Global(i int) IGlobalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobalContext); ok {
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

	return t.(IGlobalContext)
}

func (s *ProgramContext) AllFunction() []IFunctionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionContext); ok {
			len++
		}
	}

	tst := make([]IFunctionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionContext); ok {
			tst[i] = t.(IFunctionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Function(i int) IFunctionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
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

	return t.(IFunctionContext)
}

func (s *ProgramContext) AllStruct_() []IStructContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructContext); ok {
			len++
		}
	}

	tst := make([]IStructContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructContext); ok {
			tst[i] = t.(IStructContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Struct_(i int) IStructContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructContext); ok {
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

	return t.(IStructContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CParserRULE_program)
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
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&117093594606600324) != 0 {
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(68)
				p.Include()
			}

		case 2:
			{
				p.SetState(69)
				p.Comment()
			}

		case 3:
			{
				p.SetState(70)
				p.Global()
			}
			{
				p.SetState(71)
				p.Match(CParserT__0)
			}

		case 4:
			{
				p.SetState(73)
				p.Function()
			}

		case 5:
			{
				p.SetState(74)
				p.Struct_()
			}
			{
				p.SetState(75)
				p.Match(CParserT__0)
			}

		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(82)
		p.Match(CParserEOF)
	}

	return localctx
}

// IIncludeContext is an interface to support dynamic dispatch.
type IIncludeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsIncludeContext differentiates from other interfaces.
	IsIncludeContext()
}

type IncludeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncludeContext() *IncludeContext {
	var p = new(IncludeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_include
	return p
}

func (*IncludeContext) IsIncludeContext() {}

func NewIncludeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludeContext {
	var p = new(IncludeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_include

	return p
}

func (s *IncludeContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludeContext) STRING() antlr.TerminalNode {
	return s.GetToken(CParserSTRING, 0)
}

func (s *IncludeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncludeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterInclude(s)
	}
}

func (s *IncludeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitInclude(s)
	}
}

func (s *IncludeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitInclude(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Include() (localctx IIncludeContext) {
	this := p
	_ = this

	localctx = NewIncludeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CParserRULE_include)

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
		p.SetState(84)
		p.Match(CParserT__1)
	}
	{
		p.SetState(85)
		p.Match(CParserSTRING)
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	ID() antlr.TerminalNode
	Params() IParamsContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *FunctionContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FunctionContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *FunctionContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Function() (localctx IFunctionContext) {
	this := p
	_ = this

	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CParserRULE_function)
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
		p.SetState(87)
		p.Type_()
	}
	{
		p.SetState(88)
		p.Match(CParserID)
	}
	{
		p.SetState(89)
		p.Match(CParserT__2)
	}
	{
		p.SetState(90)
		p.Params()
	}
	{
		p.SetState(91)
		p.Match(CParserT__3)
	}
	{
		p.SetState(92)
		p.Match(CParserT__4)
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&117093594607640096) != 0 {
		{
			p.SetState(93)
			p.Stmt()
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(99)
		p.Match(CParserT__5)
	}

	return localctx
}

// IStructContext is an interface to support dynamic dispatch.
type IStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllField() []IFieldContext
	Field(i int) IFieldContext

	// IsStructContext differentiates from other interfaces.
	IsStructContext()
}

type StructContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructContext() *StructContext {
	var p = new(StructContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_struct
	return p
}

func (*StructContext) IsStructContext() {}

func NewStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructContext {
	var p = new(StructContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_struct

	return p
}

func (s *StructContext) GetParser() antlr.Parser { return s.parser }

func (s *StructContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *StructContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *StructContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterStruct(s)
	}
}

func (s *StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitStruct(s)
	}
}

func (s *StructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Struct_() (localctx IStructContext) {
	this := p
	_ = this

	localctx = NewStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CParserRULE_struct)
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
		p.SetState(101)
		p.Match(CParserT__6)
	}
	{
		p.SetState(102)
		p.Match(CParserID)
	}
	{
		p.SetState(103)
		p.Match(CParserT__4)
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserT__31 || _la == CParserID {
		{
			p.SetState(104)
			p.Field()
		}

		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(110)
		p.Match(CParserT__5)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	ID() antlr.TerminalNode

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FieldContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CParserRULE_field)

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
		p.SetState(112)
		p.Type_()
	}
	{
		p.SetState(113)
		p.Match(CParserID)
	}
	{
		p.SetState(114)
		p.Match(CParserT__0)
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Const_() IConstContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *TypeContext) Const_() IConstContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Type_() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CParserRULE_type)
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
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserT__31 {
		{
			p.SetState(116)
			p.Const_()
		}

	}
	{
		p.SetState(119)
		p.Match(CParserID)
	}

	return localctx
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_params
	return p
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ParamsContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
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

	return t.(IParamContext)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitParams(s)
	}
}

func (s *ParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Params() (localctx IParamsContext) {
	this := p
	_ = this

	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CParserRULE_params)
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
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserT__31 || _la == CParserID {
		{
			p.SetState(121)
			p.Param()
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CParserT__7 {
			{
				p.SetState(122)
				p.Match(CParserT__7)
			}
			{
				p.SetState(123)
				p.Param()
			}

			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	ID() antlr.TerminalNode

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParamContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitParam(s)
	}
}

func (s *ParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Param() (localctx IParamContext) {
	this := p
	_ = this

	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CParserRULE_param)

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
		p.SetState(131)
		p.Type_()
	}
	{
		p.SetState(132)
		p.Match(CParserID)
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) CopyFrom(ctx *StmtContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignSumDiffStmtContext struct {
	*StmtContext
}

func NewAssignSumDiffStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignSumDiffStmtContext {
	var p = new(AssignSumDiffStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *AssignSumDiffStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignSumDiffStmtContext) Asdexpr() IAsdexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsdexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsdexprContext)
}

func (s *AssignSumDiffStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterAssignSumDiffStmt(s)
	}
}

func (s *AssignSumDiffStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitAssignSumDiffStmt(s)
	}
}

func (s *AssignSumDiffStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitAssignSumDiffStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type SwitchStmtContext struct {
	*StmtContext
}

func NewSwitchStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStmtContext {
	var p = new(SwitchStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *SwitchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchStmtContext) AllCase_() []ICaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseContext); ok {
			len++
		}
	}

	tst := make([]ICaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseContext); ok {
			tst[i] = t.(ICaseContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStmtContext) Case_(i int) ICaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseContext); ok {
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

	return t.(ICaseContext)
}

func (s *SwitchStmtContext) Default_() IDefaultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultContext)
}

func (s *SwitchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitSwitchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclarationStmtContext struct {
	*StmtContext
}

func NewDeclarationStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclarationStmtContext {
	var p = new(DeclarationStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *DeclarationStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationStmtContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *DeclarationStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterDeclarationStmt(s)
	}
}

func (s *DeclarationStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitDeclarationStmt(s)
	}
}

func (s *DeclarationStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitDeclarationStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type LabelStmtContext struct {
	*StmtContext
}

func NewLabelStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LabelStmtContext {
	var p = new(LabelStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *LabelStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *LabelStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterLabelStmt(s)
	}
}

func (s *LabelStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitLabelStmt(s)
	}
}

func (s *LabelStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitLabelStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignStmtContext struct {
	*StmtContext
}

func NewAssignStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignStmtContext {
	var p = new(AssignStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *AssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStmtContext) Assign_expr() IAssign_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_exprContext)
}

func (s *AssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterAssignStmt(s)
	}
}

func (s *AssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitAssignStmt(s)
	}
}

func (s *AssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitAssignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type BlockStmtContext struct {
	*StmtContext
}

func NewBlockStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockStmtContext {
	var p = new(BlockStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *BlockStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *BlockStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterBlockStmt(s)
	}
}

func (s *BlockStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitBlockStmt(s)
	}
}

func (s *BlockStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitBlockStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallStmtContext struct {
	*StmtContext
}

func NewCallStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallStmtContext {
	var p = new(CallStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *CallStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallStmtContext) Call_expr() ICall_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_exprContext)
}

func (s *CallStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterCallStmt(s)
	}
}

func (s *CallStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitCallStmt(s)
	}
}

func (s *CallStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitCallStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type DoWhileStmtContext struct {
	*StmtContext
}

func NewDoWhileStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoWhileStmtContext {
	var p = new(DoWhileStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *DoWhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoWhileStmtContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *DoWhileStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DoWhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterDoWhileStmt(s)
	}
}

func (s *DoWhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitDoWhileStmt(s)
	}
}

func (s *DoWhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitDoWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type CommentStmtContext struct {
	*StmtContext
}

func NewCommentStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CommentStmtContext {
	var p = new(CommentStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *CommentStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentStmtContext) Comment() ICommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *CommentStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterCommentStmt(s)
	}
}

func (s *CommentStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitCommentStmt(s)
	}
}

func (s *CommentStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitCommentStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContinueStmtContext struct {
	*StmtContext
}

func NewContinueStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfStmtContext struct {
	*StmtContext
}

func NewIfStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStmtContext {
	var p = new(IfStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *IfStmtContext) AllElseif() []IElseifContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseifContext); ok {
			len++
		}
	}

	tst := make([]IElseifContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseifContext); ok {
			tst[i] = t.(IElseifContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Elseif(i int) IElseifContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseifContext); ok {
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

	return t.(IElseifContext)
}

func (s *IfStmtContext) Else_() IElseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseContext)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type WhileStmtContext struct {
	*StmtContext
}

func NewWhileStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStmtContext {
	var p = new(WhileStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStmtContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakStmtContext struct {
	*StmtContext
}

func NewBreakStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStmtContext {
	var p = new(BreakStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type GotoStmtContext struct {
	*StmtContext
}

func NewGotoStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GotoStmtContext {
	var p = new(GotoStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *GotoStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GotoStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *GotoStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterGotoStmt(s)
	}
}

func (s *GotoStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitGotoStmt(s)
	}
}

func (s *GotoStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitGotoStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsmStmtContext struct {
	*StmtContext
}

func NewAsmStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsmStmtContext {
	var p = new(AsmStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *AsmStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsmStmtContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(CParserSTRING)
}

func (s *AsmStmtContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(CParserSTRING, i)
}

func (s *AsmStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterAsmStmt(s)
	}
}

func (s *AsmStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitAsmStmt(s)
	}
}

func (s *AsmStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitAsmStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type DefinitionStmtContext struct {
	*StmtContext
}

func NewDefinitionStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefinitionStmtContext {
	var p = new(DefinitionStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *DefinitionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionStmtContext) Definition() IDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DefinitionStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterDefinitionStmt(s)
	}
}

func (s *DefinitionStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitDefinitionStmt(s)
	}
}

func (s *DefinitionStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitDefinitionStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStmtContext struct {
	*StmtContext
}

func NewReturnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForStmtContext struct {
	*StmtContext
}

func NewForStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStmtContext {
	var p = new(ForStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) ForInit() IForInitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForInitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForInitContext)
}

func (s *ForStmtContext) ForCondition() IForConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForConditionContext)
}

func (s *ForStmtContext) ForIter() IForIterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForIterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForIterContext)
}

func (s *ForStmtContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Stmt() (localctx IStmtContext) {
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CParserRULE_stmt)
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

	var _alt int

	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclarationStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Declaration()
		}
		{
			p.SetState(135)
			p.Match(CParserT__0)
		}

	case 2:
		localctx = NewDefinitionStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.Definition()
		}
		{
			p.SetState(138)
			p.Match(CParserT__0)
		}

	case 3:
		localctx = NewAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(140)
			p.Assign_expr()
		}
		{
			p.SetState(141)
			p.Match(CParserT__0)
		}

	case 4:
		localctx = NewAssignSumDiffStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(143)
			p.Asdexpr()
		}
		{
			p.SetState(144)
			p.Match(CParserT__0)
		}

	case 5:
		localctx = NewAsmStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(146)
			p.Match(CParserT__8)
		}
		{
			p.SetState(147)
			p.Match(CParserT__2)
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CParserSTRING {
			{
				p.SetState(148)
				p.Match(CParserSTRING)
			}

			p.SetState(153)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(154)
			p.Match(CParserT__3)
		}
		{
			p.SetState(155)
			p.Match(CParserT__0)
		}

	case 6:
		localctx = NewCallStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(156)
			p.Call_expr()
		}
		{
			p.SetState(157)
			p.Match(CParserT__0)
		}

	case 7:
		localctx = NewIfStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(159)
			p.Match(CParserT__9)
		}
		{
			p.SetState(160)
			p.Match(CParserT__2)
		}
		{
			p.SetState(161)
			p.expr(0)
		}
		{
			p.SetState(162)
			p.Match(CParserT__3)
		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(163)
				p.Match(CParserT__4)
			}
			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&117093594607640096) != 0 {
				{
					p.SetState(164)
					p.Stmt()
				}

				p.SetState(169)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(170)
				p.Match(CParserT__5)
			}

		case 2:
			{
				p.SetState(171)
				p.Stmt()
			}

		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(174)
					p.Elseif()
				}

			}
			p.SetState(179)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(180)
				p.Else_()
			}

		}

	case 8:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(183)
			p.Match(CParserT__10)
		}
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17451448559206408) != 0 {
			{
				p.SetState(184)
				p.expr(0)
			}

		}
		{
			p.SetState(187)
			p.Match(CParserT__0)
		}

	case 9:
		localctx = NewBlockStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(188)
			p.Match(CParserT__4)
		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&117093594607640096) != 0 {
			{
				p.SetState(189)
				p.Stmt()
			}

			p.SetState(194)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(195)
			p.Match(CParserT__5)
		}

	case 10:
		localctx = NewGotoStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(196)
			p.Match(CParserT__11)
		}
		{
			p.SetState(197)
			p.Match(CParserID)
		}
		{
			p.SetState(198)
			p.Match(CParserT__0)
		}

	case 11:
		localctx = NewLabelStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(199)
			p.Match(CParserID)
		}
		{
			p.SetState(200)
			p.Match(CParserT__12)
		}

	case 12:
		localctx = NewForStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(201)
			p.Match(CParserT__13)
		}
		{
			p.SetState(202)
			p.Match(CParserT__2)
		}
		{
			p.SetState(203)
			p.ForInit()
		}
		{
			p.SetState(204)
			p.Match(CParserT__0)
		}
		{
			p.SetState(205)
			p.ForCondition()
		}
		{
			p.SetState(206)
			p.Match(CParserT__0)
		}
		{
			p.SetState(207)
			p.ForIter()
		}
		{
			p.SetState(208)
			p.Match(CParserT__3)
		}
		{
			p.SetState(209)
			p.Stmt()
		}

	case 13:
		localctx = NewWhileStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(211)
			p.Match(CParserT__14)
		}
		{
			p.SetState(212)
			p.Match(CParserT__2)
		}
		{
			p.SetState(213)
			p.expr(0)
		}
		{
			p.SetState(214)
			p.Match(CParserT__3)
		}
		{
			p.SetState(215)
			p.Stmt()
		}

	case 14:
		localctx = NewDoWhileStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(217)
			p.Match(CParserT__15)
		}
		{
			p.SetState(218)
			p.Stmt()
		}
		{
			p.SetState(219)
			p.Match(CParserT__14)
		}
		{
			p.SetState(220)
			p.Match(CParserT__2)
		}
		{
			p.SetState(221)
			p.expr(0)
		}
		{
			p.SetState(222)
			p.Match(CParserT__3)
		}
		{
			p.SetState(223)
			p.Match(CParserT__0)
		}

	case 15:
		localctx = NewSwitchStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(225)
			p.Match(CParserT__16)
		}
		{
			p.SetState(226)
			p.Match(CParserT__2)
		}
		{
			p.SetState(227)
			p.expr(0)
		}
		{
			p.SetState(228)
			p.Match(CParserT__3)
		}
		{
			p.SetState(229)
			p.Match(CParserT__4)
		}
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CParserT__34 {
			{
				p.SetState(230)
				p.Case_()
			}

			p.SetState(235)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CParserT__35 {
			{
				p.SetState(236)
				p.Default_()
			}

		}
		{
			p.SetState(239)
			p.Match(CParserT__5)
		}

	case 16:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(241)
			p.Match(CParserT__17)
		}
		{
			p.SetState(242)
			p.Match(CParserT__0)
		}

	case 17:
		localctx = NewContinueStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(243)
			p.Match(CParserT__18)
		}
		{
			p.SetState(244)
			p.Match(CParserT__0)
		}

	case 18:
		localctx = NewCommentStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(245)
			p.Comment()
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AndExprContext struct {
	*ExprContext
	l IExprContext
	r IExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndExprContext) GetL() IExprContext { return s.l }

func (s *AndExprContext) GetR() IExprContext { return s.r }

func (s *AndExprContext) SetL(v IExprContext) { s.l = v }

func (s *AndExprContext) SetR(v IExprContext) { s.r = v }

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueExprContext struct {
	*ExprContext
}

func NewValueExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueExprContext {
	var p = new(ValueExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExprContext) AllValue_expr() []IValue_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValue_exprContext); ok {
			len++
		}
	}

	tst := make([]IValue_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValue_exprContext); ok {
			tst[i] = t.(IValue_exprContext)
			i++
		}
	}

	return tst
}

func (s *ValueExprContext) Value_expr(i int) IValue_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValue_exprContext); ok {
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

	return t.(IValue_exprContext)
}

func (s *ValueExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterValueExpr(s)
	}
}

func (s *ValueExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitValueExpr(s)
	}
}

func (s *ValueExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitValueExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SubscriptExprContext struct {
	*ExprContext
}

func NewSubscriptExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubscriptExprContext {
	var p = new(SubscriptExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SubscriptExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubscriptExprContext) Subscript_expr() ISubscript_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubscript_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubscript_exprContext)
}

func (s *SubscriptExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterSubscriptExpr(s)
	}
}

func (s *SubscriptExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitSubscriptExpr(s)
	}
}

func (s *SubscriptExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitSubscriptExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstantExprContext struct {
	*ExprContext
}

func NewConstantExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstantExprContext {
	var p = new(ConstantExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ConstantExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantExprContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ConstantExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterConstantExpr(s)
	}
}

func (s *ConstantExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitConstantExpr(s)
	}
}

func (s *ConstantExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitConstantExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitAndExprContext struct {
	*ExprContext
	l IExprContext
	r IExprContext
}

func NewBitAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitAndExprContext {
	var p = new(BitAndExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BitAndExprContext) GetL() IExprContext { return s.l }

func (s *BitAndExprContext) GetR() IExprContext { return s.r }

func (s *BitAndExprContext) SetL(v IExprContext) { s.l = v }

func (s *BitAndExprContext) SetR(v IExprContext) { s.r = v }

func (s *BitAndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitAndExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *BitAndExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *BitAndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterBitAndExpr(s)
	}
}

func (s *BitAndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitBitAndExpr(s)
	}
}

func (s *BitAndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitBitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqNeqExprContext struct {
	*ExprContext
	l IExprContext
	r IExprContext
}

func NewEqNeqExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqNeqExprContext {
	var p = new(EqNeqExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EqNeqExprContext) GetL() IExprContext { return s.l }

func (s *EqNeqExprContext) GetR() IExprContext { return s.r }

func (s *EqNeqExprContext) SetL(v IExprContext) { s.l = v }

func (s *EqNeqExprContext) SetR(v IExprContext) { s.r = v }

func (s *EqNeqExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqNeqExprContext) Eqneq() IEqneqContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqneqContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqneqContext)
}

func (s *EqNeqExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EqNeqExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *EqNeqExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterEqNeqExpr(s)
	}
}

func (s *EqNeqExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitEqNeqExpr(s)
	}
}

func (s *EqNeqExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitEqNeqExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitOrExprContext struct {
	*ExprContext
	l IExprContext
	r IExprContext
}

func NewBitOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitOrExprContext {
	var p = new(BitOrExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BitOrExprContext) GetL() IExprContext { return s.l }

func (s *BitOrExprContext) GetR() IExprContext { return s.r }

func (s *BitOrExprContext) SetL(v IExprContext) { s.l = v }

func (s *BitOrExprContext) SetR(v IExprContext) { s.r = v }

func (s *BitOrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitOrExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *BitOrExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *BitOrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterBitOrExpr(s)
	}
}

func (s *BitOrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitBitOrExpr(s)
	}
}

func (s *BitOrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitBitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConditionalExprContext struct {
	*ExprContext
	condition IExprContext
	true_     IExprContext
	false_    IExprContext
}

func NewConditionalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionalExprContext {
	var p = new(ConditionalExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ConditionalExprContext) GetCondition() IExprContext { return s.condition }

func (s *ConditionalExprContext) GetTrue_() IExprContext { return s.true_ }

func (s *ConditionalExprContext) GetFalse_() IExprContext { return s.false_ }

func (s *ConditionalExprContext) SetCondition(v IExprContext) { s.condition = v }

func (s *ConditionalExprContext) SetTrue_(v IExprContext) { s.true_ = v }

func (s *ConditionalExprContext) SetFalse_(v IExprContext) { s.false_ = v }

func (s *ConditionalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ConditionalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterConditionalExpr(s)
	}
}

func (s *ConditionalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitConditionalExpr(s)
	}
}

func (s *ConditionalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitConditionalExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	*ExprContext
	l IExprContext
	r IExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrExprContext) GetL() IExprContext { return s.l }

func (s *OrExprContext) GetR() IExprContext { return s.r }

func (s *OrExprContext) SetL(v IExprContext) { s.l = v }

func (s *OrExprContext) SetR(v IExprContext) { s.r = v }

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignExprContext struct {
	*ExprContext
}

func NewAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignExprContext {
	var p = new(AssignExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExprContext) Assign_expr() IAssign_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_exprContext)
}

func (s *AssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterAssignExpr(s)
	}
}

func (s *AssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitAssignExpr(s)
	}
}

func (s *AssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GroupExprContext struct {
	*ExprContext
}

func NewGroupExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupExprContext {
	var p = new(GroupExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GroupExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GroupExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterGroupExpr(s)
	}
}

func (s *GroupExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitGroupExpr(s)
	}
}

func (s *GroupExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitGroupExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivExprContext struct {
	*ExprContext
	l IExprContext
	r IExprContext
}

func NewMulDivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivExprContext {
	var p = new(MulDivExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MulDivExprContext) GetL() IExprContext { return s.l }

func (s *MulDivExprContext) GetR() IExprContext { return s.r }

func (s *MulDivExprContext) SetL(v IExprContext) { s.l = v }

func (s *MulDivExprContext) SetR(v IExprContext) { s.r = v }

func (s *MulDivExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivExprContext) Muldiv() IMuldivContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMuldivContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMuldivContext)
}

func (s *MulDivExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MulDivExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *MulDivExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterMulDivExpr(s)
	}
}

func (s *MulDivExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitMulDivExpr(s)
	}
}

func (s *MulDivExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitMulDivExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignSumDiffExprContext struct {
	*ExprContext
}

func NewAssignSumDiffExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignSumDiffExprContext {
	var p = new(AssignSumDiffExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AssignSumDiffExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignSumDiffExprContext) Asdexpr() IAsdexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsdexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsdexprContext)
}

func (s *AssignSumDiffExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterAssignSumDiffExpr(s)
	}
}

func (s *AssignSumDiffExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitAssignSumDiffExpr(s)
	}
}

func (s *AssignSumDiffExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitAssignSumDiffExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PostIncDecExprContext struct {
	*ExprContext
}

func NewPostIncDecExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostIncDecExprContext {
	var p = new(PostIncDecExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PostIncDecExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostIncDecExprContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *PostIncDecExprContext) Incdec() IIncdecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncdecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncdecContext)
}

func (s *PostIncDecExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterPostIncDecExpr(s)
	}
}

func (s *PostIncDecExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitPostIncDecExpr(s)
	}
}

func (s *PostIncDecExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitPostIncDecExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PreIncDecExprContext struct {
	*ExprContext
}

func NewPreIncDecExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreIncDecExprContext {
	var p = new(PreIncDecExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PreIncDecExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreIncDecExprContext) Incdec() IIncdecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncdecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncdecContext)
}

func (s *PreIncDecExprContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *PreIncDecExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterPreIncDecExpr(s)
	}
}

func (s *PreIncDecExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitPreIncDecExpr(s)
	}
}

func (s *PreIncDecExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitPreIncDecExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallExprContext struct {
	*ExprContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext {
	var p = new(CallExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) Call_expr() ICall_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_exprContext)
}

func (s *CallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterCallExpr(s)
	}
}

func (s *CallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitCallExpr(s)
	}
}

func (s *CallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	*ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubExprContext struct {
	*ExprContext
	l IExprContext
	r IExprContext
}

func NewAddSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubExprContext {
	var p = new(AddSubExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddSubExprContext) GetL() IExprContext { return s.l }

func (s *AddSubExprContext) GetR() IExprContext { return s.r }

func (s *AddSubExprContext) SetL(v IExprContext) { s.l = v }

func (s *AddSubExprContext) SetR(v IExprContext) { s.r = v }

func (s *AddSubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubExprContext) Addsub() IAddsubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddsubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddsubContext)
}

func (s *AddSubExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddSubExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *AddSubExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterAddSubExpr(s)
	}
}

func (s *AddSubExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitAddSubExpr(s)
	}
}

func (s *AddSubExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitAddSubExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitXorExprContext struct {
	*ExprContext
	l IExprContext
	r IExprContext
}

func NewBitXorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitXorExprContext {
	var p = new(BitXorExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BitXorExprContext) GetL() IExprContext { return s.l }

func (s *BitXorExprContext) GetR() IExprContext { return s.r }

func (s *BitXorExprContext) SetL(v IExprContext) { s.l = v }

func (s *BitXorExprContext) SetR(v IExprContext) { s.r = v }

func (s *BitXorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitXorExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *BitXorExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *BitXorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterBitXorExpr(s)
	}
}

func (s *BitXorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitBitXorExpr(s)
	}
}

func (s *BitXorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitBitXorExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MinusExprContext struct {
	*ExprContext
}

func NewMinusExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinusExprContext {
	var p = new(MinusExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MinusExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinusExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MinusExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterMinusExpr(s)
	}
}

func (s *MinusExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitMinusExpr(s)
	}
}

func (s *MinusExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitMinusExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *CParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, CParserRULE_expr, _p)

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
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPostIncDecExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(249)
			p.Match(CParserID)
		}
		{
			p.SetState(250)
			p.Incdec()
		}

	case 2:
		localctx = NewPreIncDecExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(251)
			p.Incdec()
		}
		{
			p.SetState(252)
			p.Match(CParserID)
		}

	case 3:
		localctx = NewMinusExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(254)
			p.Match(CParserT__19)
		}
		{
			p.SetState(255)
			p.expr(18)
		}

	case 4:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(256)
			p.Match(CParserT__20)
		}
		{
			p.SetState(257)
			p.expr(17)
		}

	case 5:
		localctx = NewAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(258)
			p.Assign_expr()
		}

	case 6:
		localctx = NewAssignSumDiffExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(259)
			p.Asdexpr()
		}

	case 7:
		localctx = NewConstantExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(260)
			p.Constant()
		}

	case 8:
		localctx = NewSubscriptExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(261)
			p.Subscript_expr()
		}

	case 9:
		localctx = NewValueExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(262)
			p.Value_expr()
		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(263)
					p.Match(CParserT__27)
				}
				{
					p.SetState(264)
					p.Value_expr()
				}

			}
			p.SetState(269)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
		}

	case 10:
		localctx = NewCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(270)
			p.Call_expr()
		}

	case 11:
		localctx = NewGroupExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(271)
			p.Match(CParserT__2)
		}
		{
			p.SetState(272)
			p.expr(0)
		}
		{
			p.SetState(273)
			p.Match(CParserT__3)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(310)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MulDivExprContext).l = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(277)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(278)
					p.Muldiv()
				}
				{
					p.SetState(279)

					var _x = p.expr(17)

					localctx.(*MulDivExprContext).r = _x
				}

			case 2:
				localctx = NewAddSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AddSubExprContext).l = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(281)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(282)
					p.Addsub()
				}
				{
					p.SetState(283)

					var _x = p.expr(16)

					localctx.(*AddSubExprContext).r = _x
				}

			case 3:
				localctx = NewEqNeqExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*EqNeqExprContext).l = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(285)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(286)
					p.Eqneq()
				}
				{
					p.SetState(287)

					var _x = p.expr(15)

					localctx.(*EqNeqExprContext).r = _x
				}

			case 4:
				localctx = NewBitAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BitAndExprContext).l = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(290)
					p.Match(CParserT__21)
				}
				{
					p.SetState(291)

					var _x = p.expr(14)

					localctx.(*BitAndExprContext).r = _x
				}

			case 5:
				localctx = NewBitXorExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BitXorExprContext).l = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(293)
					p.Match(CParserT__22)
				}
				{
					p.SetState(294)

					var _x = p.expr(13)

					localctx.(*BitXorExprContext).r = _x
				}

			case 6:
				localctx = NewBitOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BitOrExprContext).l = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(295)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(296)
					p.Match(CParserT__23)
				}
				{
					p.SetState(297)

					var _x = p.expr(12)

					localctx.(*BitOrExprContext).r = _x
				}

			case 7:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AndExprContext).l = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(298)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(299)
					p.Match(CParserT__24)
				}
				{
					p.SetState(300)

					var _x = p.expr(11)

					localctx.(*AndExprContext).r = _x
				}

			case 8:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OrExprContext).l = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(301)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(302)
					p.Match(CParserT__25)
				}
				{
					p.SetState(303)

					var _x = p.expr(10)

					localctx.(*OrExprContext).r = _x
				}

			case 9:
				localctx = NewConditionalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ConditionalExprContext).condition = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(304)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(305)
					p.Match(CParserT__26)
				}
				{
					p.SetState(306)

					var _x = p.expr(0)

					localctx.(*ConditionalExprContext).true_ = _x
				}
				{
					p.SetState(307)
					p.Match(CParserT__12)
				}
				{
					p.SetState(308)

					var _x = p.expr(7)

					localctx.(*ConditionalExprContext).false_ = _x
				}

			}

		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IValue_exprContext is an interface to support dynamic dispatch.
type IValue_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsValue_exprContext differentiates from other interfaces.
	IsValue_exprContext()
}

type Value_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValue_exprContext() *Value_exprContext {
	var p = new(Value_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_value_expr
	return p
}

func (*Value_exprContext) IsValue_exprContext() {}

func NewValue_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Value_exprContext {
	var p = new(Value_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_value_expr

	return p
}

func (s *Value_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Value_exprContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *Value_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Value_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterValue_expr(s)
	}
}

func (s *Value_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitValue_expr(s)
	}
}

func (s *Value_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitValue_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Value_expr() (localctx IValue_exprContext) {
	this := p
	_ = this

	localctx = NewValue_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CParserRULE_value_expr)

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
		p.SetState(315)
		p.Match(CParserID)
	}

	return localctx
}

// ISubscript_exprContext is an interface to support dynamic dispatch.
type ISubscript_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expr() IExprContext

	// IsSubscript_exprContext differentiates from other interfaces.
	IsSubscript_exprContext()
}

type Subscript_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubscript_exprContext() *Subscript_exprContext {
	var p = new(Subscript_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_subscript_expr
	return p
}

func (*Subscript_exprContext) IsSubscript_exprContext() {}

func NewSubscript_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Subscript_exprContext {
	var p = new(Subscript_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_subscript_expr

	return p
}

func (s *Subscript_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Subscript_exprContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *Subscript_exprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Subscript_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Subscript_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Subscript_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterSubscript_expr(s)
	}
}

func (s *Subscript_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitSubscript_expr(s)
	}
}

func (s *Subscript_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitSubscript_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Subscript_expr() (localctx ISubscript_exprContext) {
	this := p
	_ = this

	localctx = NewSubscript_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CParserRULE_subscript_expr)

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
		p.SetState(317)
		p.Match(CParserID)
	}
	{
		p.SetState(318)
		p.Match(CParserT__28)
	}
	{
		p.SetState(319)
		p.expr(0)
	}
	{
		p.SetState(320)
		p.Match(CParserT__29)
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SINGLE_COMMENT() antlr.TerminalNode
	MULTILINE_COMMENT() antlr.TerminalNode

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) SINGLE_COMMENT() antlr.TerminalNode {
	return s.GetToken(CParserSINGLE_COMMENT, 0)
}

func (s *CommentContext) MULTILINE_COMMENT() antlr.TerminalNode {
	return s.GetToken(CParserMULTILINE_COMMENT, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitComment(s)
	}
}

func (s *CommentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitComment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Comment() (localctx ICommentContext) {
	this := p
	_ = this

	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CParserRULE_comment)
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
		p.SetState(322)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CParserSINGLE_COMMENT || _la == CParserMULTILINE_COMMENT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) INT() antlr.TerminalNode {
	return s.GetToken(CParserINT, 0)
}

func (s *ConstantContext) STRING() antlr.TerminalNode {
	return s.GetToken(CParserSTRING, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CParserRULE_constant)
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
		p.SetState(324)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CParserINT || _la == CParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAssign_exprContext is an interface to support dynamic dispatch.
type IAssign_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	Expr() IExprContext

	// IsAssign_exprContext differentiates from other interfaces.
	IsAssign_exprContext()
}

type Assign_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_exprContext() *Assign_exprContext {
	var p = new(Assign_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_assign_expr
	return p
}

func (*Assign_exprContext) IsAssign_exprContext() {}

func NewAssign_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_exprContext {
	var p = new(Assign_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_assign_expr

	return p
}

func (s *Assign_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_exprContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CParserID)
}

func (s *Assign_exprContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CParserID, i)
}

func (s *Assign_exprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Assign_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterAssign_expr(s)
	}
}

func (s *Assign_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitAssign_expr(s)
	}
}

func (s *Assign_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitAssign_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Assign_expr() (localctx IAssign_exprContext) {
	this := p
	_ = this

	localctx = NewAssign_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CParserRULE_assign_expr)
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
		p.SetState(326)
		p.Match(CParserID)
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserT__27 {
		{
			p.SetState(327)
			p.Match(CParserT__27)
		}
		{
			p.SetState(328)
			p.Match(CParserID)
		}

		p.SetState(333)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(334)
		p.Match(CParserT__30)
	}
	{
		p.SetState(335)
		p.expr(0)
	}

	return localctx
}

// IConstContext is an interface to support dynamic dispatch.
type IConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConstContext differentiates from other interfaces.
	IsConstContext()
}

type ConstContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstContext() *ConstContext {
	var p = new(ConstContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_const
	return p
}

func (*ConstContext) IsConstContext() {}

func NewConstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstContext {
	var p = new(ConstContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_const

	return p
}

func (s *ConstContext) GetParser() antlr.Parser { return s.parser }
func (s *ConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterConst(s)
	}
}

func (s *ConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitConst(s)
	}
}

func (s *ConstContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitConst(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Const_() (localctx IConstContext) {
	this := p
	_ = this

	localctx = NewConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CParserRULE_const)

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
		p.SetState(337)
		p.Match(CParserT__31)
	}

	return localctx
}

// IAsdexprContext is an interface to support dynamic dispatch.
type IAsdexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	Asd() IAsdContext
	Expr() IExprContext

	// IsAsdexprContext differentiates from other interfaces.
	IsAsdexprContext()
}

type AsdexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsdexprContext() *AsdexprContext {
	var p = new(AsdexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_asdexpr
	return p
}

func (*AsdexprContext) IsAsdexprContext() {}

func NewAsdexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsdexprContext {
	var p = new(AsdexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_asdexpr

	return p
}

func (s *AsdexprContext) GetParser() antlr.Parser { return s.parser }

func (s *AsdexprContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CParserID)
}

func (s *AsdexprContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CParserID, i)
}

func (s *AsdexprContext) Asd() IAsdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsdContext)
}

func (s *AsdexprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsdexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsdexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsdexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterAsdexpr(s)
	}
}

func (s *AsdexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitAsdexpr(s)
	}
}

func (s *AsdexprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitAsdexpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Asdexpr() (localctx IAsdexprContext) {
	this := p
	_ = this

	localctx = NewAsdexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CParserRULE_asdexpr)
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
		p.SetState(339)
		p.Match(CParserID)
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserT__27 {
		{
			p.SetState(340)
			p.Match(CParserT__27)
		}
		{
			p.SetState(341)
			p.Match(CParserID)
		}

		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(347)
		p.Asd()
	}
	{
		p.SetState(348)
		p.expr(0)
	}

	return localctx
}

// IAsdContext is an interface to support dynamic dispatch.
type IAsdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAsdContext differentiates from other interfaces.
	IsAsdContext()
}

type AsdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsdContext() *AsdContext {
	var p = new(AsdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_asd
	return p
}

func (*AsdContext) IsAsdContext() {}

func NewAsdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsdContext {
	var p = new(AsdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_asd

	return p
}

func (s *AsdContext) GetParser() antlr.Parser { return s.parser }
func (s *AsdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterAsd(s)
	}
}

func (s *AsdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitAsd(s)
	}
}

func (s *AsdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitAsd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Asd() (localctx IAsdContext) {
	this := p
	_ = this

	localctx = NewAsdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CParserRULE_asd)
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
		p.SetState(350)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CParserT__32 || _la == CParserT__33) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICaseContext is an interface to support dynamic dispatch.
type ICaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsCaseContext differentiates from other interfaces.
	IsCaseContext()
}

type CaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseContext() *CaseContext {
	var p = new(CaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_case
	return p
}

func (*CaseContext) IsCaseContext() {}

func NewCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseContext {
	var p = new(CaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_case

	return p
}

func (s *CaseContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CaseContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *CaseContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *CaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterCase(s)
	}
}

func (s *CaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitCase(s)
	}
}

func (s *CaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Case_() (localctx ICaseContext) {
	this := p
	_ = this

	localctx = NewCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CParserRULE_case)
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
		p.SetState(352)
		p.Match(CParserT__34)
	}
	{
		p.SetState(353)
		p.expr(0)
	}
	{
		p.SetState(354)
		p.Match(CParserT__12)
	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&117093594607640096) != 0 {
		{
			p.SetState(355)
			p.Stmt()
		}

		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDefaultContext is an interface to support dynamic dispatch.
type IDefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsDefaultContext differentiates from other interfaces.
	IsDefaultContext()
}

type DefaultContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultContext() *DefaultContext {
	var p = new(DefaultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_default
	return p
}

func (*DefaultContext) IsDefaultContext() {}

func NewDefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultContext {
	var p = new(DefaultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_default

	return p
}

func (s *DefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *DefaultContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *DefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterDefault(s)
	}
}

func (s *DefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitDefault(s)
	}
}

func (s *DefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitDefault(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Default_() (localctx IDefaultContext) {
	this := p
	_ = this

	localctx = NewDefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CParserRULE_default)
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
		p.SetState(361)
		p.Match(CParserT__35)
	}
	{
		p.SetState(362)
		p.Match(CParserT__12)
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&117093594607640096) != 0 {
		{
			p.SetState(363)
			p.Stmt()
		}

		p.SetState(368)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IForInitContext is an interface to support dynamic dispatch.
type IForInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Definition() IDefinitionContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsForInitContext differentiates from other interfaces.
	IsForInitContext()
}

type ForInitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForInitContext() *ForInitContext {
	var p = new(ForInitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_forInit
	return p
}

func (*ForInitContext) IsForInitContext() {}

func NewForInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForInitContext {
	var p = new(ForInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_forInit

	return p
}

func (s *ForInitContext) GetParser() antlr.Parser { return s.parser }

func (s *ForInitContext) Definition() IDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *ForInitContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ForInitContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ForInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterForInit(s)
	}
}

func (s *ForInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitForInit(s)
	}
}

func (s *ForInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitForInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) ForInit() (localctx IForInitContext) {
	this := p
	_ = this

	localctx = NewForInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CParserRULE_forInit)
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

	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(369)
			p.Definition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(373)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17451448559206408) != 0 {
			{
				p.SetState(370)
				p.expr(0)
			}

			p.SetState(375)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IForConditionContext is an interface to support dynamic dispatch.
type IForConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsForConditionContext differentiates from other interfaces.
	IsForConditionContext()
}

type ForConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForConditionContext() *ForConditionContext {
	var p = new(ForConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_forCondition
	return p
}

func (*ForConditionContext) IsForConditionContext() {}

func NewForConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForConditionContext {
	var p = new(ForConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_forCondition

	return p
}

func (s *ForConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ForConditionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterForCondition(s)
	}
}

func (s *ForConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitForCondition(s)
	}
}

func (s *ForConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitForCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) ForCondition() (localctx IForConditionContext) {
	this := p
	_ = this

	localctx = NewForConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CParserRULE_forCondition)

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
		p.SetState(378)
		p.expr(0)
	}

	return localctx
}

// IForIterContext is an interface to support dynamic dispatch.
type IForIterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsForIterContext differentiates from other interfaces.
	IsForIterContext()
}

type ForIterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForIterContext() *ForIterContext {
	var p = new(ForIterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_forIter
	return p
}

func (*ForIterContext) IsForIterContext() {}

func NewForIterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForIterContext {
	var p = new(ForIterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_forIter

	return p
}

func (s *ForIterContext) GetParser() antlr.Parser { return s.parser }

func (s *ForIterContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ForIterContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ForIterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForIterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForIterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterForIter(s)
	}
}

func (s *ForIterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitForIter(s)
	}
}

func (s *ForIterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitForIter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) ForIter() (localctx IForIterContext) {
	this := p
	_ = this

	localctx = NewForIterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CParserRULE_forIter)
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
		p.SetState(380)
		p.expr(0)
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserT__7 {
		{
			p.SetState(381)
			p.Match(CParserT__7)
		}
		{
			p.SetState(382)
			p.expr(0)
		}

		p.SetState(387)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IGlobalContext is an interface to support dynamic dispatch.
type IGlobalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	ID() antlr.TerminalNode
	Constant() IConstantContext

	// IsGlobalContext differentiates from other interfaces.
	IsGlobalContext()
}

type GlobalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobalContext() *GlobalContext {
	var p = new(GlobalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_global
	return p
}

func (*GlobalContext) IsGlobalContext() {}

func NewGlobalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlobalContext {
	var p = new(GlobalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_global

	return p
}

func (s *GlobalContext) GetParser() antlr.Parser { return s.parser }

func (s *GlobalContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *GlobalContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *GlobalContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *GlobalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GlobalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterGlobal(s)
	}
}

func (s *GlobalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitGlobal(s)
	}
}

func (s *GlobalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitGlobal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Global() (localctx IGlobalContext) {
	this := p
	_ = this

	localctx = NewGlobalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CParserRULE_global)

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
		p.SetState(388)
		p.Type_()
	}
	{
		p.SetState(389)
		p.Match(CParserID)
	}
	{
		p.SetState(390)
		p.Match(CParserT__30)
	}
	{
		p.SetState(391)
		p.Constant()
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	ID() antlr.TerminalNode

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CParserRULE_declaration)

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
		p.SetState(393)
		p.Type_()
	}
	{
		p.SetState(394)
		p.Match(CParserID)
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	ID() antlr.TerminalNode
	Expr() IExprContext

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DefinitionContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *DefinitionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (s *DefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Definition() (localctx IDefinitionContext) {
	this := p
	_ = this

	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CParserRULE_definition)

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
		p.SetState(396)
		p.Type_()
	}
	{
		p.SetState(397)
		p.Match(CParserID)
	}
	{
		p.SetState(398)
		p.Match(CParserT__30)
	}
	{
		p.SetState(399)
		p.expr(0)
	}

	return localctx
}

// IElseifContext is an interface to support dynamic dispatch.
type IElseifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsElseifContext differentiates from other interfaces.
	IsElseifContext()
}

type ElseifContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseifContext() *ElseifContext {
	var p = new(ElseifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_elseif
	return p
}

func (*ElseifContext) IsElseifContext() {}

func NewElseifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseifContext {
	var p = new(ElseifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_elseif

	return p
}

func (s *ElseifContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseifContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ElseifContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ElseifContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *ElseifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterElseif(s)
	}
}

func (s *ElseifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitElseif(s)
	}
}

func (s *ElseifContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitElseif(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Elseif() (localctx IElseifContext) {
	this := p
	_ = this

	localctx = NewElseifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CParserRULE_elseif)
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
		p.SetState(401)
		p.Match(CParserT__36)
	}
	{
		p.SetState(402)
		p.Match(CParserT__2)
	}
	{
		p.SetState(403)
		p.expr(0)
	}
	{
		p.SetState(404)
		p.Match(CParserT__3)
	}
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(405)
			p.Match(CParserT__4)
		}
		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&117093594607640096) != 0 {
			{
				p.SetState(406)
				p.Stmt()
			}

			p.SetState(411)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(412)
			p.Match(CParserT__5)
		}

	case 2:
		{
			p.SetState(413)
			p.Stmt()
		}

	}

	return localctx
}

// IElseContext is an interface to support dynamic dispatch.
type IElseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsElseContext differentiates from other interfaces.
	IsElseContext()
}

type ElseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseContext() *ElseContext {
	var p = new(ElseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_else
	return p
}

func (*ElseContext) IsElseContext() {}

func NewElseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseContext {
	var p = new(ElseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_else

	return p
}

func (s *ElseContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ElseContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *ElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterElse(s)
	}
}

func (s *ElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitElse(s)
	}
}

func (s *ElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitElse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Else_() (localctx IElseContext) {
	this := p
	_ = this

	localctx = NewElseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CParserRULE_else)
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
		p.SetState(416)
		p.Match(CParserT__37)
	}
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(417)
			p.Match(CParserT__4)
		}
		p.SetState(421)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&117093594607640096) != 0 {
			{
				p.SetState(418)
				p.Stmt()
			}

			p.SetState(423)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(424)
			p.Match(CParserT__5)
		}

	case 2:
		{
			p.SetState(425)
			p.Stmt()
		}

	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArgsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (s *ArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Args() (localctx IArgsContext) {
	this := p
	_ = this

	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CParserRULE_args)
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
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17451448559206408) != 0 {
		{
			p.SetState(428)
			p.expr(0)
		}
		p.SetState(433)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CParserT__7 {
			{
				p.SetState(429)
				p.Match(CParserT__7)
			}
			{
				p.SetState(430)
				p.expr(0)
			}

			p.SetState(435)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// ICall_exprContext is an interface to support dynamic dispatch.
type ICall_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	Args() IArgsContext

	// IsCall_exprContext differentiates from other interfaces.
	IsCall_exprContext()
}

type Call_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_exprContext() *Call_exprContext {
	var p = new(Call_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_call_expr
	return p
}

func (*Call_exprContext) IsCall_exprContext() {}

func NewCall_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_exprContext {
	var p = new(Call_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_call_expr

	return p
}

func (s *Call_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_exprContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CParserID)
}

func (s *Call_exprContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CParserID, i)
}

func (s *Call_exprContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *Call_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterCall_expr(s)
	}
}

func (s *Call_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitCall_expr(s)
	}
}

func (s *Call_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitCall_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Call_expr() (localctx ICall_exprContext) {
	this := p
	_ = this

	localctx = NewCall_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, CParserRULE_call_expr)
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
		p.SetState(438)
		p.Match(CParserID)
	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserT__27 {
		{
			p.SetState(439)
			p.Match(CParserT__27)
		}
		{
			p.SetState(440)
			p.Match(CParserID)
		}

		p.SetState(445)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(446)
		p.Match(CParserT__2)
	}
	{
		p.SetState(447)
		p.Args()
	}
	{
		p.SetState(448)
		p.Match(CParserT__3)
	}

	return localctx
}

// IMuldivContext is an interface to support dynamic dispatch.
type IMuldivContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMuldivContext differentiates from other interfaces.
	IsMuldivContext()
}

type MuldivContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMuldivContext() *MuldivContext {
	var p = new(MuldivContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_muldiv
	return p
}

func (*MuldivContext) IsMuldivContext() {}

func NewMuldivContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MuldivContext {
	var p = new(MuldivContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_muldiv

	return p
}

func (s *MuldivContext) GetParser() antlr.Parser { return s.parser }
func (s *MuldivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MuldivContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MuldivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterMuldiv(s)
	}
}

func (s *MuldivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitMuldiv(s)
	}
}

func (s *MuldivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitMuldiv(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Muldiv() (localctx IMuldivContext) {
	this := p
	_ = this

	localctx = NewMuldivContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CParserRULE_muldiv)
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
		p.SetState(450)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3848290697216) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAddsubContext is an interface to support dynamic dispatch.
type IAddsubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAddsubContext differentiates from other interfaces.
	IsAddsubContext()
}

type AddsubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddsubContext() *AddsubContext {
	var p = new(AddsubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_addsub
	return p
}

func (*AddsubContext) IsAddsubContext() {}

func NewAddsubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddsubContext {
	var p = new(AddsubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_addsub

	return p
}

func (s *AddsubContext) GetParser() antlr.Parser { return s.parser }
func (s *AddsubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddsubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddsubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterAddsub(s)
	}
}

func (s *AddsubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitAddsub(s)
	}
}

func (s *AddsubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitAddsub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Addsub() (localctx IAddsubContext) {
	this := p
	_ = this

	localctx = NewAddsubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CParserRULE_addsub)
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
		p.SetState(452)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CParserT__19 || _la == CParserT__41) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEqneqContext is an interface to support dynamic dispatch.
type IEqneqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEqneqContext differentiates from other interfaces.
	IsEqneqContext()
}

type EqneqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqneqContext() *EqneqContext {
	var p = new(EqneqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_eqneq
	return p
}

func (*EqneqContext) IsEqneqContext() {}

func NewEqneqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqneqContext {
	var p = new(EqneqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_eqneq

	return p
}

func (s *EqneqContext) GetParser() antlr.Parser { return s.parser }
func (s *EqneqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqneqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqneqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterEqneq(s)
	}
}

func (s *EqneqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitEqneq(s)
	}
}

func (s *EqneqContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitEqneq(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Eqneq() (localctx IEqneqContext) {
	this := p
	_ = this

	localctx = NewEqneqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CParserRULE_eqneq)
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
		p.SetState(454)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&554153860399104) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIncdecContext is an interface to support dynamic dispatch.
type IIncdecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIncdecContext differentiates from other interfaces.
	IsIncdecContext()
}

type IncdecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncdecContext() *IncdecContext {
	var p = new(IncdecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_incdec
	return p
}

func (*IncdecContext) IsIncdecContext() {}

func NewIncdecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncdecContext {
	var p = new(IncdecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_incdec

	return p
}

func (s *IncdecContext) GetParser() antlr.Parser { return s.parser }
func (s *IncdecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncdecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncdecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterIncdec(s)
	}
}

func (s *IncdecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitIncdec(s)
	}
}

func (s *IncdecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitIncdec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Incdec() (localctx IIncdecContext) {
	this := p
	_ = this

	localctx = NewIncdecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CParserRULE_incdec)
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
		p.SetState(456)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CParserT__48 || _la == CParserT__49) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *CParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

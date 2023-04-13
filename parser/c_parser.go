// Code generated from ../C.g4 by ANTLR 4.12.0. DO NOT EDIT.

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
		"'.'", "'='", "'if'", "'return'", "'goto'", "':'", "'for'", "'while'",
		"'do'", "'switch'", "'break'", "'continue'", "'-'", "'!'", "'&'", "'^'",
		"'|'", "'&&'", "'||'", "'const'", "'+='", "'-='", "'case'", "'default'",
		"'else if'", "'else'", "'*'", "'/'", "'%'", "'+'", "'=='", "'!='", "'<'",
		"'>'", "'<='", "'>='", "'++'", "'--'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "INT", "STRING",
		"ID", "WS", "SINGLE_COMMENT", "MULTILINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"program", "include", "function", "struct", "field", "type", "params",
		"param", "stmt", "expr", "alt", "const", "asdexpr", "asd", "case", "default",
		"forInit", "forCondition", "forIter", "declaration", "definition", "elseif",
		"else", "args", "call_expr", "muldiv", "addsub", "eqneq", "incdec",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 52, 423, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0,
		64, 8, 0, 10, 0, 12, 0, 67, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 81, 8, 2, 10, 2, 12, 2, 84, 9, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 92, 8, 3, 10, 3, 12, 3, 95, 9,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 3, 5, 104, 8, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 5, 6, 111, 8, 6, 10, 6, 12, 6, 114, 9, 6, 3, 6, 116,
		8, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 5, 8, 130, 8, 8, 10, 8, 12, 8, 133, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		5, 8, 151, 8, 8, 10, 8, 12, 8, 154, 9, 8, 1, 8, 1, 8, 3, 8, 158, 8, 8,
		1, 8, 5, 8, 161, 8, 8, 10, 8, 12, 8, 164, 9, 8, 1, 8, 3, 8, 167, 8, 8,
		1, 8, 1, 8, 3, 8, 171, 8, 8, 1, 8, 1, 8, 1, 8, 5, 8, 176, 8, 8, 10, 8,
		12, 8, 179, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 5, 8, 217, 8, 8, 10, 8, 12, 8, 220, 9, 8, 1, 8, 3, 8,
		223, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 231, 8, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 251, 8, 9, 10, 9, 12, 9, 254, 9, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 261, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 4, 9, 290,
		8, 9, 11, 9, 12, 9, 291, 5, 9, 294, 8, 9, 10, 9, 12, 9, 297, 9, 9, 1, 10,
		1, 10, 1, 10, 5, 10, 302, 8, 10, 10, 10, 12, 10, 305, 9, 10, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 5, 12, 312, 8, 12, 10, 12, 12, 12, 315, 9, 12,
		1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 326,
		8, 14, 10, 14, 12, 14, 329, 9, 14, 1, 15, 1, 15, 1, 15, 5, 15, 334, 8,
		15, 10, 15, 12, 15, 337, 9, 15, 1, 16, 1, 16, 5, 16, 341, 8, 16, 10, 16,
		12, 16, 344, 9, 16, 3, 16, 346, 8, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18,
		5, 18, 353, 8, 18, 10, 18, 12, 18, 356, 9, 18, 1, 19, 1, 19, 1, 19, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		5, 21, 372, 8, 21, 10, 21, 12, 21, 375, 9, 21, 1, 21, 1, 21, 3, 21, 379,
		8, 21, 1, 22, 1, 22, 1, 22, 5, 22, 384, 8, 22, 10, 22, 12, 22, 387, 9,
		22, 1, 22, 1, 22, 3, 22, 391, 8, 22, 1, 23, 1, 23, 1, 23, 5, 23, 396, 8,
		23, 10, 23, 12, 23, 399, 9, 23, 3, 23, 401, 8, 23, 1, 24, 1, 24, 1, 24,
		5, 24, 406, 8, 24, 10, 24, 12, 24, 409, 9, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 0, 1,
		18, 29, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 0, 6, 1, 0, 47, 48, 1,
		0, 29, 30, 1, 0, 35, 37, 2, 0, 21, 21, 38, 38, 1, 0, 39, 44, 1, 0, 45,
		46, 458, 0, 65, 1, 0, 0, 0, 2, 70, 1, 0, 0, 0, 4, 73, 1, 0, 0, 0, 6, 87,
		1, 0, 0, 0, 8, 98, 1, 0, 0, 0, 10, 103, 1, 0, 0, 0, 12, 115, 1, 0, 0, 0,
		14, 117, 1, 0, 0, 0, 16, 230, 1, 0, 0, 0, 18, 260, 1, 0, 0, 0, 20, 298,
		1, 0, 0, 0, 22, 306, 1, 0, 0, 0, 24, 308, 1, 0, 0, 0, 26, 319, 1, 0, 0,
		0, 28, 321, 1, 0, 0, 0, 30, 330, 1, 0, 0, 0, 32, 345, 1, 0, 0, 0, 34, 347,
		1, 0, 0, 0, 36, 349, 1, 0, 0, 0, 38, 357, 1, 0, 0, 0, 40, 360, 1, 0, 0,
		0, 42, 365, 1, 0, 0, 0, 44, 380, 1, 0, 0, 0, 46, 400, 1, 0, 0, 0, 48, 402,
		1, 0, 0, 0, 50, 414, 1, 0, 0, 0, 52, 416, 1, 0, 0, 0, 54, 418, 1, 0, 0,
		0, 56, 420, 1, 0, 0, 0, 58, 64, 3, 2, 1, 0, 59, 64, 3, 4, 2, 0, 60, 61,
		3, 6, 3, 0, 61, 62, 5, 1, 0, 0, 62, 64, 1, 0, 0, 0, 63, 58, 1, 0, 0, 0,
		63, 59, 1, 0, 0, 0, 63, 60, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63, 1,
		0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 68, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 68,
		69, 5, 0, 0, 1, 69, 1, 1, 0, 0, 0, 70, 71, 5, 2, 0, 0, 71, 72, 5, 48, 0,
		0, 72, 3, 1, 0, 0, 0, 73, 74, 3, 10, 5, 0, 74, 75, 5, 49, 0, 0, 75, 76,
		5, 3, 0, 0, 76, 77, 3, 12, 6, 0, 77, 78, 5, 4, 0, 0, 78, 82, 5, 5, 0, 0,
		79, 81, 3, 16, 8, 0, 80, 79, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1,
		0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 85, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85,
		86, 5, 6, 0, 0, 86, 5, 1, 0, 0, 0, 87, 88, 5, 7, 0, 0, 88, 89, 5, 49, 0,
		0, 89, 93, 5, 5, 0, 0, 90, 92, 3, 8, 4, 0, 91, 90, 1, 0, 0, 0, 92, 95,
		1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 96, 1, 0, 0, 0,
		95, 93, 1, 0, 0, 0, 96, 97, 5, 6, 0, 0, 97, 7, 1, 0, 0, 0, 98, 99, 3, 10,
		5, 0, 99, 100, 5, 49, 0, 0, 100, 101, 5, 1, 0, 0, 101, 9, 1, 0, 0, 0, 102,
		104, 3, 22, 11, 0, 103, 102, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105,
		1, 0, 0, 0, 105, 106, 5, 49, 0, 0, 106, 11, 1, 0, 0, 0, 107, 112, 3, 14,
		7, 0, 108, 109, 5, 8, 0, 0, 109, 111, 3, 14, 7, 0, 110, 108, 1, 0, 0, 0,
		111, 114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113,
		116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 107, 1, 0, 0, 0, 115, 116,
		1, 0, 0, 0, 116, 13, 1, 0, 0, 0, 117, 118, 3, 10, 5, 0, 118, 119, 5, 49,
		0, 0, 119, 15, 1, 0, 0, 0, 120, 121, 3, 38, 19, 0, 121, 122, 5, 1, 0, 0,
		122, 231, 1, 0, 0, 0, 123, 124, 3, 40, 20, 0, 124, 125, 5, 1, 0, 0, 125,
		231, 1, 0, 0, 0, 126, 131, 5, 49, 0, 0, 127, 128, 5, 9, 0, 0, 128, 130,
		5, 49, 0, 0, 129, 127, 1, 0, 0, 0, 130, 133, 1, 0, 0, 0, 131, 129, 1, 0,
		0, 0, 131, 132, 1, 0, 0, 0, 132, 134, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0,
		134, 135, 5, 10, 0, 0, 135, 136, 3, 18, 9, 0, 136, 137, 5, 1, 0, 0, 137,
		231, 1, 0, 0, 0, 138, 139, 3, 24, 12, 0, 139, 140, 5, 1, 0, 0, 140, 231,
		1, 0, 0, 0, 141, 142, 3, 48, 24, 0, 142, 143, 5, 1, 0, 0, 143, 231, 1,
		0, 0, 0, 144, 145, 5, 11, 0, 0, 145, 146, 5, 3, 0, 0, 146, 147, 3, 18,
		9, 0, 147, 157, 5, 4, 0, 0, 148, 152, 5, 5, 0, 0, 149, 151, 3, 16, 8, 0,
		150, 149, 1, 0, 0, 0, 151, 154, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152,
		153, 1, 0, 0, 0, 153, 155, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 155, 158,
		5, 6, 0, 0, 156, 158, 3, 16, 8, 0, 157, 148, 1, 0, 0, 0, 157, 156, 1, 0,
		0, 0, 158, 162, 1, 0, 0, 0, 159, 161, 3, 42, 21, 0, 160, 159, 1, 0, 0,
		0, 161, 164, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163,
		166, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 165, 167, 3, 44, 22, 0, 166, 165,
		1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 231, 1, 0, 0, 0, 168, 170, 5, 12,
		0, 0, 169, 171, 3, 18, 9, 0, 170, 169, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0,
		171, 172, 1, 0, 0, 0, 172, 231, 5, 1, 0, 0, 173, 177, 5, 5, 0, 0, 174,
		176, 3, 16, 8, 0, 175, 174, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 175,
		1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 180, 1, 0, 0, 0, 179, 177, 1, 0,
		0, 0, 180, 231, 5, 6, 0, 0, 181, 182, 5, 13, 0, 0, 182, 183, 5, 49, 0,
		0, 183, 231, 5, 1, 0, 0, 184, 185, 5, 49, 0, 0, 185, 231, 5, 14, 0, 0,
		186, 187, 5, 15, 0, 0, 187, 188, 5, 3, 0, 0, 188, 189, 3, 32, 16, 0, 189,
		190, 5, 1, 0, 0, 190, 191, 3, 34, 17, 0, 191, 192, 5, 1, 0, 0, 192, 193,
		3, 36, 18, 0, 193, 194, 5, 4, 0, 0, 194, 195, 3, 16, 8, 0, 195, 231, 1,
		0, 0, 0, 196, 197, 5, 16, 0, 0, 197, 198, 5, 3, 0, 0, 198, 199, 3, 18,
		9, 0, 199, 200, 5, 4, 0, 0, 200, 201, 3, 16, 8, 0, 201, 231, 1, 0, 0, 0,
		202, 203, 5, 17, 0, 0, 203, 204, 3, 16, 8, 0, 204, 205, 5, 16, 0, 0, 205,
		206, 5, 3, 0, 0, 206, 207, 3, 18, 9, 0, 207, 208, 5, 4, 0, 0, 208, 209,
		5, 1, 0, 0, 209, 231, 1, 0, 0, 0, 210, 211, 5, 18, 0, 0, 211, 212, 5, 3,
		0, 0, 212, 213, 3, 18, 9, 0, 213, 214, 5, 4, 0, 0, 214, 218, 5, 5, 0, 0,
		215, 217, 3, 28, 14, 0, 216, 215, 1, 0, 0, 0, 217, 220, 1, 0, 0, 0, 218,
		216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 222, 1, 0, 0, 0, 220, 218,
		1, 0, 0, 0, 221, 223, 3, 30, 15, 0, 222, 221, 1, 0, 0, 0, 222, 223, 1,
		0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 225, 5, 6, 0, 0, 225, 231, 1, 0, 0,
		0, 226, 227, 5, 19, 0, 0, 227, 231, 5, 1, 0, 0, 228, 229, 5, 20, 0, 0,
		229, 231, 5, 1, 0, 0, 230, 120, 1, 0, 0, 0, 230, 123, 1, 0, 0, 0, 230,
		126, 1, 0, 0, 0, 230, 138, 1, 0, 0, 0, 230, 141, 1, 0, 0, 0, 230, 144,
		1, 0, 0, 0, 230, 168, 1, 0, 0, 0, 230, 173, 1, 0, 0, 0, 230, 181, 1, 0,
		0, 0, 230, 184, 1, 0, 0, 0, 230, 186, 1, 0, 0, 0, 230, 196, 1, 0, 0, 0,
		230, 202, 1, 0, 0, 0, 230, 210, 1, 0, 0, 0, 230, 226, 1, 0, 0, 0, 230,
		228, 1, 0, 0, 0, 231, 17, 1, 0, 0, 0, 232, 233, 6, 9, -1, 0, 233, 234,
		5, 49, 0, 0, 234, 261, 3, 56, 28, 0, 235, 236, 3, 56, 28, 0, 236, 237,
		5, 49, 0, 0, 237, 261, 1, 0, 0, 0, 238, 239, 5, 21, 0, 0, 239, 261, 3,
		18, 9, 16, 240, 241, 5, 22, 0, 0, 241, 261, 3, 18, 9, 15, 242, 261, 3,
		24, 12, 0, 243, 261, 5, 49, 0, 0, 244, 261, 7, 0, 0, 0, 245, 246, 5, 49,
		0, 0, 246, 247, 5, 9, 0, 0, 247, 252, 5, 49, 0, 0, 248, 249, 5, 9, 0, 0,
		249, 251, 5, 49, 0, 0, 250, 248, 1, 0, 0, 0, 251, 254, 1, 0, 0, 0, 252,
		250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 261, 1, 0, 0, 0, 254, 252,
		1, 0, 0, 0, 255, 261, 3, 48, 24, 0, 256, 257, 5, 3, 0, 0, 257, 258, 3,
		18, 9, 0, 258, 259, 5, 4, 0, 0, 259, 261, 1, 0, 0, 0, 260, 232, 1, 0, 0,
		0, 260, 235, 1, 0, 0, 0, 260, 238, 1, 0, 0, 0, 260, 240, 1, 0, 0, 0, 260,
		242, 1, 0, 0, 0, 260, 243, 1, 0, 0, 0, 260, 244, 1, 0, 0, 0, 260, 245,
		1, 0, 0, 0, 260, 255, 1, 0, 0, 0, 260, 256, 1, 0, 0, 0, 261, 295, 1, 0,
		0, 0, 262, 263, 10, 14, 0, 0, 263, 264, 3, 50, 25, 0, 264, 265, 3, 18,
		9, 15, 265, 294, 1, 0, 0, 0, 266, 267, 10, 13, 0, 0, 267, 268, 3, 52, 26,
		0, 268, 269, 3, 18, 9, 14, 269, 294, 1, 0, 0, 0, 270, 271, 10, 12, 0, 0,
		271, 272, 3, 54, 27, 0, 272, 273, 3, 18, 9, 13, 273, 294, 1, 0, 0, 0, 274,
		275, 10, 11, 0, 0, 275, 276, 5, 23, 0, 0, 276, 294, 3, 18, 9, 12, 277,
		278, 10, 10, 0, 0, 278, 279, 5, 24, 0, 0, 279, 294, 3, 18, 9, 11, 280,
		281, 10, 9, 0, 0, 281, 282, 5, 25, 0, 0, 282, 294, 3, 18, 9, 10, 283, 284,
		10, 8, 0, 0, 284, 285, 5, 26, 0, 0, 285, 294, 3, 18, 9, 9, 286, 289, 10,
		7, 0, 0, 287, 288, 5, 27, 0, 0, 288, 290, 3, 20, 10, 0, 289, 287, 1, 0,
		0, 0, 290, 291, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0,
		292, 294, 1, 0, 0, 0, 293, 262, 1, 0, 0, 0, 293, 266, 1, 0, 0, 0, 293,
		270, 1, 0, 0, 0, 293, 274, 1, 0, 0, 0, 293, 277, 1, 0, 0, 0, 293, 280,
		1, 0, 0, 0, 293, 283, 1, 0, 0, 0, 293, 286, 1, 0, 0, 0, 294, 297, 1, 0,
		0, 0, 295, 293, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 19, 1, 0, 0, 0,
		297, 295, 1, 0, 0, 0, 298, 303, 3, 18, 9, 0, 299, 300, 5, 26, 0, 0, 300,
		302, 3, 18, 9, 0, 301, 299, 1, 0, 0, 0, 302, 305, 1, 0, 0, 0, 303, 301,
		1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 21, 1, 0, 0, 0, 305, 303, 1, 0,
		0, 0, 306, 307, 5, 28, 0, 0, 307, 23, 1, 0, 0, 0, 308, 313, 5, 49, 0, 0,
		309, 310, 5, 9, 0, 0, 310, 312, 5, 49, 0, 0, 311, 309, 1, 0, 0, 0, 312,
		315, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 316,
		1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 316, 317, 3, 26, 13, 0, 317, 318, 3,
		18, 9, 0, 318, 25, 1, 0, 0, 0, 319, 320, 7, 1, 0, 0, 320, 27, 1, 0, 0,
		0, 321, 322, 5, 31, 0, 0, 322, 323, 3, 18, 9, 0, 323, 327, 5, 14, 0, 0,
		324, 326, 3, 16, 8, 0, 325, 324, 1, 0, 0, 0, 326, 329, 1, 0, 0, 0, 327,
		325, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 29, 1, 0, 0, 0, 329, 327, 1,
		0, 0, 0, 330, 331, 5, 32, 0, 0, 331, 335, 5, 14, 0, 0, 332, 334, 3, 16,
		8, 0, 333, 332, 1, 0, 0, 0, 334, 337, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0,
		335, 336, 1, 0, 0, 0, 336, 31, 1, 0, 0, 0, 337, 335, 1, 0, 0, 0, 338, 346,
		3, 40, 20, 0, 339, 341, 3, 18, 9, 0, 340, 339, 1, 0, 0, 0, 341, 344, 1,
		0, 0, 0, 342, 340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 346, 1, 0, 0,
		0, 344, 342, 1, 0, 0, 0, 345, 338, 1, 0, 0, 0, 345, 342, 1, 0, 0, 0, 346,
		33, 1, 0, 0, 0, 347, 348, 3, 18, 9, 0, 348, 35, 1, 0, 0, 0, 349, 354, 3,
		18, 9, 0, 350, 351, 5, 8, 0, 0, 351, 353, 3, 18, 9, 0, 352, 350, 1, 0,
		0, 0, 353, 356, 1, 0, 0, 0, 354, 352, 1, 0, 0, 0, 354, 355, 1, 0, 0, 0,
		355, 37, 1, 0, 0, 0, 356, 354, 1, 0, 0, 0, 357, 358, 3, 10, 5, 0, 358,
		359, 5, 49, 0, 0, 359, 39, 1, 0, 0, 0, 360, 361, 3, 10, 5, 0, 361, 362,
		5, 49, 0, 0, 362, 363, 5, 10, 0, 0, 363, 364, 3, 18, 9, 0, 364, 41, 1,
		0, 0, 0, 365, 366, 5, 33, 0, 0, 366, 367, 5, 3, 0, 0, 367, 368, 3, 18,
		9, 0, 368, 378, 5, 4, 0, 0, 369, 373, 5, 5, 0, 0, 370, 372, 3, 16, 8, 0,
		371, 370, 1, 0, 0, 0, 372, 375, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 373,
		374, 1, 0, 0, 0, 374, 376, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 376, 379,
		5, 6, 0, 0, 377, 379, 3, 16, 8, 0, 378, 369, 1, 0, 0, 0, 378, 377, 1, 0,
		0, 0, 379, 43, 1, 0, 0, 0, 380, 390, 5, 34, 0, 0, 381, 385, 5, 5, 0, 0,
		382, 384, 3, 16, 8, 0, 383, 382, 1, 0, 0, 0, 384, 387, 1, 0, 0, 0, 385,
		383, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 388, 1, 0, 0, 0, 387, 385,
		1, 0, 0, 0, 388, 391, 5, 6, 0, 0, 389, 391, 3, 16, 8, 0, 390, 381, 1, 0,
		0, 0, 390, 389, 1, 0, 0, 0, 391, 45, 1, 0, 0, 0, 392, 397, 3, 18, 9, 0,
		393, 394, 5, 8, 0, 0, 394, 396, 3, 18, 9, 0, 395, 393, 1, 0, 0, 0, 396,
		399, 1, 0, 0, 0, 397, 395, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 401,
		1, 0, 0, 0, 399, 397, 1, 0, 0, 0, 400, 392, 1, 0, 0, 0, 400, 401, 1, 0,
		0, 0, 401, 47, 1, 0, 0, 0, 402, 407, 5, 49, 0, 0, 403, 404, 5, 9, 0, 0,
		404, 406, 5, 49, 0, 0, 405, 403, 1, 0, 0, 0, 406, 409, 1, 0, 0, 0, 407,
		405, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408, 410, 1, 0, 0, 0, 409, 407,
		1, 0, 0, 0, 410, 411, 5, 3, 0, 0, 411, 412, 3, 46, 23, 0, 412, 413, 5,
		4, 0, 0, 413, 49, 1, 0, 0, 0, 414, 415, 7, 2, 0, 0, 415, 51, 1, 0, 0, 0,
		416, 417, 7, 3, 0, 0, 417, 53, 1, 0, 0, 0, 418, 419, 7, 4, 0, 0, 419, 55,
		1, 0, 0, 0, 420, 421, 7, 5, 0, 0, 421, 57, 1, 0, 0, 0, 36, 63, 65, 82,
		93, 103, 112, 115, 131, 152, 157, 162, 166, 170, 177, 218, 222, 230, 252,
		260, 291, 293, 295, 303, 313, 327, 335, 342, 345, 354, 373, 378, 385, 390,
		397, 400, 407,
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
	CParserINT               = 47
	CParserSTRING            = 48
	CParserID                = 49
	CParserWS                = 50
	CParserSINGLE_COMMENT    = 51
	CParserMULTILINE_COMMENT = 52
)

// CParser rules.
const (
	CParserRULE_program      = 0
	CParserRULE_include      = 1
	CParserRULE_function     = 2
	CParserRULE_struct       = 3
	CParserRULE_field        = 4
	CParserRULE_type         = 5
	CParserRULE_params       = 6
	CParserRULE_param        = 7
	CParserRULE_stmt         = 8
	CParserRULE_expr         = 9
	CParserRULE_alt          = 10
	CParserRULE_const        = 11
	CParserRULE_asdexpr      = 12
	CParserRULE_asd          = 13
	CParserRULE_case         = 14
	CParserRULE_default      = 15
	CParserRULE_forInit      = 16
	CParserRULE_forCondition = 17
	CParserRULE_forIter      = 18
	CParserRULE_declaration  = 19
	CParserRULE_definition   = 20
	CParserRULE_elseif       = 21
	CParserRULE_else         = 22
	CParserRULE_args         = 23
	CParserRULE_call_expr    = 24
	CParserRULE_muldiv       = 25
	CParserRULE_addsub       = 26
	CParserRULE_eqneq        = 27
	CParserRULE_incdec       = 28
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
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562950221856900) != 0 {
		p.SetState(63)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case CParserT__1:
			{
				p.SetState(58)
				p.Include()
			}

		case CParserT__27, CParserID:
			{
				p.SetState(59)
				p.Function()
			}

		case CParserT__6:
			{
				p.SetState(60)
				p.Struct_()
			}
			{
				p.SetState(61)
				p.Match(CParserT__0)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
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
		p.SetState(70)
		p.Match(CParserT__1)
	}
	{
		p.SetState(71)
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
		p.SetState(73)
		p.Type_()
	}
	{
		p.SetState(74)
		p.Match(CParserID)
	}
	{
		p.SetState(75)
		p.Match(CParserT__2)
	}
	{
		p.SetState(76)
		p.Params()
	}
	{
		p.SetState(77)
		p.Match(CParserT__3)
	}
	{
		p.SetState(78)
		p.Match(CParserT__4)
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562950223935520) != 0 {
		{
			p.SetState(79)
			p.Stmt()
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(85)
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
		p.SetState(87)
		p.Match(CParserT__6)
	}
	{
		p.SetState(88)
		p.Match(CParserID)
	}
	{
		p.SetState(89)
		p.Match(CParserT__4)
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserT__27 || _la == CParserID {
		{
			p.SetState(90)
			p.Field()
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(96)
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
		p.SetState(98)
		p.Type_()
	}
	{
		p.SetState(99)
		p.Match(CParserID)
	}
	{
		p.SetState(100)
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
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserT__27 {
		{
			p.SetState(102)
			p.Const_()
		}

	}
	{
		p.SetState(105)
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
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserT__27 || _la == CParserID {
		{
			p.SetState(107)
			p.Param()
		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CParserT__7 {
			{
				p.SetState(108)
				p.Match(CParserT__7)
			}
			{
				p.SetState(109)
				p.Param()
			}

			p.SetState(114)
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
		p.SetState(117)
		p.Type_()
	}
	{
		p.SetState(118)
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

type AssignmentStmtContext struct {
	*StmtContext
}

func NewAssignmentStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStmtContext {
	var p = new(AssignmentStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *AssignmentStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStmtContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CParserID)
}

func (s *AssignmentStmtContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CParserID, i)
}

func (s *AssignmentStmtContext) Expr() IExprContext {
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

func (s *AssignmentStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterAssignmentStmt(s)
	}
}

func (s *AssignmentStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitAssignmentStmt(s)
	}
}

func (s *AssignmentStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitAssignmentStmt(s)

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

	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclarationStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.Declaration()
		}
		{
			p.SetState(121)
			p.Match(CParserT__0)
		}

	case 2:
		localctx = NewDefinitionStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.Definition()
		}
		{
			p.SetState(124)
			p.Match(CParserT__0)
		}

	case 3:
		localctx = NewAssignmentStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(126)
			p.Match(CParserID)
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CParserT__8 {
			{
				p.SetState(127)
				p.Match(CParserT__8)
			}
			{
				p.SetState(128)
				p.Match(CParserID)
			}

			p.SetState(133)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(134)
			p.Match(CParserT__9)
		}
		{
			p.SetState(135)
			p.expr(0)
		}
		{
			p.SetState(136)
			p.Match(CParserT__0)
		}

	case 4:
		localctx = NewAssignSumDiffStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(138)
			p.Asdexpr()
		}
		{
			p.SetState(139)
			p.Match(CParserT__0)
		}

	case 5:
		localctx = NewCallStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(141)
			p.Call_expr()
		}
		{
			p.SetState(142)
			p.Match(CParserT__0)
		}

	case 6:
		localctx = NewIfStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(144)
			p.Match(CParserT__10)
		}
		{
			p.SetState(145)
			p.Match(CParserT__2)
		}
		{
			p.SetState(146)
			p.expr(0)
		}
		{
			p.SetState(147)
			p.Match(CParserT__3)
		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(148)
				p.Match(CParserT__4)
			}
			p.SetState(152)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562950223935520) != 0 {
				{
					p.SetState(149)
					p.Stmt()
				}

				p.SetState(154)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(155)
				p.Match(CParserT__5)
			}

		case 2:
			{
				p.SetState(156)
				p.Stmt()
			}

		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(159)
					p.Elseif()
				}

			}
			p.SetState(164)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(165)
				p.Else_()
			}

		}

	case 7:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(168)
			p.Match(CParserT__11)
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1090715541045256) != 0 {
			{
				p.SetState(169)
				p.expr(0)
			}

		}
		{
			p.SetState(172)
			p.Match(CParserT__0)
		}

	case 8:
		localctx = NewBlockStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(173)
			p.Match(CParserT__4)
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562950223935520) != 0 {
			{
				p.SetState(174)
				p.Stmt()
			}

			p.SetState(179)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(180)
			p.Match(CParserT__5)
		}

	case 9:
		localctx = NewGotoStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(181)
			p.Match(CParserT__12)
		}
		{
			p.SetState(182)
			p.Match(CParserID)
		}
		{
			p.SetState(183)
			p.Match(CParserT__0)
		}

	case 10:
		localctx = NewLabelStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(184)
			p.Match(CParserID)
		}
		{
			p.SetState(185)
			p.Match(CParserT__13)
		}

	case 11:
		localctx = NewForStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(186)
			p.Match(CParserT__14)
		}
		{
			p.SetState(187)
			p.Match(CParserT__2)
		}
		{
			p.SetState(188)
			p.ForInit()
		}
		{
			p.SetState(189)
			p.Match(CParserT__0)
		}
		{
			p.SetState(190)
			p.ForCondition()
		}
		{
			p.SetState(191)
			p.Match(CParserT__0)
		}
		{
			p.SetState(192)
			p.ForIter()
		}
		{
			p.SetState(193)
			p.Match(CParserT__3)
		}
		{
			p.SetState(194)
			p.Stmt()
		}

	case 12:
		localctx = NewWhileStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(196)
			p.Match(CParserT__15)
		}
		{
			p.SetState(197)
			p.Match(CParserT__2)
		}
		{
			p.SetState(198)
			p.expr(0)
		}
		{
			p.SetState(199)
			p.Match(CParserT__3)
		}
		{
			p.SetState(200)
			p.Stmt()
		}

	case 13:
		localctx = NewDoWhileStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(202)
			p.Match(CParserT__16)
		}
		{
			p.SetState(203)
			p.Stmt()
		}
		{
			p.SetState(204)
			p.Match(CParserT__15)
		}
		{
			p.SetState(205)
			p.Match(CParserT__2)
		}
		{
			p.SetState(206)
			p.expr(0)
		}
		{
			p.SetState(207)
			p.Match(CParserT__3)
		}
		{
			p.SetState(208)
			p.Match(CParserT__0)
		}

	case 14:
		localctx = NewSwitchStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(210)
			p.Match(CParserT__17)
		}
		{
			p.SetState(211)
			p.Match(CParserT__2)
		}
		{
			p.SetState(212)
			p.expr(0)
		}
		{
			p.SetState(213)
			p.Match(CParserT__3)
		}
		{
			p.SetState(214)
			p.Match(CParserT__4)
		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CParserT__30 {
			{
				p.SetState(215)
				p.Case_()
			}

			p.SetState(220)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CParserT__31 {
			{
				p.SetState(221)
				p.Default_()
			}

		}
		{
			p.SetState(224)
			p.Match(CParserT__5)
		}

	case 15:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(226)
			p.Match(CParserT__18)
		}
		{
			p.SetState(227)
			p.Match(CParserT__0)
		}

	case 16:
		localctx = NewContinueStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(228)
			p.Match(CParserT__19)
		}
		{
			p.SetState(229)
			p.Match(CParserT__0)
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
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

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

func (s *ConstantExprContext) INT() antlr.TerminalNode {
	return s.GetToken(CParserINT, 0)
}

func (s *ConstantExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(CParserSTRING, 0)
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
}

func NewBitAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitAndExprContext {
	var p = new(BitAndExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

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
}

func NewEqNeqExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqNeqExprContext {
	var p = new(EqNeqExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EqNeqExprContext) GetRuleContext() antlr.RuleContext {
	return s
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
}

func NewBitOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitOrExprContext {
	var p = new(BitOrExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

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

type OrExprContext struct {
	*ExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) Expr() IExprContext {
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

func (s *OrExprContext) AllAlt() []IAltContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAltContext); ok {
			len++
		}
	}

	tst := make([]IAltContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAltContext); ok {
			tst[i] = t.(IAltContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) Alt(i int) IAltContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAltContext); ok {
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

	return t.(IAltContext)
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
}

func NewMulDivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivExprContext {
	var p = new(MulDivExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MulDivExprContext) GetRuleContext() antlr.RuleContext {
	return s
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

type MemberExprContext struct {
	*ExprContext
}

func NewMemberExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberExprContext {
	var p = new(MemberExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MemberExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberExprContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CParserID)
}

func (s *MemberExprContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CParserID, i)
}

func (s *MemberExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterMemberExpr(s)
	}
}

func (s *MemberExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitMemberExpr(s)
	}
}

func (s *MemberExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitMemberExpr(s)

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

type VariableExprContext struct {
	*ExprContext
}

func NewVariableExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableExprContext {
	var p = new(VariableExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *VariableExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableExprContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *VariableExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterVariableExpr(s)
	}
}

func (s *VariableExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitVariableExpr(s)
	}
}

func (s *VariableExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitVariableExpr(s)

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
}

func NewAddSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubExprContext {
	var p = new(AddSubExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddSubExprContext) GetRuleContext() antlr.RuleContext {
	return s
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
}

func NewBitXorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitXorExprContext {
	var p = new(BitXorExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

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
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPostIncDecExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(233)
			p.Match(CParserID)
		}
		{
			p.SetState(234)
			p.Incdec()
		}

	case 2:
		localctx = NewPreIncDecExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(235)
			p.Incdec()
		}
		{
			p.SetState(236)
			p.Match(CParserID)
		}

	case 3:
		localctx = NewMinusExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(238)
			p.Match(CParserT__20)
		}
		{
			p.SetState(239)
			p.expr(16)
		}

	case 4:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(240)
			p.Match(CParserT__21)
		}
		{
			p.SetState(241)
			p.expr(15)
		}

	case 5:
		localctx = NewAssignSumDiffExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(242)
			p.Asdexpr()
		}

	case 6:
		localctx = NewVariableExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(243)
			p.Match(CParserID)
		}

	case 7:
		localctx = NewConstantExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(244)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CParserINT || _la == CParserSTRING) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 8:
		localctx = NewMemberExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(245)
			p.Match(CParserID)
		}
		{
			p.SetState(246)
			p.Match(CParserT__8)
		}
		{
			p.SetState(247)
			p.Match(CParserID)
		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(248)
					p.Match(CParserT__8)
				}
				{
					p.SetState(249)
					p.Match(CParserID)
				}

			}
			p.SetState(254)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
		}

	case 9:
		localctx = NewCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(255)
			p.Call_expr()
		}

	case 10:
		localctx = NewGroupExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(256)
			p.Match(CParserT__2)
		}
		{
			p.SetState(257)
			p.expr(0)
		}
		{
			p.SetState(258)
			p.Match(CParserT__3)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(293)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(262)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(263)
					p.Muldiv()
				}
				{
					p.SetState(264)
					p.expr(15)
				}

			case 2:
				localctx = NewAddSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(266)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(267)
					p.Addsub()
				}
				{
					p.SetState(268)
					p.expr(14)
				}

			case 3:
				localctx = NewEqNeqExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(270)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(271)
					p.Eqneq()
				}
				{
					p.SetState(272)
					p.expr(13)
				}

			case 4:
				localctx = NewBitAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(274)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(275)
					p.Match(CParserT__22)
				}
				{
					p.SetState(276)
					p.expr(12)
				}

			case 5:
				localctx = NewBitXorExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(277)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(278)
					p.Match(CParserT__23)
				}
				{
					p.SetState(279)
					p.expr(11)
				}

			case 6:
				localctx = NewBitOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(280)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(281)
					p.Match(CParserT__24)
				}
				{
					p.SetState(282)
					p.expr(10)
				}

			case 7:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(283)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(284)
					p.Match(CParserT__25)
				}
				{
					p.SetState(285)
					p.expr(9)
				}

			case 8:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(286)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				p.SetState(289)
				p.GetErrorHandler().Sync(p)
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(287)
							p.Match(CParserT__26)
						}
						{
							p.SetState(288)
							p.Alt()
						}

					default:
						panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					}

					p.SetState(291)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
				}

			}

		}
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}

	return localctx
}

// IAltContext is an interface to support dynamic dispatch.
type IAltContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsAltContext differentiates from other interfaces.
	IsAltContext()
}

type AltContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAltContext() *AltContext {
	var p = new(AltContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_alt
	return p
}

func (*AltContext) IsAltContext() {}

func NewAltContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AltContext {
	var p = new(AltContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_alt

	return p
}

func (s *AltContext) GetParser() antlr.Parser { return s.parser }

func (s *AltContext) AllExpr() []IExprContext {
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

func (s *AltContext) Expr(i int) IExprContext {
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

func (s *AltContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AltContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AltContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterAlt(s)
	}
}

func (s *AltContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitAlt(s)
	}
}

func (s *AltContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitAlt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Alt() (localctx IAltContext) {
	this := p
	_ = this

	localctx = NewAltContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CParserRULE_alt)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.expr(0)
	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(299)
				p.Match(CParserT__25)
			}
			{
				p.SetState(300)
				p.expr(0)
			}

		}
		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 22, CParserRULE_const)

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
		p.SetState(306)
		p.Match(CParserT__27)
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
	p.EnterRule(localctx, 24, CParserRULE_asdexpr)
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
		p.SetState(308)
		p.Match(CParserID)
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserT__8 {
		{
			p.SetState(309)
			p.Match(CParserT__8)
		}
		{
			p.SetState(310)
			p.Match(CParserID)
		}

		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(316)
		p.Asd()
	}
	{
		p.SetState(317)
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
	p.EnterRule(localctx, 26, CParserRULE_asd)
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
		p.SetState(319)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CParserT__28 || _la == CParserT__29) {
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
	p.EnterRule(localctx, 28, CParserRULE_case)
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
		p.SetState(321)
		p.Match(CParserT__30)
	}
	{
		p.SetState(322)
		p.expr(0)
	}
	{
		p.SetState(323)
		p.Match(CParserT__13)
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562950223935520) != 0 {
		{
			p.SetState(324)
			p.Stmt()
		}

		p.SetState(329)
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
	p.EnterRule(localctx, 30, CParserRULE_default)
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
		p.SetState(330)
		p.Match(CParserT__31)
	}
	{
		p.SetState(331)
		p.Match(CParserT__13)
	}
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562950223935520) != 0 {
		{
			p.SetState(332)
			p.Stmt()
		}

		p.SetState(337)
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
	p.EnterRule(localctx, 32, CParserRULE_forInit)
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

	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(338)
			p.Definition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1090715541045256) != 0 {
			{
				p.SetState(339)
				p.expr(0)
			}

			p.SetState(344)
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
	p.EnterRule(localctx, 34, CParserRULE_forCondition)

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
		p.SetState(347)
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
	p.EnterRule(localctx, 36, CParserRULE_forIter)
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
		p.SetState(349)
		p.expr(0)
	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserT__7 {
		{
			p.SetState(350)
			p.Match(CParserT__7)
		}
		{
			p.SetState(351)
			p.expr(0)
		}

		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 38, CParserRULE_declaration)

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
		p.SetState(357)
		p.Type_()
	}
	{
		p.SetState(358)
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
	p.EnterRule(localctx, 40, CParserRULE_definition)

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
		p.SetState(360)
		p.Type_()
	}
	{
		p.SetState(361)
		p.Match(CParserID)
	}
	{
		p.SetState(362)
		p.Match(CParserT__9)
	}
	{
		p.SetState(363)
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
	p.EnterRule(localctx, 42, CParserRULE_elseif)
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
		p.SetState(365)
		p.Match(CParserT__32)
	}
	{
		p.SetState(366)
		p.Match(CParserT__2)
	}
	{
		p.SetState(367)
		p.expr(0)
	}
	{
		p.SetState(368)
		p.Match(CParserT__3)
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(369)
			p.Match(CParserT__4)
		}
		p.SetState(373)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562950223935520) != 0 {
			{
				p.SetState(370)
				p.Stmt()
			}

			p.SetState(375)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(376)
			p.Match(CParserT__5)
		}

	case 2:
		{
			p.SetState(377)
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
	p.EnterRule(localctx, 44, CParserRULE_else)
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
		p.Match(CParserT__33)
	}
	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(381)
			p.Match(CParserT__4)
		}
		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562950223935520) != 0 {
			{
				p.SetState(382)
				p.Stmt()
			}

			p.SetState(387)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(388)
			p.Match(CParserT__5)
		}

	case 2:
		{
			p.SetState(389)
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
	p.EnterRule(localctx, 46, CParserRULE_args)
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
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1090715541045256) != 0 {
		{
			p.SetState(392)
			p.expr(0)
		}
		p.SetState(397)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CParserT__7 {
			{
				p.SetState(393)
				p.Match(CParserT__7)
			}
			{
				p.SetState(394)
				p.expr(0)
			}

			p.SetState(399)
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
	p.EnterRule(localctx, 48, CParserRULE_call_expr)
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
		p.SetState(402)
		p.Match(CParserID)
	}
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserT__8 {
		{
			p.SetState(403)
			p.Match(CParserT__8)
		}
		{
			p.SetState(404)
			p.Match(CParserID)
		}

		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(410)
		p.Match(CParserT__2)
	}
	{
		p.SetState(411)
		p.Args()
	}
	{
		p.SetState(412)
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
	p.EnterRule(localctx, 50, CParserRULE_muldiv)
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
		p.SetState(414)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&240518168576) != 0) {
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
	p.EnterRule(localctx, 52, CParserRULE_addsub)
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
		_la = p.GetTokenStream().LA(1)

		if !(_la == CParserT__20 || _la == CParserT__37) {
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
	p.EnterRule(localctx, 54, CParserRULE_eqneq)
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
		p.SetState(418)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34634616274944) != 0) {
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
	p.EnterRule(localctx, 56, CParserRULE_incdec)
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
		p.SetState(420)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CParserT__44 || _la == CParserT__45) {
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
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

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
		"'if'", "'return'", "'goto'", "':'", "'for'", "'while'", "'do'", "'switch'",
		"'break'", "'continue'", "'-'", "'!'", "'&'", "'^'", "'|'", "'&&'",
		"'||'", "'?'", "'.'", "'='", "'const'", "'+='", "'-='", "'case'", "'default'",
		"'else if'", "'else'", "'*'", "'/'", "'%'", "'+'", "'=='", "'!='", "'<'",
		"'>'", "'<='", "'>='", "'++'", "'--'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "INT", "STRING",
		"ID", "WS", "SINGLE_COMMENT", "MULTILINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"program", "include", "function", "struct", "field", "type", "params",
		"param", "stmt", "expr", "constant", "assign_expr", "const", "asdexpr",
		"asd", "case", "default", "forInit", "forCondition", "forIter", "global",
		"declaration", "definition", "elseif", "else", "args", "call_expr",
		"muldiv", "addsub", "eqneq", "incdec",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 53, 434, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 71, 8, 0, 10, 0, 12, 0, 74,
		9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 5, 2, 88, 8, 2, 10, 2, 12, 2, 91, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 5, 3, 99, 8, 3, 10, 3, 12, 3, 102, 9, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 5, 3, 5, 111, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5,
		6, 118, 8, 6, 10, 6, 12, 6, 121, 9, 6, 3, 6, 123, 8, 6, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 149, 8,
		8, 10, 8, 12, 8, 152, 9, 8, 1, 8, 1, 8, 3, 8, 156, 8, 8, 1, 8, 5, 8, 159,
		8, 8, 10, 8, 12, 8, 162, 9, 8, 1, 8, 3, 8, 165, 8, 8, 1, 8, 1, 8, 3, 8,
		169, 8, 8, 1, 8, 1, 8, 1, 8, 5, 8, 174, 8, 8, 10, 8, 12, 8, 177, 9, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		5, 8, 215, 8, 8, 10, 8, 12, 8, 218, 9, 8, 1, 8, 3, 8, 221, 8, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 229, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 5, 9, 250, 8, 9, 10, 9, 12, 9, 253, 9, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 3, 9, 260, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 5, 9, 295, 8, 9, 10, 9, 12, 9, 298, 9, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 5, 11, 305, 8, 11, 10, 11, 12, 11, 308, 9, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 5, 13, 318, 8, 13,
		10, 13, 12, 13, 321, 9, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 5, 15, 332, 8, 15, 10, 15, 12, 15, 335, 9, 15, 1, 16,
		1, 16, 1, 16, 5, 16, 340, 8, 16, 10, 16, 12, 16, 343, 9, 16, 1, 17, 1,
		17, 5, 17, 347, 8, 17, 10, 17, 12, 17, 350, 9, 17, 3, 17, 352, 8, 17, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 19, 5, 19, 359, 8, 19, 10, 19, 12, 19, 362,
		9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23,
		383, 8, 23, 10, 23, 12, 23, 386, 9, 23, 1, 23, 1, 23, 3, 23, 390, 8, 23,
		1, 24, 1, 24, 1, 24, 5, 24, 395, 8, 24, 10, 24, 12, 24, 398, 9, 24, 1,
		24, 1, 24, 3, 24, 402, 8, 24, 1, 25, 1, 25, 1, 25, 5, 25, 407, 8, 25, 10,
		25, 12, 25, 410, 9, 25, 3, 25, 412, 8, 25, 1, 26, 1, 26, 1, 26, 5, 26,
		417, 8, 26, 10, 26, 12, 26, 420, 9, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 0, 1, 18, 31,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
		38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 0, 6, 1, 0, 48, 49, 1,
		0, 30, 31, 1, 0, 36, 38, 2, 0, 19, 19, 39, 39, 1, 0, 40, 45, 1, 0, 46,
		47, 468, 0, 72, 1, 0, 0, 0, 2, 77, 1, 0, 0, 0, 4, 80, 1, 0, 0, 0, 6, 94,
		1, 0, 0, 0, 8, 105, 1, 0, 0, 0, 10, 110, 1, 0, 0, 0, 12, 122, 1, 0, 0,
		0, 14, 124, 1, 0, 0, 0, 16, 228, 1, 0, 0, 0, 18, 259, 1, 0, 0, 0, 20, 299,
		1, 0, 0, 0, 22, 301, 1, 0, 0, 0, 24, 312, 1, 0, 0, 0, 26, 314, 1, 0, 0,
		0, 28, 325, 1, 0, 0, 0, 30, 327, 1, 0, 0, 0, 32, 336, 1, 0, 0, 0, 34, 351,
		1, 0, 0, 0, 36, 353, 1, 0, 0, 0, 38, 355, 1, 0, 0, 0, 40, 363, 1, 0, 0,
		0, 42, 368, 1, 0, 0, 0, 44, 371, 1, 0, 0, 0, 46, 376, 1, 0, 0, 0, 48, 391,
		1, 0, 0, 0, 50, 411, 1, 0, 0, 0, 52, 413, 1, 0, 0, 0, 54, 425, 1, 0, 0,
		0, 56, 427, 1, 0, 0, 0, 58, 429, 1, 0, 0, 0, 60, 431, 1, 0, 0, 0, 62, 71,
		3, 2, 1, 0, 63, 64, 3, 40, 20, 0, 64, 65, 5, 1, 0, 0, 65, 71, 1, 0, 0,
		0, 66, 71, 3, 4, 2, 0, 67, 68, 3, 6, 3, 0, 68, 69, 5, 1, 0, 0, 69, 71,
		1, 0, 0, 0, 70, 62, 1, 0, 0, 0, 70, 63, 1, 0, 0, 0, 70, 66, 1, 0, 0, 0,
		70, 67, 1, 0, 0, 0, 71, 74, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 72, 73, 1,
		0, 0, 0, 73, 75, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 75, 76, 5, 0, 0, 1, 76,
		1, 1, 0, 0, 0, 77, 78, 5, 2, 0, 0, 78, 79, 5, 49, 0, 0, 79, 3, 1, 0, 0,
		0, 80, 81, 3, 10, 5, 0, 81, 82, 5, 50, 0, 0, 82, 83, 5, 3, 0, 0, 83, 84,
		3, 12, 6, 0, 84, 85, 5, 4, 0, 0, 85, 89, 5, 5, 0, 0, 86, 88, 3, 16, 8,
		0, 87, 86, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90,
		1, 0, 0, 0, 90, 92, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 93, 5, 6, 0, 0,
		93, 5, 1, 0, 0, 0, 94, 95, 5, 7, 0, 0, 95, 96, 5, 50, 0, 0, 96, 100, 5,
		5, 0, 0, 97, 99, 3, 8, 4, 0, 98, 97, 1, 0, 0, 0, 99, 102, 1, 0, 0, 0, 100,
		98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 103, 1, 0, 0, 0, 102, 100, 1,
		0, 0, 0, 103, 104, 5, 6, 0, 0, 104, 7, 1, 0, 0, 0, 105, 106, 3, 10, 5,
		0, 106, 107, 5, 50, 0, 0, 107, 108, 5, 1, 0, 0, 108, 9, 1, 0, 0, 0, 109,
		111, 3, 24, 12, 0, 110, 109, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112,
		1, 0, 0, 0, 112, 113, 5, 50, 0, 0, 113, 11, 1, 0, 0, 0, 114, 119, 3, 14,
		7, 0, 115, 116, 5, 8, 0, 0, 116, 118, 3, 14, 7, 0, 117, 115, 1, 0, 0, 0,
		118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120,
		123, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 114, 1, 0, 0, 0, 122, 123,
		1, 0, 0, 0, 123, 13, 1, 0, 0, 0, 124, 125, 3, 10, 5, 0, 125, 126, 5, 50,
		0, 0, 126, 15, 1, 0, 0, 0, 127, 128, 3, 42, 21, 0, 128, 129, 5, 1, 0, 0,
		129, 229, 1, 0, 0, 0, 130, 131, 3, 44, 22, 0, 131, 132, 5, 1, 0, 0, 132,
		229, 1, 0, 0, 0, 133, 134, 3, 22, 11, 0, 134, 135, 5, 1, 0, 0, 135, 229,
		1, 0, 0, 0, 136, 137, 3, 26, 13, 0, 137, 138, 5, 1, 0, 0, 138, 229, 1,
		0, 0, 0, 139, 140, 3, 52, 26, 0, 140, 141, 5, 1, 0, 0, 141, 229, 1, 0,
		0, 0, 142, 143, 5, 9, 0, 0, 143, 144, 5, 3, 0, 0, 144, 145, 3, 18, 9, 0,
		145, 155, 5, 4, 0, 0, 146, 150, 5, 5, 0, 0, 147, 149, 3, 16, 8, 0, 148,
		147, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151,
		1, 0, 0, 0, 151, 153, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 156, 5, 6,
		0, 0, 154, 156, 3, 16, 8, 0, 155, 146, 1, 0, 0, 0, 155, 154, 1, 0, 0, 0,
		156, 160, 1, 0, 0, 0, 157, 159, 3, 46, 23, 0, 158, 157, 1, 0, 0, 0, 159,
		162, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 164,
		1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 163, 165, 3, 48, 24, 0, 164, 163, 1,
		0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 229, 1, 0, 0, 0, 166, 168, 5, 10, 0,
		0, 167, 169, 3, 18, 9, 0, 168, 167, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169,
		170, 1, 0, 0, 0, 170, 229, 5, 1, 0, 0, 171, 175, 5, 5, 0, 0, 172, 174,
		3, 16, 8, 0, 173, 172, 1, 0, 0, 0, 174, 177, 1, 0, 0, 0, 175, 173, 1, 0,
		0, 0, 175, 176, 1, 0, 0, 0, 176, 178, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0,
		178, 229, 5, 6, 0, 0, 179, 180, 5, 11, 0, 0, 180, 181, 5, 50, 0, 0, 181,
		229, 5, 1, 0, 0, 182, 183, 5, 50, 0, 0, 183, 229, 5, 12, 0, 0, 184, 185,
		5, 13, 0, 0, 185, 186, 5, 3, 0, 0, 186, 187, 3, 34, 17, 0, 187, 188, 5,
		1, 0, 0, 188, 189, 3, 36, 18, 0, 189, 190, 5, 1, 0, 0, 190, 191, 3, 38,
		19, 0, 191, 192, 5, 4, 0, 0, 192, 193, 3, 16, 8, 0, 193, 229, 1, 0, 0,
		0, 194, 195, 5, 14, 0, 0, 195, 196, 5, 3, 0, 0, 196, 197, 3, 18, 9, 0,
		197, 198, 5, 4, 0, 0, 198, 199, 3, 16, 8, 0, 199, 229, 1, 0, 0, 0, 200,
		201, 5, 15, 0, 0, 201, 202, 3, 16, 8, 0, 202, 203, 5, 14, 0, 0, 203, 204,
		5, 3, 0, 0, 204, 205, 3, 18, 9, 0, 205, 206, 5, 4, 0, 0, 206, 207, 5, 1,
		0, 0, 207, 229, 1, 0, 0, 0, 208, 209, 5, 16, 0, 0, 209, 210, 5, 3, 0, 0,
		210, 211, 3, 18, 9, 0, 211, 212, 5, 4, 0, 0, 212, 216, 5, 5, 0, 0, 213,
		215, 3, 30, 15, 0, 214, 213, 1, 0, 0, 0, 215, 218, 1, 0, 0, 0, 216, 214,
		1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 220, 1, 0, 0, 0, 218, 216, 1, 0,
		0, 0, 219, 221, 3, 32, 16, 0, 220, 219, 1, 0, 0, 0, 220, 221, 1, 0, 0,
		0, 221, 222, 1, 0, 0, 0, 222, 223, 5, 6, 0, 0, 223, 229, 1, 0, 0, 0, 224,
		225, 5, 17, 0, 0, 225, 229, 5, 1, 0, 0, 226, 227, 5, 18, 0, 0, 227, 229,
		5, 1, 0, 0, 228, 127, 1, 0, 0, 0, 228, 130, 1, 0, 0, 0, 228, 133, 1, 0,
		0, 0, 228, 136, 1, 0, 0, 0, 228, 139, 1, 0, 0, 0, 228, 142, 1, 0, 0, 0,
		228, 166, 1, 0, 0, 0, 228, 171, 1, 0, 0, 0, 228, 179, 1, 0, 0, 0, 228,
		182, 1, 0, 0, 0, 228, 184, 1, 0, 0, 0, 228, 194, 1, 0, 0, 0, 228, 200,
		1, 0, 0, 0, 228, 208, 1, 0, 0, 0, 228, 224, 1, 0, 0, 0, 228, 226, 1, 0,
		0, 0, 229, 17, 1, 0, 0, 0, 230, 231, 6, 9, -1, 0, 231, 232, 5, 50, 0, 0,
		232, 260, 3, 60, 30, 0, 233, 234, 3, 60, 30, 0, 234, 235, 5, 50, 0, 0,
		235, 260, 1, 0, 0, 0, 236, 237, 5, 19, 0, 0, 237, 260, 3, 18, 9, 18, 238,
		239, 5, 20, 0, 0, 239, 260, 3, 18, 9, 17, 240, 260, 3, 22, 11, 0, 241,
		260, 3, 26, 13, 0, 242, 260, 5, 50, 0, 0, 243, 260, 3, 20, 10, 0, 244,
		245, 5, 50, 0, 0, 245, 246, 5, 27, 0, 0, 246, 251, 5, 50, 0, 0, 247, 248,
		5, 27, 0, 0, 248, 250, 5, 50, 0, 0, 249, 247, 1, 0, 0, 0, 250, 253, 1,
		0, 0, 0, 251, 249, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 260, 1, 0, 0,
		0, 253, 251, 1, 0, 0, 0, 254, 260, 3, 52, 26, 0, 255, 256, 5, 3, 0, 0,
		256, 257, 3, 18, 9, 0, 257, 258, 5, 4, 0, 0, 258, 260, 1, 0, 0, 0, 259,
		230, 1, 0, 0, 0, 259, 233, 1, 0, 0, 0, 259, 236, 1, 0, 0, 0, 259, 238,
		1, 0, 0, 0, 259, 240, 1, 0, 0, 0, 259, 241, 1, 0, 0, 0, 259, 242, 1, 0,
		0, 0, 259, 243, 1, 0, 0, 0, 259, 244, 1, 0, 0, 0, 259, 254, 1, 0, 0, 0,
		259, 255, 1, 0, 0, 0, 260, 296, 1, 0, 0, 0, 261, 262, 10, 16, 0, 0, 262,
		263, 3, 54, 27, 0, 263, 264, 3, 18, 9, 17, 264, 295, 1, 0, 0, 0, 265, 266,
		10, 15, 0, 0, 266, 267, 3, 56, 28, 0, 267, 268, 3, 18, 9, 16, 268, 295,
		1, 0, 0, 0, 269, 270, 10, 14, 0, 0, 270, 271, 3, 58, 29, 0, 271, 272, 3,
		18, 9, 15, 272, 295, 1, 0, 0, 0, 273, 274, 10, 13, 0, 0, 274, 275, 5, 21,
		0, 0, 275, 295, 3, 18, 9, 14, 276, 277, 10, 12, 0, 0, 277, 278, 5, 22,
		0, 0, 278, 295, 3, 18, 9, 13, 279, 280, 10, 11, 0, 0, 280, 281, 5, 23,
		0, 0, 281, 295, 3, 18, 9, 12, 282, 283, 10, 10, 0, 0, 283, 284, 5, 24,
		0, 0, 284, 295, 3, 18, 9, 11, 285, 286, 10, 9, 0, 0, 286, 287, 5, 25, 0,
		0, 287, 295, 3, 18, 9, 10, 288, 289, 10, 6, 0, 0, 289, 290, 5, 26, 0, 0,
		290, 291, 3, 18, 9, 0, 291, 292, 5, 12, 0, 0, 292, 293, 3, 18, 9, 7, 293,
		295, 1, 0, 0, 0, 294, 261, 1, 0, 0, 0, 294, 265, 1, 0, 0, 0, 294, 269,
		1, 0, 0, 0, 294, 273, 1, 0, 0, 0, 294, 276, 1, 0, 0, 0, 294, 279, 1, 0,
		0, 0, 294, 282, 1, 0, 0, 0, 294, 285, 1, 0, 0, 0, 294, 288, 1, 0, 0, 0,
		295, 298, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297,
		19, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 299, 300, 7, 0, 0, 0, 300, 21, 1,
		0, 0, 0, 301, 306, 5, 50, 0, 0, 302, 303, 5, 27, 0, 0, 303, 305, 5, 50,
		0, 0, 304, 302, 1, 0, 0, 0, 305, 308, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0,
		306, 307, 1, 0, 0, 0, 307, 309, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 309,
		310, 5, 28, 0, 0, 310, 311, 3, 18, 9, 0, 311, 23, 1, 0, 0, 0, 312, 313,
		5, 29, 0, 0, 313, 25, 1, 0, 0, 0, 314, 319, 5, 50, 0, 0, 315, 316, 5, 27,
		0, 0, 316, 318, 5, 50, 0, 0, 317, 315, 1, 0, 0, 0, 318, 321, 1, 0, 0, 0,
		319, 317, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 322, 1, 0, 0, 0, 321,
		319, 1, 0, 0, 0, 322, 323, 3, 28, 14, 0, 323, 324, 3, 18, 9, 0, 324, 27,
		1, 0, 0, 0, 325, 326, 7, 1, 0, 0, 326, 29, 1, 0, 0, 0, 327, 328, 5, 32,
		0, 0, 328, 329, 3, 18, 9, 0, 329, 333, 5, 12, 0, 0, 330, 332, 3, 16, 8,
		0, 331, 330, 1, 0, 0, 0, 332, 335, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 333,
		334, 1, 0, 0, 0, 334, 31, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0, 336, 337, 5,
		33, 0, 0, 337, 341, 5, 12, 0, 0, 338, 340, 3, 16, 8, 0, 339, 338, 1, 0,
		0, 0, 340, 343, 1, 0, 0, 0, 341, 339, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0,
		342, 33, 1, 0, 0, 0, 343, 341, 1, 0, 0, 0, 344, 352, 3, 44, 22, 0, 345,
		347, 3, 18, 9, 0, 346, 345, 1, 0, 0, 0, 347, 350, 1, 0, 0, 0, 348, 346,
		1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 352, 1, 0, 0, 0, 350, 348, 1, 0,
		0, 0, 351, 344, 1, 0, 0, 0, 351, 348, 1, 0, 0, 0, 352, 35, 1, 0, 0, 0,
		353, 354, 3, 18, 9, 0, 354, 37, 1, 0, 0, 0, 355, 360, 3, 18, 9, 0, 356,
		357, 5, 8, 0, 0, 357, 359, 3, 18, 9, 0, 358, 356, 1, 0, 0, 0, 359, 362,
		1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 39, 1, 0,
		0, 0, 362, 360, 1, 0, 0, 0, 363, 364, 3, 10, 5, 0, 364, 365, 5, 50, 0,
		0, 365, 366, 5, 28, 0, 0, 366, 367, 3, 20, 10, 0, 367, 41, 1, 0, 0, 0,
		368, 369, 3, 10, 5, 0, 369, 370, 5, 50, 0, 0, 370, 43, 1, 0, 0, 0, 371,
		372, 3, 10, 5, 0, 372, 373, 5, 50, 0, 0, 373, 374, 5, 28, 0, 0, 374, 375,
		3, 18, 9, 0, 375, 45, 1, 0, 0, 0, 376, 377, 5, 34, 0, 0, 377, 378, 5, 3,
		0, 0, 378, 379, 3, 18, 9, 0, 379, 389, 5, 4, 0, 0, 380, 384, 5, 5, 0, 0,
		381, 383, 3, 16, 8, 0, 382, 381, 1, 0, 0, 0, 383, 386, 1, 0, 0, 0, 384,
		382, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 387, 1, 0, 0, 0, 386, 384,
		1, 0, 0, 0, 387, 390, 5, 6, 0, 0, 388, 390, 3, 16, 8, 0, 389, 380, 1, 0,
		0, 0, 389, 388, 1, 0, 0, 0, 390, 47, 1, 0, 0, 0, 391, 401, 5, 35, 0, 0,
		392, 396, 5, 5, 0, 0, 393, 395, 3, 16, 8, 0, 394, 393, 1, 0, 0, 0, 395,
		398, 1, 0, 0, 0, 396, 394, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 399,
		1, 0, 0, 0, 398, 396, 1, 0, 0, 0, 399, 402, 5, 6, 0, 0, 400, 402, 3, 16,
		8, 0, 401, 392, 1, 0, 0, 0, 401, 400, 1, 0, 0, 0, 402, 49, 1, 0, 0, 0,
		403, 408, 3, 18, 9, 0, 404, 405, 5, 8, 0, 0, 405, 407, 3, 18, 9, 0, 406,
		404, 1, 0, 0, 0, 407, 410, 1, 0, 0, 0, 408, 406, 1, 0, 0, 0, 408, 409,
		1, 0, 0, 0, 409, 412, 1, 0, 0, 0, 410, 408, 1, 0, 0, 0, 411, 403, 1, 0,
		0, 0, 411, 412, 1, 0, 0, 0, 412, 51, 1, 0, 0, 0, 413, 418, 5, 50, 0, 0,
		414, 415, 5, 27, 0, 0, 415, 417, 5, 50, 0, 0, 416, 414, 1, 0, 0, 0, 417,
		420, 1, 0, 0, 0, 418, 416, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 421,
		1, 0, 0, 0, 420, 418, 1, 0, 0, 0, 421, 422, 5, 3, 0, 0, 422, 423, 3, 50,
		25, 0, 423, 424, 5, 4, 0, 0, 424, 53, 1, 0, 0, 0, 425, 426, 7, 2, 0, 0,
		426, 55, 1, 0, 0, 0, 427, 428, 7, 3, 0, 0, 428, 57, 1, 0, 0, 0, 429, 430,
		7, 4, 0, 0, 430, 59, 1, 0, 0, 0, 431, 432, 7, 5, 0, 0, 432, 61, 1, 0, 0,
		0, 34, 70, 72, 89, 100, 110, 119, 122, 150, 155, 160, 164, 168, 175, 216,
		220, 228, 251, 259, 294, 296, 306, 319, 333, 341, 348, 351, 360, 384, 389,
		396, 401, 408, 411, 418,
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
	CParserINT               = 48
	CParserSTRING            = 49
	CParserID                = 50
	CParserWS                = 51
	CParserSINGLE_COMMENT    = 52
	CParserMULTILINE_COMMENT = 53
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
	CParserRULE_constant     = 10
	CParserRULE_assign_expr  = 11
	CParserRULE_const        = 12
	CParserRULE_asdexpr      = 13
	CParserRULE_asd          = 14
	CParserRULE_case         = 15
	CParserRULE_default      = 16
	CParserRULE_forInit      = 17
	CParserRULE_forCondition = 18
	CParserRULE_forIter      = 19
	CParserRULE_global       = 20
	CParserRULE_declaration  = 21
	CParserRULE_definition   = 22
	CParserRULE_elseif       = 23
	CParserRULE_else         = 24
	CParserRULE_args         = 25
	CParserRULE_call_expr    = 26
	CParserRULE_muldiv       = 27
	CParserRULE_addsub       = 28
	CParserRULE_eqneq        = 29
	CParserRULE_incdec       = 30
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
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1125900443713668) != 0 {
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(62)
				p.Include()
			}

		case 2:
			{
				p.SetState(63)
				p.Global()
			}
			{
				p.SetState(64)
				p.Match(CParserT__0)
			}

		case 3:
			{
				p.SetState(66)
				p.Function()
			}

		case 4:
			{
				p.SetState(67)
				p.Struct_()
			}
			{
				p.SetState(68)
				p.Match(CParserT__0)
			}

		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(75)
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
		p.SetState(77)
		p.Match(CParserT__1)
	}
	{
		p.SetState(78)
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
		p.SetState(80)
		p.Type_()
	}
	{
		p.SetState(81)
		p.Match(CParserID)
	}
	{
		p.SetState(82)
		p.Match(CParserT__2)
	}
	{
		p.SetState(83)
		p.Params()
	}
	{
		p.SetState(84)
		p.Match(CParserT__3)
	}
	{
		p.SetState(85)
		p.Match(CParserT__4)
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1125900444233248) != 0 {
		{
			p.SetState(86)
			p.Stmt()
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(92)
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
		p.SetState(94)
		p.Match(CParserT__6)
	}
	{
		p.SetState(95)
		p.Match(CParserID)
	}
	{
		p.SetState(96)
		p.Match(CParserT__4)
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserT__28 || _la == CParserID {
		{
			p.SetState(97)
			p.Field()
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(103)
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
		p.SetState(105)
		p.Type_()
	}
	{
		p.SetState(106)
		p.Match(CParserID)
	}
	{
		p.SetState(107)
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
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserT__28 {
		{
			p.SetState(109)
			p.Const_()
		}

	}
	{
		p.SetState(112)
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
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserT__28 || _la == CParserID {
		{
			p.SetState(114)
			p.Param()
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CParserT__7 {
			{
				p.SetState(115)
				p.Match(CParserT__7)
			}
			{
				p.SetState(116)
				p.Param()
			}

			p.SetState(121)
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
		p.SetState(124)
		p.Type_()
	}
	{
		p.SetState(125)
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

	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclarationStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Declaration()
		}
		{
			p.SetState(128)
			p.Match(CParserT__0)
		}

	case 2:
		localctx = NewDefinitionStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.Definition()
		}
		{
			p.SetState(131)
			p.Match(CParserT__0)
		}

	case 3:
		localctx = NewAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(133)
			p.Assign_expr()
		}
		{
			p.SetState(134)
			p.Match(CParserT__0)
		}

	case 4:
		localctx = NewAssignSumDiffStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(136)
			p.Asdexpr()
		}
		{
			p.SetState(137)
			p.Match(CParserT__0)
		}

	case 5:
		localctx = NewCallStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(139)
			p.Call_expr()
		}
		{
			p.SetState(140)
			p.Match(CParserT__0)
		}

	case 6:
		localctx = NewIfStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(142)
			p.Match(CParserT__8)
		}
		{
			p.SetState(143)
			p.Match(CParserT__2)
		}
		{
			p.SetState(144)
			p.expr(0)
		}
		{
			p.SetState(145)
			p.Match(CParserT__3)
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(146)
				p.Match(CParserT__4)
			}
			p.SetState(150)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1125900444233248) != 0 {
				{
					p.SetState(147)
					p.Stmt()
				}

				p.SetState(152)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(153)
				p.Match(CParserT__5)
			}

		case 2:
			{
				p.SetState(154)
				p.Stmt()
			}

		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(157)
					p.Elseif()
				}

			}
			p.SetState(162)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(163)
				p.Else_()
			}

		}

	case 7:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(166)
			p.Match(CParserT__9)
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2181431071080456) != 0 {
			{
				p.SetState(167)
				p.expr(0)
			}

		}
		{
			p.SetState(170)
			p.Match(CParserT__0)
		}

	case 8:
		localctx = NewBlockStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(171)
			p.Match(CParserT__4)
		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1125900444233248) != 0 {
			{
				p.SetState(172)
				p.Stmt()
			}

			p.SetState(177)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(178)
			p.Match(CParserT__5)
		}

	case 9:
		localctx = NewGotoStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(179)
			p.Match(CParserT__10)
		}
		{
			p.SetState(180)
			p.Match(CParserID)
		}
		{
			p.SetState(181)
			p.Match(CParserT__0)
		}

	case 10:
		localctx = NewLabelStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(182)
			p.Match(CParserID)
		}
		{
			p.SetState(183)
			p.Match(CParserT__11)
		}

	case 11:
		localctx = NewForStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(184)
			p.Match(CParserT__12)
		}
		{
			p.SetState(185)
			p.Match(CParserT__2)
		}
		{
			p.SetState(186)
			p.ForInit()
		}
		{
			p.SetState(187)
			p.Match(CParserT__0)
		}
		{
			p.SetState(188)
			p.ForCondition()
		}
		{
			p.SetState(189)
			p.Match(CParserT__0)
		}
		{
			p.SetState(190)
			p.ForIter()
		}
		{
			p.SetState(191)
			p.Match(CParserT__3)
		}
		{
			p.SetState(192)
			p.Stmt()
		}

	case 12:
		localctx = NewWhileStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(194)
			p.Match(CParserT__13)
		}
		{
			p.SetState(195)
			p.Match(CParserT__2)
		}
		{
			p.SetState(196)
			p.expr(0)
		}
		{
			p.SetState(197)
			p.Match(CParserT__3)
		}
		{
			p.SetState(198)
			p.Stmt()
		}

	case 13:
		localctx = NewDoWhileStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(200)
			p.Match(CParserT__14)
		}
		{
			p.SetState(201)
			p.Stmt()
		}
		{
			p.SetState(202)
			p.Match(CParserT__13)
		}
		{
			p.SetState(203)
			p.Match(CParserT__2)
		}
		{
			p.SetState(204)
			p.expr(0)
		}
		{
			p.SetState(205)
			p.Match(CParserT__3)
		}
		{
			p.SetState(206)
			p.Match(CParserT__0)
		}

	case 14:
		localctx = NewSwitchStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(208)
			p.Match(CParserT__15)
		}
		{
			p.SetState(209)
			p.Match(CParserT__2)
		}
		{
			p.SetState(210)
			p.expr(0)
		}
		{
			p.SetState(211)
			p.Match(CParserT__3)
		}
		{
			p.SetState(212)
			p.Match(CParserT__4)
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CParserT__31 {
			{
				p.SetState(213)
				p.Case_()
			}

			p.SetState(218)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CParserT__32 {
			{
				p.SetState(219)
				p.Default_()
			}

		}
		{
			p.SetState(222)
			p.Match(CParserT__5)
		}

	case 15:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(224)
			p.Match(CParserT__16)
		}
		{
			p.SetState(225)
			p.Match(CParserT__0)
		}

	case 16:
		localctx = NewContinueStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(226)
			p.Match(CParserT__17)
		}
		{
			p.SetState(227)
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
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPostIncDecExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(231)
			p.Match(CParserID)
		}
		{
			p.SetState(232)
			p.Incdec()
		}

	case 2:
		localctx = NewPreIncDecExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(233)
			p.Incdec()
		}
		{
			p.SetState(234)
			p.Match(CParserID)
		}

	case 3:
		localctx = NewMinusExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(236)
			p.Match(CParserT__18)
		}
		{
			p.SetState(237)
			p.expr(18)
		}

	case 4:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(238)
			p.Match(CParserT__19)
		}
		{
			p.SetState(239)
			p.expr(17)
		}

	case 5:
		localctx = NewAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(240)
			p.Assign_expr()
		}

	case 6:
		localctx = NewAssignSumDiffExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(241)
			p.Asdexpr()
		}

	case 7:
		localctx = NewVariableExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(242)
			p.Match(CParserID)
		}

	case 8:
		localctx = NewConstantExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(243)
			p.Constant()
		}

	case 9:
		localctx = NewMemberExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(244)
			p.Match(CParserID)
		}
		{
			p.SetState(245)
			p.Match(CParserT__26)
		}
		{
			p.SetState(246)
			p.Match(CParserID)
		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(247)
					p.Match(CParserT__26)
				}
				{
					p.SetState(248)
					p.Match(CParserID)
				}

			}
			p.SetState(253)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
		}

	case 10:
		localctx = NewCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(254)
			p.Call_expr()
		}

	case 11:
		localctx = NewGroupExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(255)
			p.Match(CParserT__2)
		}
		{
			p.SetState(256)
			p.expr(0)
		}
		{
			p.SetState(257)
			p.Match(CParserT__3)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(294)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(261)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(262)
					p.Muldiv()
				}
				{
					p.SetState(263)
					p.expr(17)
				}

			case 2:
				localctx = NewAddSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(265)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(266)
					p.Addsub()
				}
				{
					p.SetState(267)
					p.expr(16)
				}

			case 3:
				localctx = NewEqNeqExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(269)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(270)
					p.Eqneq()
				}
				{
					p.SetState(271)
					p.expr(15)
				}

			case 4:
				localctx = NewBitAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(273)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(274)
					p.Match(CParserT__20)
				}
				{
					p.SetState(275)
					p.expr(14)
				}

			case 5:
				localctx = NewBitXorExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(276)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(277)
					p.Match(CParserT__21)
				}
				{
					p.SetState(278)
					p.expr(13)
				}

			case 6:
				localctx = NewBitOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(279)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(280)
					p.Match(CParserT__22)
				}
				{
					p.SetState(281)
					p.expr(12)
				}

			case 7:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AndExprContext).l = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(282)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(283)
					p.Match(CParserT__23)
				}
				{
					p.SetState(284)

					var _x = p.expr(11)

					localctx.(*AndExprContext).r = _x
				}

			case 8:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OrExprContext).l = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(285)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(286)
					p.Match(CParserT__24)
				}
				{
					p.SetState(287)

					var _x = p.expr(10)

					localctx.(*OrExprContext).r = _x
				}

			case 9:
				localctx = NewConditionalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ConditionalExprContext).condition = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expr)
				p.SetState(288)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(289)
					p.Match(CParserT__25)
				}
				{
					p.SetState(290)

					var _x = p.expr(0)

					localctx.(*ConditionalExprContext).true_ = _x
				}
				{
					p.SetState(291)
					p.Match(CParserT__11)
				}
				{
					p.SetState(292)

					var _x = p.expr(7)

					localctx.(*ConditionalExprContext).false_ = _x
				}

			}

		}
		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 20, CParserRULE_constant)
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
		p.SetState(299)
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
	p.EnterRule(localctx, 22, CParserRULE_assign_expr)
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
		p.SetState(301)
		p.Match(CParserID)
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserT__26 {
		{
			p.SetState(302)
			p.Match(CParserT__26)
		}
		{
			p.SetState(303)
			p.Match(CParserID)
		}

		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(309)
		p.Match(CParserT__27)
	}
	{
		p.SetState(310)
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
	p.EnterRule(localctx, 24, CParserRULE_const)

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
		p.SetState(312)
		p.Match(CParserT__28)
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
	p.EnterRule(localctx, 26, CParserRULE_asdexpr)
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
		p.SetState(314)
		p.Match(CParserID)
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserT__26 {
		{
			p.SetState(315)
			p.Match(CParserT__26)
		}
		{
			p.SetState(316)
			p.Match(CParserID)
		}

		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(322)
		p.Asd()
	}
	{
		p.SetState(323)
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
	p.EnterRule(localctx, 28, CParserRULE_asd)
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
		p.SetState(325)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CParserT__29 || _la == CParserT__30) {
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
	p.EnterRule(localctx, 30, CParserRULE_case)
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
		p.SetState(327)
		p.Match(CParserT__31)
	}
	{
		p.SetState(328)
		p.expr(0)
	}
	{
		p.SetState(329)
		p.Match(CParserT__11)
	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1125900444233248) != 0 {
		{
			p.SetState(330)
			p.Stmt()
		}

		p.SetState(335)
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
	p.EnterRule(localctx, 32, CParserRULE_default)
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
		p.SetState(336)
		p.Match(CParserT__32)
	}
	{
		p.SetState(337)
		p.Match(CParserT__11)
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1125900444233248) != 0 {
		{
			p.SetState(338)
			p.Stmt()
		}

		p.SetState(343)
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
	p.EnterRule(localctx, 34, CParserRULE_forInit)
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

	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(344)
			p.Definition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2181431071080456) != 0 {
			{
				p.SetState(345)
				p.expr(0)
			}

			p.SetState(350)
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
	p.EnterRule(localctx, 36, CParserRULE_forCondition)

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
		p.SetState(353)
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
	p.EnterRule(localctx, 38, CParserRULE_forIter)
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
		p.SetState(355)
		p.expr(0)
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserT__7 {
		{
			p.SetState(356)
			p.Match(CParserT__7)
		}
		{
			p.SetState(357)
			p.expr(0)
		}

		p.SetState(362)
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
	p.EnterRule(localctx, 40, CParserRULE_global)

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
		p.SetState(363)
		p.Type_()
	}
	{
		p.SetState(364)
		p.Match(CParserID)
	}
	{
		p.SetState(365)
		p.Match(CParserT__27)
	}
	{
		p.SetState(366)
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
	p.EnterRule(localctx, 42, CParserRULE_declaration)

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
		p.SetState(368)
		p.Type_()
	}
	{
		p.SetState(369)
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
	p.EnterRule(localctx, 44, CParserRULE_definition)

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
		p.SetState(371)
		p.Type_()
	}
	{
		p.SetState(372)
		p.Match(CParserID)
	}
	{
		p.SetState(373)
		p.Match(CParserT__27)
	}
	{
		p.SetState(374)
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
	p.EnterRule(localctx, 46, CParserRULE_elseif)
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
		p.SetState(376)
		p.Match(CParserT__33)
	}
	{
		p.SetState(377)
		p.Match(CParserT__2)
	}
	{
		p.SetState(378)
		p.expr(0)
	}
	{
		p.SetState(379)
		p.Match(CParserT__3)
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(380)
			p.Match(CParserT__4)
		}
		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1125900444233248) != 0 {
			{
				p.SetState(381)
				p.Stmt()
			}

			p.SetState(386)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(387)
			p.Match(CParserT__5)
		}

	case 2:
		{
			p.SetState(388)
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
	p.EnterRule(localctx, 48, CParserRULE_else)
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
		p.SetState(391)
		p.Match(CParserT__34)
	}
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(392)
			p.Match(CParserT__4)
		}
		p.SetState(396)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1125900444233248) != 0 {
			{
				p.SetState(393)
				p.Stmt()
			}

			p.SetState(398)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(399)
			p.Match(CParserT__5)
		}

	case 2:
		{
			p.SetState(400)
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
	p.EnterRule(localctx, 50, CParserRULE_args)
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
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2181431071080456) != 0 {
		{
			p.SetState(403)
			p.expr(0)
		}
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CParserT__7 {
			{
				p.SetState(404)
				p.Match(CParserT__7)
			}
			{
				p.SetState(405)
				p.expr(0)
			}

			p.SetState(410)
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
	p.EnterRule(localctx, 52, CParserRULE_call_expr)
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
		p.SetState(413)
		p.Match(CParserID)
	}
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserT__26 {
		{
			p.SetState(414)
			p.Match(CParserT__26)
		}
		{
			p.SetState(415)
			p.Match(CParserID)
		}

		p.SetState(420)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(421)
		p.Match(CParserT__2)
	}
	{
		p.SetState(422)
		p.Args()
	}
	{
		p.SetState(423)
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
	p.EnterRule(localctx, 54, CParserRULE_muldiv)
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
		p.SetState(425)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&481036337152) != 0) {
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
	p.EnterRule(localctx, 56, CParserRULE_addsub)
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
		p.SetState(427)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CParserT__18 || _la == CParserT__38) {
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
	p.EnterRule(localctx, 58, CParserRULE_eqneq)
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
		p.SetState(429)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&69269232549888) != 0) {
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
	p.EnterRule(localctx, 60, CParserRULE_incdec)
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
		p.SetState(431)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CParserT__45 || _la == CParserT__46) {
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

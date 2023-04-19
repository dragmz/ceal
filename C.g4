grammar C;

program:
    (
        include
        | global ';'
        | function
        | struct ';'
    )* EOF;

include: '#include' STRING;
function: type ID '(' params ')' '{' stmt* '}';

struct: 'struct' ID '{' field* '}';
field: type ID ';';

type: const? ID;
params: (param (',' param)*)?;
param: type ID;
stmt:
    declaration ';'                                                 # DeclarationStmt
    | definition ';'                                                # DefinitionStmt
    | assign_expr ';'                                               # AssignStmt
    | asdexpr ';'                                                   # AssignSumDiffStmt
    | 'asm' '(' STRING* ')' ';'                                     # AsmStmt
    | call_expr ';'                                                 # CallStmt
    | 'if' '(' expr ')' (('{' stmt* '}') | stmt) elseif* else?      # IfStmt
    | 'return' expr? ';'                                            # ReturnStmt
    | '{' stmt* '}'                                                 # BlockStmt
    | 'goto' ID ';'                                                 # GotoStmt
    | ID ':'                                                        # LabelStmt
    | 'for' '(' forInit ';' forCondition ';' forIter ')' stmt       # ForStmt
    | 'while' '(' expr ')' stmt                                     # WhileStmt
    | 'do' stmt 'while' '(' expr ')' ';'                            # DoWhileStmt
    | 'switch' '(' expr ')' '{' case* default? '}'                  # SwitchStmt
    | 'break' ';'                                                   # BreakStmt
    | 'continue' ';'                                                # ContinueStmt
    | (SINGLE_COMMENT | MULTILINE_COMMENT)                          # CommentStmt
    ;

expr:
    ID incdec                   # PostIncDecExpr
    | incdec ID                 # PreIncDecExpr
    | '-' expr                  # MinusExpr
    | '!' expr                  # NotExpr
    | expr muldiv expr          # MulDivExpr
    | expr addsub expr          # AddSubExpr
    | expr eqneq expr           # EqNeqExpr
    | expr '&' expr             # BitAndExpr
    | expr '^' expr             # BitXorExpr
    | expr '|' expr             # BitOrExpr
    | l=expr '&&' r=expr        # AndExpr
    | l=expr '||' r=expr        # OrExpr
    | assign_expr               # AssignExpr
    | asdexpr                   # AssignSumDiffExpr
    | condition=expr '?' true=expr ':' false=expr    # ConditionalExpr
    | ID                        # VariableExpr
    | constant                  # ConstantExpr
    | ID '[' expr ']'           # SubscriptExpr
    | ID '.' ID ('.' ID)*       # MemberExpr
    | call_expr                 # CallExpr
    | '(' expr ')'              # GroupExpr
    ;

constant: (INT | STRING);
assign_expr: ID ('.' ID)* '=' expr;
const: 'const';
asdexpr: ID ('.' ID)* asd expr;
asd: '+=' | '-=';
case: 'case' expr ':' stmt*;
default: 'default' ':' stmt*;

forInit:
    definition
    | expr*
    ;

forCondition:
    expr
    ;

forIter:
    expr (',' expr)*
    ;

global: type ID '=' constant;
declaration: type ID;
definition: type ID '=' expr;
elseif: 'else if' '(' expr ')' (('{' stmt* '}') | stmt);
else: 'else' (('{' stmt* '}') | stmt);
args: (expr (',' expr)*)?;
call_expr: ID ('.' ID)* '(' args ')';
muldiv: '*' | '/' | '%';
addsub: ('+' | '-');
eqneq: ('==' | '!=' | '<' | '>' | '<=' | '>=');
incdec: ('++' | '--');

INT: [0-9]+;
STRING: '"' ( ~["\\] | '\\' . )* '"' ;
ID: [a-zA-Z_][a-zA-Z0-9_]*;
WS: [ \t\r\n]+ -> skip;
SINGLE_COMMENT: '//' .*? '\n';
MULTILINE_COMMENT: '/*' .*? '*/';

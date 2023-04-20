grammar C;

program:
    (
        include
        | comment
        | global ';'
        | function
        | struct ';'
    )* EOF;

stmt:
    declaration ';'                                                 # DeclarationStmt
    | definition ';'                                                # DefinitionStmt
    | pre_incdec_expr ';'                                           # PreIncDecStmt
    | post_incdec_expr ';'                                          # PostIncDecStmt
    | assign_expr ';'                                               # AssignStmt
    | asd_expr ';'                                                  # AssignSumDiffStmt
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
    | comment                                                       # CommentStmt
    ;

expr:
    post_incdec_expr            # PostIncDecExpr
    | pre_incdec_expr           # PreIncDecExpr
    | '-' expr                  # MinusExpr
    | '!' expr                  # NotExpr
    | l=expr muldiv r=expr      # MulDivExpr
    | l=expr addsub r=expr      # AddSubExpr
    | l=expr eqneq r=expr       # EqNeqExpr
    | l=expr '&' r=expr         # BitAndExpr
    | l=expr '^' r=expr         # BitXorExpr
    | l=expr '|' r=expr         # BitOrExpr
    | l=expr '&&' r=expr        # AndExpr
    | l=expr '||' r=expr        # OrExpr
    | assign_expr               # AssignExpr
    | asd_expr                  # AssignSumDiffExpr
    | condition=expr '?' true=expr ':' false=expr # ConditionalExpr
    | constant                  # ConstantExpr
    | subscript_expr            # SubscriptExpr
    | dot_expr                  # DotExpr
    | call_expr                 # CallExpr
    | '(' expr ')'              # GroupExpr
    ;

include: '#include' STRING;
function: type ID '(' params ')' '{' stmt* '}';
struct: 'struct' ID '{' field* '}';
global: type ID '=' constant;
field: type ID ';';
type: const? ID;
params: (param (',' param)*)?;
param: type ID;
value_access_expr: (dot_expr | subscript_expr);
pre_incdec_expr: incdec value_access_expr;
post_incdec_expr: value_access_expr incdec;
dot_expr: ID ('.' ID)*;
subscript_expr: dot_expr '[' expr ']';
comment: (SINGLE_COMMENT | MULTILINE_COMMENT);
constant: (INT | STRING);
assign_expr: value_access_expr '=' expr;
const: 'const';
asd_expr: value_access_expr asd expr;
asd: '+=' | '-=';
case: 'case' expr ':' stmt*;
default: 'default' ':' stmt*;
forInit: definition | expr*;
forCondition: expr ;
forIter: expr (',' expr)*;
declaration: type ID;
definition: type ID '=' expr;
elseif: 'else if' '(' expr ')' (('{' stmt* '}') | stmt);
else: 'else' (('{' stmt* '}') | stmt);
args: (expr (',' expr)*)?;
call_expr: value_access_expr '(' args ')';
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

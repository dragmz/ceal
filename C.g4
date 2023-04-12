grammar C;

program:
    (
        include
        | function
        | struct ';'
    )* EOF;

include: '#include' STRING;
function: type ID '(' params ')' '{' stmt* '}';

struct: 'struct' ID '{' field* '}';
field: type ID ';';

type: 'const'? ID;
params: (param (',' param)*)?;
param: type ID;
stmt:
    declaration ';'                                                 # DeclarationStmt
    | definition ';'                                                # DefinitionStmt
    | ID ('.' ID)* '=' expr ';'                                     # AssignmentStmt
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
    | expr '&&' expr            # AndExpr
    | expr '||' expr            # OrExpr
    | ID                        # VariableExpr
    | (INT | STRING)            # ConstantExpr
    | ID '.' ID ('.' ID)*       # MemberExpr
    | call_expr                 # CallExpr
    ;

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
SINGLE_COMMENT: '//' .*? '\n' -> skip;
MULTILINE_COMMENT: '/*' .*? '*/' -> skip;

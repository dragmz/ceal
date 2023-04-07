grammar C;

program:
    (
        include
        | function
        | struct
    )* EOF;

include: '#include' STRING;
function: type ID '(' params ')' '{' stmt* '}';

struct: 'struct' ID '{' field* '}' ';';
field: type ID ';';

type: 'const'? ID;
params: (param (',' param)*)?;
param: type ID;
stmt:
    type ID ';'                                         # Declaration
    | type ID '=' expr ';'                              # Definition
    | ID ('.' ID)* '=' expr ';'                         # Assignment
    | call_expr ';'                                     # Call
    | 'if' '(' expr ')' (('{' stmt* '}') | stmt) elseif* else?        # If
    | 'return' expr? ';'                                 # Return
    | '{' stmt* '}'                                     # Block
    ;
expr:
    expr muldiv expr            # MulDivExpr
    | expr addsub expr          # AddSubExpr
    | '-' expr                  # MinusExpr
    | expr eqneq expr           # EqNeqExpr
    | ID                        # VariableExpr
    | (INT | STRING)            # ConstantExpr
    | ID '.' ID ('.' ID)*       # MemberExpr
    | call_expr                 # CallExpr
    ;
elseif: 'else if' '(' expr ')' (('{' stmt* '}') | stmt);
else: 'else' (('{' stmt* '}') | stmt);
args: (expr (',' expr)*)?;
call_expr: ID ('.' ID)* '(' args ')';
muldiv: '*' | '/';
addsub: ('+' | '-');
eqneq: ('==' | '!=');

INT: [0-9]+;
STRING: '"' ( ~["\\] | '\\' . )* '"' ;
ID: [a-zA-Z_][a-zA-Z0-9_]*;
WS: [ \t\r\n]+ -> skip;
SINGLE_COMMENT: '//' .*? '\n' -> skip;
MULTILINE_COMMENT: '/*' .*? '*/' -> skip;

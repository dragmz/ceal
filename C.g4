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
    ;
expr:
    expr '||' expr              # OrExpr
    | expr '&&' expr            # AndExpr
    | expr muldiv expr          # MulDivExpr
    | expr addsub expr          # AddSubExpr
    | '-' expr                  # MinusExpr
    | expr eqneq expr           # EqNeqExpr
    | '!' expr                  # NotExpr
    | ID                        # VariableExpr
    | (INT | STRING)            # ConstantExpr
    | ID '.' ID ('.' ID)*       # MemberExpr
    | call_expr                 # CallExpr
    ;

declaration: type ID;
definition: type ID '=' expr;
elseif: 'else if' '(' expr ')' (('{' stmt* '}') | stmt);
else: 'else' (('{' stmt* '}') | stmt);
args: (expr (',' expr)*)?;
call_expr: ID ('.' ID)* '(' args ')';
muldiv: '*' | '/';
addsub: ('+' | '-');
eqneq: ('==' | '!=' | '<' | '>' | '<=' | '>=');

INT: [0-9]+;
STRING: '"' ( ~["\\] | '\\' . )* '"' ;
ID: [a-zA-Z_][a-zA-Z0-9_]*;
WS: [ \t\r\n]+ -> skip;
SINGLE_COMMENT: '//' .*? '\n' -> skip;
MULTILINE_COMMENT: '/*' .*? '*/' -> skip;

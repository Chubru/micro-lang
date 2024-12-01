grammar MicroLang;

prog:   stat* EOF;

stat:   ID ASSIGN expr ';'         # assignStat
    |   expr ';'                   # printExpr
    |   'while' '(' expr ')' '{' stat* '}' # whileStat
    |   'if' '(' expr ')' '{' stat* '}'    # ifStat
    ;

expr:   expr op=('>' | '<' | '==') expr
    |   expr op=('*' | '/') expr
    |   expr op=('+' | '-') expr
    |   INT
    |   ID
    |   '(' expr ')'
    ;

MULT  : '*' ;
DIV   : '/' ;
PLUS  : '+' ;
MINUS : '-' ;
LPAREN: '(' ;
RPAREN: ')' ;
ASSIGN: '=' ;
GREAT : '>' ;
LESS  : '<' ;
EQUAL : '==';

ID  : [a-zA-Z_][a-zA-Z0-9_]* ;
INT : [0-9]+ ;
WS  : [ \t\r\n]+ -> skip ;
COMMENT : '/*' .*? '*/' -> skip ;
LINE_COMMENT : '//' ~[\r\n]* -> skip ;
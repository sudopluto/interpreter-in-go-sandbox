package token

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // identifiers
    IDENT = "IDENT"
    INT = "INT"

    // operators
    ASSIGN = "="
    PLUS = "+"
    MINUS = "-"
    MULT = "*"
    DIV = "/"
    BANG = "!"

    LT = "<"
    GT = ">"
    EQ = "=="
    NEQ = "!="

    // delims
    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // keywords
    FUNCTION = "FUNCTION"
    LET =      "LET"
    TRUE =     "TRUE"
    FALSE =    "FALSE"
    IF =       "IF"
    ELSE =     "ELSE"
    RETURN =   "RETURN"
)

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

var keywords = map[string]TokenType{
    "fn":     FUNCTION,
    "let":    LET,
    "true":   TRUE,
    "false":  FALSE,
    "if":     IF,
    "else":   ELSE,
    "return": RETURN,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}

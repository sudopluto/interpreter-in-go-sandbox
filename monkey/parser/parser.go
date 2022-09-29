package parser

import (
    "monkey/ast"
    "monkey/lexer"
    "monkey/token"
)

type Parser struct {
    l *lexer.Lexer

    curToken  token.Token
    peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l}

    // read inital two tokens (if exist) so cur and peek tokens are set
    p.nextToken()
    p.nextToken()

    return p
}

func (p *Parser) nextToken {
    p.curToken = p.nextToken
    p.nextToken = p.l.NextToken()
}

func (p *Parser) ParseProgram *ast.Program {
    return nil
}

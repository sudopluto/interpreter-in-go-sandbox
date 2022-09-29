package parser

import (
    "testing"
    "monkey/ast"
    "moneky/lexer"
)

func TestLetStatements(t *testing.T) {
    input := `
    let x = 5;
    let y = 10;
    let foobar = 1234;
    `

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()

    if program == nil {
        t.Fatalf("ParseProgram() returned nil")
    }
}

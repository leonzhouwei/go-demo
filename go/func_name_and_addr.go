package main

import (
	"fmt"
)

type stateFn func() string

func lexFnAlice() string {
	return "alice"
}

func lexFnBob() string {
	return "bob"
}

type lexer struct {
	stateFnName string
	stateFn stateFn
}

func (e *lexer) setStateFnByName() {
	switch {
	case e.stateFnName == "lexFnAlice":
		e.stateFn = lexFnAlice
	case e.stateFnName == "lexFnBob":
		e.stateFn = lexFnBob
	default:
		e.stateFn = nil
	}
}

func main() {
	fmt.Println("lexFnAlice() addr:", lexFnAlice)
	fmt.Println("lexFnBob() addr:", lexFnBob)

	lexerAlice := lexer{
		stateFnName: "lexFnAlice",
		stateFn: lexFnAlice,
	}
	lexerBob := lexer{
		stateFnName: "lexFnBob",
		stateFn: lexFnBob,
	}
	lexerBob2 := lexer{
		stateFnName: "lexFnBob",
	}
	lexerBob2.setStateFnByName()

	fmt.Println("alice:", lexerAlice)
	fmt.Println("  bob:", lexerBob)
	fmt.Println(" bob2:", lexerBob)
}

package main

import (
	"fmt"
	"strconv"
)

type Interpreter struct {
	stack     Stack
	dictStack *DictStack
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		dictStack: NewDictStack(),
	}
}

func (interp *Interpreter) SetLexical(on bool) {
	interp.dictStack.SetLexical(on)
}

func (interp *Interpreter) Execute(tokens []Token) error {
	i := 0
	for i < len(tokens) {
		tok := tokens[i]

		switch tok.Type {

		case TokInt:
			val, _ := strconv.Atoi(tok.Value)
			interp.stack.Push(PSObject{Type: TypeInt, IVal: val})

		case TokFloat:
			val, _ := strconv.ParseFloat(tok.Value, 64)
			interp.stack.Push(PSObject{Type: TypeFloat, FVal: val})

		case TokBool:
			interp.stack.Push(PSObject{Type: TypeBool, BVal: tok.Value == "true"})

		case TokString:
			interp.stack.Push(PSObject{Type: TypeString, SVal: tok.Value})

		case TokName:
			interp.stack.Push(PSObject{Type: TypeName, SVal: tok.Value})

		case TokProcStart:
			procTokens, consumed, err := collectProc(tokens, i+1)
			if err != nil {
				return err
			}
			interp.stack.Push(procTokens)
			i += consumed + 1 // Skip past the procedure tokens and the closing '}'

		case TokOperator:
			obj := tokenToObject(Token{Type: TokOperator, Value: tok.Value})
			if err := interp.evalObject(obj); err != nil {
				return err
			}
		}

		i++
	}
	return nil
}

func (interp *Interpreter) ExecuteProc(proc PSObject) error {
	return interp.ExecuteObjects(proc.AVal)
}

func collectProc(tokens []Token, start int) (PSObject, int, error) {
	var body []PSObject
	depth := 1
	i := start

	for i < len(tokens) {
		tok := tokens[i]

		if tok.Type == TokProcStart {
			nested, consumed, err := collectProc(tokens, i+1)
			if err != nil {
				return PSObject{}, 0, err
			}
			body = append(body, nested)
			i += consumed + 1
			depth++
		} else if tok.Type == TokProcEnd {
			depth--
			if depth == 0 {
				return PSObject{Type: TypeArray, AVal: body}, i - start, nil
			}
		} else {
			body = append(body, tokenToObject(tok))
		}
		i++
	}
	return PSObject{}, 0, fmt.Errorf("unmatched '{'")
}

func tokenToObject(tok Token) PSObject {
	switch tok.Type {
	case TokInt:
		n, _ := strconv.Atoi(tok.Value)
		return PSObject{Type: TypeInt, IVal: n}
	case TokFloat:
		f, _ := strconv.ParseFloat(tok.Value, 64)
		return PSObject{Type: TypeFloat, FVal: f}
	case TokBool:
		return PSObject{Type: TypeBool, BVal: tok.Value == "true"}
	case TokString:
		return PSObject{Type: TypeString, SVal: tok.Value}
	case TokName:
		return PSObject{Type: TypeName, SVal: tok.Value}
	default:
		return PSObject{Type: TypeName, SVal: tok.Value}
	}

}

// ExecuteObjects runs a procedure body stored as []PSObject
func (interp *Interpreter) ExecuteObjects(objects []PSObject) error {
	for _, obj := range objects {
		if err := interp.evalObject(obj); err != nil {
			return err
		}
	}
	return nil
}

// evalObject evaluates a single PSObject — the shared core logic
func (interp *Interpreter) evalObject(obj PSObject) error {
	switch obj.Type {
	case TypeInt, TypeFloat, TypeBool, TypeString:
		interp.stack.Push(obj)

	case TypeArray:
		interp.stack.Push(obj)

	case TypeName:
		if fn, ok := builtins[obj.SVal]; ok {
			if err := fn(interp); err != nil {
				if err == ErrQuit {
					return err
				}
				return fmt.Errorf("error in %s: %w", obj.SVal, err)
			}
		} else {
			found, ok := interp.dictStack.Lookup(obj.SVal)
			if !ok {
				return fmt.Errorf("undefined: %s", obj.SVal)
			}
			if found.Type == TypeArray {
				if interp.dictStack.lexical {
					scope := found.Scope
					if scope == nil {
						// defined before lexical was toggled on
						// fall back to global dict only
						scope = []map[string]PSObject{interp.dictStack.dicts[0]}
					}
					saved := interp.dictStack.dicts
					interp.dictStack.dicts = scope
					err := interp.ExecuteObjects(found.AVal)
					interp.dictStack.dicts = saved
					return err
				}
				return interp.ExecuteObjects(found.AVal)
			} else {
				interp.stack.Push(found)
			}
		}
	}
	return nil
}

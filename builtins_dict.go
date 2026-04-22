package main

import (
	"fmt"
)

// opDict creates new dictionary with given capcity
func opDict(interp *Interpreter) error {
	nObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if nObj.Type != TypeInt {
		return fmt.Errorf("Dict: Expected int, got %v", nObj.Type)
	}
	newDict := make(map[string]PSObject, nObj.IVal)
	interp.stack.Push(PSObject{Type: TypeDict, DVal: newDict})
	return nil
}

// opLength returns # of key-value pairs in dict or # chars in string
func opLength(interp *Interpreter) error {
	obj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	switch obj.Type {
	case TypeDict:
		interp.stack.Push(PSObject{Type: TypeInt, IVal: len(obj.DVal)})
	case TypeString:
		interp.stack.Push(PSObject{Type: TypeInt, IVal: len(obj.SVal)})
	default:
		return fmt.Errorf("Length: Expected dict or string, got %v", obj.Type)
	}
	return nil
}

// opMaxLength returns max capacity of dict
func opMaxLength(interp *Interpreter) error {
	obj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if obj.Type != TypeDict {
		return fmt.Errorf("MaxLength: Expected dict, got %v", obj.Type)
	}
	interp.stack.Push(PSObject{Type: TypeInt, IVal: len(obj.DVal)})
	return nil
}

// opBegin pushes dict onto dictstack, now current scope
func opBegin(interp *Interpreter) error {
	obj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if obj.Type != TypeDict {
		return fmt.Errorf("Begin: Expected dict, got %v", obj.Type)
	}
	interp.dictStack.Begin(obj.DVal)
	return nil
}

// opEnd pops current dictionary off dictstack
func opEnd(interp *Interpreter) error {
	return interp.dictStack.End()
}

// opDef associates key with value in current dict
func opDef(interp *Interpreter) error {
	value, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	key, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if key.Type != TypeName {
		return fmt.Errorf("def: key must be a name, got %v", key.Type)
	}

	if interp.dictStack.lexical && value.Type == TypeArray {
		scope := interp.dictStack.CaptureScope()
		value = PSObject{
			Type:  TypeArray,
			AVal:  value.AVal,
			Scope: scope,
		}
	}

	interp.dictStack.Define(key.SVal, value)
	return nil
}

package main

import (
	"fmt"
)

func opExch(interp *Interpreter) error {
	a, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	b, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	interp.stack.Push(a)
	interp.stack.Push(b)
	return nil
}

func opPop(interp *Interpreter) error {
	_, err := interp.stack.Pop()
	return err
}

func opDup(interp *Interpreter) error {
	top, err := interp.stack.Peek()
	if err != nil {
		return err
	}
	interp.stack.Push(top)
	return nil
}

func opCopy(interp *Interpreter) error {
	nObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if nObj.Type != TypeInt {
		return fmt.Errorf("Copy: Expected int, got %v", nObj.Type)
	}
	n := nObj.IVal
	if n < 0 {
		return fmt.Errorf("Copy: Negative count %d", n)
	}
	if n > interp.stack.Len() {
		return fmt.Errorf("Copy: Not enough elements on stack to copy %d", n)
	}

	items := make([]PSObject, n)
	for i := 0; i < n; i++ {
		obj, err := interp.stack.Index(n - 1 - i)
		if err != nil {
			return err
		}
		items[i] = obj
	}

	for _, obj := range items {
		interp.stack.Push(obj)
	}
	return nil
}

func opClear(interp *Interpreter) error {
	interp.stack.Clear()
	return nil
}

func opCount(interp *Interpreter) error {
	n := interp.stack.Len()
	interp.stack.Push(PSObject{Type: TypeInt, IVal: n})
	return nil
}

package main

import (
	"fmt"
)

// opIf executes procedure if condition is true
func opIf(interp *Interpreter) error {
	procObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	condObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if condObj.Type != TypeBool {
		return fmt.Errorf("If: Expected bool condition, got %v", condObj.Type)
	}
	if condObj.BVal {
		return interp.ExecuteProc(procObj)
	}
	return nil
}

// opIfElse executes first procedure if condition is true, otherwise executes second procedure
func opIfElse(interp *Interpreter) error {
	elseProcObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	ifProcObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	condObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if condObj.Type != TypeBool {
		return fmt.Errorf("IfElse: Expected bool condition, got %v", condObj.Type)
	}
	if ifProcObj.Type != TypeArray || elseProcObj.Type != TypeArray {
		return fmt.Errorf("IfElse: Expected procedures (arrays), got %v and %v", ifProcObj.Type, elseProcObj.Type)
	}
	if condObj.BVal {
		return interp.ExecuteProc(ifProcObj)
	}
	return interp.ExecuteProc(elseProcObj)
}

// opFor executes procedure for range of values
func opFor(interp *Interpreter) error {
	procObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	limitObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	incrObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	initObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if procObj.Type != TypeArray {
		return fmt.Errorf("For: Expected procedure (array), got %v", procObj.Type)
	}
	initial, err := toFloat(initObj)
	if err != nil {
		return fmt.Errorf("for: invalid initial value: %v", err)
	}
	incr, err := toFloat(incrObj)
	if err != nil {
		return fmt.Errorf("for: invalid increment value: %v", err)
	}
	limit, err := toFloat(limitObj)
	if err != nil {
		return fmt.Errorf("for: invalid limit value: %v", err)
	}
	if incr == 0 {
		return fmt.Errorf("for: increment cannot be zero")
	}

	allInts := initObj.Type == TypeInt && incrObj.Type == TypeInt && limitObj.Type == TypeInt

	for i := initial; (incr > 0 && i <= limit) || (incr < 0 && i >= limit); i += incr {
		if allInts {
			interp.stack.Push(PSObject{Type: TypeInt, IVal: int(i)})
		} else {
			interp.stack.Push(PSObject{Type: TypeFloat, FVal: i})
		}
		if err := interp.ExecuteProc(procObj); err != nil {
			return err
		}
	}
	return nil
}

// opRepeat executes procedure n times
func opRepeat(interp *Interpreter) error {
	procObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	nObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if nObj.Type != TypeInt {
		return fmt.Errorf("Repeat: Expected int count, got %v", nObj.Type)
	}
	if procObj.Type != TypeArray {
		return fmt.Errorf("Repeat: Expected procedure (array), got %v", procObj.Type)
	}
	if nObj.IVal < 0 {
		return fmt.Errorf("Repeat: Negative count %d", nObj.IVal)
	}
	for i := 0; i < nObj.IVal; i++ {
		if err := interp.ExecuteProc(procObj); err != nil {
			return err
		}
	}
	return nil
}

var ErrQuit = fmt.Errorf("quit")

func opQuit(interp *Interpreter) error {
	return ErrQuit
}

package main

import (
	"fmt"
)

// compare helper function to compare two PSObjects
// Returns: -1 if a < b, 0 if a == b, 1 if a > b
func compare(a, b PSObject) (int, error) {
	if (a.Type == TypeInt || a.Type == TypeFloat) &&
		(b.Type == TypeInt || b.Type == TypeFloat) {
		aVal, _ := toFloat(a)
		bVal, _ := toFloat(b)
		if aVal < bVal {
			return -1, nil
		}
		if aVal > bVal {
			return 1, nil
		}
		return 0, nil
	}
	if a.Type == TypeString && b.Type == TypeString {
		if a.SVal < b.SVal {
			return -1, nil
		} else if a.SVal > b.SVal {
			return 1, nil
		}
		return 0, nil
	}
	return 0, fmt.Errorf("Cannot compare types %v and %v", a.Type, b.Type)
}

// opEq checks if two values are equal, returns true or false
func opEq(interp *Interpreter) error {
	bObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	aObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	result, err := compare(aObj, bObj)
	if err != nil {
		return err
	}
	interp.stack.Push(PSObject{Type: TypeBool, BVal: result == 0})
	return nil
}

// opNe checks if two values are not equal
func opNe(interp *Interpreter) error {
	bObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	aObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	result, err := compare(aObj, bObj)
	if err != nil {
		return err
	}
	interp.stack.Push(PSObject{Type: TypeBool, BVal: result != 0})
	return nil
}

// opGe checks if a (second) >= b (top)
func opGe(interp *Interpreter) error {
	bObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	aObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	result, err := compare(aObj, bObj)
	if err != nil {
		return err
	}
	interp.stack.Push(PSObject{Type: TypeBool, BVal: result >= 0})
	return nil
}

// opGt checks if a (second) > b (top)
func opGt(interp *Interpreter) error {
	bObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	aObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	result, err := compare(aObj, bObj)
	if err != nil {
		return err
	}
	interp.stack.Push(PSObject{Type: TypeBool, BVal: result > 0})
	return nil
}

// opLe tests if a <= b
func opLe(interp *Interpreter) error {
	bObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	aObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	result, err := compare(aObj, bObj)
	if err != nil {
		return err
	}
	interp.stack.Push(PSObject{Type: TypeBool, BVal: result <= 0})
	return nil
}

// opLt tests if a < b
func opLt(interp *Interpreter) error {
	bObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	aObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	result, err := compare(aObj, bObj)
	if err != nil {
		return err
	}
	interp.stack.Push(PSObject{Type: TypeBool, BVal: result < 0})
	return nil
}

// opAnd logical (bool) or bitwise(int) AND
func opAnd(interp *Interpreter) error {
	bObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	aObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if aObj.Type == TypeBool && bObj.Type == TypeBool {
		interp.stack.Push(PSObject{Type: TypeBool, BVal: aObj.BVal && bObj.BVal})
	} else if aObj.Type == TypeInt && bObj.Type == TypeInt {
		interp.stack.Push(PSObject{Type: TypeInt, IVal: aObj.IVal & bObj.IVal})
	} else {
		return fmt.Errorf("And: Expected both operands to be bool or int, got %v and %v", aObj.Type, bObj.Type)
	}
	return nil
}

// opOr logical (bool) or bitwise (int) OR
func opOr(interp *Interpreter) error {
	bObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	aObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if aObj.Type == TypeBool && bObj.Type == TypeBool {
		interp.stack.Push(PSObject{Type: TypeBool, BVal: aObj.BVal || bObj.BVal})
	} else if aObj.Type == TypeInt && bObj.Type == TypeInt {
		interp.stack.Push(PSObject{Type: TypeInt, IVal: aObj.IVal | bObj.IVal})
	} else {
		return fmt.Errorf("or: Expected both operands to be bool or int, got %v and %v", aObj.Type, bObj.Type)
	}
	return nil
}

// opNot logical (bool) or bitwise (int) NOT
func opNot(interp *Interpreter) error {
	obj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if obj.Type == TypeBool {
		interp.stack.Push(PSObject{Type: TypeBool, BVal: !obj.BVal})
	} else {
		return fmt.Errorf("not: Expected operand to be bool, got %v", obj.Type)
	}
	return nil
}

// opTrue pushes true onto stack
func opTrue(interp *Interpreter) error {
	interp.stack.Push(PSObject{Type: TypeBool, BVal: true})
	return nil
}

// opFalse pushes false onto stack
func opFalse(interp *Interpreter) error {
	interp.stack.Push(PSObject{Type: TypeBool, BVal: false})
	return nil
}

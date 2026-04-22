package main

import (
	"fmt"
	"math"
)

// Convert int or float PSObject to Go float64
func toFloat(obj PSObject) (float64, error) {
	switch obj.Type {
	case TypeInt:
		return float64(obj.IVal), nil
	case TypeFloat:
		return obj.FVal, nil
	default:
		return 0, fmt.Errorf("Expected number, got %v", obj.Type)
	}
}

// Helper function called at end of other operations
// If both inputs are ints, return result as an int, otherwise return as float
func numericResult(a, b PSObject, result float64) PSObject {
	if a.Type == TypeInt && b.Type == TypeInt {
		return PSObject{Type: TypeInt, IVal: int(result)}
	}
	return PSObject{Type: TypeFloat, FVal: result}
}

// opAdd adds the top two numbers
func opAdd(interp *Interpreter) error {
	bObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	aObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}

	a, err := toFloat(aObj)
	if err != nil {
		return err
	}
	b, err := toFloat(bObj)
	if err != nil {
		return err
	}
	interp.stack.Push(numericResult(aObj, bObj, a+b))
	return nil
}

// opSub subtracts top number from one below it
func opSub(interp *Interpreter) error {
	bObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	aObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	a, err := toFloat(aObj)
	if err != nil {
		return err
	}
	b, err := toFloat(bObj)
	if err != nil {
		return err
	}
	interp.stack.Push(numericResult(aObj, bObj, a-b))
	return nil
}

// opMul multiplies top two numbers on stack
func opMul(interp *Interpreter) error {
	bObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	aObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	a, err := toFloat(aObj)
	if err != nil {
		return err
	}
	b, err := toFloat(bObj)
	if err != nil {
		return err
	}
	interp.stack.Push(numericResult(aObj, bObj, a*b))
	return nil
}

// opDiv divides second element by top element, returns float always
func opDiv(interp *Interpreter) error {
	bObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	aObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	a, err := toFloat(aObj)
	if err != nil {
		return err
	}
	b, err := toFloat(bObj)
	if err != nil {
		return err
	}
	if b == 0 {
		return fmt.Errorf("Division by zero")
	}
	interp.stack.Push(PSObject{Type: TypeFloat, FVal: a / b})
	return nil
}

// opIDiv divides second element by top element, returns int result
func opIdiv(interp *Interpreter) error {
	bObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	aObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if aObj.Type != TypeInt || bObj.Type != TypeInt {
		return fmt.Errorf("Idiv: Both operands must be integers")
	}
	if bObj.IVal == 0 {
		return fmt.Errorf("IDiv by zero")
	}
	interp.stack.Push(PSObject{Type: TypeInt, IVal: aObj.IVal / bObj.IVal})
	return nil
}

// opMod computes a mod b, b is the top element, a is the next element down
// Both must be ints, int is returned
func opMod(interp *Interpreter) error {
	bObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	aObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if aObj.Type != TypeInt || bObj.Type != TypeInt {
		return fmt.Errorf("Mod: Both operands must be integers")
	}
	if bObj.IVal == 0 {
		return fmt.Errorf("Mod by zero")
	}
	interp.stack.Push(PSObject{Type: TypeInt, IVal: aObj.IVal % bObj.IVal})
	return nil
}

// opAbs returns absolute value of top element, keeps Int/Float type
func opAbs(interp *Interpreter) error {
	obj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	switch obj.Type {
	case TypeInt:
		val := obj.IVal
		if val < 0 {
			val = -val
			interp.stack.Push(PSObject{Type: TypeInt, IVal: val})
		}
	case TypeFloat:
		val := obj.FVal
		if val < 0 {
			val = -val
			interp.stack.Push(PSObject{Type: TypeFloat, FVal: val})
		}
	default:
		return fmt.Errorf("Abs: Expected number, got %v", obj.Type)
	}
	return nil
}

// opNeg negates top element, keeps original type
func opNeg(interp *Interpreter) error {
	obj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	switch obj.Type {
	case TypeInt:
		interp.stack.Push(PSObject{Type: TypeInt, IVal: -obj.IVal})
	case TypeFloat:
		interp.stack.Push(PSObject{Type: TypeFloat, FVal: -obj.FVal})
	default:
		return fmt.Errorf("Neg: Expected number, got %v", obj.Type)
	}
	return nil
}

// opCeiling rounds top element up to nearest int, returns float
func opCeiling(interp *Interpreter) error {
	obj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	objF, err := toFloat(obj)
	if err != nil {
		return err
	}
	interp.stack.Push(PSObject{Type: TypeFloat, FVal: math.Ceil(objF)})
	return nil
}

// opFloor rounds top element down to nearest int, returns float
func opFloor(interp *Interpreter) error {
	obj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	objF, err := toFloat(obj)
	if err != nil {
		return err
	}
	interp.stack.Push(PSObject{Type: TypeFloat, FVal: math.Floor(objF)})
	return nil
}

// opRound rounds top element to nearest int, returns float
func opRound(interp *Interpreter) error {
	obj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	objF, err := toFloat(obj)
	if err != nil {
		return err
	}
	interp.stack.Push(PSObject{Type: TypeFloat, FVal: math.Round(objF)})
	return nil
}

// opSqrt computes square root of top element, returns float
func opSqrt(interp *Interpreter) error {
	obj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	objF, err := toFloat(obj)
	if err != nil {
		return err
	}
	if objF < 0 {
		return fmt.Errorf("Sqrt: Argument must be a non-negative number")
	}
	interp.stack.Push(PSObject{Type: TypeFloat, FVal: math.Sqrt(objF)})
	return nil
}

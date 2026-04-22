package main

import (
	"fmt"
)

// opGet returns character code at given index in string
func opGet(interp *Interpreter) error {
	indexObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	strObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if strObj.Type != TypeString {
		return fmt.Errorf("Get: Expected string, got %v", strObj.Type)
	}
	if indexObj.Type != TypeInt {
		return fmt.Errorf("Get: Expected int index, got %v", indexObj.Type)
	}
	index := indexObj.IVal
	if index < 0 || index >= len(strObj.SVal) {
		return fmt.Errorf("Get: Index %d out of bounds for string of length %d", index, len(strObj.SVal))
	}
	interp.stack.Push(PSObject{Type: TypeInt, IVal: int(strObj.SVal[index])})
	return nil
}

// opGetInterval returns substring from index to n characters
func opGetInterval(interp *Interpreter) error {
	countObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	indexObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	strObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if strObj.Type != TypeString {
		return fmt.Errorf("GetInterval: Expected string, got %v", strObj.Type)
	}
	if indexObj.Type != TypeInt {
		return fmt.Errorf("GetInterval: Expected int index, got %v", indexObj.Type)
	}
	if countObj.Type != TypeInt {
		return fmt.Errorf("GetInterval: Expected int count, got %v", countObj.Type)
	}
	index := indexObj.IVal
	count := countObj.IVal
	if index < 0 || count < 0 || index+count > len(strObj.SVal) {
		return fmt.Errorf("GetInterval: Index %d and count %d out of bounds for string of length %d", index, count, len(strObj.SVal))
	}
	substring := strObj.SVal[index : index+count]
	interp.stack.Push(PSObject{Type: TypeString, SVal: substring})
	return nil
}

// opPutInterval replaces part of string1 with string2 starting at index
// Input: string1, index, string2
func opPutInterval(interp *Interpreter) error {
	str2Obj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	indexObj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	str1Obj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if str1Obj.Type != TypeString {
		return fmt.Errorf("PutInterval: Expected string1, got %v", str1Obj.Type)
	}
	if indexObj.Type != TypeInt {
		return fmt.Errorf("PutInterval: Expected int index, got %v", indexObj.Type)
	}
	if str2Obj.Type != TypeString {
		return fmt.Errorf("PutInterval: Expected string2, got %v", str2Obj.Type)
	}

	index := indexObj.IVal
	str1 := []byte(str1Obj.SVal)
	str2 := []byte(str2Obj.SVal)

	if index < 0 || index+len(str2) > len(str1) {
		return fmt.Errorf("PutInterval: Index %d out of bounds for string of length %d", index, len(str1))
	}

	copy(str1[index:], str2)
	_ = str1
	return nil
}

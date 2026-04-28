package main

import (
	"fmt"
)

// opPrint prints top stirng on stack without newline
func opPrint(interp *Interpreter) error {
	obj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	if obj.Type != TypeString {
		return fmt.Errorf("Print: Expected string, got %v", obj.Type)
	}
	fmt.Print(obj.SVal)
	return nil
}

// opEqual prints top object with newline and pops it
func opEqual(interp *Interpreter) error {
	obj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	fmt.Println(psObjectToString(obj))
	return nil
}

// opEqualEqual prints top object in full detail with newline and pops
func opEqualEqual(interp *Interpreter) error {
	obj, err := interp.stack.Pop()
	if err != nil {
		return err
	}
	fmt.Println(psObjectToDetail(obj))
	return nil
}

// psObjectToString converts PSObject to string for printing
func psObjectToString(obj PSObject) string {
	switch obj.Type {
	case TypeInt:
		return fmt.Sprintf("%d", obj.IVal)
	case TypeFloat:
		return fmt.Sprintf("%g", obj.FVal)
	case TypeBool:
		return fmt.Sprintf("%t", obj.BVal)
	case TypeString:
		return obj.SVal
	case TypeName:
		return "/" + obj.SVal
	case TypeArray:
		return "--proc--"
	case TypeDict:
		return "--dict--"
	default:
		return "--unknown--"
	}
}

// psObjectToDetail converts PSObject to detailed string representation
func psObjectToDetail(obj PSObject) string {
	switch obj.Type {
	case TypeInt:
		return fmt.Sprintf("int: %d", obj.IVal)
	case TypeFloat:
		return fmt.Sprintf("float: %g", obj.FVal)
	case TypeBool:
		return fmt.Sprintf("bool: %t", obj.BVal)
	case TypeString:
		return fmt.Sprintf("string: \"%s\"", obj.SVal)
	case TypeName:
		return fmt.Sprintf("name: /%s", obj.SVal)
	case TypeArray:
		result := "{ "
		for _, o := range obj.AVal {
			result += psObjectToString(o) + " "
		}
		result += "}"
		return result
	case TypeDict:
		result := "<< \n"
		for k, v := range obj.DVal {
			result += fmt.Sprintf("  /%s %s\n", k, psObjectToString(v))
		}
		result += ">>"
		return result
	default:
		return "--unknown--"
	}
}

func opSetDynamic(interp *Interpreter) error {
	interp.dictStack.SetLexical(false)
	fmt.Println("scoping: dynamic")
	return nil
}

func opSetLexical(interp *Interpreter) error {
	interp.dictStack.SetLexical(true)
	fmt.Println("scoping: lexical")
	return nil
}

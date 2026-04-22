// interpreter_test.go
package main

import (
	"testing"
)

// helper to create a fresh interpreter for each test
func newTestInterp() *Interpreter {
	return NewInterpreter()
}

// helper to run a postscript string and return the interpreter
func run(t *testing.T, interp *Interpreter, code string) {
	t.Helper()
	tokens := Tokenize(code)
	err := interp.Execute(tokens)
	if err != nil && err.Error() != "quit" {
		t.Fatalf("execution error: %v", err)
	}
}

// helper to pop and check an int from the stack
func expectInt(t *testing.T, interp *Interpreter, expected int) {
	t.Helper()
	obj, err := interp.stack.Pop()
	if err != nil {
		t.Fatalf("stack underflow: %v", err)
	}
	if obj.Type != TypeInt {
		t.Fatalf("expected TypeInt, got %v", obj.Type)
	}
	if obj.IVal != expected {
		t.Fatalf("expected %d, got %d", expected, obj.IVal)
	}
}

// helper to pop and check a float from the stack
func expectFloat(t *testing.T, interp *Interpreter, expected float64) {
	t.Helper()
	obj, err := interp.stack.Pop()
	if err != nil {
		t.Fatalf("stack underflow: %v", err)
	}
	if obj.Type != TypeFloat {
		t.Fatalf("expected TypeFloat, got %v", obj.Type)
	}
	if obj.FVal != expected {
		t.Fatalf("expected %f, got %f", expected, obj.FVal)
	}
}

// helper to pop and check a bool from the stack
func expectBool(t *testing.T, interp *Interpreter, expected bool) {
	t.Helper()
	obj, err := interp.stack.Pop()
	if err != nil {
		t.Fatalf("stack underflow: %v", err)
	}
	if obj.Type != TypeBool {
		t.Fatalf("expected TypeBool, got %v", obj.Type)
	}
	if obj.BVal != expected {
		t.Fatalf("expected %v, got %v", expected, obj.BVal)
	}
}

// helper to pop and check a string from the stack
func expectString(t *testing.T, interp *Interpreter, expected string) {
	t.Helper()
	obj, err := interp.stack.Pop()
	if err != nil {
		t.Fatalf("stack underflow: %v", err)
	}
	if obj.Type != TypeString {
		t.Fatalf("expected TypeString, got %v", obj.Type)
	}
	if obj.SVal != expected {
		t.Fatalf("expected %q, got %q", expected, obj.SVal)
	}
}

// helper to check stack is empty
func expectEmpty(t *testing.T, interp *Interpreter) {
	t.Helper()
	if interp.stack.Len() != 0 {
		t.Fatalf("expected empty stack, got %d items", interp.stack.Len())
	}
}

// ── Stack Manipulation ────────────────────────────────────────────────────────

func TestDup(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "5 dup")
	expectInt(t, interp, 5)
	expectInt(t, interp, 5)
}

func TestExch(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "1 2 exch")
	expectInt(t, interp, 1)
	expectInt(t, interp, 2)
}

func TestPop(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "1 2 pop")
	expectInt(t, interp, 1)
	expectEmpty(t, interp)
}

func TestCopy(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "1 2 3 2 copy")
	expectInt(t, interp, 3)
	expectInt(t, interp, 2)
	expectInt(t, interp, 3)
	expectInt(t, interp, 2)
	expectInt(t, interp, 1)
}

func TestClear(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "1 2 3 clear")
	expectEmpty(t, interp)
}

func TestCount(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "1 2 3 count")
	expectInt(t, interp, 3)
}

// ── Arithmetic ────────────────────────────────────────────────────────────────

func TestAdd(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "3 4 add")
	expectInt(t, interp, 7)
}

func TestAddFloat(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "3 4.0 add")
	expectFloat(t, interp, 7.0)
}

func TestSub(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "10 3 sub")
	expectInt(t, interp, 7)
}

func TestMul(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "3 4 mul")
	expectInt(t, interp, 12)
}

func TestDiv(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "10 2 div")
	expectFloat(t, interp, 5.0)
}

func TestIdiv(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "10 3 idiv")
	expectInt(t, interp, 3)
}

func TestMod(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "10 3 mod")
	expectInt(t, interp, 1)
}

func TestAbs(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "-5 abs")
	expectInt(t, interp, 5)
}

func TestNeg(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "5 neg")
	expectInt(t, interp, -5)
}

func TestCeiling(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "3.2 ceiling")
	expectFloat(t, interp, 4.0)
}

func TestFloor(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "3.9 floor")
	expectFloat(t, interp, 3.0)
}

func TestFloorNegative(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "-4.2 floor")
	expectFloat(t, interp, -5.0)
}

func TestRound(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "3.5 round")
	expectFloat(t, interp, 4.0)
}

func TestSqrt(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "9 sqrt")
	expectFloat(t, interp, 3.0)
}

// ── Boolean ───────────────────────────────────────────────────────────────────

func TestEqTrue(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "3 3 eq")
	expectBool(t, interp, true)
}

func TestEqFalse(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "3 4 eq")
	expectBool(t, interp, false)
}

func TestNe(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "3 4 ne")
	expectBool(t, interp, true)
}

func TestGt(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "5 3 gt")
	expectBool(t, interp, true)
}

func TestLt(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "3 5 lt")
	expectBool(t, interp, true)
}

func TestAnd(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "true true and")
	expectBool(t, interp, true)
}

func TestAndFalse(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "true false and")
	expectBool(t, interp, false)
}

func TestOr(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "false true or")
	expectBool(t, interp, true)
}

func TestNot(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "true not")
	expectBool(t, interp, false)
}

// ── Dictionary ────────────────────────────────────────────────────────────────

func TestDef(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "/x 42 def x")
	expectInt(t, interp, 42)
}

func TestDefProc(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "/double { 2 mul } def 5 double")
	expectInt(t, interp, 10)
}

func TestBeginEnd(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "2 dict begin /x 10 def x end")
	expectInt(t, interp, 10)
}

// ── Strings ───────────────────────────────────────────────────────────────────

func TestStringLength(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "(hello) length")
	expectInt(t, interp, 5)
}

func TestGet(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "(hello) 0 get")
	expectInt(t, interp, 104) // ASCII for 'h'
}

func TestGetInterval(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "(hello) 1 3 getinterval")
	expectString(t, interp, "ell")
}

// ── Control Flow ──────────────────────────────────────────────────────────────

func TestIf(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "true { 42 } if")
	expectInt(t, interp, 42)
}

func TestIfFalse(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "false { 42 } if")
	expectEmpty(t, interp)
}

func TestIfelse(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "true { 1 } { 2 } ifelse")
	expectInt(t, interp, 1)
}

func TestIfelsefalse(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "false { 1 } { 2 } ifelse")
	expectInt(t, interp, 2)
}

func TestRepeat(t *testing.T) {
	interp := newTestInterp()
	run(t, interp, "0 3 { 1 add } repeat")
	expectInt(t, interp, 3)
}

func TestFor(t *testing.T) {
	interp := newTestInterp()
	// sums 1+2+3+4+5 = 15
	run(t, interp, "0 1 1 5 { add } for")
	expectInt(t, interp, 15)
}

func TestForNegative(t *testing.T) {
	interp := newTestInterp()
	// counts down 5+4+3+2+1 = 15
	run(t, interp, "0 5 -1 1 { add } for")
	expectInt(t, interp, 15)
}

// ── Scoping ───────────────────────────────────────────────────────────────────

func TestDynamicScoping(t *testing.T) {
	interp := newTestInterp()
	interp.SetLexical(false)
	// In dynamic scoping, getX sees the x defined in the calling scope
	run(t, interp, `
		/x 1 def
		/getX { x } def
		/x 2 def
		getX
	`)
	expectInt(t, interp, 2)
}

func TestLexicalScoping(t *testing.T) {
	interp := newTestInterp()
	interp.SetLexical(true)
	run(t, interp, "/x 1 def /getX { x } def /x 2 def getX")
	expectInt(t, interp, 1)
}

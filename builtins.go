// builtins.go
package main

type BuiltinFn func(interp *Interpreter) error

var builtins map[string]BuiltinFn

func init() {
	builtins = map[string]BuiltinFn{
	// Stack
	"exch":  opExch,
	"pop":   opPop,
	"dup":   opDup,
	"copy":  opCopy,
	"clear": opClear,
	"count": opCount,

	// Arithmetic
	"add":     opAdd,
	"sub":     opSub,
	"mul":     opMul,
	"div":     opDiv,
	"idiv":    opIdiv,
	"mod":     opMod,
	"abs":     opAbs,
	"neg":     opNeg,
	"ceiling": opCeiling,
	"floor":   opFloor,
	"round":   opRound,
	"sqrt":    opSqrt,

	// Dictionary
	"dict":      opDict,
	"length":    opLength,
	"maxlength": opMaxLength,
	"begin":     opBegin,
	"end":       opEnd,
	"def":       opDef,

	// String
	"get":         opGet,
	"getinterval": opGetInterval,
	"putinterval": opPutInterval,

	// Boolean
	"eq":    opEq,
	"ne":    opNe,
	"ge":    opGe,
	"gt":    opGt,
	"le":    opLe,
	"lt":    opLt,
	"and":   opAnd,
	"or":    opOr,
	"not":   opNot,
	"true":  opTrue,
	"false": opFalse,

	// Control
	"if":     opIf,
	"ifelse": opIfElse,
	"for":    opFor,
	"repeat": opRepeat,
	"quit":   opQuit,

	// IO
	"print": opPrint,
	"=":     opEqual,
	"==":    opEqualEqual,
	}
}
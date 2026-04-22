//All defined types live here

package main

type PSType int

const (
	TypeInt PSType = iota
	TypeFloat
	TypeBool
	TypeString
	TypeName
	TypeArray
	TypeDict
)

type PSObject struct {
	Type  PSType
	IVal  int
	FVal  float64
	BVal  bool
	SVal  string
	AVal  []PSObject
	DVal  map[string]PSObject
	Scope []map[string]PSObject
}

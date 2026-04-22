// Dictionary stack and its methods

package main

import "fmt"

type DictStack struct {
	dicts   []map[string]PSObject
	lexical bool // false for dynamic (default), true for lexical(static) scoping
}

func NewDictStack() *DictStack {
	ds := &DictStack{}
	ds.dicts = append(ds.dicts, make(map[string]PSObject))
	return ds
}

func (ds *DictStack) SetLexical(on bool) {
	ds.lexical = on
}

func (ds *DictStack) Begin(dict map[string]PSObject) {
	ds.dicts = append(ds.dicts, dict)
}

func (ds *DictStack) End() error {
	if len(ds.dicts) <= 1 {
		return fmt.Errorf("dictstack underflow error: cannot pop the global dictionary")
	}
	ds.dicts = ds.dicts[:len(ds.dicts)-1]
	return nil
}

func (ds *DictStack) Define(key string, value PSObject) {
	ds.dicts[len(ds.dicts)-1][key] = value
}

func (ds *DictStack) Lookup(key string) (PSObject, bool) {
	for i := len(ds.dicts) - 1; i >= 0; i-- {
		if val, ok := ds.dicts[i][key]; ok {
			return val, true
		}
	}
	return PSObject{}, false
}

func (ds *DictStack) LookupInScope(key string, scope []map[string]PSObject) (PSObject, bool) {
	for i := len(scope) - 1; i >= 0; i-- {
		if val, ok := scope[i][key]; ok {
			return val, true
		}
	}
	return PSObject{}, false
}

func (ds *DictStack) CaptureScope() []map[string]PSObject {
	snapshot := make([]map[string]PSObject, len(ds.dicts))
	for i, dict := range ds.dicts {
		copied := make(map[string]PSObject, len(dict))
		for k, v := range dict {
			copied[k] = v
		}
		snapshot[i] = copied
	}
	return snapshot
}

func (ds *DictStack) Current() map[string]PSObject {
	return ds.dicts[len(ds.dicts)-1]
}

func (ds *DictStack) Depth() int {
	return len(ds.dicts)
}

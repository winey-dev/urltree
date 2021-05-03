package urltree

import (
	_ "strings"
)

func (t *Tree) SetHookHandler(handler func(method, path string, data interface{})) {
	t.setHook = handler
}
func (t *Tree) SetDataWithHook(method, path string, data interface{}) error {
	// find Node Point and Run SetHookHandelr
	return nil
}

func (t *Tree) SetData(method, path string, data interface{}) error {
	return nil
}

/*
func (t *Tree) DeleteData(method, path string) error {
	return nil
}
*/

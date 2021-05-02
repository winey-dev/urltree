package tree

import (
	_ "strings"
)

// /em-ui/v1/*      	:: 모든 Method 에도 true 설정
// /em-ui/v1/*/GET      :: GET 모든 URL의 GET에 true 설정
// /em-ui/v1/*/* 		:: 모든 URL과 모든 Method에 true 설정

func (t *Tree) SetData(path string, data interface{}) error {
	return nil
}

func setdata(node []Node, p []string, data interface{}) error {
	return nil

}

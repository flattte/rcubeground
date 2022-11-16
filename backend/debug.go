package main

import (
	"fmt"
)

func noDebug(s string, f ...interface{}) {}

func okDebug(s string, f ...interface{}) {
	fmt.Printf(s, f...)
}

var Debugf func(s string, f ...interface{}) = okDebug

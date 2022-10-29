package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

var target int64
var count int64

func init() {
	flag.Int64Var(&target, "target", 0, "Initial key number")
	flag.Int64Var(&count, "count", 1, "Number of keys requested")
	flag.Parse()
}

type File struct {
	Name       string
	Start, End int64
	f          *os.File
}

func main() {
	list := NewList(".")
	defer list.CloseAll()
	data, tail, _ := list.GetKeys(target, count)
	fmt.Println(data)
	if tail > 0 {
		data, _, _ = list.GetKeys(target+count-tail, tail)
		fmt.Println(data)
	}

}

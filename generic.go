package main

import (
	"fmt"
)

func foo[T any](s T) string {
	return fmt.Sprint(s, "::", s)
}

func main() {
	println(foo[int](1))
	println(foo[string]("abc"))
	println(foo[bool](true))
	type T struct {
		Id   int64
		Name string
	}
	println(foo[T](T{
		Id:   22,
		Name: "Name",
	}))
}

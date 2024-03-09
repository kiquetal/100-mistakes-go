package main

import (
	"fmt"
	"runtime"
)

type Foo struct {
	v []byte
}

func main() {
	foos := make([]Foo, 1_000)
	printAlloc()
	for i := 0; i < len(foos); i++ {
		foos[i] = Foo{
			v: make([]byte, 1024*1024),
		}
	}
	printAlloc()
	two := keepFirstTwoElementsOnly2(foos)
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(two)
}

func keepFirstTwoElementsOnly(foos []Foo) interface{} {

	return foos[:2]
}

func keepFirstTwoElementsOnly2(foos []Foo) interface{} {

	res := make([]Foo, 2)
	copy(res, foos)
	return res

}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}

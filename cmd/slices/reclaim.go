package main

import (
	"bufio"
	"fmt"
	"golang_mentoring/internal"
	"os"
	"runtime"
	"strings"
)

// This program demonstrates memory allocation and garbage collection in Go.
// It creates a slice of Foo structs, each containing a large byte slice.
// After populating the slice, it retains only the first two elements and
// forces garbage collection to reclaim memory used by the rest of the elements.
// The memory usage is printed before and after the operations to show the effect of garbage collection.
// As a result, memory footprint should decrease after the garbage collection, but in fact it won't unless unused elements are explicitly set to nil.
// Therefore, even though these 998 elements can’t be accessed, they stay in memory as long as the variable
// returned by `keepFirstTwoElementsOnly` is referenced.  

func main() {
	foos := make([]Foo, 1_000)
	internal.PrintAlloc()

	for i := 0; i < len(foos); i++ {
		foos[i] = Foo{
			v: make([]byte, 1024*1024),
		}
	}
    fmt.Println("Before garbage collection:")
	internal.PrintAlloc()

	// Ask user if they want to use the improved version
	fmt.Println("Choose which method to use:")
	fmt.Println("1. Regular (keepFirstTwoElementsOnly)")
	fmt.Println("2. Improved (keepFirstTwoElementsOnlyImproved)")
	fmt.Print("Enter your choice (1 or 2): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := strings.TrimSpace(scanner.Text())

	var two []Foo
	if choice == "2" {
		fmt.Println("Using improved version...")
		two = keepFirstTwoElementsOnlyImproved(foos)
	} else {
		fmt.Println("Using regular version...")
		two = keepFirstTwoElementsOnly(foos)
	}

	runtime.GC()
    fmt.Println("After garbage collection:")
	internal.PrintAlloc()
	runtime.KeepAlive(two)
}

func keepFirstTwoElementsOnly(foos []Foo) []Foo {
	return foos[:2]
}

func keepFirstTwoElementsOnlyImproved(foos []Foo) []Foo {
	for i := 2; i < len(foos); i++ {
		foos[i].v = nil
	}
	return foos[:2]
}

type Foo struct {
	v []byte
}

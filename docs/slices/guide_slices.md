All code snippets are checked

1. 
### Slices   
https://go.dev/tour/moretypes/7
An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type `[]T` is a slice with elements of type `T`.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

```go
a[low : high]
```

This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of `a`:
```go
a[1:4]
```

2. 
### Slices are like references to arrays
https://go.dev/tour/moretypes/8
A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

E.g.
```go
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // prints [John Paul] \n [Paul George]

	b[0] = "XXX" // alters the underlying array and affects both slices a & b!
	fmt.Println(a, b) // prints prints [John XXX] \n [XXX George]
	fmt.Println(names) // prints [John XXX George Ringo]
}
```

3. 
   ### Slice literals
A slice literal is like an array literal without the length.

This is an array literal:
```go
[3]bool{true, true, false}
```

And this creates the same array as above, then builds a slice that references it:
```go
[]bool{true, true, false}
```

NOTE: literal means that the value "literally" presents what we wrote, like `foo := 42` is literally `42`

Usage of slice literals in case of anonymous structs:

```go
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
```

4. 
   ### Slice defaults
https://go.dev/tour/moretypes/10  
When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

For the array
```go
var a [10]int
```

these slice expressions are equivalent:

```go
a[0:10]
a[:10]
a[0:]
a[:]
```

Example of sequentially overriding slice's bounds:

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s) // [3 5 7]

	s = s[:2]
	fmt.Println(s) // [3 5]

	s = s[1:]
	fmt.Println(s) // [5]
}
```

5. 
   ### Slice length and capacity
https://go.dev/tour/moretypes/11
A slice has both a _length_ and a _capacity_.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, **counting from the first element in the slice** (**NOTE**).

The length and capacity of a slice `s` can be obtained using the expressions `len(s)` and `cap(s)`

You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

Example showing that capacity is not altering strongly couples with the length:

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // len=2 cap=4 [5 7]

	// "Redeem" the last 2 elements storing in the underlying array
	s = s[:4]
	printSlice(s) // len=4 cap=4 [5 7 11 13]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

Note: in the penultimate example the capacity was cut due to the cut of 1st two elements, but we still can "redeem" last elements storing in the underlying array. 

6. 
   ### Nil slices
https://go.dev/tour/moretypes/12
The zero value of a slice is `nil`.

A `nil` slice has a length and capacity of 0 and has no underlying array.

```go
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
```

7. 
   ### Creating a slice with make
https://go.dev/tour/moretypes/13
Slices can be created with the built-in `make` function; this is how you create dynamically-sized arrays.

The `make` function allocates a zeroed array and returns a slice that refers to that array:
```go
a := make([]int, 5)  // len(a)=5
```

To specify a capacity, pass a third argument to `make`:
```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

Example showing how we can access the underlying array of the slice (note slice `c` and `d`)

```go
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a) // a len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	printSlice("b", b) // b len=0 cap=5 []

	c := b[:2]
	printSlice("c", c) // c len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice("d", d) // d len=3 cap=3 [0 0 0] NB!
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
```

8. 
   ## Slices of slices
https://go.dev/tour/moretypes/14
Slices can contain any type, including other slices.

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
```

9.
### Appending to a slice
https://go.dev/tour/moretypes/15

It is common to append new elements to a slice, and so Go provides a built-in `append` function. The [documentation](https://go.dev/pkg/builtin/#append) of the built-in package describes `append`.

```go
func append(s []T, vs ...T) []T
```

The first parameter `s` of `append` is a slice of type `T`, and the rest are `T` values to append to the slice.

The resulting value of `append` is a slice containing all the elements of the original slice plus the provided values.

If the backing array of `s` is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

(To learn more about slices, read the [Slices: usage and internals](https://go.dev/blog/go-slices-usage-and-internals) article.)

Example demonstrating growth of the slice's `len` and `cap`:

```go
package main

import "fmt"

func main() {
	var s []int
	printSlice(s) // len=0 cap=0 []

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s) // len=1 cap=1 [0]

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s) // len=2 cap=2 [0 1]

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s) // len=5 cap=6 [0 1 2 3 4]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

10.
### Range
https://go.dev/tour/moretypes/16
The `range` form of the `for` loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

11.
### Range continued
https://go.dev/tour/moretypes/17
You can skip the index or value by assigning to `_`.
```go
for i, _ := range pow
for _, value := range pow
```

If you only want the index, you can omit the second variable.
```go
for i := range pow
```

```go
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```

### MISC & GOTCHAS

12. Transform array to slice:
```go
arr := [3]int{1, 2, 3}
slice := arr[:]
```

13. If we sort the slice, the underlying array also alters:
```go
arr := [3]string{"c", "a", "b"}
slice := arr[:]
sort.Strings(slice)
fmt.Println(arr) // [a b c]
```

That's because sorting alters a slice.

14. When using `copy(dst, src)` we need to make sure that `dst` has the same length as `src` otherwise not elements of `src` will be copied:
```go
src := []int{0, 1, 2}  
var dst []int  
copy(dst, src)  
fmt.Println(dst) // []
```


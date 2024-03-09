#### Data types

- Common mistkaes related to basic types
- Fundamental concepts for slices and maps to prevent possible bugs, leaks, or inaccuracies
- Comparing values


#### 3.19 Not understanding floatng points.

Using 0o 
Binary ob or oB
Hexadecimal 0x or 0X
Imaginary 3i
use underscore for readability



#### 3.20 Not understanding slice length and capacity

Capacity vs length: 
length:  actual elements
capacity: backing array


#### 3.21 Ineficient slice initialization

bars := make([]Bar,0)

bars := make([]Bar, 0, n)  //capacity equals to n


#### 3.22 Being confused about nil vs empty slices

- One of the main differences between a nil and a empty slice regards allocations. Initializing a nil slice doesn't 
require any allocation, which isn't the case for a empty slice.

var s1 [] string
fmt.Printf(append(s1,"foo"))

Option 2: s := []string(nil) === 
Option 3: s := []string {} === only for initializing elements


#### 2.23 Not properly checking if a slice is empty

func handleOperations(id string) {
	operations:= getOperations(id)
	if operations != nill {
		handle(operations)
	}

}

func getOperations(id string) [] float32 { 
	operations:= make([]float32,0)
	if id == "" {
	return operations
}
	return operations

==> returns a empty slice

change the return for nil, so the comparission will work.
==> check the len() nil services always are empty.

the same strategy should be applied for maps.


#### 2.24 Not making slice copies correctly.

src := []int{0,1,2}
dst := make([]int, len(src))

#### 2.25 Unexpected side effects using slice append

```
func main() {
	s := []int{1,2,3}
  f(s[:2])
  fmt.Println(s)
}
func f(s []int) {
	s = append(s, 4)
}
```
- You should use copy() to avoid side effects.

- Can use [:2:2] - full slice expression to avoid side effects.

#### 3.10.1 ### Leaking capacity when slicing a slice

Using the full slice option is not a option.

As a rule of thumb, remember that slicing a large slice or array can lead 
to potential high memory consumption. The remaining space won't be reclaimed 
by the GC, and we can keep a large backing array despite using only a few
elements. Using a slice is the solution to preven such case.



### Code and project organization

- Organizing our code idiomatically
- Dealing efficiently with abstractions: interfaces and genercis
- Best practices regarding how the project is structured

#### 2.1 Unintended variable shadowing
#### 2.2 Unnecesary nested code.
#### 2.3 Missting init functions

```go
package main
import "fmt"

var a = func() int {
	fmt.Println("var")
	return 0
}()

func init() {
	fmt.Println("init")
}
func main() {
	fmt.Println("main")
}

``` 
Variables are evaluated first, then init, and later main method.

Using logic in init , prevents the caller function to handle retry for example. [UNIT TESTING is more complicated when logic resides in init]

#### 2.4 Overusing getters and settes
We use Balance -> not GetBalancer
We use SetBalance
#### 2.5 Interface pollution

func copySourceToDest( source io.Reader, dest io.Write) error {
} ---> Reusability for interfaces.
"The bigger the interface, the weaker the abstractions"

Using when:

- Common behavior: {swapping 2 elements, retrieving elements from collection}
- Decoupling: {Follow the Liskov }
- Restrict behavior:{}

We should create an interface when we need it, not when foresee that we could need it.

#### 2.6 Interface on producer

Interface on the producer side :  not a go idiom

#### 2.7 Returning interface

Returnng structs instead of interfaces
Accept interfaces

#### 2.8 Any says nothing

empty interface interface{} -> any

#### 2.9 Being confused about when to use generics

func foo[T any] (t T) {}

Supplying a type argument is called initialization.
K as comparable
Resctricting type argument to match specific requirements is called a constraint.
~int -> underlying int
int -> only int

generics for data structures [binary tree,heap]

#### 2.10 NOt being aware of the possible problems with type embedding

Embedding are useful for interfaces.

type InMeM struct {
	sync.Mutex
	m map[string]int
}

func New() *InMeM {
	return &InMeM{m: make(m[string]int)}
}


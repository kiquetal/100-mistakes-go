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
===Resolve the previous adding a name as m sync.Mutex
func New() *InMeM {
	return &InMeM{m: make(m[string]int)}
}

[It shouldn't be used soley as some syntatic sugar to simplify accessing a field]
[It shouldn't promoted data]


#### 2.11 Not using functional options patterns

Builder Pattern: so we can use method chaining. Return the object using in the builder.

Functional options patterns: 
An unexported struct holds the configuration options
Each options is a funciton that reutnrs the same type

type options struct {
  port *int
}
type Option func(options *options) error

#### 2.12 Project misorganization

Better consistent.

#### 2.13 Creating utility packages

Better a name for thoughtfull.


#### 2.14 Ignoring package name collisions

Be carefull with the name 


#### 2.15 Missing code documentation


If is exported must be documented.

#### 2.16 Not using linters.

A linter is an automatic tool to analyze code and catch errors.

vet
errcheck
gocyclo
goconst

code formatter
gofmt
goimports




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


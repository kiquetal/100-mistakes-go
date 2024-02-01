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



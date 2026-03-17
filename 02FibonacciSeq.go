//Print Fibonacci sequence
package main
import "fmt"

func FibonacciSeq(){
    a := 0                     
    b := 1
    for i:=1;i<10;i++ {
        fmt.Print(a)
        a , b = b, a + b
    }
}
// Main Func would be same no difference
func main(){
    FibonacciSeq()
}

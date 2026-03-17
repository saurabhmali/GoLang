/*
#Function Return in GO

NOTE: When you return something from go func, You are supposed to mention the dtype of the 
return variable in the function declaration.
*/
//import
package main
import "fmt"

func ReturnVariable(x int, y int) int {   // Note here
    sum := 0
    sum = x + y
    return sum
}

func main(){
    op := ReturnVariable(10,12)
    fmt.Println(op)
}

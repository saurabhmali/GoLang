
// Find Largest Number/////////////WRONG CODE

package main
import "fmt"

func LargestNumber(arr []int, lnum int){
    
    for i:=0; i < len(arr); i++ {
        if arr[i] > lnum {
            lnum := arr[i]
        }
    }
    
}


func main(){
    arr1 := []int{99, 22, 44, 86, 10, 32}
    lnum := arr1[0]
    maxnum := LargestNumber(arr1,lnum)
    fmt.Println(maxnum)
}

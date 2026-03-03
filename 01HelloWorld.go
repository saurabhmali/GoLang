//1. Hello World in Go

package main     // to declare that this is the main package among others
import ("fmt")   // way to import multiple packages. (comma-seperated)


// 1st way using return/////////////////
func PrintHW1() string{
    return "Hello World !"
}

// 2nd way ////////////////////////////
func PrintHW2(){
    fmt.Println("Hellow World !!")
}

// function calling////////////////////
func main(){
    output := PrintHW1()     // Calling 1st func
    fmt.Println(output)
    
    PrintHW2()               // Calling 2nd func
}





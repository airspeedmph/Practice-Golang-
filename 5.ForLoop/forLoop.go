//only for loop is there , there is no while loop in Golang

package main

import "fmt"


//simple for loop in golang 

// func main (){
// 	var i int =0
// 	for  ; i<=5 ; i++ {
// 		fmt.Println(i)
// 	}
// }



//while loop using for 

// func main (){
// 	var x int = 0 
// 	for x <10 {
// 		fmt.Println(x)
// 		x++
// 	}
// }

//range 
func main (){
	for x:=range 10 {
		fmt.Println(x)

	}
}

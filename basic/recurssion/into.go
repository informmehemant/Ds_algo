// recursion => russian doll
package main

import (
	"errors"
	"fmt"
)
func main() {
    // openRussianDoll(4)
    // power := powerOfTwoIt(2)
    // fmt.Println(power)
    fmt.Println(factorial(4))
}
func openRussianDoll(doll int ) {
    if(doll == 1) {
        fmt.Println("ALl dolls are opened")
    } else {
        openRussianDoll(doll -1)
    }
    
}

// when do we use recurssion 
/*
 1. if u can divide the problem into similar sub problems
 2. Design an algorith to compute nth
 3. Write code to list the n...
 4. practice
 5. The prominent usage of recursion in data structure like trees and graphs.
 6. interviews
 7. it is used in many algorithm(divide and conquer, greedy and dynamic programming)
*/

// recursion vs iterative Solutions
// recurion method 

func powerOfTwo(n int) int{
    if(n == 0 || n == 1 ){
        return 1
    } else{
        power := powerOfTwo(n-1)
        return power * 2
    }
    
}
   
// iterative solutions
func powerOfTwoIt(n int) int{
    i := 0
    power := 1
    for i <= n {
      power = power * 2
      i = i + 1
    }
    return power
}
  
//  writing recursion in 3 steps

// func factorial(n int)  (int, error) {
//     if n < 0 {
//        return 0, errors.New("Please enter Positive number")
//     }
//     if n == 0 || n == 1   {
//       return 1 , nil
//     }
//     f, err := factorial(n-1)
//     if err != nil {
//         return 0, err
//     }
//     return n * f, nil    
// }

func factorial(n int) (int , error) {
    if n< 0 {
        return 0, errors.New("Please enter Postive number!")
    }
    if n == 0 || n == 1 {
        return 1, nil
    }
    f,err := factorial(n-1)

    if err != nil {
        return 0, err
    }
    return f * n, nil

}

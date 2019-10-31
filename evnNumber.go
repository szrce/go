/* this variable
 * var x string = "hello"
 * var x int = 1
 * */

/*
 * x := "short variable"
 * y := 1
 * x = x + "short2"  string plus
 */

/* multiple variable
 *	var (
 *		x = "hello0"
 *		b = "hello1"
 *	)
 * const(
 *		x = "hello0"
 *		b = "hello1"
 * )
 * switch i {
 * case 0: fmt.Println("Sıfır")
 * case 1: fmt.Println("Bir")
 * case 2: fmt.Println("İki")
 * default: fmt.Println("Bilinmeyen sayı")
 * }
 */

package main

import "fmt"

func main() {
 i := 1
 
 for i <= 100{
  if( i % 2 == 0){
    fmt.Println(i)
  }
	i = i + 1
	}
	
}

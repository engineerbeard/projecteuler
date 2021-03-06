/*
	Project Euler Question 1:
	If we list all the natural numbers below 10 that are multiples of 3 or 5,
	we get 3, 5, 6 and 9. The sum of these multiples is 23. Find the sum of
	all the multiples of 3 or 5 below 1000.

	Author: thebeard@engineerbeard.com
*/

package main
import "fmt"

func main() {
	fmt.Println(mult3and5(1000))
}

func mult3and5(max int) int {
	sum := 0;
	for i:= 0; i < max; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum = sum + i;
		}
	}
	return sum
}

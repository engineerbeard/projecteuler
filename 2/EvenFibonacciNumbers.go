/*
	Project Euler Question 2:
	Each new term in the Fibonacci sequence is generated by adding the previous
	two terms. By starting with 1 and 2, the first 10 terms will be:

	1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

	By considering the terms in the Fibonacci sequence whose values do not exceed
	four million, find the sum of the even-valued terms.

	Author: thebeard@engineerbeard.com
*/
package main

import (
	"fmt"
	"time"
)

const MAX_FIB_VAL = 4000000

func main() {
	start := time.Now()
	fmt.Println(fibNumbers(MAX_FIB_VAL))
	elapsed := time.Since(start)
    	fmt.Printf("Elapsed: %s", elapsed)
}

func fibNumbers(max int) int {
	f1 := 1;
	f2 := 1;
	fibNum := 0;
	evenSum := 0;
	for i := 0; i < max; i++ {
		fibNum = f1 + f2;
		f1 = f2;
		f2 = fibNum;
		if fibNum % 2 == 0 {
			evenSum = evenSum + fibNum;
		}
		if fibNum >= MAX_FIB_VAL {
			break;
		}
	}
	return evenSum;
}
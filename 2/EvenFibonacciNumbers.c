/*
	Project Euler Question 2:
	Each new term in the Fibonacci sequence is generated by adding the previous 
	two terms. By starting with 1 and 2, the first 10 terms will be:

	1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

	By considering the terms in the Fibonacci sequence whose values do not exceed 
	four million, find the sum of the even-valued terms.

	Author: thebeard@engineerbeard.com
*/
#include <stdio.h>
int main()
{
	int n1 = 1;
	int n2 = 2;
	int newFibNum = 0;
	long evenSums = 2;
	int const MAX = 4000000;
	for(int i = 2; i < MAX; i++)
	{
		newFibNum = n1 + n2;
		n1 = n2;
		n2 = newFibNum;
		
		if(newFibNum % 2 == 0)
			evenSums = evenSums + newFibNum;
		
		if(newFibNum > MAX)
			break;
	}

	printf("Total: %i", evenSums);
	return 0;
}
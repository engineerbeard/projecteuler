"""
The sequence of triangle numbers is generated by adding the natural numbers.
So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28. The first ten terms would be:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

Let us list the factors of the first seven triangle numbers:

 1: 1
 3: 1,3
 6: 1,2,3,6
10: 1,2,5,10
15: 1,3,5,15
21: 1,3,7,21
28: 1,2,4,7,14,28
We can see that 28 is the first triangle number to have over five divisors.

What is the value of the first triangle number to have over five hundred divisors?
"""

import time
import math

def fastComputeFactors(n):
    i = 2
    fCount = 0
    while i <= math.sqrt(n):
        if n % i == 0:
            fCount += 1
            if i != (n / i):
                fCount += 1
        i = i + 1
    return fCount + 2 # add 1 and i back in to total count


def genTriangleNums(n):
    sum, i = 0, 1
    while (True):
        sum, i = sum + i, i + 1
        if fastComputeFactors(sum) > n: return sum

tStart = time.time()
print genTriangleNums(500)
print time.time() - tStart

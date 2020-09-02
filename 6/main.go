package main

import "fmt"

// source exercise: codility 3.tape_equilibrium
/*
For example, consider array A such that:
  A[0] = 3
  A[1] = 1
  A[2] = 2
  A[3] = 4
  A[4] = 3

We can split this tape in four places:
  P = 1, difference = |3 − 10| = 7
  P = 2, difference = |4 − 9| = 5
  P = 3, difference = |6 − 7| = 1
  P = 4, difference = |10 − 3| = 7

that, given a non-empty array A of N integers, returns the minimal difference that can be achieved.
the function should return 1, as explained above.

Write an efficient algorithm for the following assumptions:
  N is an integer within the range [2..100,000];
  each element of array A is an integer within the range [−1,000..1,000].
*/

func main() {
  array := []int{3, 1, 2, 4, 3}
  result := Solution(array)
  fmt.Println("result:",result)

  array = []int{-2000, 2000}
  result = Solution(array)
  fmt.Println("result:",result)
}

func Solution(A []int) int {
    sumDec := 0
    sumInc := 0
    min := 0

    if len(A) == 0 {
        return -1
    }

    for _, value := range A {
        sumDec += value
    }

    min = sumDec - (2 * A[0])
    if min < 0 {
        min *= -1
    }

    for i := 0; i < len(A) - 1; i++ {
        sumDec -= A[i]
        sumInc += A[i]
        difference := sumDec - sumInc

        if difference < 0 {
            difference *= -1
        }
        if difference < min {
            min = difference
        }
    }

    return min
}

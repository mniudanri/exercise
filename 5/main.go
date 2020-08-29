package main
import "fmt"

// source exercise: codility 3.perm_missing_elem
/*
given an array A, returns the value of the missing element.
For example, given array A such that:
  A[0] = 2
  A[1] = 3
  A[2] = 1
  A[3] = 5

the function should return 4, as it is the missing element.

Write an efficient algorithm for the following assumptions:
  - N is an integer within the range [0..100,000];
  - the elements of A are all distinct;
  - each element of array A is an integer within the range [1..(N + 1)].
*/

func main() {
  array := []int{2,3,1,5}
  result := Solution(array)
  fmt.Println("result:",result)

  array = []int{15, 3, 18, 5, 6, 7, 1, 8, 20, 9, 19, 17, 2, 10, 11, 13, 14, 16, 4}
  result = Solution(array)
  fmt.Println("result:",result)
}

func Solution(A []int) int {
    // holds sum total of number in A
    actualSum := 0

    // length array before missing element
    lengthArray := len(A) + 1

    // sum total of n formula: 1/2n(n+1)
    // expected-sum of number sequence n
    expectedSum := lengthArray * (lengthArray + 1) / 2

    for _, value := range A {
        actualSum += value
    }

    // get exact 1 missing sequence number by differencing expected-sum and actual sum
    return expectedSum - actualSum
}

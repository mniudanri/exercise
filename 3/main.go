package main
import "fmt"

// source exercise: codility 2.odd_occurrences_in_array
/*
  For example, given array A which contains odd numbers such that:
    A[0] = 9  A[1] = 3  A[2] = 9
    A[3] = 3  A[4] = 9  A[5] = 7
    A[6] = 9

  - the elements at indexes 0 and 2 have value 9,
  - the elements at indexes 1 and 3 have value 3,
  - the elements at indexes 4 and 6 have value 9,
  - the element at index 5 has value 7 and is unpaired.

  the function should return 7 as unpaired-number.

  Write an efficient algorithm for the following assumptions:
    - N is an odd integer within the range [1..1,000,000];
    - each element of array A is an integer within the range [1..1,000,000,000];
    - all but one of the values in A occur an even number of times.
*/

func main() {
  arr := []int{9, 3, 9, 3, 9, 7, 9}
  result := Solution(arr)
  fmt.Println("result:",result)

  arr = []int{1000000, 9, 1, 1000000, 1, 9, 11, 101, 1, 11, 1}
  result = Solution(arr)
  fmt.Println("result:",result)
}

func Solution(A []int) int {
    // NOTE:
    // Step: array should be has add length, return fail if even
    //       use bitwise xor to get unique/unpaired number in array

    if (len(A) % 2 == 0) {
      return -1
    }

    result := A[0]

    for i := 1; i < len(A); i++ {
        // debugging
        // fmt.Printf("\t%020b\n\t%020b %d\n\n", result, A[i], A[i])
        result ^= A[i]
    }
    return result
}

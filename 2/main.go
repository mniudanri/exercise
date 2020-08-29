package main
import "fmt"

// source exercise: codility 2.cyclic_rotation
/*
  For example, given
    A = [3, 8, 9, 7, 6]
    K = 3

  the function should return [9, 7, 6, 3, 8]. Three rotations were made:
    [3, 8, 9, 7, 6] -> [6, 3, 8, 9, 7]
    [6, 3, 8, 9, 7] -> [7, 6, 3, 8, 9]
    [7, 6, 3, 8, 9] -> [9, 7, 6, 3, 8]
*/

func main() {
  numRotation := 3
  arr := []int{1, 2, 3, 4, 5}
  result := Solution(arr, numRotation)

  fmt.Printf("%v, %d\n", arr, numRotation)
  fmt.Println(result)

  numRotation = 99
  arr = []int{8, 5, 19, 2, 11}
  result = Solution(arr, numRotation)

  fmt.Printf("\n%v, %d\n", arr, numRotation)
  fmt.Println(result)
}

func Solution(A []int, K int) []int {
    // NOTE: format A to new array equals to rotated-array with center index is remainder of K devided by A's length
    boundaryIdx := len(A) - K % len(A)
    firstArr := A[boundaryIdx:]
    lastArr := A[:boundaryIdx]

    return append(firstArr, lastArr...)
}

package main
import "fmt"

// source exercise: codility 3.frog_jump
/*
For example, given:
  X = 10, as stating position of frog
  Y = 85, as minimum position should frog will be positioned at
  D = 30, fix length of frog jump

the function should return 3, because the frog will be positioned as follows:
  - after the first jump, at position 10 + 30 = 40
  - after the second jump, at position 10 + 30 + 30 = 70
  - after the third jump, at position 10 + 30 + 30 + 30 = 100

Write an efficient algorithm for the following assumptions:
  - X, Y and D are integers within the range [1..1,000,000,000];
  - X ≤ Y.
*/

func main() {
  result := Solution(10, 85, 30)
  fmt.Println("result:",result)

  result := Solution(2, 52, 1000000)
  fmt.Println("result:",result)
}

func Solution(X int, Y int, D int) int {
    // NOTE:
    // Step: Y is supposed to be more than X, return fail if Y lower than equals to X
    // get total jump equals to round of difference of Y and X divided by D

    if Y <= X {
      return -1
    }

    totalJumpInDecimal := float64(Y-X) / float64(D)
    fixJumpCount := Round(totalJumpInDecimal)

    return fixJumpCount
}

func Round(number float64) int {
    lowerInt := int(number)
    remainder := (float64(lowerInt + 1) - number)

    if remainder == float64(1.0) {
        return lowerInt
    } else {
        return lowerInt + 1
    }
}

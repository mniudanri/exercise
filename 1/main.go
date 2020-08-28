package main
import "fmt"

// source exercise: codility 1.binary_gap

func main() {
  binary := 1025
  fmt.Printf("%08b\n", binary)
  fmt.Println(Solution(binary))

  binary = 654
  fmt.Printf("%08b\n", binary)
  fmt.Println(Solution(binary))
}

func Solution(N int) int {
	// NOTE:
  // step: remove all 0's in tail binary
	//        loop the binary length. if binary digit = 0, increase gap counter,
  //          else, get the highest gap counter value and reset counter value to 0

	maxGap := 0
	currentGap := 0

	for N > 0 && N%2 == 0 {
		N = N>>1
	}
	for N > 0 {
		if N%2 == 0 {
			currentGap++
		} else {
			if currentGap > maxGap {
				maxGap = currentGap
			}
			currentGap = 0
		}
		N = N>>1
	}

	return maxGap
}

package main

// import "fmt"
// func CubeSum(num int) int {
// 	sum := 0
// 	for num != 0 {
// 		digit := num % 10
// 		sum += digit * digit * digit
// 		num /= 10
// 	}
// 	return sum
// }

// func main() {
// 	Result := CubeSum(123)
// 	fmt.Println(Result)
// }

import (
	"fmt"
	"time"
)

func main() {
	myChannel := make(chan int)

	// Sending data into the channel (in a separate goroutine)
	go func() {
		for i := 1; i <= 5; i++ {
			myChannel <- i
		}
		// Close the channel to signal that no more data will be sent
		// close(myChannel)
	}()
	go func() {
		for j := 6; j <= 10; j++ {
			myChannel <- j
		}
		time.Sleep(1 * time.Second)
		close(myChannel)
	}()

	// Receiving and printing data from the channel
	for {
		num, ok := <-myChannel
		if !ok {
			// The channel is closed
			fmt.Println("Channel is closed.")
			break
		}
		fmt.Printf("Received: %d\n", num)
	}
}

package main

// import "fmt"
func CubeSum(num int) int {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num /= 10
	}
	return sum
}

// func main() {
// 	Result := CubeSum(123)
// 	fmt.Println(Result)
// }

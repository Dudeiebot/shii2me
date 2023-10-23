package Customerr

import (
	"errors"
	"fmt"
)

// error is the ididomatic way of handling abnormal behaviors sha
// custom error with structs and method
type areaError struct {
	err    string  // error description
	length float64 // length which caused the error
	width  float64 // width which caused the error
}

func (e *areaError) Error() string {
	return e.err
}

func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero" // append our err message here so that it will not overwrite the other
		}
	}
	if err != "" {
		return 0, &areaError{
			err:    err,
			length: length,
			width:  width,
		}
	}
	return length * width, nil
}

func main() {
	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		var areaError *areaError
		if errors.As(
			err,
			&areaError,
		) { // check out these more representation of error with AS and IS
			if areaError.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", areaError.length)
			}
			if areaError.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", areaError.width)
			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", area)
}

package main

// just having these repo here for my random testing and main also available here

import (
	"fmt"
	"testing"
)

func TestCubeSum(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{123, 36},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.num), func(t *testing.T) {
			got := CubeSum(tc.num)
			if got != tc.want {
				t.Fatalf("CubeSum() = %v; want %v", got, tc.want)
			}
		})
	}
}

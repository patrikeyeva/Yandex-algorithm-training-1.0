package main

import (
	"fmt"
)

func min(a, b int8) int8 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

func main() {

	var troom, tcond int8
	var mode string
	fmt.Scan(&troom, &tcond, &mode)
	switch mode {
	case "freeze":
		fmt.Print(min(troom, tcond))
	case "heat":
		fmt.Print(max(troom, tcond))
	case "auto":
		fmt.Print(tcond)
	case "fan":
		fmt.Print(troom)

	}

}

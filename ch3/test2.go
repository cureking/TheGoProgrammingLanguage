package main

import (
	"fmt"
)

func main() {
	const (
		a = 2
		b
		c = "110"
		d
		f = 15
	)

	fmt.Println(a, b, c, d, f)

	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	type Flags uint
	const (
		FlagUp Flags = 1 << iota
		FlagBroadcast
		FlagLoopback
		FlagPointToPoint
		FlagMulticast
	)
	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)

}

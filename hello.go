package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Print("Siddharth Mishra\n")
	fmt.Println("This is the time: ", time.Now())
	fmt.Println("My favorite number is", rand.Intn(13))
	fmt.Printf("Now you have %g problems\n", math.Sqrt(234))
	fmt.Println(math.Pi)
	fmt.Println(add(4, 5))
	fmt.Println(swap("siddharth", "mishra"))
	fmt.Println(split(13))
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

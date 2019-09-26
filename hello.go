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
	arrayExample()
	dictionaryExample()
	forloopExample()
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

func arrayExample() {
	arr := [10]int{1, 2}
	arr[7] = 34
	fmt.Println(arr)

	arr01 := []int{1, 2}
	arr01 = append(arr01, 45)
	fmt.Println(arr01)
}

func dictionaryExample() {
	vertices := make(map[string]int)
	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["dodecagon"] = 12
	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])
	fmt.Println(vertices["sad"])
	delete(vertices, "square")
	fmt.Println(vertices)
	delete(vertices, "sad")
	fmt.Println(vertices)
}

func forloopExample() {
	for index := 0; index < 5; index++ {
		fmt.Println(index)
	}

	fmt.Println("***While Loop")
	i := 0
	for i < 6 {
		fmt.Println(i)
		i++
	}

	fmt.Println("**** For loop with Array")
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}

	fmt.Println("**** For loop with Array")
	dictionary := make(map[string]string)
	dictionary["firstName"] = "siddharth"
	dictionary["last name"] = "mishra"

	for key, value := range dictionary {
		fmt.Println("key", key, "value", value)
	}
}

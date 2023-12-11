package main

import "fmt"

func fibonacci(num int64, calcs map[int64]int64) int64 {
	if num == 0 || num == 1 {
		return 1
	}
	calc, exists := calcs[num]

	if exists {
		return calc
	} else {
		calcs[num] = fibonacci(num-2, calcs) + fibonacci(num-1, calcs)
		return calcs[num]
	}
}

func fibonacciSerie(num int64) []int64 {
	var serie []int64
	var i int64 = 0
	for ; i <= num; i++ {
		next := fibonacci(i, make(map[int64]int64))
		serie = append(serie, next)
		//serie[i] = fibonacci(num, make(map[int]int))
	}
	return serie
}

func main() {
	//fmt.Printf("%v", fibonacciSerie(5))
	fmt.Println(fibonacciSerie(80))

}

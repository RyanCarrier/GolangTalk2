package main

import "sync"

func fibDyn(i int) int {
	//0 0
	//1 1
	//2 1
	a := 1
	b := 0
	for j := 0; j < i; j++ {
		a, b = a+b, a
	}
	return b
}

func fibRec(i int) int {
	//0 0
	//1 1
	//2 1
	if i < 2 {
		return i
	}
	return fibRec(i-1) + fibRec(i-2)
}

func seq(i int) int {
	a := fibDyn(i)
	b := fibDyn(i + 1)
	c := fibDyn(i + 2)
	return (a % b) % c
}

func multi(i, n int) int {
	values := make([]int, n)
	wg := sync.WaitGroup{}
	wg.Add(n)
	final := 0
	for j := 0; j < n; j++ {
		go func(k int) {
			values[k] = fibDyn(i + k)
			defer wg.Done()
		}(j)
	}
	wg.Wait()
	for j := 0; j < n; j++ {
		final %= values[j]
	}
	return final
}

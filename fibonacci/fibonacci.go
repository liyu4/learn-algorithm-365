// Package fibonacci provide implement
package fibonacci

func fibonacci(n int) int {
	if n == 1 || n== 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// if n ==  6
// 1  1  2  3  5  8{我们要找的数据}  13  21  34
/*
	f(5)+ f(4)

	f(4)  + f(3) +  f(3) + f(2)

	f(3) + f(2)  + f(1) + f(2) + f(2) + f(1) + f(2）

 	f(2) +  f(1) + f(2)  + f(1) + f(2) + f(2) + f(1) + f(2）

 	f(2) +  f(1) + f(2)  + f(1) + f(2) + f(2) + f(1) + 1

 	f(2) +  f(1) + f(2)  + f(1) + f(2) + f(2) + 2

 	f(2) +  f(1) + f(2)  + f(1) + f(2) + 3

 	f(2) +  f(1) + f(2)  + f(1) + 4

 	f(2) +  f(1) + f(2)  + 5

 	f(2) +  f(1) + 6

 	f(2) +  7

 	return 8
/*


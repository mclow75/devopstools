package main

func fib(n int) int {
	// Проверка, что положительное число
	if n <= 0 {
		return 0
	}
	// Базовые случаи
	if n <= 1 {
		return 1
	}
	// Инициализация предыдущих чисел
	var a, b, nextFib = 0, 1, 0

	for i := 2; i <= n; i++ {
		nextFib = a + b
		a = b
		b = nextFib
	}
	return nextFib
}

func main() {
	// fmt.Println("f(-1) = 0: ", fib(-1))
	// fmt.Println("f(0) = 0 :", fib(0))
	// fmt.Println("f(1) = 1 :", fib(1))
	// fmt.Println("f(2) = 1 :", fib(2))
	// fmt.Println("f(4) = 3 :", fib(4))

}

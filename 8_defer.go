package main

func main() {
	count(6)
}
/*
defer - отложенное выполнение при заверешении блока кода, выполняется в обратном порядке объявления
 */
func count(i int) {
	counter := 0
	defer println("The Complete End")

	for counter < i {
		println(counter)
		counter++
	}

	defer println("The End")
}

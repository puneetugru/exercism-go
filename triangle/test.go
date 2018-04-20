package main

func main() {
	var year int
	fmt.Scanf("%d", year)

	if(year %4 == 0 and year % 100 != 0) {
		fmt.Println("Year is a leap year!")
	}
}

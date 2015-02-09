package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func is_perfect_number(num int) {
	temp := 0
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			temp += i
		}
	}

	if temp == num {
		fmt.Println(num, " is perfect number")
	} else {
		fmt.Println(num, " is not perfect number")
	}
}

func main() {

	// open files
	file, Error := os.Open("numbers.txt")
	if Error != nil {
		fmt.Println("Open File Error.")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		is_perfect_number(number)
	}
}

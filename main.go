package main

import (
	"fmt"
	as "myproject/assignment"
	"myproject/service"
	"os"
	"strconv"
)

func main() {

	// validate user input
	if len(os.Args) != 2 {
		fmt.Println("Invalid Arguments")
		os.Exit(1)
	} else {
		// Convert count to an integer
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(2)
		}
		// Generate number
		for i := 1; i <= num; i++ {
			evenNumber := i * 2
			_ = as.FetchData(service.HttpAPIClient, evenNumber)
		}
	}
	fmt.Println("Done")
}

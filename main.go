package main

import (
	"fmt"
	"log"
	"os"

	"loyalty/loyal"
)

func main() {
	// Check command-line arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: loyalty <log_file_day1> <log_file_day2>")
		fmt.Println("Example: loyalty mock_data/log_day0.csv mock_data/log_day1.csv")
		os.Exit(1)
	}

	day1Path := os.Args[1]
	day2Path := os.Args[2]

	f0, err := os.Open(day1Path)
	if err != nil {
		log.Fatalf("failed to open %s: %v", day1Path, err)
	}
	defer f0.Close()

	f1, err := os.Open(day2Path)
	if err != nil {
		log.Fatalf("failed to open %s: %v", day2Path, err)
	}
	defer f1.Close()

	loyals, err := loyal.FindLoyalCustomers(f0, f1)
	if err != nil {
		log.Fatalf("FindLoyalCustomers error: %v", err)
	}
	fmt.Println("Loyal customers:", loyals)
}

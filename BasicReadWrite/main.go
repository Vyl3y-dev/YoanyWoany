package main

import (
	"BasicReadWrite/input"
	"BasicReadWrite/models"

	//"YoanyWoany/read"
	//"YoanyWoany/write"
	"fmt"
)

func main() {
	fmt.Println("YOANY IT IS WORKING 🍉❤️")

	//var allData [][]string // a slice of rows
	var people []models.Person

	for {
		cleaned := input.GetUserInput()
		person := models.Person{
			Name:  cleaned[0],
			Role:  cleaned[1],
			Title: cleaned[2],
		}
		people = append(people, person)
		//allData = append(allData, cleaned) // collect but don't write yet

		fmt.Print("Add another? (y/n): ")
		var choice string
		fmt.Scanln(&choice)
		if choice != "y" {
			break
		}
	}

	// Write all rows once, after loop
	// for _, row := range allData {
	// 	if err := write.WriteCSV(row); err != nil {
	// 		fmt.Println("Error writing CSV:", err)
	// 		return
	// 	}
	// }

	for _, p := range people {
		fmt.Printf("%s %s %s\n", p.Name, p.Role, p.Title)
	}

	//fmt.Println("File Written!")
	//read.ReadCSV("data/test.csv")
}

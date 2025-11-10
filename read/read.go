package read

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSV(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for _, row := range records {
		fmt.Println(row)
	}
}

package write

import (
	"encoding/csv"
	"os"
)

func WriteCSV(data []string) error {
	path := "data/test.csv"

	_, err := os.Stat(path)
	fileExists := !os.IsNotExist(err)

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if !fileExists {
		writer.Write([]string{"Name", "Role", "Title"})
	}

	return writer.Write(data)
}

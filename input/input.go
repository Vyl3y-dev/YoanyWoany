package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput() []string {
	fmt.Print("Enter comma-separated values: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	raw := scanner.Text()
	parts := strings.Split(raw, ", ")

	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	return parts

}

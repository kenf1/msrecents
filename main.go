package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kenf1/msrecents/logic"
)

func main() {
	//not req
	var status string

	err := godotenv.Load(".env")
	if err != nil {
		status = "production"
	} else {
		status = os.Getenv("STATUS")
	}

	//available options
	appDict := map[string]string{
		"word":       "Word",
		"excel":      "Excel",
		"powerpoint": "Powerpoint",
	}

	//prompt user select option
	selectedOption, err := logic.ShowSelect()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	switch selectedOption {
	case "word":
		logic.ProcessApp(status, appDict["word"])
	case "excel":
		logic.ProcessApp(status, appDict["excel"])
	case "powerpoint":
		logic.ProcessApp(status, appDict["powerpoint"])
	case "all":
		logic.ProcessAllApps(status, appDict)
	case "exit":
		fmt.Println("Goodbye...")
	default:
		fmt.Println(fmt.Errorf("invalid input: %s", selectedOption))
		return
	}
}

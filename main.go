package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kenf1/msrecents/logic"
)

func main() {
	//allow debug
	var status string

	if err := godotenv.Load(".env"); err != nil {
		status = "production"
	} else {
		status = os.Getenv("STATUS")
	}

	appDict := map[string]string{
		"word":       "Word",
		"excel":      "Excel",
		"powerpoint": "Powerpoint",
	}

	selectedOption, err := logic.ShowSelect()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//reduce repeat chunks
	wrapperfn := func(input string) {
		res, err := logic.GetFullPath([]string{appDict[input]})
		if err != nil {
			fmt.Println("error occurred", err)
		}

		if status == "production" {
			logic.PlistRemove(res)
			if err != nil {
				fmt.Println("error occurred", err)
			}
		} else {
			fmt.Println(res)
		}
	}

	switch selectedOption {
	case "word":
		wrapperfn("word")
	case "excel":
		wrapperfn("excel")
	case "powerpoint":
		wrapperfn("powerpoint")
	case "all":
		res, err := logic.GetFullPath(
			[]string{
				appDict["word"],
				appDict["excel"],
				appDict["powerpoint"],
			},
		)
		if err != nil {
			fmt.Println("error occurred", err)
		}

		if status == "production" {
			logic.PlistRemove(res)
			if err != nil {
				fmt.Println("error occurred", err)
			}
		} else {
			fmt.Println(res)
		}
	case "exit":
		fmt.Println("Goodbye...")
		return
	default:
		fmt.Println(fmt.Errorf("invalid input: %s", selectedOption))
		return
	}
}

package main

import (
	"fmt"

	logic "github.com/kenf1/msrecents/logic"
)

func main() {
	res, err := logic.GetFullPath([]string{"Word", "Excel"})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(res)
}

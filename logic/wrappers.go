package logic

import (
	"fmt"
)

// single app
func ProcessApp(status string, appName string) {
	res, err := GetFullPath([]string{appName})
	if err != nil {
		fmt.Println(err)
		return
	}

	HandleResult(status, res)
}

// all
func ProcessAllApps(status string, appDict map[string]string) {
	var paths []string

	for _, appName := range appDict {
		path, err := GetFullPath([]string{appName})
		if err != nil {
			fmt.Println(err)
			continue
		}
		paths = append(paths, path...)
	}

	HandleResult(status, paths)
}

// delete if status = production, else print chosen option full plist path
func HandleResult(status string, res interface{}) {
	if status == "production" {
		if PromptBool() {
			strPaths, ok := res.([]string)
			if !ok {
				fmt.Println("Error: expected []string but got", res)
				return
			}
			err := PlistRemove(strPaths)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Cancelled")
		}
	} else {
		fmt.Println(res)
	}
}

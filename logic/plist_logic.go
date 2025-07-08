package logic

import (
	"fmt"
	"os"
)

// accept list of app names, return full path for each accepted app name
func GetFullPath(appNames []string) ([]string, error) {
	//create full path helper function
	fpHelper := func(appname string) string {
		return fmt.Sprintf(
			"%s/Library/Containers/com.microsoft.%s/Data/Library/Preferences/com.microsoft.%s.securebookmarks.plist",
			os.Getenv("HOME"),
			appname,
			appname,
		)
	}

	var paths []string

	for _, app := range appNames {
		switch app {
		case "Word":
			paths = append(paths, fpHelper("Word"))
		case "Excel":
			paths = append(paths, fpHelper("Excel"))
		case "Powerpoint":
			paths = append(paths, fpHelper("Powerpoint"))
		default:
			return nil, fmt.Errorf("invalid input: %s", app)
		}
	}
	return paths, nil
}

// remove plist files (result of GetFullPath)
func PlistRemove(plistPaths []string) error {
	for _, path := range plistPaths {
		err := os.Remove(path)
		if err != nil {
			return fmt.Errorf("failed to remove %s: %w", path, err)
		}
	}
	return nil
}

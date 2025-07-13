package logic_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/kenf1/msrecents/logic"
	"github.com/stretchr/testify/assert"
)

// unit
func TestGetFullPath(t *testing.T) {
	type testCase struct {
		//input
		input []string

		//expected
		output []string
		err    string
	}

	//depends on user + test env
	curEnv := os.Getenv("HOME")

	t.Run("valid app names", func(t *testing.T) {
		tests := []testCase{
			{
				input:  []string{"Word"},
				output: []string{fmt.Sprintf("%s/Library/Containers/com.microsoft.Word/Data/Library/Preferences/com.microsoft.Word.securebookmarks.plist", curEnv)},
			},
			{
				input:  []string{"Excel"},
				output: []string{fmt.Sprintf("%s/Library/Containers/com.microsoft.Excel/Data/Library/Preferences/com.microsoft.Excel.securebookmarks.plist", curEnv)},
			},
			{
				input:  []string{"Powerpoint"},
				output: []string{fmt.Sprintf("%s/Library/Containers/com.microsoft.Powerpoint/Data/Library/Preferences/com.microsoft.Powerpoint.securebookmarks.plist", curEnv)},
			},
			{
				input: []string{"Word", "Excel", "Powerpoint"},
				output: []string{
					fmt.Sprintf("%s/Library/Containers/com.microsoft.Word/Data/Library/Preferences/com.microsoft.Word.securebookmarks.plist", curEnv),
					fmt.Sprintf("%s/Library/Containers/com.microsoft.Excel/Data/Library/Preferences/com.microsoft.Excel.securebookmarks.plist", curEnv),
					fmt.Sprintf("%s/Library/Containers/com.microsoft.Powerpoint/Data/Library/Preferences/com.microsoft.Powerpoint.securebookmarks.plist", curEnv),
				},
			},
		}

		for _, test := range tests {
			actual, err := logic.GetFullPath(test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.output, actual)
		}
	})

	t.Run("invalid app names", func(t *testing.T) {
		tests := []testCase{
			{input: []string{"neovim"}, err: "invalid input: neovim"}, //not accepted option
			{input: []string{"word"}, err: "invalid input: word"},     //lowercase name
		}

		for _, test := range tests {
			_, err := logic.GetFullPath(test.input)
			assert.Error(t, err)                //error present
			assert.EqualError(t, err, test.err) //match error message
		}
	})
}

// helper: create temp file
func createTempFile(t *testing.T, pattern string) string {
	t.Helper()

	f, err := os.CreateTemp("", pattern)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	// schedule cleanup if test fails before removal
	t.Cleanup(func() { os.Remove(f.Name()) })
	return f.Name()
}

func TestPlistRemoveSuccess(t *testing.T) {
	file1 := createTempFile(t, "file1.plist")
	file2 := createTempFile(t, "file2.plist")

	paths := []string{file1, file2}

	err := logic.PlistRemove(paths)
	assert.NoError(t, err, "PlistRemove returned error")

	for _, path := range paths {
		_, err := os.Stat(path)
		assert.True(t, os.IsNotExist(err), "File %s was not removed", path)
	}
}

func TestPlistRemoveFail(t *testing.T) {
	invalidFile := "/tmp/invalid-file.txt"

	paths := []string{invalidFile}

	err := logic.PlistRemove(paths)
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to remove /tmp/invalid-file.txt: remove /tmp/invalid-file.txt: no such file or directory")
}

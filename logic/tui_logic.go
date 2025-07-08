package logic

import huh "github.com/charmbracelet/huh"

// entry select box
func ShowSelect() (string, error) {
	var appOption string

	selectForm := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().Title("msrecents by github.com/kenf1"),
			huh.NewSelect[string]().
				Title("Select option:").
				Options(
					huh.NewOption("Word", "word"),
					huh.NewOption("Excel", "excel"),
					huh.NewOption("Powerpoint", "powerpoint"),
					huh.NewOption("All", "all"),
					huh.NewOption("Exit", "exit"),
				).
				Value(&appOption),
		),
	)

	if err := selectForm.Run(); err != nil {
		return "", err
	}

	return appOption, nil
}

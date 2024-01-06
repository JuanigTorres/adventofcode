package resources

import (
	"embed"
	"fmt"
)

//go:embed inputs/day/*.txt
var inputs embed.FS

func InputByDay(day int) []byte {
	filename := fmt.Sprintf("inputs/day/%d.txt", day)

	file, err := inputs.ReadFile(filename)
	if err != nil {
		fmt.Println("there was an error", err, "returning nil instead")
		return nil
	}

	return file
}

func InputStringByDay(day int) string {
	return string(InputByDay(day))
}

package gget

import (
	"io"
	"os"
	"sort"
)

// Join joins multiple parts of a file to 1 file
func Join(parts []string, output string) error {
	sort.Strings(parts)

	outputFile, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, defaultFileMode)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	for _, part := range parts {
		if err = copy(part, outputFile); err != nil {
			return err
		}
	}
	return nil
}

func copy(from string, to io.Writer) error {
	file, err := os.OpenFile(from, os.O_RDONLY, defaultFileMode)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(to, file)
	return err
}

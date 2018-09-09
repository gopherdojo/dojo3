package gget

import (
	"io"
	"os"
	"sort"
)

// Join joins multiple parts of a file to 1 file
func (g *GGet) Join() error {
	sort.Strings(g.partFiles)

	outputFile, err := os.OpenFile(g.fileName, os.O_CREATE|os.O_WRONLY, defaultFileMode)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	for _, part := range g.partFiles {
		if err = copy(part, outputFile); err != nil {
			return err
		}
	}
	for _, name := range g.partFiles {
		os.Remove(name)
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

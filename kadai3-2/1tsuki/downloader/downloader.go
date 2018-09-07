package downloader

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func Download(url *url.URL) (string, error) {
	// Get the data
	resp, err := http.Get(url.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	filename := filepath.Base(url.Path)
	// Create the file
	out, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return filename, nil
}

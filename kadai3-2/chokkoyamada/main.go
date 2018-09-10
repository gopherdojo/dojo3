package chokkoyamada

import (
	"os"
	"fmt"
	"path/filepath"
	"log"
	"context"

	"github.com/gopherdojo/dojo3/kadai3-2/chokkoyamada/download"
)

func main() {
	switch len(os.Args) {
	case 2:
		doDownload(os.Args[1])
	default:
		fmt.Fprintf(os.Stderr, "usage: %s URL\n", os.Args[0])
		os.Exit(1)
	}
}

func doDownload(url string) {
	filename := filepath.Base(url)
	if filename == "" {
		filename = "file"
	}
	w, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Could not create file %s: %s", filename, err)
	}
	defer w.Close()

	log.Printf("Downloading %s to %s", url, filename)
	d := download.New(url)
	ctx := context.Background()
	rng, err := d.GetContent(ctx, w)
	if err != nil {
		log.Fatalf("Could not donwload %s: %s", url, err)
	}
	log.Printf("Wrote %d bytes", rng.Length())
}

package cli_test

import (
	"image"
	"image/jpeg"
	"os"
	"testing"

	"github.com/gopherdojo/dojo3/kadai1/daikurosawa/cli"
	"github.com/gopherdojo/dojo3/kadai1/daikurosawa/option"
)

const (
	testImageFileName = "test_image.jpg"
	testDirName       = "test_dir"
	fileMode          = 0777
	exitCodeOK        = 0
)

type Mock struct {
}

// Convert mock
func (Mock) Convert(path string) error {
	return nil
}

var testFixtures = struct {
	dirName       string
	fromExtension string
	toExtension   string
}{
	testDirName,
	"jpg",
	"png",
}

func init() {
	var (
		x      = 0
		y      = 0
		width  = 100
		height = 50
	)

	// make directory
	if err := os.Mkdir(testDirName, fileMode); err != nil {
		panic(err)
	}

	// make test image
	img := image.NewRGBA(image.Rect(x, y, width, height))
	file, _ := os.Create(testDirName + "/" + testImageFileName)
	defer file.Close()
	if err := jpeg.Encode(file, img, nil); err != nil {
		panic(err)
	}
}

func TestCli_Run(t *testing.T) {
	defer func() {
		if exist := isExist(testDirName); exist {
			if err := os.RemoveAll(testDirName); err != nil {
				panic(err)
			}
		}
	}()

	option := &option.Option{
		DirName:       testFixtures.dirName,
		FromExtension: testFixtures.fromExtension,
		ToExtension:   testFixtures.toExtension,
	}
	convert := &Mock{}
	cli := cli.NewCli(convert, option)
	if exitCode := cli.Run(); exitCode != exitCodeOK {
		t.Fatal("failed: exit code is not success")
	}
}

func isExist(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

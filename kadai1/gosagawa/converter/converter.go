//Package converter provides convert image
package converter

import (
	"io/ioutil"
	"path/filepath"
)

//Converter is information for convert images
type Converter struct {
	inType   string
	outType  string
	dir      string
	dispLog  bool
	inPaths  []string
	outPaths []string
}

//NewConverter create new converter
func NewConverter(inType string, outType string, dir string, dispLog bool) *Converter {
	c := Converter{}
	c.inType = inType
	c.outType = outType
	c.dispLog = dispLog
	c.dir = dir

	c.setPath()
	return &c
}

//ConvertImage convert image
func (c *Converter) ConvertImage() {

	for k, path := range c.inPaths {
		ci := ConvertImage{c.outType, path, c.outPaths[k], c.dispLog}
		ci.ConvertImage()
	}
}

func (c *Converter) setPath() {

	c.inPaths = getConvertList(c.inType, c.dir)
	for _, path := range c.inPaths {
		outPath := getConvertToPath(c.outType, path)
		c.outPaths = append(c.outPaths, outPath)
	}
}

func getConvertList(imageType string, dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, getConvertList(imageType, filepath.Join(dir, file.Name()))...)
			continue
		}
		path := filepath.Join(dir, file.Name())
		if imageType == GetImageType(path) {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}

	return paths
}

func getConvertToPath(outType string, path string) string {
	return path[:len(path)-len(filepath.Ext(path))] + "." + outType
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"sync"
)

const (
	dlDoingName = ".download"
	outputDir   = "../_data/"
	ouputFile   = "logo.png"
	outputPath  = outputDir + ouputFile + dlDoingName
	dlSize      = 1024 * 5
)

var (
	urls = []string{
		"https://www.google.co.jp/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
	}
)

type DlFileHeader struct {
	ContentType string
	Size        int
}

type DlStatus struct {
	curNum   int
	maxNum   int
	fromSize int
	toSize   int
}

type File struct {
	url        string
	dlFileInfo DlFileHeader
	dlStatus   DlStatus
	path       string
	file       *os.File
	fileInfo   *os.FileInfo
	resp       *http.Response
	data       []byte
}

func main() {
	for _, url := range urls {
		run(url)
	}
}

var wg sync.WaitGroup

func run(url string) {

	f := File{
		path: outputPath,
	}

	f.url = url
	f.dlFileInfo = getDlFileInfo(url)
	f.getWriter().getFileInfo()
	defer f.file.Close()

	dlCurNum := 0
	for i := 0; i <= f.dlStatus.maxNum; i++ {
		wg.Add(1)
		dlCurNum++

		fmt.Println(f)
		go func(dlCurNum int) {
			fmt.Println("++++++++++>", f)
			f.incr(dlCurNum)
			f.req().save(f.dl())
			wg.Done()
		}(dlCurNum)
	}

	f.closing()
}

func getDlFileInfo(url string) DlFileHeader {
	resp, err := http.Head(url)
	if err != nil {
		fmt.Println(err)
	}

	maps := resp.Header
	length, err := strconv.Atoi(maps["Content-Length"][0])
	if err != nil {
		fmt.Println(err)
	}

	dlFileInfo := DlFileHeader{
		maps["Content-Type"][0],
		length,
	}

	return dlFileInfo
}

func (f *File) incr(dlCurNum int) {
	dlFileSize := f.dlFileInfo.Size
	fromSize := (dlCurNum - 1) * dlSize
	toSize := dlCurNum * dlSize
	if toSize > dlFileSize {
		toSize = dlFileSize
	}

	f.dlStatus = DlStatus{
		dlCurNum,
		int(dlFileSize / dlSize),
		fromSize,
		toSize,
	}
}

func (f *File) getFileInfo() *File {
	stat, err := f.file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	if stat.Name() != "" {
		if f.dlStatus.maxNum < int(stat.Size()) {
			log.Fatal("DL済みか、ファイルデータが取得できません")
		}
	}

	return f
}

func (f *File) req() *File {
	req, _ := http.NewRequest("GET", f.url, nil)
	req.Header.Add("Contents-Type", f.dlFileInfo.ContentType)
	// req.Header.Add("Pragma", "no-cache")
	// req.Header.Add("Cache-Control", "no-cache")
	// req.Header.Add("Expires", "-1")
	rangeHeader := "bytes=" + strconv.Itoa(f.dlStatus.fromSize) + "-" + strconv.Itoa(f.dlStatus.toSize)
	req.Header.Add("Range", rangeHeader)

	client := new(http.Client)
	resp, err := client.Do(req)

	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	f.resp = resp

	return f
}

func (f *File) dl() string {
	fmt.Println(f.resp.Body)
	b, err := ioutil.ReadAll(f.resp.Body)
	if err == nil {
		fmt.Println(err)
	}

	return string(b)
}

func (f *File) getWriter() *File {
	fl, err := os.OpenFile(f.path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	f.file = fl

	return f
}
func (f *File) save(data string) *File {

	ioutil.WriteFile(f.path, []byte(data), 0x777) // Write to the file i as a byte array
	wg.Done()
	// if f.dlStatus.curNum == 1 {
	// 	f.file.Seek(0, io.SeekStart)
	// } else {
	// 	f.file.Seek(int64(f.dlStatus.curNum*dlSize), io.SeekEnd)
	// }
	// f.file.WriteString(data)
	f.getFileInfo()

	return f
}

func cmd(header string, cmd string, args ...string) {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(header)
	fmt.Printf("%s", string(out))
}

func (f *File) closing() {
	err := os.Rename(outputPath, outputPath[:len(outputPath)-len(dlDoingName)])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("\n\n")
	cmd("<========= DataDir =========>", "ls", "-l", outputDir)
	fmt.Print("\n\n")
}

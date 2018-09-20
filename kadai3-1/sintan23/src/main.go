package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

const (
	debugFlg    = false
	Yes         = "y"
	No          = "n"
	src         = "../_data/word.txt"
	listMaxLine = 398
	viewMaxLine = 10
	timeUpSec   = 10
)

var (
	score = 0
)

func main() {
	fmt.Print("Would you like to start?(y/n): ")

	t := ""
	if debugFlg {
		t = Yes
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		t = scanner.Text()
	}
	fmt.Println("")

	if t != Yes {
		fmt.Println("input is ", t, ". Please enter y.")
		return
	}

	r := open(src)

	nlist := getRandLine()
	if debugFlg {
		fmt.Println("nlist -> ", nlist)
	}

	wlist := getList(r, nlist)
	if debugFlg {
		fmt.Println("wlist -> ", wlist)
	}

	fmt.Println("\n\n")
	fmt.Println("Ready Go!")
	fmt.Println("\n\n\n\n")
	timer0 := time.NewTimer(time.Second * 1)
	<-timer0.C

	timer()
	for {
		sw := selectWord(wlist)
		fmt.Println("▼ Please enter the displayed word.")
		fmt.Print(sw, " -> ")
		iw := ""
		for {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			iw = scanner.Text()
			if iw != "" {
				if check(sw, iw) {
					fmt.Println("  -> result: OK!", "\n")
					score++
				} else {
					fmt.Println("  -> result: NG", "\n")
				}
				break
			}
		}
	}
}

func timer() {
	timer := time.NewTimer(time.Second * timeUpSec)
	go func() {
		<-timer.C
		fmt.Println("\n\n")
		fmt.Println("Time Up!", "\n")
		timer2 := time.NewTimer(time.Second * 1)
		<-timer2.C
		fmt.Println("======================")
		fmt.Println("Score: ", score)
		os.Exit(0)
	}()
}

func check(selectWord string, inputWord string) bool {
	return selectWord == inputWord
}

func open(src string) io.Reader {
	r, err := os.Open(src)
	if err != nil {
		fmt.Println("Can't Open! src=", src)
	}

	return r
}

func getList(r io.Reader, nlist []int) []string {
	scanner := bufio.NewScanner(r)

	ln := 0
	vn := 0
	list := make([]string, 0, viewMaxLine-1)
	for scanner.Scan() {
		for _, nln := range nlist {
			// fmt.Println(nln)
			if ln == nln {
				// fmt.Println("break !! vn=", vn)
				list = append(list, scanner.Text())
				vn++
				break
			}
		}
		if vn == viewMaxLine {
			break
		}
		ln++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Scan error: %v\n", err)
	}

	return list
}

func getRandLine() []int {
	rand.Seed(time.Now().UnixNano())

	list := make([]int, 0, viewMaxLine-1)
	for i := 0; i < viewMaxLine; i++ {
		list = append(list, rand.Intn(listMaxLine))
	}

	return list
}

func selectWord(wordList []string) string {
	ln := len(wordList)
	rand.Seed(time.Now().UnixNano())

	word := ""
	n := rand.Intn(ln)
	for i := 0; i < ln; i++ {
		if n == i {
			word = wordList[i]
			break
		}
	}

	return word
}

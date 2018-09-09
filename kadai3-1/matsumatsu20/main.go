package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
	"math/rand"
	"time"
)

func main() {
	ch := input(os.Stdin)
	var count int64

	rand.Seed(time.Now().UnixNano())
	prefectures := []string{"hokkaido", "aomori", "iwate", "miyagi", "akita", "yamagata", "fukushima", "ibaraki", "tochigi", "gunma", "saitama", "chiba", "tokyo", "kanagawa", "niigata", "toyama", "ishikawa", "fukui", "yamanashi", "nagano", "gifu", "shizuoka", "aichi", "mie", "shiga", "kyoto", "osaka", "hyogo", "nara", "wakayama", "tottori", "shimane", "okayama", "hiroshima", "yamaguchi", "tokushima", "kagawa", "ehime", "kochi", "fukuoka", "saga", "nagasaki", "kumamoto", "oita", "miyazaki", "kagoshima", "okinawa"}

	go func() {
		<-time.After(30 * time.Second)

		fmt.Println("終了しました")
		fmt.Printf("あなたのスコア: %v\n", count)
		os.Exit(0)
	}()

	for {
		str := prefectures[rand.Intn(len(prefectures))]
		fmt.Print("> ")
		fmt.Println(str)

		select {
		case inputStr := <-ch:
			if str == inputStr {
				fmt.Println("ok!")
				count += 1
			} else {
				fmt.Println("miss!")
			}
		}
	}

}

func input(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		s := bufio.NewScanner(r)

		for s.Scan() {
			ch <- s.Text()
		}
	}()
	return ch
}

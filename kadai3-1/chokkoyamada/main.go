package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
	"time"
	"math/rand"
)

func input(r io.Reader) <-chan string {
	c := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			c <- s.Text()
		}
		close(c)
	}()
	return c

}

func main() {
	//制限時間
	limit := 15
	//問題の単語はこの中から毎回ランダムで抽出される
	question := []string{"apple", "orange", "pine", "strawberry", "banana"}
	rand.Seed(time.Now().Unix())

	score := 0

	fmt.Printf("タイピングゲーム開始! 表示された単語を入力してください。制限時間は%d秒です。\n", limit)
	t := time.After(time.Duration(limit) * time.Second)

	ch := input(os.Stdin)

M:
	for {
		s := question[rand.Intn(len(question))]
		fmt.Printf(">%s\n", s)
		select {
		case v := <-ch:
			fmt.Println(v)
			if s == v {
				fmt.Println("正解！")
				score += 1
			} else {
				fmt.Println("間違い")
			}

		case <-t:
			fmt.Println("タイムアップ！")
			break M
		}
	}

	fmt.Printf("タイピングゲーム終了です。スコアは%d点です。\n", score)

}

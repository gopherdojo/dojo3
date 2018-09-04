package shimastripe

import (
	"math/rand"
	"time"
)

var list = []string{"apple", "banana", "peach", "orange", "grape", "kiwi"}

func init() {
	rand.Seed(time.Now().Unix())
}

func shuffle(data []string) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

func GetWordList() []string {
	shuffle(list)
	return list
}

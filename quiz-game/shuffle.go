package quizgame

import (
	"math/rand"
	"time"
)

func shuffleSlice(slice [][]string) {
	rand.Seed(time.Now().UnixNano())
	for i := len(slice) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			slice[i], slice[j] = slice[j], slice[i]
	}
}
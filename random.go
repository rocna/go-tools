package go_tools

import (
	"math/rand"
	"time"
)

type random struct {
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var numberRunes = []rune("0123456789")

func (r random) RandString(n int) string {
	b := make([]rune, n)
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (r random) RandNumber(n int) string {
	b := make([]rune, n)
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = numberRunes[rand.Intn(len(numberRunes))]
	}
	return string(b)
}

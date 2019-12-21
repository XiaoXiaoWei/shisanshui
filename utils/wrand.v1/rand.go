package wrand

import (
	"math/rand"
	"sync"
	"time"
)

var pool sync.Pool

func init() {
	pool.New = func() interface{} {
		return rand.New(rand.NewSource(time.Now().UnixNano()))
	}
}

//GetInt 获取一个Int的随机数
func GetInt(n int) (num int) {
	rands := pool.Get().(*rand.Rand)
	num = rands.Intn(n)
	pool.Put(rands)
	return num
}

//GetInt32 获取一个Int的随机数
func GetInt32(n int32) (num int32) {
	rands := pool.Get().(*rand.Rand)
	num = rands.Int31n(n)
	pool.Put(rands)
	return num
}

//GetInt64 获取一个Int的随机数
func GetInt64(n int64) (num int64) {
	rands := pool.Get().(*rand.Rand)
	num = rands.Int63n(n)
	pool.Put(rands)
	return num
}

//OutOfOrder 乱序
func OutOfOrder(s []int) {
	for i := 0; i < len(s); i++ {
		j := GetInt(len(s) - 1)
		s[i%len(s)], s[j] = s[j], s[i%len(s)]
	}
}

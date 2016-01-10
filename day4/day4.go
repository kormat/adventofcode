package main

import (
	"container/list"
	"crypto/md5"
	"fmt"
	"github.com/kormat/adventofcode/util"
	"os"
)

type Result struct {
	prefixLen int
	coin      string
	idx       int
}

func main() {
	lines, ok := util.ReadFileArg(os.Args[1:])
	if !ok {
		os.Exit(1)
	}

	for i, key := range lines {
		c := make(chan Result)
		go mineCoin(key, c)
		fmt.Printf("%d. Key: %s.\n", i, key)
		for res := range c {
			fmt.Printf("  Prefix len: %d Coin: %s Coin idx: %d\n", res.prefixLen, res.coin, res.idx)
		}
	}
}

func mineCoin(key string, c chan Result) {
	targets := list.New()
	targets.PushFront("00000")
	targets.PushFront("000000")
	for i := 0; targets.Len() > 0; i++ {
		full_key := fmt.Sprintf("%s%d", key, i)
		byte_hash := md5.Sum([]byte(full_key))
		hash := fmt.Sprintf("%x", byte_hash)
		for t := targets.Front(); t != nil; t = t.Next() {
			tlen := len(t.Value.(string))
			if hash[:tlen] == t.Value {
				c <- Result{tlen, hash, i}
				targets.Remove(t)
			}
		}
	}
	close(c)
}

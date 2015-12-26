package main

import (
	"crypto/md5"
	"fmt"
	"github.com/kormat/adventofcode/util"
	"os"
)

func main() {
	lines, err := util.ReadFileArg()
	if err {
		os.Exit(1)
	}

	for i, key := range lines {
		coin, idx := mineCoin(key)
		fmt.Printf("%d. Key: %s. Coin: %s Coin idx: %d\n", i, key, coin, idx)
	}
}

func mineCoin(key string) (string, int) {
	target := "000000"
	for i := 0; ; i++ {
		full_key := fmt.Sprintf("%s%d", key, i)
		byte_hash := md5.Sum([]byte(full_key))
		hash := fmt.Sprintf("%x", byte_hash)
		if hash[:len(target)] == target {
			return hash, i
		}
	}
}

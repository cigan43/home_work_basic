package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func genData(gChannel chan int, limit int64) {
	t := time.Now()
	bInt := big.NewInt(limit)
	for {
		if time.Since(t).Seconds() > 60.0 {
			break
		}
		intRand, err := rand.Int(rand.Reader, bInt)
		if err != nil {
			break
		}
		gChannel <- int(intRand.Int64())
	}
	close(gChannel)
}

func aVg(inChan <-chan int, outChan chan float32) {
	var data, count int
	for i := range inChan {
		data += i
		count++
		if count > 10 {
			outChan <- float32(data) / 10
			data = 0
			count = 0
		}
	}
	close(outChan)
}

func main() {
	gChan := make(chan int)
	counted := make(chan float32)
	go genData(gChan, 1000)
	go aVg(gChan, counted)

	for b := range counted {
		fmt.Println(b)
	}
}

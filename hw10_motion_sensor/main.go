package main

import (
	"fmt"
	"math/rand"
	"time"
)

// - Создайте горутину для имитации считывания данных с сенсора и передачи их в канал. Данные должны считываться в течение 1 минуты.
// - Создайте горутину для обработки данных. Для каждых 10 полученных значений вычисляется среднее арифметическое и отправляется в канал с обработанными данными.
// - Главная горутина будет получать обработанные данные из канала и выводить их на экран.
// - Напишите юнит тесты на реализованные функции;

func genData(gChannel chan int) {
	t := time.Now()
	for {
		// fmt.Println(time.Since(t).Seconds(), "--------")
		if time.Since(t).Seconds() > 60.0 {
			break
		}
		gChannel <- rand.Intn(1000)
	}
	close(gChannel)

}

func aVg(inChan <-chan int, outChan chan float32) {
	var data, count int
	for i := range inChan {
		if count == 10 {
			outChan <- float32(data) / 10
			data = 0
			count = 0
		} else {
			data += i
			count += 1
		}
	}
	close(outChan)

}

func main() {
	gChan := make(chan int)
	counted := make(chan float32)
	go genData(gChan)
	go aVg(gChan, counted)

	for b := range counted {
		fmt.Println(b, "!!!!!!!")
	}

}

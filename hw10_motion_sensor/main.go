package main

import (
	"math/rand"
	"time"
)

// - Создайте горутину для имитации считывания данных с сенсора и передачи их в канал. Данные должны считываться в течение 1 минуты.
// - Создайте горутину для обработки данных. Для каждых 10 полученных значений вычисляется среднее арифметическое и отправляется в канал с обработанными данными.
// - Главная горутина будет получать обработанные данные из канала и выводить их на экран.
// - Напишите юнит тесты на реализованные функции;

func genData() <-chan int {
	gChannel := make(chan int)
	for {
		gChannel <- rand.Intn(1000)
		time.Sleep(60 * time.Second)
		close(gChannel)
	}
}

func aVg(in <-chan int) <-chan int {
	out := make(chan float32)
	var data, count int
	for i := range in {
		if count == 10 {
			out <- float32(data) / 10
			data = 0
			count = 0
		} else {
			data += i
		}
		close(out)
	}
}

func main() {
	out := make(chan float32)
	go genData()
	go aVg(out)

	for 
}

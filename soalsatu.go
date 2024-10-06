package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	//input batasawal dari user dan batasi untuk range atas
	var batasawal, batasi int
	fmt.Println("Input the lower limit: ")
	fmt.Scan(&batasawal)
	fmt.Println("Input the limit: ")
	fmt.Scan(&batasi)

	//membuat waitgroup untuk menyinkronisasi goroutines
	var wg sync.WaitGroup

	//set threads menjadi 7
	threads := 7

	//menghitung total range untuk tiap goroutines
	totalrange := batasi - batasawal + 1
	chunkSize := totalrange / threads

	//channel untuk collect prime numbers dari goroutines
	primeChannel := make(chan int, batasi-batasawal+1)

	//fungsi untuk mengecek pada range tersebut apakah ada angka prima atau tidak
	isPrime := func(n int) bool {
		if n <= 1 {
			return false
		}
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	//goroutines untuk mencari angka prima dalam range
	for i := 0; i < threads; i++ {
		wg.Add(1)

		//menghitung range awal dan akhir dari tiap goroutine
		start := batasawal + i*chunkSize
		end := batasawal + (i+1)*chunkSize - 1

		if i == threads-1 {
			end = batasi
		}

		go func(start, end int) {
			//menandai goroutine selesai ketika seluruh task terpenuhi
			defer wg.Done()
			for j := start; j <= end; j++ {
				if isPrime(j) {
					//menyimpan bilangan prima ke channel
					primeChannel <- j
				}
			}
		}(start, end)
	}
	//menutup channel ketika seluruh task selesai
	go func() {
		wg.Wait()
		close(primeChannel)
	}()

	//collect dan meprint bilangan prima
	fmt.Println("Prime numbers found:")
	for prime := range primeChannel {
		fmt.Println(prime)
	}

}

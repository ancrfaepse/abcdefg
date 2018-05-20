package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("処理開始")
	fmt.Println("通常処理コール")

	fmt.Println("通常コール")
	serialno()

	fmt.Println("並列処理コール")
	go serialno()
	time.Sleep(1 * time.Second)

	fmt.Println("処理終了")

}

func serialno() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

package text

import "fmt"

func T3() {
	ch := make(chan int, 10)
	for i := 1; i < 11; i++ {
		go func() {
			fmt.Println("这是输出的话。")
			ch <- 0
		}()
	}
	for i := 1; i < 11; i++ {
		<-ch
	}
}

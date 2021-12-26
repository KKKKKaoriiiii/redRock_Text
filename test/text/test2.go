package text

import "fmt"

func T2() {
	var x = 0
	ch := make(chan int)
	go func() {
		fmt.Println("下山的路又堵起了")
		ch <- x
	}()
	<-ch
}

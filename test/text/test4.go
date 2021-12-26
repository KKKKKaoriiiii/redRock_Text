package text

import "fmt"

func T4() {
	n := 1000000
	var flag [1000001]bool
	ch := make(chan int, 3)
	f := true
	for t := range flag {
		flag[t] = true
	}
	for i := 0; i < 2; i++ {
		flag[i] = false
	}
	go func() {
		for i := 2; i <= n/3; i++ {
			for j := 2; j < i; j++ {
				if flag[j] == true {
					if i%j == 0 {
						flag[i] = false
						f = false
						break
					}
				}
			}
			if f {
				flag[i] = true
			}
			f = true
		}
		ch <- 0
	}()
	go func() {
		for i := n/3 + 1; i <= n/3*2; i++ {
			for j := 2; j < i; j++ {
				if flag[j] == true {
					if i%j == 0 {
						flag[i] = false
						f = false
						break
					}
				}
			}
			if f {
				flag[i] = true
			}
			f = true
		}
		ch <- 0
	}()
	go func() {
		for i := n/3*2 + 1; i <= n; i++ {
			for j := 2; j < i; j++ {
				if flag[j] == true {
					if i%j == 0 {
						flag[i] = false
						f = false
						break
					}
				}
			}
			if f {
				flag[i] = true
			}
			f = true
		}
		ch <- 0
	}()
	for i := 0; i < 3; i++ {
		<-ch
	}
	for i, t := range flag {
		if t {
			fmt.Println(i)
		}
	}
}

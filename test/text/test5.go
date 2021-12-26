package text

import "fmt"

type door struct {
	ladder bool
	sign   int
}

func T5() {

	var (
		aWei struct {
			floor int
			room  int
		}
		n, m int
	)

	fmt.Scan(&n, &m)
	room := make([][]door, n)
	for i := range room {
		room[i] = make([]door, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&room[i][j].ladder, &room[i][j].sign)
		}
	}
	aWei.floor = 0
	fmt.Scan(&aWei.room)

	sum := room[aWei.floor][aWei.room].sign
	for {
		aWeiRoom:=room[aWei.floor][aWei.room]
		s:=aWeiRoom.sign
		for {
			if aWeiRoom.ladder {
				s--
				if s == 0{
					break
				}
			}
			if aWei.room != m-1 {
				aWei.room++
				aWeiRoom = room[aWei.floor][aWei.room]
			}else {
				aWei.room = 0
				aWeiRoom = room[aWei.floor][0]
			}
		}
		aWei.floor++
		if aWei.floor == n {
			break
		}
		sum += room[aWei.floor][aWei.room].sign
	}
	fmt.Println(sum)
}

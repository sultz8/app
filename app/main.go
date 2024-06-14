package main

import (
    "fmt"
)

func main() {
	l := &List{
		nil,
		nil,
		0,
	}
	
	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		err := l.PushFront(v)
		if err != nil {
			fmt.Println(err)
		}
	}
	
	l.Print()
	
}
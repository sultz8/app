package main

import (
    "errors"
    "fmt"
	"github.com/sultz8/app/models"
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
	fmt.Println(l.size)

	for l.size != 0 {
		fmt.Println(l.popBack())
	}

	_, err := l.popBack()
	var emptyListErr *EmptyListError
	if err != nil {
		if errors.As(err, &emptyListErr) {
			fmt.Println(err)
		}
	}

	s := &models.Stack{}
	fmt.Println(s)
}
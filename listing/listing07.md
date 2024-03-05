```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Output:
Числа от 1 до 8 в произвольном порядке, затем будет выводится 0 бесконечно.
Каналы a и b будут закрыты, после того как в функциях asChan все значения будут переданы в каналы.
Каналы a и b передаются функции merge, в функции горутины, созданной в функции merge, в бесконечном
цикле, с помощью оператора select, читаются данные из каналов a и b, и передаются в канал c.
При попытке получения данных из закрытого канала будет немедленно получено нулевое значение канала.
Поэтому после закрытия каналов a и b, в бесконечном цикле с оператором select в канал c будут
передаваться нулевые значения каналов a и b.
```
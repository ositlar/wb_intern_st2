```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ
```
Output:
nil
false

Функция foo() возвращает значение типа интерфейса, возвращаемое значение равно nil, но переменная типа интерфейса будет равна nil, только когда nil будет равны и значение переменной и динамический тип интерфейс
```
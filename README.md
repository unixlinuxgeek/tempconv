#### Упражнение 2.1.

Добавьте в пакет ```tempconv``` типы, константы и функции для работы с температурой по шкале Кельвина,
в которой нуль градусов соответствует температуре ```-273.15 °C```,
а разница температур в 1К имеет ту же величину, что и ```1°С```.

Решение задачи:

Устанавливаем нашу библитеку:
```shell
go get github.com/unixlinuxgeek/tempconv
```

Создаем файл go (например app.go):
```go
package main

import (
	"fmt"
	"github.com/unixlinuxgeek/tempconv"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		temp, err := strconv.Atoi(os.Args[1])
		if err != nil {
			os.Exit(1)
		}
		c := tempconv.FToC(tempconv.Fahrenheit(temp))
		fmt.Fprintf(os.Stdout, "%d градуса Фаренгейта равно %.2f градусов по Цельсию\n", temp, c)
	} else {
		fmt.Println("Не зедан аргумент!!!")
	}
}
```

Запускаем наш пример с аргументом 1:
```shell
$ go run ./app.go 1
1 градуса Фаренгейта равно -17.22 градусов по Цельсию
```

Запускаем наш пример с аргументом 5:
```shell
$ go run ./app.go 5
5 градуса Фаренгейта равно -15.00 градусов по Цельсию
```

##### Запуск теста с подробным выводом: 

```bash
go test -v
```

Увидим что то похожее на это:
```shell
$ go test -v
=== RUN   TestKelvinToC
    tempconv_test.go:16: Test TestKelvinToC is Passed!!!  2.044943 °K == -271.105057 °C
--- PASS: TestKelvinToC (0.00s)
PASS
ok      a       0.001s
```

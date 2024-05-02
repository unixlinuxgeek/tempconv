### Упражнение 2.1.

Добавьте в пакет ```tempconv``` типы, константы и функции для работы с температурой по шкале Кельвина,
в которой нуль градусов соответствует температуре ```-273.15 °C```,
а разница температур в 1К имеет ту же величину, что и ```1°С```.

Решение:

#### Шаг 1 

Создаем модуль:
```shell
go mod init abc
```

Выведет:
```
go: creating new go.mod: module abc
```


#### Шаг 2:

Устанавливаем нашу библитеку:
```shell
go get github.com/unixlinuxgeek/tempconv
```

#### Шаг 3:

Создаем файл go (например app.go) со следующим содержимым:
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
		c := tempconv.FahToCel(tempconv.Fahrenheit(temp))
		fmt.Fprintf(os.Stdout, "%d градуса Фаренгейта равно %.2f градусов по Цельсию\n", temp, c)
	} else {
		fmt.Println("Не зедан аргумент!!!")
	}
}
```

#### Шаг 4

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

#### Шаг 5 (необязательное)

```shell
git clone https://github.com/unixlinuxgeek/tenpconv
```

```shell
cd tenpconv
```

#### Запуск теста с подробным выводом: 

```bash
go test -v
```

`Увидим` что то похожее на это:
```shell
$ go test -v
=== RUN   TestKelToCel
    tempconv_test.go:19: Test TestKelToCel is Passed!!!  5.13 °K equal -268.02 °C
--- PASS: TestKelToCel (0.00s)
=== RUN   TestKelToFah
    tempconv_test.go:34: Test TestKelToFah is Passed!!!  3.89 °K equal -452.67 °F
--- PASS: TestKelToFah (0.00s)
=== RUN   TestCelToFah
    tempconv_test.go:48: Test TestCelToFah is Passed!!! 1.36 °C equal 34.45 °F
--- PASS: TestCelToFah (0.00s)
=== RUN   TestCelToKel
    tempconv_test.go:61: Test TestCelToKel is Passed!!! 8.90 °C equal 282.05 °K
--- FAIL: TestCelToKel (0.00s)
=== RUN   TestFahToCel
    tempconv_test.go:76: Test TestFahToCel is Passed!!! 5.63 °F equal -14.65 °C
--- PASS: TestFahToCel (0.00s)
=== RUN   TestFahToKel
    tempconv_test.go:90: Test TestFahToKel is Passed!!! 6.79 °F equal 259.14 °K
--- PASS: TestFahToKel (0.00s)
FAIL
exit status 1
FAIL    github.com/unixlinuxgeek/tempconv       0.002s
```

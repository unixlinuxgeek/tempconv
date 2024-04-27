![The Go Programming Language](https://raw.githubusercontent.com/unixlinuxgeek/logos/main/ISBN/9780134190570/128x128.png)

### Упражнение 2.1.

Добавьте в пакет ```tempconv``` типы, константы и функции для работы с температурой по шкале Кельвина,
в которой нуль градусов соответствует температуре ```-273.15 °C```,
а разница температур в 1К имеет ту же величину, что и ```1°С```.

### Запуск теста с подробным выводом: 

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
// Запуск теста с подробным выводом: go test -v

package tempconv

import (
	"github.com/unixlinuxgeek/floatgen" // Используем вспомогательную библиотеку для генерации чисел с плав. запятой
	"testing"
)

func TestKelvinToC(t *testing.T) {
	rnd := floatgen.GenRan(1, 9)
	o := rnd - 273.15

	kc := KToC(Kelvin(rnd))
	c := Celsius(o)

	if kc == c {
		t.Logf("Test TestKelvinToC is Passed!!!  %f °K == %f °C", rnd, o)
	} else {
		t.Fatal("Test TestKelvinToC Not Passed")
	}
}

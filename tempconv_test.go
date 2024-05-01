// Запуск теста с подробным выводом: go test -v

package tempconv

import (
	"fmt"
	"github.com/unixlinuxgeek/floatgen" // Используем вспомогательную библиотеку для генерации чисел с плав. запятой
	"os"
	"testing"
)

// Checking  Kelvin to Celsius (KelToC) converting
func TestKelToC(t *testing.T) {
	rnd := getFltRnd()
	o := rnd - 273.15

	kc := KToC(Kelvin(rnd))
	c := Celsius(o)

	if kc == c {
		t.Logf("Test TestKelvinToC is Passed!!!  %.2f °K == %.2f °C", rnd, o)
	} else {
		t.Fatal("Test TestKelvinToC Not Passed")
	}
}

// Checking Celsius to Fahrenheit (CToF) converting
func TestCToFar(t *testing.T) {
	rnd := getFltRnd()

	f := Fahrenheit((rnd * 9 / 5) + 32)

	cf := CToF(Celsius(rnd))
	if cf == f {
		fmt.Fprintf(os.Stdout, "Test Passed!!! %.2f °C equal %.2f °F", rnd, f)
	} else {
		fmt.Fprintf(os.Stdout, "Test Failed: %.2f °C not equal %f.2 °F", rnd, f)
	}
}

// Generation random value from 1 to 9
func getFltRnd() float64 {
	return floatgen.GenRan(1, 9)
}

// Запуск теста с подробным выводом: go test -v

package tempconv

import (
	"fmt"
	"github.com/unixlinuxgeek/floatgen" // Используем вспомогательную библиотеку для генерации чисел с плав. запятой
	"os"
	"testing"
)

// Checking  Kelvin to Celsius (KelToCel) converting
func TestKelToCel(t *testing.T) {
	rnd := getFltRnd()
	o := rnd - 273.15

	kc := KelToCel(Kelvin(rnd))
	c := Celsius(o)

	if kc == c {
		t.Logf("Test TestKelvinToC is Passed!!!  %.2f °K == %.2f °C", rnd, o)
	} else {
		t.Fatal("Test TestKelvinToC Not Passed")
	}
}

// Checking Celsius to Fahrenheit (CelToFah) converting
func TestCelToFar(t *testing.T) {
	rnd := getFltRnd()

	f := Fahrenheit((rnd * 9 / 5) + 32)

	cf := CelToFar(Celsius(rnd))
	if cf == f {
		fmt.Fprintf(os.Stdout, "Test Passed!!! %.2f °C equal %.2f °F", rnd, f)
	} else {
		fmt.Fprintf(os.Stdout, "Test Failed: %.2f °C not equal %f.2 °F", rnd, f)
	}
}

// Checking Fahrenheit to Celsius (FahToCel) converting
func TestFahToCel(t *testing.T) {
	rnd := getFltRnd()
	f := Fahrenheit(rnd)

	// converting Fahrenheit to Celsius formula
	fToc := Celsius((f - 32) * 5 / 9)

	if fToc == FahToCel(Fahrenheit(rnd)) {
		fmt.Fprintf(os.Stdout, "Test Passed!!! %.2f °F equal %.2f °C", rnd, fToc)
	} else {
		fmt.Fprintf(os.Stdout, "Test Failed: %.2f °F not equal %.2f.2 °C", rnd, fToc)
	}
}

// Generation random value from 1 to 9 (without zero!!!)
func getFltRnd() float64 {
	return floatgen.GenRan(1, 9)
}

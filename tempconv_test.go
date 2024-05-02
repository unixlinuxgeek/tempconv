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
		t.Logf("Test TestKelToC is Passed!!!  %.2f °K == %.2f °C", rnd, o)
	} else {
		t.Fatal("Test TestKelToC Not Passed")
	}
}

// Kel -> Fah
// (32K − 273.15) × 9/5 + 32 = -402.1°F
func TestKelToFah(t *testing.T) {
	rnd := getFltRnd()
	o := (rnd-273.15)*9/5 + 32

	conv := KelToFah(Kelvin(rnd))
	f := Fahrenheit(o)

	if f == conv {
		t.Logf("Test TestKelToFah is Passed!!!  %.2f °K == %.2f °F\n", rnd, conv)
	} else {
		t.Fatalf("Test Failed %.2f °K  not equal %.2f °F \n", f, conv)
	}
}

// Checking Celsius to Fahrenheit (CelToFah) converting
func TestCelToFah(t *testing.T) {
	rnd := getFltRnd()

	f := Fahrenheit((rnd * 9 / 5) + 32)

	cf := CelToFah(Celsius(rnd))
	if cf == f {
		fmt.Fprintf(os.Stdout, "Test Passed!!! %.2f °C equal %.2f °F", rnd, f)
	} else {
		t.Fatalf("Test Failed: %.2f °C not equal %f.2 °F", rnd, f)
	}
}

// Checking Celsius to Kelvin (CelToKel) converting
func TestCelToKel(t *testing.T) {
	rnd := getFltRnd()
	k := Kelvin((rnd-273.15)*9/5 + 32)
	ck := CelToKel(Celsius(rnd))
	if k == ck {
		fmt.Fprintf(os.Stdout, "Test Passed!!! %.2f °C equal %.2f °K", rnd, k)
	} else {
		t.Fatalf("Test Failed: %.2f °C not equal %f.2 °K", rnd, k)
	}
}

// Checking Fahrenheit to Celsius (FahToCel) converting
func TestFahToCel(t *testing.T) {
	rnd := getFltRnd()
	f := Fahrenheit(rnd)

	// converting Fahrenheit to Celsius formula
	fToC := Celsius((f - 32) * 5 / 9)

	if fToC == FahToCel(Fahrenheit(rnd)) {
		fmt.Fprintf(os.Stdout, "Test Passed!!! %.2f °F equal %.2f °C", rnd, fToC)
	} else {
		t.Fatalf("Test Failed: %.2f °F not equal %.2f.2 °C", rnd, fToC)
	}
}

// Checking Fahrenheit to Kelvin (FahToKel) converting
func TestFahToKel(t *testing.T) {
	rnd := getFltRnd()

	//  far to kel formula
	k := Kelvin((rnd-32)*5/9 + 273.15)

	if k == FahToKel(Fahrenheit(rnd)) {

	} else {
		t.Fatalf("Test Failed: %.2f °F not equal %.2f °K", rnd, k)
	}
}

// Generation random value from 1 to 9 (without zero!!!)
func getFltRnd() float64 {
	return floatgen.GenRan(1, 9)
}

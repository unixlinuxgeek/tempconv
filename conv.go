package tempconv

// CelToFar преобразует температуру по Цельсию в температуру по Фаренгейту.
func CelToFar(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// CellToKel преобразует температуру по Цельсию в температуру по Кельвину
func CelToKel(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

// FahToCel преобразует температуру по Фаренгейту в температуру по Цельсию
func FahToCel(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// FahToKel преобразует температуру по Фаренгейту в температуру по Кельвину
func FahToKel(f Fahrenheit) Kelvin {
	return Kelvin((f-32)*5/9 + 273.15)
}

// KelToCel преобразует температуру по Кельвину в температуру по Цельсию
func KelToCel(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// KelToFah преобразует температуру по Кельвину в температуру по Фаренгейту
func KelToFah(k Kelvin) Fahrenheit {
	return Fahrenheit((k-273.15)*9/5 + 32)
}

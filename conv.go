package tempconv

// CelToFar преобразует температуру по Цельсию в температуру по Фаренгейту.
func CelToFar(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FahToCel преобразует температуру по Фаренгейту в температуру по Цельсию
func FahToCel(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// KToC преобразует температуру по Кельвину в температуру по Цельсию
func KelToCel(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

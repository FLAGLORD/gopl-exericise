package tempconv

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

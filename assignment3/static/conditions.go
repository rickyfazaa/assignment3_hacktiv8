package static

func WaterCondition(waterValue int) string {
	var waterStatus string
	if waterValue <= 5 {
		waterStatus = "Aman"
	} else if waterValue >= 6 && waterValue <= 8 {
		waterStatus = "Siaga"
	} else {
		waterStatus = "Bahaya"
	}
	return waterStatus
}

func WindCondition(windValue int) string {
	var windStatus string
	if windValue <= 6 {
		windStatus = "Aman"
	} else if windValue >= 7 && windValue <= 15 {
		windStatus = "Siaga"
	} else {
		windStatus = "Bahaya"
	}
	return windStatus
}

package utils

func Hitungratarata(scores []float64) float64 {
	var total float64
	for _, score := range scores {
		total += score
	}
	return total / float64(len(scores))
}

func Celstofahrenheit(cel float64) float64 {
	celsi := 1.8*cel + 32
	return celsi
}

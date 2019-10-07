package teste

import (
	"fmt"
	"strconv"
)

//Media e uma funcao para calcular media
func Media(nums ...float64) float64 {
	total := 0.0

	for _, num := range nums {
		total += num
	}

	media := total / float64(len(nums))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)

	return mediaArredondada
}

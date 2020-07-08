package utils

import (
	"math"
)

// 平均数
func Avg(input ...float64) float64 {
	if len(input) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range input {
		sum += v
	}

	return sum / float64(len(input))
}

// 计算方差
func Variance(input ...float64) float64{
	avg := Avg(input...)
	sum := 0.0
	for _,v := range input{
		sum += math.Pow((v-avg),2.0)
	}
	return sum/float64(len(input))
}

// 计算标准差
func Stdev(input ...float64) float64 {
	// 算方差
	vari := Variance(input...)
	return math.Sqrt(vari)
}

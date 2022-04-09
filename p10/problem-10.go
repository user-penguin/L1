/*Package p10
Дана последовательность температурных колебаний:
-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/
package p10

import (
	"strconv"
	"strings"
)

func Run() {
	allTemp := []float64{-25.4, -27.0, 13.0, -8, 19.0, 15.5, 24.5, 0, -21.0, 32.5, 8, 1, 11, -11}
	groups := groupByInterval(allTemp, 10)
	var result string
	for id, temps := range groups {
		strId := strconv.Itoa(id)
		result = result + strId + ":{"
		for _, temp := range temps {
			result = result + strconv.FormatFloat(temp, 'f', 1, 64) + ", "
		}
		result = strings.TrimSuffix(result, ", ")
		result = result + "}, "
	}
	result = strings.TrimSuffix(result, ", ")
	println(result)
}

func groupByInterval(items []float64, interval float64) map[int][]float64 {
	group := make(map[int][]float64)
	for _, item := range items {
		switch {
		case item < 0:
			key := int((item-10)/interval) * 10
			group[key] = append(group[key], item)
		default:
			key := int(item/interval) * 10
			group[key] = append(group[key], item)
		}
	}
	return group
}

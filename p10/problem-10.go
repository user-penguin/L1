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
	allTemp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)
	for _, temp := range allTemp {
		key := int(temp/10) * 10
		groups[key] = append(groups[key], temp)
	}
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

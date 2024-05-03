package main

import "strconv"

func main() {
	operations := []string{"5", "2", "C", "D", "+"}

	println(calPoints(operations))
}

func calPoints(operations []string) int {
	record := []int{}
	for i := 0; i < len(operations); i++ {
		if operations[i] == "+" {
			record = append(record, record[len(record)-1]+record[len(record)-2])
		} else if operations[i] == "D" {
			record = append(record, 2*record[len(record)-1])
		} else if operations[i] == "C" {
			record = record[:len(record)-1]
		} else {
			x, _ := strconv.Atoi(operations[i])
			record = append(record, x)
		}
	}
	ans := 0
	for _, v := range record {
		ans += v
	}
	return ans
}

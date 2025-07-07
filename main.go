package main

import "fmt"

var intesection_list []int
var a []int
var b []int
var c []int
var result_a []int
var result_b []int

func uniq_value(list []int) []int {
	sort_list := []int{}
	if len(sort_list) == 0 {
		sort_list = append(sort_list, list[0])
	}
	for i := 1; i < len(list); i++ {
		var count int = 0
		for j := 0; j < len(sort_list); j++ {
			if sort_list[j] == list[i] {
				count += 1
			}
		}
		if count == 0 {
			sort_list = append(sort_list, list[i])
		}
	}
	return sort_list
}

func intesection(list_a []int, list_b []int) []int {
	for i := 0; i < len(list_a); i++ {
		for j := 0; j < len(list_b); j++ {
			if list_a[i] == list_b[j] {
				intesection_list = append(intesection_list, list_b[j])
			}

		}
	}
	return uniq_value(intesection_list)
}

func main() {
	a = []int{5, 1, 2, 5}
	b = []int{4, 2, 5, 1, 1, 2}

	result_a = append(uniq_value(a))
	result_b = append(uniq_value(b))
	c = append(result_a, result_b...)

	fmt.Println(result_a, result_b)
	fmt.Println(intesection(a, b))
	fmt.Println(uniq_value(c))
}

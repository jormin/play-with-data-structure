package main

import "fmt"

func main() {
	// params := []int{1, 3, 4, 1, 2, 3}
	params := []interface{}{1, "3", 4, 1, 2, 3}
	i := []interface{}{}
	for _, v := range params {
		i = append(i, v)
	}
	i = removeRepeat(i...)
	fmt.Printf("%+v\n", i)
}

func removeRepeat(params ...interface{}) []interface{} {
	m := map[interface{}]interface{}{}
	var res []interface{}
	for _, v := range params {
		if _, ok := m[v]; !ok {
			res = append(res, v)
			m[v] = 1
		}
	}
	return res
}

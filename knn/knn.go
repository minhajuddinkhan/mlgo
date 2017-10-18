package knn

import "fmt"

type knn struct {
	classes  []int
	dataSets map[int][]int
	k        int
	input    int
}

func Initialize(classes []int, dataSets map[int][]int, K int, input int) *knn {

	return &knn{classes, dataSets, K, input}
}

func (k *knn) Run() {

	classes := make(map[int][]int)
	for key, v := range k.dataSets {
		classes[key] = euc(k.input, v)

	}

	fmt.Println("classes", classes)
	results := []map[string]int{}

	for key, v := range classes {
		for _, i := range v {
			obj := make(map[string]int)
			obj["value"] = i
			obj["key"] = key
			results = append(results, obj)
		}
	}

	for index := (len(results) - 1); index >= 0; index-- {
		for j := 1; j <= index; j++ {
			if results[j-1]["value"] > results[j]["value"] {
				temp := results[j-1]
				results[j-1] = results[j]
				results[j] = temp
			}
		}
	}

	i := 0
	candidates := make(map[int]int)
	for i < k.k {
		if candidates[results[i]["key"]] == 0 {
			candidates[results[i]["key"]] = 1
		} else {
			candidates[results[i]["key"]]++
		}
		i++
	}

	var max, clusterKey int
	for key, v := range candidates {
		max = v
		clusterKey = key
		if max > v {
			max = v
			clusterKey = key

		}
	}
	k.dataSets[clusterKey] = append(k.dataSets[clusterKey], k.input)
	fmt.Println(k.dataSets)

}

func min(x []int) int {
	m := 0
	for i := 0; i < len(x); i++ {
		if x[i] < x[m] {
			m = i
		}
	}
	return m
}

func euc(data int, dataSet []int) []int {

	differences := []int{}

	for _, ctr := range dataSet {
		diff := data - ctr
		if diff < 0 {
			diff = -diff
		}
		differences = append(differences, diff)
	}
	return differences
}

func sort(ar []int) {
	for i := 0; i < len(ar); i++ {
		min := i
		for j := i + 1; j < len(ar); j++ {
			if ar[j] < ar[min] {
				min = j
			}
		}
		temp := ar[i]
		ar[i] = ar[min]
		ar[min] = temp
	}
}

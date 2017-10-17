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
	fmt.Println("cluster", k.dataSets)
	fmt.Println("classes ", (classes))

	candidates := make(map[int][]int)
	selected := []int{}
	for x := 0; x < k.k; x++ {
		for key, v := range classes {
			if len(v) > 0 {

				i := min(v)
				candidates[key] = append(candidates[key], v[i])
				v = append(v[:i], v[i+1:]...)
			}
			classes[key] = v
		}
		fmt.Println("cdd", candidates)

		//j := min(candidates)
		//selected = append(selected, candidates[j])
		//candidates = append(candidates[:j], candidates[j+1:]...)
	}

	fmt.Println("selected", selected)

	//fmt.Println(classes)

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

func copyClasses(classes map[int][]int) map[int][]int {

	nm := make(map[int][]int)
	for k, v := range classes {
		nm[k] = v
	}
	return nm
}

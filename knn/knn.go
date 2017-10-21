package knn

type knn struct {
	classes    []int
	DataSets   map[int][]int
	k          int
	input      int
	Neighbours []int
}

func Initialize(classes []int, dataSets map[int][]int, K int, input int) *knn {

	return &knn{classes, dataSets, K, input, nil}
}

func (k *knn) Run() *knn {

	classes := make(map[int][]int)
	for key, v := range k.DataSets {
		classes[key] = euc(k.input, v)

	}

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

		k.Neighbours = append(k.Neighbours, results[i]["key"])

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

	k.DataSets[clusterKey] = append(k.DataSets[clusterKey], k.input)
	return k
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

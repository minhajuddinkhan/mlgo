package kmeans

import (
	"fmt"
)

//KmeansAlgo structure
type KmeansAlgo struct {
	k         int
	dataSet   []int
	centroids []int
	clusters  map[int]int
}

// 1, 7
//dataSet = []int{1, 5, 7, 9}

//Initialize inits the algo.
func Initialize(k int, dataSet []int) *KmeansAlgo {

	return &KmeansAlgo{k, dataSet, nil, nil}
}

//RandomCentroids selects random centroids
func (kma *KmeansAlgo) RandomCentroids() *KmeansAlgo {
	kma.centroids = []int{
		kma.dataSet[1],
		kma.dataSet[2],
	}
	return kma
}

//SetCentroids set user given centroids
func (kma *KmeansAlgo) SetCentroids(centroids []int) (*KmeansAlgo, error) {
	if len(centroids) != kma.k {
		return nil, fmt.Errorf("cannot set more centroids")
	}
	kma.centroids = centroids
	return kma, nil
}

func (kma *KmeansAlgo) Iterate() {

	for _, data := range kma.dataSet {
		diff := euc(data, kma.centroids)
		fmt.Println(diff)
	}
}

func euc(data int, centroids []int) []int {

	differences := []int{}

	for _, ctr := range centroids {
		diff := data - ctr
		if diff < 0 {
			diff = -diff
		}
		differences = append(differences, diff)
	}
	return differences
}

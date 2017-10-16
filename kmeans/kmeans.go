package kmeans

import (
	"fmt"
	"math/rand"
	"time"
)

//KmeansAlgo structure
type KmeansAlgo struct {
	k         int
	dataSet   []int
	centroids []int
	clusters  map[int][]int
}

//Initialize inits the algo.
func Initialize(k int, dataSet []int) *KmeansAlgo {

	return &KmeansAlgo{k, dataSet, nil, make(map[int][]int)}
}

//RandomCentroids selects random centroids
func (kma *KmeansAlgo) RandomCentroids(num int) *KmeansAlgo {
	var i int
	centroids := []int{}

	for i < num {
		centroids = append(centroids, getRandomNumber(kma.dataSet))
		i++
	}
	kma.centroids = centroids
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
		index := returnSmallestIndex(diff)
		centroid := kma.centroids[index]
		cluster := kma.clusters[centroid]
		cluster = append(cluster, data)
		kma.clusters[centroid] = cluster
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

func returnSmallestIndex(x []int) int {
	m := 0
	for i := 0; i < len(x); i++ {
		if x[i] < x[m] {
			m = i
		}
	}
	return m
}

func (kma *KmeansAlgo) reArrageCentroids() []int {
	centroids := kma.centroids
	for _, ctr := range kma.clusters {
		sum := 0
		for _, v := range ctr {
			sum += v
		}
		centroids = append(centroids, sum/len(ctr))
	}
	return centroids
}

func (kma *KmeansAlgo) Run(max int) (map[int][]int, error) {

	//	done := false
	if len(kma.centroids) == 0 {
		return nil, fmt.Errorf("cannot execute with zero centroids")
	}
	for i := 0; i < max; i++ {
		kma.Iterate()
		newCentroids := kma.reArrageCentroids()
		fmt.Println("kmacentroids", kma.centroids)
		fmt.Println("newCentroids", newCentroids)
		// for i, v := range kma.centroids {
		// 	if v == newCentroids[i] {
		// 		done = true

		// 	}

		// }

		// if !done {
		// 	kma.centroids = newCentroids
		// } else {
		// 	break
		// }
	}

	return kma.clusters, nil
}

func getRandomNumber(ds []int) int {

	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(ds) - 1)
}

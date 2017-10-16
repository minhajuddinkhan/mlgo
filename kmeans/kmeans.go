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

//InitialRandomCentroids selects random centroids
func (kma *KmeansAlgo) InitialRandomCentroids(num int) *KmeansAlgo {
	i := 0
	centroids := []int{}

	for i < num {
		seed := getRandomNumber(kma.dataSet)
		centroids = append(centroids, seed)
		i++
	}
	kma.centroids = centroids
	return kma
}

//SetInitialCentroids set user given centroids
func (kma *KmeansAlgo) SetInitialCentroids(centroids []int) (*KmeansAlgo, error) {
	if len(centroids) != kma.k {
		return nil, fmt.Errorf("cannot set more centroids")
	}
	kma.centroids = centroids
	return kma, nil
}

func (kma *KmeansAlgo) emptyClusters() {

	for k := range kma.clusters {
		delete(kma.clusters, k)
	}

}

func (kma *KmeansAlgo) iterate() {

	kma.emptyClusters()
	for _, data := range kma.dataSet {
		diff := euc(data, kma.centroids)
		index := min(diff)
		centroid := kma.centroids[index]
		cluster := kma.clusters[centroid]
		cluster = append(cluster, data)
		kma.clusters[centroid] = cluster
	}

}

// finds eucledian distance. mock for now.
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

func min(x []int) int {
	m := 0
	for i := 0; i < len(x); i++ {
		if x[i] < x[m] {
			m = i
		}
	}
	return m
}

func (kma *KmeansAlgo) reArrageCentroids() []int {
	centroids := []int{}
	for _, ctr := range kma.clusters {
		sum := 0
		for _, v := range ctr {
			sum += v
		}
		centroids = append(centroids, sum/len(ctr))
	}
	return centroids
}

//Run executes the algorithm.
func (kma *KmeansAlgo) Run(max int) (map[int][]int, error) {

	done := true
	if len(kma.centroids) == 0 {
		return nil, fmt.Errorf("cannot execute with zero centroids")
	}
	for i := 0; i < max; i++ {
		kma.iterate()

		newCentroids := kma.reArrageCentroids()
		for i, v := range newCentroids {
			if v != kma.centroids[i] {
				done = false
			}
		}
		if done {
			break
		}
	}

	return kma.clusters, nil
}

func getRandomNumber(ds []int) int {

	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(ds) - 1)
}

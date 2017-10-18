package main

import (
	km "github.com/minhajuddinkhan/mlgo/kmeans"
	"github.com/minhajuddinkhan/mlgo/knn"
)

var (
	dataSet = []int{1, 2, 34, 14, 26, 2, 42, 53, 4, 64, 6, 7, 21, 57, 67, 4, 4, 2, 43, 56, 6, 76, 7, 87, 8, 9}
)

func main() {
	algo := km.Initialize(5, dataSet).InitialRandomCentroids()
	clusters, _ := algo.Run(1)

	classes := []int{}
	for k, _ := range clusters {
		classes = append(classes, k)
		//		fmt.Println(k, v)
	}

	knn.Initialize(classes, clusters, 3, 15).Run()

}

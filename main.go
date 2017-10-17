package main

import (
	km "github.com/minhajuddinkhan/mlgo/kmeans"
	"github.com/minhajuddinkhan/mlgo/knn"
)

var (
	dataSet = []int{1, 3, 4, 6, 7, 8, 9}
)

func main() {
	algo := km.Initialize(3, dataSet).InitialRandomCentroids()
	clusters, _ := algo.Run(20)

	classes := []int{}
	for k, _ := range clusters {
		classes = append(classes, k)
		//		fmt.Println(k, v)
	}

	knn.Initialize(classes, clusters, 3, 15).Run()

}

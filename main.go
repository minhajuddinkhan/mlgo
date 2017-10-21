package main

import (
	"fmt"

	km "github.com/minhajuddinkhan/mlgo/kmeans"
	"github.com/minhajuddinkhan/mlgo/knn"
)

var (
	dataSet = []int{1, 2, 34, 14, 26, 2, 42, 53, 4, 64, 6, 7, 21, 57, 67, 4, 4, 2, 43, 56, 6, 76, 7, 87, 8, 9}
	input   = 15
)

func main() {

	fmt.Println("\nDATA SET:: \n", dataSet)
	algo := km.Initialize(5, dataSet).InitialRandomCentroids()
	clusters, _ := algo.Run(1)

	classes := []int{}
	for k, v := range clusters {
		classes = append(classes, k)
		fmt.Println("CLUSTERS FORMED:: \n", k, v)
	}
	fmt.Println("\n\nINPUT:: ", input)

	kNN := knn.Initialize(classes, clusters, 3, input).Run()
	fmt.Println("\n\n ============\n K NEAREST NEIGHBOURS:: \n", kNN.Neighbours)

	fmt.Println("\n\n=======UPDATED DATA SET")
	for k, v := range kNN.DataSets {
		fmt.Println(k, v)
	}

}

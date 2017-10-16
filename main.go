package main

import (
	"fmt"

	km "github.com/minhajuddinkhan/mlgo/kmeans"
)

var (
	dataSet = []int{1, 5, 43, 54, 65, 6, 432, 3, 2, 6, 1, 5, 67, 7, 43, 7, 9}
)

func main() {
	algo := km.Initialize(3, dataSet).RandomCentroids(3)
	clusters, _ := algo.Run(9)
	fmt.Println(clusters)
}

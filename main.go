package main

import (
	"fmt"

	km "github.com/minhajuddinkhan/mlgo/kmeans"
)

var (
	dataSet = []int{1, 3, 5, 7, 543, 3, 12, 54, 1, 3, 123, 54, 65, 4, 2, 12, 32, 43, 54, 9}
)

func main() {
	algo := km.Initialize(3, dataSet).RandomCentroids(3)
	clusters, _ := algo.Run(2)
	fmt.Println(clusters)
}

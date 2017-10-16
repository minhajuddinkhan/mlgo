package main

import (
	"fmt"

	km "github.com/minhajuddinkhan/mlgo/kmeans"
)

var (
	dataSet = []int{1, 3, 5, 7, 43, 3, 12, 54, 1, 3, 23, 54, 65, 4, 2, 12, 32, 43, 54, 9, 21, 1, 23, 43, 11}
)

func main() {
	algo := km.Initialize(3, dataSet).InitialRandomCentroids(4)
	clusters, _ := algo.Run(50)
	for k, v := range clusters {
		fmt.Println(k, v)
	}
}

package main

import (
	km "github.com/minhajuddinkhan/ml/kmeans"
)

var (
	dataSet = []int{1, 5, 7, 9}
)

func main() {
	algo := km.Initialize(3, dataSet).RandomCentroids()
	algo.Iterate()
}

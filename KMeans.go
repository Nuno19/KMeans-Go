package kmeans

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sync"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

//KMeans - Struct for KMeans Algorithm
type KMeans struct {
	//number of clusters
	K int
	//max number of iterations
	MaxIter int
	//centroids of the clustering
	Centroids []Point
	//points for clustering
	Points []Point
	//labels of the corresponding point
	Labels []int
}

//InitCentroids - Inits the Centroids for random
func (kmeans *KMeans) InitCentroids() bool {
	if len(kmeans.points) <= kmeans.k {
		return false
	}
	kmeans.centroids = make([]Point, kmeans.k)
	var perm = rand.Perm(len(kmeans.points))
	for i := 0; i < kmeans.k; i++ {
		kmeans.centroids[i] = kmeans.points[perm[i]]
	}
	return true
}

//ComputeSSE - computes Sum of Squared Erros
func (kmeans *KMeans) ComputeSSE() float64 {
	var wg1 sync.WaitGroup
	wg1.Add(len(kmeans.centroids))
	distances := make(Point, len(kmeans.centroids))
	for i := range kmeans.centroids {
		go func(i int) {
			for j, point := range kmeans.points {
				if kmeans.labels[j] == i {
					distances[i] += point.Subtract(kmeans.centroids[i]).Norm()
				}
			}
			defer wg1.Done()
		}(i)
	}
	wg1.Wait()
	distance := 0.0
	for _, dist := range distances {
		distance += math.Pow(dist, 2)
	}
	return distance
}

//ComputeLabels - Computes labels of the points
func (kmeans *KMeans) ComputeLabels() []int {
	var labels = make([]int, len(kmeans.points))
	var wg1 sync.WaitGroup
	wg1.Add(len(kmeans.points))

	for i := range kmeans.points {
		go func(i int) {
			min := math.MaxFloat64
			minIdx := -1
			for j, cent := range kmeans.centroids {
				distance := kmeans.points[i].PointDist(cent)
				if distance < min {
					min = distance
					minIdx = j
				}
			}

			labels[i] = minIdx
			defer wg1.Done()
		}(i)
	}
	wg1.Wait()
	return labels
}

//ComputeCentroids - Recalculates centroids position
func ComputeCentroids(pointList []Point, distIndex []int, k int) []Point {
	clusters := make([][]Point, k)
	for i, point := range distIndex {
		clusters[point] = append(clusters[point], pointList[i])
	}

	centroids := make([]Point, k)
	for c, clu := range clusters {
		sums := make([]float64, len(pointList[0]))
		for _, point := range clu {
			for j, val := range point {
				sums[j] += val
			}
		}
		centroids[c] = make([]float64, len(sums))
		for i, s := range sums {
			centroids[c][i] = s / float64(len(clu))
		}
	}
	return centroids
}
func (kmeans *KMeans) update() bool {
	oldCentroids := make([]Point, len(kmeans.centroids))
	for i := 0; i < kmeans.maxIter; i++ {
		oldCentroids = kmeans.centroids
		kmeans.labels = kmeans.ComputeLabels()

		kmeans.centroids = ComputeCentroids(kmeans.points, kmeans.labels, kmeans.k)

		if equals(kmeans.centroids, oldCentroids) {

			return true
		}
	}

	return false
}

//Fit - Fits the points into k centroids
func (kmeans *KMeans) Fit(pointList []Point) bool {
	if kmeans.k == 0 {
		fmt.Printf("K not defined")
		return false
	}

	kmeans.points = pointList
	init := kmeans.InitCentroids()
	if !init {
		return false
	}
	found := kmeans.update()
	if found {
		return true
	}
	return false

}

//AddPoint - Add a new Point to the list and recalculate centroids
func (kmeans *KMeans) AddPoint(point Point) {
	kmeans.points = append(kmeans.points, point)
	found := kmeans.update()
	if found {
		fmt.Printf("Centroids Found!\n")
	} else {
		fmt.Printf("Centroids Not Found!\n")
	}
}

//AddPointList - Add a list of Points to the list and recalculate centroids
func (kmeans *KMeans) AddPointList(points []Point) {
	for _, p := range points {
		kmeans.points = append(kmeans.points, p)
	}
	found := kmeans.update()
	if found {
		fmt.Printf("Centroids Found!\n")
	} else {
		fmt.Printf("Centroids Not Found!\n")
	}
}

//PointsToString - Returns the String of the Points
func (kmeans *KMeans) PointsToString() string {
	var toReturn string

	for i, point := range kmeans.points {
		for j, p := range point {
			if j < len(point)-1 {
				toReturn += fmt.Sprintf("%f,", p)
			} else {
				toReturn += fmt.Sprintf("%f", p)
			}
		}
		if i < len(kmeans.points)-1 {
			toReturn += fmt.Sprintf(";")
		}
	}
	return toReturn
}

//CentroidsToString - returns the string of centroids
func (kmeans *KMeans) CentroidsToString() string {
	var toReturn string

	for i, cent := range kmeans.centroids {
		for j, p := range cent {
			if j < len(cent)-1 {
				toReturn += fmt.Sprintf("%f,", p)
			} else {
				toReturn += fmt.Sprintf("%f", p)
			}
		}
		if i < len(kmeans.centroids)-1 {
			toReturn += fmt.Sprintf(";")
		}
	}
	return toReturn
}

//LabelsToString - Returns the string of Labels
func (kmeans *KMeans) LabelsToString() string {
	var toReturn string

	for i, label := range kmeans.labels {

		if i < len(kmeans.labels)-1 {
			toReturn += fmt.Sprintf("%d,", label)
		} else {
			toReturn += fmt.Sprintf("%d", label)
		}
	}
	return toReturn
}

//LabelCount - Returns the number of points per centroid
func (kmeans *KMeans) LabelCount() []int {
	count := make([]int, len(kmeans.centroids))
	for _, label := range kmeans.labels {
		count[label]++
	}
	return count
}

//LabelCount - Returns the number of points per centroid
func LabelCount(labels []int, k int) []int {
	count := make([]int, k)
	for _, label := range labels {
		count[label]++
	}
	return count
}

//GetPointIdxOfCentroid - return indexes of points of a certain centroid
func (kmeans *KMeans) GetPointIdxOfCentroid(centroidIdx int) []int {
	var idxs []int

	for i, label := range kmeans.labels {
		if label == centroidIdx {
			idxs = append(idxs, i)
		}
	}
	return idxs
}

func (kmeans *KMeans) saveImage(size vg.Length, fileName string) {
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	for i, point := range kmeans.centroids {
		addScatters(p, i, "", convertPointToPlotterXY(point), true)
	}
	for i, point := range kmeans.points {
		addScatters(p, kmeans.labels[i], "", convertPointToPlotterXY(point), false)
	}

	p.Save(size*vg.Centimeter, size*vg.Centimeter, fileName)
}

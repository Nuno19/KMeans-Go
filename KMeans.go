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

type Point []float64

type KMeans struct {
	k         int
	maxIter   int
	centroids []Point
	points    []Point
	labels    []int
}

func (point Point) print() {

	for i, num := range point {

		if i == len(point)-1 {
			fmt.Printf("%f", num)
		} else {
			fmt.Printf("%f,", num)
		}
	}
	fmt.Printf("\n")
}

func (point Point) pointDist(p2 Point) float64 {
	var sum float64
	for i := 0; i < len(point); i++ {
		sum += math.Pow(point[i]-p2[i], 2.0)
	}
	return math.Sqrt(sum)
}

func (point Point) pointEqual(p2 Point) bool {
	if len(point) != len(p2) {
		return false
	}
	for i := 0; i < len(point); i++ {
		if point[i] != p2[i] {
			return false
		}
	}
	return true
}

func pointDist(p1 Point, p2 Point) {
	var sum float64
	for i := 0; i < len(p1); i++ {
		sum += math.Pow(p1[i]-p2[i], 2.0)
	}

}

func (kmeans *KMeans) initCentroids() {
	kmeans.centroids = make([]Point, kmeans.k)
	var perm = rand.Perm(len(kmeans.points))
	for i := 0; i < kmeans.k; i++ {
		kmeans.centroids[i] = kmeans.points[perm[i]]
	}
}

func computeLabels(pointList []Point, centroids []Point) []int {
	var labels = make([]int, len(pointList))

	var wg1 sync.WaitGroup
	wg1.Add(len(pointList))

	for i := range pointList {
		go func(i int) {
			//distances = make([]float64, len(centroids))
			min := math.MaxFloat64
			minIdx := -1
			for j, cent := range centroids {

				distance := pointList[i].pointDist(cent)
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

func computeCentroids(pointList []Point, distIndex []int, k int) []Point {
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
		//start := time.Now().UnixNano()
		oldCentroids = kmeans.centroids
		kmeans.labels = computeLabels(kmeans.points, kmeans.centroids)

		kmeans.centroids = computeCentroids(kmeans.points, kmeans.labels, kmeans.k)
		//fmt.Printf("IDX: %d |", i)
		//fmt.Printf("%f\n", float64((end-start))/1000000000)
		//end := time.Now().UnixNano()
		if equals(kmeans.centroids, oldCentroids) {

			return true
		}
	}

	return false
}

func equals(points1 []Point, points2 []Point) bool {
	if len(points1) != len(points2) {
		return false
	}
	for i := range points1 {
		if !points1[i].pointEqual(points2[i]) {
			return false
		}
	}
	return true
}

func (kmeans *KMeans) fit(pointList []Point) {
	if kmeans.k == 0 {
		fmt.Printf("K not defined")
		return
	}

	kmeans.points = pointList
	kmeans.initCentroids()

	found := kmeans.update()
	if found {
		fmt.Printf("Centroids Found!\n")
	} else {
		fmt.Printf("Centroids Not Found!\n")
	}
}

func (kmeans *KMeans) addPoint(point Point) {
	kmeans.points = append(kmeans.points, point)
	found := kmeans.update()
	if found {
		fmt.Printf("Centroids Found!\n")
	} else {
		fmt.Printf("Centroids Not Found!\n")
	}
}

func (kmeans *KMeans) addPointList(points []Point) {
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
func (kmeans *KMeans) pointsToString() string {
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

func (kmeans *KMeans) centroidsToString() string {
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

func (kmeans *KMeans) labelsToString() string {
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
func printList(list []int) {
	for _, el := range list {
		fmt.Printf("%d,", el)
	}
	fmt.Println()
}

func (kmeans *KMeans) labelCount() []int {
	count := make([]int, len(kmeans.centroids))
	for _, label := range kmeans.labels {
		count[label]++
	}
	return count
}

func Hello() string {
	return "Hello, world."
}

/*
func main() {

	maxIter := 200
	k := 50
	dataSize := 1000

	rand.Seed(time.Now().Unix())

	pointArr := make([]Point, dataSize)

	for i := 0; i < len(pointArr); i++ {
		pointArr[i] = Point{rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64()}
		//pointArr[i].print()
	}
	fmt.Println(len(pointArr[0]))
	kmeans := KMeans{k: k, maxIter: maxIter}
	start := time.Now().UnixNano()
	kmeans.fit(pointArr)
	end := time.Now().UnixNano()
	fmt.Printf("%f\n", float64((end-start))/1000000000)
	fmt.Println("COUNTS: ")
	kmeans.saveImage(8, "file.pdf")
	printList(kmeans.labelCount())
}
*/

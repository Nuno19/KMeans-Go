package kmeans

import (
	"reflect"
	"testing"

	"gonum.org/v1/plot/vg"
)

func TestPoint_pointDist(t *testing.T) {
	type args struct {
		p2 Point
	}
	tests := []struct {
		name  string
		point Point
		args  args
		want  float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.point.pointDist(tt.args.p2); got != tt.want {
				t.Errorf("Point.pointDist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pointDist(t *testing.T) {
	type args struct {
		p1 Point
		p2 Point
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pointDist(tt.args.p1, tt.args.p2)
		})
	}
}

func TestPoint_print(t *testing.T) {
	tests := []struct {
		name  string
		point Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.point.print()
		})
	}
}

func TestPoint_pointEqual(t *testing.T) {
	type args struct {
		p2 Point
	}
	tests := []struct {
		name  string
		point Point
		args  args
		want  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.point.pointEqual(tt.args.p2); got != tt.want {
				t.Errorf("Point.pointEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMeans_initCentroids(t *testing.T) {
	tests := []struct {
		name   string
		kmeans *KMeans
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.kmeans.initCentroids()
		})
	}
}

func Test_computeLabels(t *testing.T) {
	type args struct {
		pointList []Point
		centroids []Point
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeLabels(tt.args.pointList, tt.args.centroids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeCentroids(t *testing.T) {
	type args struct {
		pointList []Point
		distIndex []int
		k         int
	}
	tests := []struct {
		name string
		args args
		want []Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeCentroids(tt.args.pointList, tt.args.distIndex, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeCentroids() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMeans_update(t *testing.T) {
	tests := []struct {
		name   string
		kmeans *KMeans
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kmeans.update(); got != tt.want {
				t.Errorf("KMeans.update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_equals(t *testing.T) {
	type args struct {
		points1 []Point
		points2 []Point
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equals(tt.args.points1, tt.args.points2); got != tt.want {
				t.Errorf("equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMeans_fit(t *testing.T) {
	type args struct {
		pointList []Point
	}
	tests := []struct {
		name   string
		kmeans *KMeans
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.kmeans.fit(tt.args.pointList)
		})
	}
}

func TestKMeans_addPoint(t *testing.T) {
	type args struct {
		point Point
	}
	tests := []struct {
		name   string
		kmeans *KMeans
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.kmeans.addPoint(tt.args.point)
		})
	}
}

func TestKMeans_addPointList(t *testing.T) {
	type args struct {
		points []Point
	}
	tests := []struct {
		name   string
		kmeans *KMeans
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.kmeans.addPointList(tt.args.points)
		})
	}
}

func TestKMeans_saveImage(t *testing.T) {
	type args struct {
		size     vg.Length
		fileName string
	}
	tests := []struct {
		name   string
		kmeans *KMeans
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.kmeans.saveImage(tt.args.size, tt.args.fileName)
		})
	}
}

func TestKMeans_pointsToString(t *testing.T) {
	tests := []struct {
		name   string
		kmeans *KMeans
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kmeans.pointsToString(); got != tt.want {
				t.Errorf("KMeans.pointsToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMeans_centroidsToString(t *testing.T) {
	tests := []struct {
		name   string
		kmeans *KMeans
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kmeans.centroidsToString(); got != tt.want {
				t.Errorf("KMeans.centroidsToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMeans_labelsToString(t *testing.T) {
	tests := []struct {
		name   string
		kmeans *KMeans
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kmeans.labelsToString(); got != tt.want {
				t.Errorf("KMeans.labelsToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printList(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printList(tt.args.list)
		})
	}
}

func TestKMeans_labelCount(t *testing.T) {
	tests := []struct {
		name   string
		kmeans *KMeans
		want   []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kmeans.labelCount(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KMeans.labelCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}

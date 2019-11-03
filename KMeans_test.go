package kmeans

import (
	"reflect"
	"testing"

	"gonum.org/v1/plot/vg"
)

func TestKMeans_initCentroids(t *testing.T) {
	tests := []struct {
		name   string
		kmeans *KMeans
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kmeans.initCentroids(); got != tt.want {
				t.Errorf("KMeans.initCentroids() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMeans_computeSSE(t *testing.T) {
	tests := []struct {
		name   string
		kmeans *KMeans
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kmeans.computeSSE(); got != tt.want {
				t.Errorf("KMeans.computeSSE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMeans_computeLabels(t *testing.T) {
	tests := []struct {
		name   string
		kmeans *KMeans
		want   []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kmeans.computeLabels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KMeans.computeLabels() = %v, want %v", got, tt.want)
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

func TestKMeans_fit(t *testing.T) {
	type args struct {
		pointList []Point
	}
	tests := []struct {
		name   string
		kmeans *KMeans
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kmeans.fit(tt.args.pointList); got != tt.want {
				t.Errorf("KMeans.fit() = %v, want %v", got, tt.want)
			}
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

func Test_labelCount(t *testing.T) {
	type args struct {
		labels []int
		k      int
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
			if got := labelCount(tt.args.labels, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("labelCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMeans_getPointIdxOfCentroid(t *testing.T) {
	type args struct {
		centroidIdx int
	}
	tests := []struct {
		name   string
		kmeans *KMeans
		args   args
		want   []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kmeans.getPointIdxOfCentroid(tt.args.centroidIdx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KMeans.getPointIdxOfCentroid() = %v, want %v", got, tt.want)
			}
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

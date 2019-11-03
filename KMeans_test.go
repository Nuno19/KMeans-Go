package kmeans

import (
	"reflect"
	"testing"

	"gonum.org/v1/plot/vg"
)

func TestKMeans_InitCentroids(t *testing.T) {
	tests := []struct {
		name   string
		kmeans *KMeans
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kmeans.InitCentroids(); got != tt.want {
				t.Errorf("KMeans.InitCentroids() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMeans_ComputeSSE(t *testing.T) {
	tests := []struct {
		name   string
		kmeans *KMeans
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kmeans.ComputeSSE(); got != tt.want {
				t.Errorf("KMeans.ComputeSSE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMeans_ComputeLabels(t *testing.T) {
	tests := []struct {
		name   string
		kmeans *KMeans
		want   []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kmeans.ComputeLabels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KMeans.ComputeLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeCentroids(t *testing.T) {
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
			if got := ComputeCentroids(tt.args.pointList, tt.args.distIndex, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComputeCentroids() = %v, want %v", got, tt.want)
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

func TestKMeans_Fit(t *testing.T) {
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
			if got := tt.kmeans.Fit(tt.args.pointList); got != tt.want {
				t.Errorf("KMeans.Fit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMeans_AddPoint(t *testing.T) {
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
			tt.kmeans.AddPoint(tt.args.point)
		})
	}
}

func TestKMeans_AddPointList(t *testing.T) {
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
			tt.kmeans.AddPointList(tt.args.points)
		})
	}
}

func TestKMeans_PointsToString(t *testing.T) {
	tests := []struct {
		name   string
		kmeans *KMeans
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kmeans.PointsToString(); got != tt.want {
				t.Errorf("KMeans.PointsToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMeans_CentroidsToString(t *testing.T) {
	tests := []struct {
		name   string
		kmeans *KMeans
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kmeans.CentroidsToString(); got != tt.want {
				t.Errorf("KMeans.CentroidsToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMeans_LabelsToString(t *testing.T) {
	tests := []struct {
		name   string
		kmeans *KMeans
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kmeans.LabelsToString(); got != tt.want {
				t.Errorf("KMeans.LabelsToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMeans_LabelCount(t *testing.T) {
	tests := []struct {
		name   string
		kmeans *KMeans
		want   []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kmeans.LabelCount(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KMeans.LabelCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLabelCount(t *testing.T) {
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
			if got := LabelCount(tt.args.labels, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LabelCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMeans_GetPointIdxOfCentroid(t *testing.T) {
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
			if got := tt.kmeans.GetPointIdxOfCentroid(tt.args.centroidIdx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KMeans.GetPointIdxOfCentroid() = %v, want %v", got, tt.want)
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

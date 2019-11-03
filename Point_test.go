package kmeans

import (
	"reflect"
	"testing"
)

func TestPoint_Print(t *testing.T) {
	tests := []struct {
		name  string
		point Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.point.Print()
		})
	}
}

func TestPoint_PointDist(t *testing.T) {
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
			if got := tt.point.PointDist(tt.args.p2); got != tt.want {
				t.Errorf("Point.PointDist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_PointEqual(t *testing.T) {
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
			if got := tt.point.PointEqual(tt.args.p2); got != tt.want {
				t.Errorf("Point.PointEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Subtract(t *testing.T) {
	type args struct {
		p2 Point
	}
	tests := []struct {
		name  string
		point Point
		args  args
		want  Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.point.Subtract(tt.args.p2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point.Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Norm(t *testing.T) {
	tests := []struct {
		name  string
		point Point
		want  float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.point.Norm(); got != tt.want {
				t.Errorf("Point.Norm() = %v, want %v", got, tt.want)
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

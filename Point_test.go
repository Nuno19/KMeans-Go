package kmeans

import (
	"reflect"
	"testing"
)

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

func TestPoint_subtract(t *testing.T) {
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
			if got := tt.point.subtract(tt.args.p2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point.subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_norm(t *testing.T) {
	tests := []struct {
		name  string
		point Point
		want  float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.point.norm(); got != tt.want {
				t.Errorf("Point.norm() = %v, want %v", got, tt.want)
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

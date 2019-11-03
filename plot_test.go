package kmeans

import (
	"reflect"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func Test_convertPointToPlotterXY(t *testing.T) {
	type args struct {
		v Point
	}
	tests := []struct {
		name string
		args args
		want plotter.XYs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertPointToPlotterXY(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertPointToPlotterXY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addScatters(t *testing.T) {
	type args struct {
		plt   *plot.Plot
		index int
		name  string
		xyers plotter.XYs
		cent  bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := addScatters(tt.args.plt, tt.args.index, tt.args.name, tt.args.xyers, tt.args.cent); (err != nil) != tt.wantErr {
				t.Errorf("addScatters() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sa(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa()
		})
	}
}

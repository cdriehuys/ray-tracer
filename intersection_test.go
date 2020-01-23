package main

import (
	"reflect"
	"testing"
)

func TestMakeIntersection(t *testing.T) {
	s := &Sphere{}
	i := MakeIntersection(3.5, s)

	if !Float64Equal(3.5, i.T) {
		t.Errorf("Expected i.T == 3.5, got %v", i.T)
	}

	if !reflect.DeepEqual(i.Object, s) {
		t.Errorf("Expected i.Object == %v, got %v", s, i.Object)
	}
}

func TestIntersections_Hit(t *testing.T) {
	sphere := Sphere{}
	testCases := []struct {
		name          string
		intersections Intersections
		want          Intersection
		wantNothing   bool
	}{
		{
			name: "all positive t",
			intersections: Intersections{
				MakeIntersection(1, sphere),
				MakeIntersection(2, sphere),
			},
			want: MakeIntersection(1, sphere),
		},
		{
			name: "some negative t",
			intersections: Intersections{
				MakeIntersection(-1, sphere),
				MakeIntersection(1, sphere),
			},
			want: MakeIntersection(1, sphere),
		},
		{
			name: "all negative t",
			intersections: Intersections{
				MakeIntersection(-2, sphere),
				MakeIntersection(-1, sphere),
			},
			wantNothing: true,
		},
		{
			name: "lowest non-negative",
			intersections: Intersections{
				MakeIntersection(5, sphere),
				MakeIntersection(7, sphere),
				MakeIntersection(-3, sphere),
				MakeIntersection(2, sphere),
			},
			want: MakeIntersection(2, sphere),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, exists := tt.intersections.Hit()

			if tt.wantNothing && exists {
				t.Errorf("Wanted no hits but exists is true")
			}

			emptyIntersection := Intersection{}
			if !exists && got != emptyIntersection {
				t.Errorf("Exists is false but hit is non-empty: %v", got)
			}

			if !tt.wantNothing && got != tt.want {
				t.Errorf("Expected hit to be %v, got %v", tt.want, got)
			}
		})
	}
}

func TestIntersections_Sort(t *testing.T) {
	sphere := Sphere{}
	testCases := []struct {
		name          string
		intersections Intersections
		want          Intersections
	}{
		{
			"unordered intersections",
			Intersections{
				MakeIntersection(5, sphere),
				MakeIntersection(7, sphere),
				MakeIntersection(-3, sphere),
				MakeIntersection(2, sphere),
			},
			Intersections{
				MakeIntersection(-3, sphere),
				MakeIntersection(2, sphere),
				MakeIntersection(5, sphere),
				MakeIntersection(7, sphere),
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.intersections.Sort()
			if !reflect.DeepEqual(tt.intersections, tt.want) {
				t.Errorf("Expected different sort results:\nExpected: %v\nReceived: %v", tt.want, tt.intersections)
			}
		})
	}
}

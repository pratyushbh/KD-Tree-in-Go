package main

import (
	"fmt"
	"sort"
)

type Point struct {
	x int
	y int
}
type KDTree struct {
	points []Point
	n      int
	m      int
	median Point
	axis   string
	left   *KDTree
	right  *KDTree
}

func (k *KDTree) Sort() {
	if k.axis == "x" {
		sort.Slice(k.points, func(i, j int) bool {
			return k.points[i].x < k.points[j].x
		})
	}
	if k.axis == "y" {
		sort.Slice(k.points, func(i, j int) bool {
			return k.points[i].y < k.points[j].y
		})
	}
	//fmt.Println(k.points)
}

func Build(points []Point, axis string) *KDTree {
	k := &KDTree{points: []Point{}}
	k.points = append(k.points, points...)
	k.n = len(k.points)
	k.axis = axis
	k.Sort()
	k.m = int(k.n / 2)
	k.median = k.points[k.m]
	if k.m > 0 {
		if axis == "x" {
			k.left = Build(k.points[:k.m], "y")
		}
		if axis == "y" {
			k.left = Build(k.points[:k.m], "x")
		}
	}
	if k.n-(k.m+1) > 0 {
		if axis == "x" {
			k.right = Build(k.points[k.m+1:], "y")
		}
		if axis == "y" {
			k.right = Build(k.points[k.m+1:], "x")
		}
	}
	return k
}

func (k *KDTree) Insert(point *Point, axis string) *KDTree {
	if k == nil {
		kd := Build([]Point{*point}, axis)
		return kd
	}
	if k.axis == "x" {

	}

}
func main() {
	points := []Point{Point{x: 1, y: 9}, Point{x: 2, y: 3}, Point{x: 4, y: 1}, Point{x: 3, y: 7}, Point{x: 5, y: 4}, Point{x: 6, y: 8}, Point{x: 7, y: 2}, Point{x: 8, y: 8}, Point{x: 7, y: 9}, Point{x: 9, y: 6}}
	kd := Build(points, "x")
	fmt.Println(kd.right.right)
}

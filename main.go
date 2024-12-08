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
func (k *KDTree) inbox(box *[4]Point) bool {
	return true
}
func (k *KDTree) rangeSearch(box *[4]Point, sol []Point) []Point {
	if k.inbox(box) {
		sol = append(sol, k.median)
	}
	if k.left != nil {
		sol = k.left.rangeSearch(box, sol)
	}
	if k.right != nil {
		sol = k.right.rangeSearch(box, sol)
	}
	return sol
}

func main() {
	points := []Point{{x: 1, y: 9}, {x: 2, y: 3}, {x: 4, y: 1}, {x: 3, y: 7}, {x: 5, y: 4}, {x: 6, y: 8}, {x: 7, y: 2}, {x: 8, y: 8}, {x: 7, y: 9}, {x: 9, y: 6}}
	kd := Build(points, "x")
	inboxPoints := []Point{}
	//fmt.Println(kd.points)
	//fmt.Println(kd.left.points)
	//fmt.Println(kd.right.points)
	//fmt.Println(kd.left.left.points)
	//fmt.Println(kd.left.right.points)
	//fmt.Println(kd.right.left.points)
	//fmt.Println(kd.right.right.points)
	inboxPoints = kd.rangeSearch(&[4]Point{{x: 2, y: 2}, {x: 2, y: 4}, {x: 4, y: 2}, {x: 4, y: 4}}, inboxPoints)
	fmt.Println(inboxPoints)
}

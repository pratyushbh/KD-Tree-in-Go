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
	//points []Point
	n      int
	m      int
	median Point
	axis   string
	left   *KDTree
	right  *KDTree
}

func (k *KDTree) Sort(points []Point) []Point {
	if k.axis == "x" {
		sort.Slice(points, func(i, j int) bool {
			return points[i].x < points[j].x
		})
	}
	if k.axis == "y" {
		sort.Slice(points, func(i, j int) bool {
			return points[i].y < points[j].y
		})
	}
	return points
	//fmt.Println(k.points)
}

func Build(points []Point, axis string) *KDTree {
	k := &KDTree{}
	//k.points = append(k.points, points...)
	k.n = len(points)
	k.axis = axis
	k.Sort(points)
	k.m = int(k.n / 2)
	k.median = points[k.m]
	if k.m > 0 {
		if axis == "x" {
			k.left = Build(points[:k.m], "y")
		}
		if axis == "y" {
			k.left = Build(points[:k.m], "x")
		}
	}
	if k.n-(k.m+1) > 0 {
		if axis == "x" {
			k.right = Build(points[k.m+1:], "y")
		}
		if axis == "y" {
			k.right = Build(points[k.m+1:], "x")
		}
	}
	return k
}
func (k *KDTree) inbox(box *[4]Point) bool {
	var result bool = true
	fmt.Println("CHECKING FOR POINT:", k.median)
	if !(box[0].x <= k.median.x && box[0].y <= k.median.y) {
		//fmt.Println("CONDITION 1 FAILED!", box[0].x, box[0].y)
		result = false
		return result
	}
	if !(box[1].x <= k.median.x && box[1].y >= k.median.y) {
		//fmt.Println("CONDITION 2 FAILED!", box[1].x, box[1].y)
		result = false
		return result
	}
	if !(box[2].x >= k.median.x && box[2].y <= k.median.y) {
		//fmt.Println("CONDITION 3 FAILED!", box[2].x, box[2].y)
		result = false
		return result
	}
	if !(box[3].x >= k.median.x && box[3].y >= k.median.y) {
		//fmt.Println("CONDITION 4 FAILED!", box[3].x, box[3].y)
		result = false
		return result
	}
	//fmt.Println(" RESULT:", result)
	return result
}
func (k *KDTree) rangeSearch(box *[4]Point, sol []Point) []Point {
	if k.inbox(box) {
		sol = append(sol, k.median)
	}
	var min Point
	var max Point
	if k.axis == "x" {
		min = box[0]
		max = box[2]
		if k.left != nil && k.median.x >= min.x {
			sol = k.left.rangeSearch(box, sol)
		}
		if k.right != nil && k.median.x <= max.x {
			sol = k.right.rangeSearch(box, sol)
		}
	}
	if k.axis == "y" {
		min = box[1]
		max = box[3]
		if k.left != nil && k.median.x >= min.x {
			sol = k.left.rangeSearch(box, sol)
		}
		if k.right != nil && k.median.x <= max.x {
			sol = k.right.rangeSearch(box, sol)
		}
	}
	return sol
}

func min(axis string, a *Point, b *Point, c *Point) *Point {
	if axis == "x" {
		if a.x < b.x {
			if a.x < c.x {
				return a
			} else {
				return c
			}
		} else {
			if b.x < c.x {
				return b
			} else {
				return c
			}
		}
	} else if axis == "y" {
		if a.y < b.y {
			if a.y < c.y {
				return a
			} else {
				return c
			}
		} else {
			if b.y < c.y {
				return b
			} else {
				return c
			}
		}
	}
	return nil
}

func (k *KDTree) findMin(axis string) *Point {
	if k == nil {
		return nil
	}
	if k.axis == axis {
		if k.left == nil {
			return &k.median
		} else {
			return k.left.findMin(axis)
		}
	} else {
		ls := k.left.findMin(axis)
		rs := k.right.findMin(axis)
		if ls == nil {
			if rs == nil {
				return &k.median
			} else {
				if axis == "x" {
					if rs.x < k.median.x {
						return rs
					} else {
						return &k.median
					}
				} else if axis == "y" {
					if rs.y < k.median.y {
						return rs
					} else {
						return &k.median
					}
				}
			}
		} else if rs == nil {
			if axis == "x" {
				if ls.x < k.median.x {
					return ls
				} else {
					return &k.median
				}
			} else if axis == "y" {
				if ls.y < k.median.y {
					return ls
				} else {
					return &k.median
				}
			}

		} else {
			MinPoint := min(axis, &k.median, ls, rs)
			return MinPoint
		}
	}
	return nil
}

func (k *KDTree) Delete(x Point) *KDTree {
	if k == nil {
		return nil
	}
	if x == k.median {
		if k.right != nil {
			k.median = *k.right.findMin(k.axis)
			k.right = k.right.Delete(k.median)
		} else if k.left != nil {
			k.median = *k.left.findMin(k.axis)
			k.right = k.left.Delete(k.median)
		} else {
			k = nil
		}
	} else {
		if k.axis == "x" {
			if x.x < k.median.x {
				k.left = k.left.Delete(x)
			} else {
				k.right = k.right.Delete(x)
			}
		} else if k.axis == "y" {
			if x.y < k.median.y {
				k.left = k.left.Delete(x)
			} else {
				k.right = k.right.Delete(x)
			}
		}

	}
	return k
}

func (k *KDTree) insert(x *Point, axis string) *KDTree {
	if k == nil {
		return Build([]Point{*x}, axis)
	} else if *x == k.median {
		return k
	} else {
		if axis == "x" {
			if x.x < k.median.x {
				k.left = k.left.insert(x, "y")
			} else {
				k.right = k.right.insert(x, "y")
			}
		} else if axis == "y" {
			if x.x < k.median.x {
				k.left = k.left.insert(x, "x")
			} else {
				k.right = k.right.insert(x, "x")
			}
		}
	}
	return k
}

func main() {
	points := []Point{{x: 1, y: 9}, {x: 2, y: 3}, {x: 4, y: 1}, {x: 3, y: 7}, {x: 5, y: 4}, {x: 6, y: 8}, {x: 7, y: 2}, {x: 8, y: 8}, {x: 7, y: 9}, {x: 9, y: 6}}
	kd := Build(points, "x")
	inboxPoints := []Point{}
	inboxPoints = kd.rangeSearch(&[4]Point{{x: 2, y: 2}, {x: 2, y: 4}, {x: 4, y: 2}, {x: 4, y: 4}}, inboxPoints)
	fmt.Println(inboxPoints)
	fmt.Println(kd.findMin("y"))
	kd = kd.Delete(Point{x: 6, y: 8})
	kd = kd.insert(&Point{x: 6, y: 8}, kd.axis)
	fmt.Println(kd)
	fmt.Println(kd.left)
	fmt.Println(kd.right)
	fmt.Println(kd.left.left)
	fmt.Println(kd.left.right.right)
	fmt.Println(kd.right.left)
	fmt.Println(kd.right.right)
}

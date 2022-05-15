package main

import "math"

func largestTriangleArea(points [][]int) float64 {
	ans := 0.0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				if points[i][0] == points[j][0] && points[i][0] == points[k][0] {
					continue
				}
				if points[i][1] == points[j][1] && points[i][0] == points[k][1] {
					continue
				}
				e1 := getEdgeLength(points[i], points[j])
				e2 := getEdgeLength(points[i], points[k])
				e3 := getEdgeLength(points[j], points[k])
				p := (e1 + e2 + e3) / 2.0
				area := math.Sqrt(p * (p - e1) * (p - e2) * (p - e3))
				//ans = math.Max(ans, area) // 使用这个有问题
				if area > ans {
					ans = area
				}
			}
		}
	}
	return ans
}

func getEdgeLength(a, b []int) float64 {
	x := a[0] - b[0]
	y := a[1] - b[1]

	return math.Sqrt(float64(x*x + y*y))
}

func main() {

}

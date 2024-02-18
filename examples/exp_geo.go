package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := GeoPoint{
		29.490295,
		106.486654,
	}

	p2 := GeoPoint{
		29.615467,
		106.581515,
	}

	fmt.Println(EarthDistance(p1, p2))
}

type GeoPoint struct {
	lat float64 // 纬度
	lon float64 // 经度
}

func EarthDistance(p1, p2 GeoPoint) float64 {
	radius := 6371000.0    // 地球半径，单位米
	rad := math.Pi / 180.0 // 弧度=度数*π/180

	p1.lat = p1.lat * rad
	p1.lon = p1.lon * rad
	p2.lat = p2.lat * rad
	p2.lon = p2.lon * rad

	theta := p2.lon - p1.lon
	dist := math.Acos(math.Sin(p1.lat)*math.Sin(p2.lat) + math.Cos(p1.lat)*math.Cos(p2.lat)*math.Cos(theta))

	return dist * radius
}

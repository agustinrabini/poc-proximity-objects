package main

import (
	"fmt"
	"math"
)

func Programa1(shop domain.Shop) {
	//https://stuff.21zoo.com/posts/golang-calculate-distance-between-two-lat-long-points/
	fmt.Println(" ====== Formula 1 ====== \n\n ")

	radlat1 := float64(math.Pi * loc1.lat / 180)
	radlat2 := float64(math.Pi * loc2.lat / 180)

	theta := float64(loc1.lng - loc2.lng)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	dist = dist * 1.60934 * 1000

	fmt.Printf("Distancia de %s a %s en metros. Segun programa 1: %f \n", loc1.Name, loc2.Name, dist)
}

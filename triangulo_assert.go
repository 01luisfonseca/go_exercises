package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/01luisfonseca/go_exercises/readinput"
)

type segmento struct {
	a [2]int64
	b [2]int64
}

func main_old() {
	fmt.Println("Programa de prueba de triangulos")
	pointA1, _ := readinput.Read("Punto 1 Segmento A. Ej: 1,2")
	pointA2, _ := readinput.Read("Punto 2 Segmento A. Ej: 1,2")
	pointB1, _ := readinput.Read("Punto 1 Segmento B. Ej: 1,2")
	pointB2, _ := readinput.Read("Punto 2 Segmento B. Ej: 1,2")
	pointC1, _ := readinput.Read("Punto 1 Segmento C. Ej: 1,2")
	pointC2, _ := readinput.Read("Punto 2 Segmento C. Ej: 1,2")

	// First segment
	pointA1Split := strings.Split(pointA1, ",")
	pointA1_0, _ := strconv.ParseInt(pointA1Split[0], 10, 0)
	pointA1_1, _ := strconv.ParseInt(pointA1Split[1], 10, 0)
	pointA1Int := [2]int64{pointA1_0, pointA1_1}
	pointA2Split := strings.Split(pointA2, ",")
	pointA2_0, _ := strconv.ParseInt(pointA2Split[0], 10, 0)
	pointA2_1, _ := strconv.ParseInt(pointA2Split[1], 10, 0)
	pointA2Int := [2]int64{pointA2_0, pointA2_1}
	segment1 := segmento{a: pointA1Int, b: pointA2Int}

	// Second segment
	pointB1Split := strings.Split(pointB1, ",")
	pointB1_0, _ := strconv.ParseInt(pointB1Split[0], 10, 0)
	pointB1_1, _ := strconv.ParseInt(pointB1Split[1], 10, 0)
	pointB1Int := [2]int64{pointB1_0, pointB1_1}
	pointB2Split := strings.Split(pointB2, ",")
	pointB2_0, _ := strconv.ParseInt(pointB2Split[0], 10, 0)
	pointB2_1, _ := strconv.ParseInt(pointB2Split[1], 10, 0)
	pointB2Int := [2]int64{pointB2_0, pointB2_1}
	segment2 := segmento{a: pointB1Int, b: pointB2Int}

	// Thirth segment
	pointC1Split := strings.Split(pointC1, ",")
	pointC1_0, _ := strconv.ParseInt(pointC1Split[0], 10, 0)
	pointC1_1, _ := strconv.ParseInt(pointC1Split[1], 10, 0)
	pointC1Int := [2]int64{pointC1_0, pointC1_1}
	pointC2Split := strings.Split(pointC2, ",")
	pointC2_0, _ := strconv.ParseInt(pointC2Split[0], 10, 0)
	pointC2_1, _ := strconv.ParseInt(pointC2Split[1], 10, 0)
	pointC2Int := [2]int64{pointC2_0, pointC2_1}
	segment3 := segmento{a: pointC1Int, b: pointC2Int}

	fmt.Println("--- POINTS OF SEGMENTS ---")
	fmt.Println("SEGMENT 1", pointA1, pointA2)
	fmt.Println("SEGMENT 2", pointB1, pointB2)
	fmt.Println("SEGMENT 3", pointC1, pointC2)
	fmt.Println("---")

	result := isTriangle(segment1, segment2, segment3)
	fmt.Println("Result", result)
}

func isTriangle(segment1, segment2, segment3 segmento) bool {
	if segment1.a == segment2.a {
		if segment2.b == segment3.a {
			if segment3.b == segment1.b {
				return true
			}
		} else {
			if segment2.b == segment3.b {
				if segment3.a == segment1.b {
					return true
				}
			}
		}
	} else {
		if segment1.a == segment2.b {
			if segment2.a == segment3.a {
				if segment3.b == segment1.b {
					return true
				}
			} else {
				if segment2.a == segment3.b {
					if segment3.a == segment1.b {
						return true
					}
				}
			}
		}
	}
	if segment1.b == segment2.a {
		if segment2.b == segment3.a {
			if segment3.b == segment1.a {
				return true
			}
		} else {
			if segment2.b == segment3.b {
				if segment3.a == segment1.a {
					return true
				}
			}
		}
	} else {
		if segment1.b == segment2.b {
			if segment2.a == segment3.a {
				if segment3.b == segment1.a {
					return true
				}
			} else {
				if segment2.a == segment3.b {
					if segment3.a == segment1.a {
						return true
					}
				}
			}
		}
	}
	return false
}

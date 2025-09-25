package main

import "fmt"

func main() {
	array := []int{2, 4, 5, 7, 1, 2, 3, 6}
	merge(array, 0, len(array)-1)
	fmt.Println(array)
}

func merge(a []int, bajo, alto int) {
	longitud := alto - bajo + 1
	pivote := (bajo + alto) / 2
	temporal := make([]int, longitud)

	for i := range longitud {
		temporal[i] = a[bajo+i]
	}

	m1 := 0
	m2 := pivote - bajo + 1

	for i := range longitud {
		if m2 <= alto-bajo {
			if m1 <= pivote-bajo {
				if temporal[m1] > temporal[m2] {
					a[i+bajo] = temporal[m2]
					m2++
				} else {
					a[i+bajo] = temporal[m1]
					m1++
				}
			} else {
				a[i+bajo] = temporal[m2]
				m2++
			}
		} else {
			a[i+bajo] = temporal[m1]
			m1++
		}
	}
}

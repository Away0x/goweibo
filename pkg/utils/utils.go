package utils

import (
	"github.com/teris-io/shortid"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func UniqueList(arr1 []uint, arr2 []uint) (arr3 []uint) {
	l := append(arr1, arr2...)
	m := make(map[uint]bool, 0)
	arr3 = make([]uint, 0)

	for _, v := range l {
		m[v] = true
	}

	for k := range m {
		arr3 = append(arr3, k)
	}

	return
}

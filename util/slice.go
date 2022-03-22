package util

import (
	"bytes"
	"fmt"
	"golang.org/x/exp/constraints"
)

type Comparable[T any] interface {
	Equals(other T) bool
}

func InSliceI[T Comparable[T]](a T, list []T) (bool, int) {
	for i, b := range list {
		if b.Equals(a) {
			return true, i
		}
	}
	return false, -1
}

func InSlice[T comparable](a T, list []T) (bool, int) {
	for i, b := range list {
		if b == a {
			return true, i
		}
	}
	return false, -1
}

func MapToStrings[T any](a []T) []string {
	result := make([]string, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = fmt.Sprint(a[i])
	}

	return result
}

func RemoveFromSlice[T any](slice []T, idx int) []T {
	return append(slice[:idx:idx], slice[idx+1:]...)
}

func CloneSlice[T any](slice []T) []T {
	cloned := make([]T, len(slice))
	for i := 0; i < len(slice); i++ {
		cloned[i] = slice[i]
	}

	return cloned
}

func Max[T constraints.Ordered](array []T) (max T, index int) {
	result := array[0]
	index = 0
	for i, v := range array[1:] {
		if v > result {
			result, index = v, i+1
		}
	}
	return result, index
}

func Min[T constraints.Ordered](array []T) (max T, index int) {
	result := array[0]
	index = 0
	for i, v := range array[1:] {
		if v < result {
			result, index = v, i+1
		}
	}
	return result, index
}

func Joint[T any](array []T, sep string) string {
	var buf bytes.Buffer
	first := false
	for _, val := range array {
		if !first {
			buf.WriteString(sep)
		}
		first = false
		buf.WriteString(fmt.Sprint(val))
	}

	return buf.String()
}

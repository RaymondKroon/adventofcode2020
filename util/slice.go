package util

import "fmt"

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "SliceType=String,Int"
//type ValueType = generic.Type

func SliceTypeInSlice(a SliceType, list []SliceType) bool {
	for _, b := range list {
		if b.Equals(a) {
			return true
		}
	}
	return false
}

func MapSliceTypesToStrings(a []SliceType) []string {
	result := make([]string, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = fmt.Sprint(a[i])
	}

	return result
}

func RemoveFromSliceTypeSlice(slice []SliceType, idx int) []SliceType {
	return append(slice[:idx:idx], slice[idx+1:]...)
}

func CloneSliceTypeSlice(slice []SliceType) []SliceType {
	cloned := make([]SliceType, len(slice))
	for i := 0; i < len(slice); i++ {
		cloned[i] = slice[i]
	}

	return cloned
}

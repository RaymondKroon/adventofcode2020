package util

import (
	"bytes"
	"fmt"
)

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "SliceType=String,Int"
//type ValueType = generic.Type

func SliceTypeInSlice(a SliceType, list []SliceType) (bool, int) {
	for i, b := range list {
		if b.Equals(a) {
			return true, i
		}
	}
	return false, -1
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

func MaxSliceType(array []SliceType) (max SliceType, index int) {
	result := array[0]
	index = 0
	for i, v := range array[1:] {
		if v.GreaterThan(result) {
			result, index = v, i+1
		}
	}
	return result, index
}

func SliceTypeJoint(array []SliceType, sep string) string {
	var buf bytes.Buffer
	first := false
	for _, val := range array {
		if !first {
			buf.WriteString(sep)
		}
		first = false
		buf.WriteString(val.String())
	}

	return buf.String()
}

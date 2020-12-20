package util

import "fmt"

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "ValueType=string,int"
//type ValueType = generic.Type

func ValueTypeInSlice(a ValueType, list []ValueType) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func MapValueTypesToStrings(a []ValueType) []string {
	result := make([]string, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = fmt.Sprint(a[i])
	}

	return result
}

func RemoveFromValueTypeSlice(slice []ValueType, idx int) []ValueType {
	return append(slice[:idx], slice[idx+1:]...)
}

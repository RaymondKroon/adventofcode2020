package util

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

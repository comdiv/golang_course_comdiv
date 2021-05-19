package main

func main() {
	_ = retSomeSlice()
	_ = retSomeSlicePointer()
	s1 := retSomeSlice()
	println(s1)
	s2 := retSomeSlicePointer()
	println(s2)
}

//go:noinline
func retSomeSlice() []int {
	var result []int
	result = append(result, 1)
	return result
}

//go:noinline
func retSomeSlicePointer() *[]int {
	var result []int
	result = append(result, 1)
	return &result
}

package main

import "fmt"

// groupTempOscillations returns groups of temperature oscillations with range equal to 10
func groupTempOscillations(temps []float32) map[int32][]float32 {
	res := map[int32][]float32{}
	for _, val := range temps {
		res[int32(val)/10*10] = append(res[int32(val)/10*10], val)
	}
	return res
}

func main() {
	tempOscillations := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(groupTempOscillations(tempOscillations))
}

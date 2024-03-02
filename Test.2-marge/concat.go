package main

import (
	"fmt"
)

func mergeArray(slice1, slice2 []string) ([]string, []string) {
	// สร้างตัวแปรเพื่อเก็บข้อมูลใน array เชื่อมกัน
	mergeSlice := make([]string, len(slice1)+len(slice2))
	for i, values := range slice1 {
		mergeSlice[i] = values
	}
	for i, values := range slice2 {
		mergeSlice[len(slice1)+i] = values
	}
	return mergeSlice[:len(mergeSlice)-1], mergeSlice[len(mergeSlice)-2:]
	// ----------------------------------->รีเทิร์นค่าของ mergeSlice หรือ slice2 ให้มีค่าว่าง
}

func main() {
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{"b", "c", "d"}

	// เพิ่มค่า d ให้กับ slice1
	slice1 = append(slice1, "d")
	// เพิ่มค่า a ให้กับ slice2
	slice2 = append(slice2, "a")
	// เพิ่มค่า d ให้กับ slice2
	slice2 = append(slice2, "d")

	mergeSlices1, mergeSlices2 := mergeArray(slice1, slice2)

	fmt.Println(mergeSlices1[0:4]) //แสดงข้อมูลตำแหน่งข้อมูลที่ 0 ถึง 4

	fmt.Println(mergeSlices2[0:2]) //แสดงข้อมูลตำแหน่งข้อมูลที่ 0 ถึง 2

}

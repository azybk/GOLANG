package main

import "fmt"

func main() {
	// membuat dan inisialisasi Array
	// months := [...]string{
	// 	"Januari", "Pebruari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "Nopember", "Desember",
	// }

	// membuat slice dan reference ke array months
	// slice1 := months[4:7]
	// fmt.Println(slice1)
	// fmt.Println(len(slice1))
	// fmt.Println(cap(slice1))

	// slice2 := months[2:4]
	// slice2 := months[10:]
	// fmt.Println(slice2)

	// slice3 := append(slice2, "Saya")
	// fmt.Println(slice3)

	// slice3[0] = "Ubah"
	// fmt.Println(slice3)
	// fmt.Println(slice2)
	// fmt.Println(months)

	// membuat slice baru tanpa reference ke array lain
	// newSlice := make([]string, 2, 5)
	// newSlice[0] = "Bahasa"
	// newSlice[1] = "Golang"
	// fmt.Println(len(newSlice))
	// fmt.Println(cap(newSlice))
	// fmt.Println(newSlice)

	// mengcopy slice
	// copySlice := make([]string, len(newSlice), cap(newSlice))
	// copy(copySlice, newSlice)
	// fmt.Println(copySlice)
	// fmt.Println(newSlice)

	// perbedaan pembuatan array dan slice
	iniArray := [5]int{1, 2, 3, 4, 5}
	iniArrayLagi := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniArrayLagi)
	fmt.Println(iniSlice)

}

package main

import "fmt"

func main() {
	// Membuat slice bernama sliceData
	sliceData := []int{10, 20, 30, 40, 50}

	// Cetak sliceData 
	fmt.Printf("===[mySlice]===\n")
	fmt.Printf("Data: %d\n", sliceData)
	fmt.Printf("Len: %d\n",len(sliceData))
	fmt.Printf("Cap: %d\n",cap(sliceData))
	fmt.Println()

	// Membuat slice baru dari sliceData
	sliceData1 := sliceData[0:3]

	// Cetak sliceData1
	fmt.Printf("===[Slice1]===\n")
	fmt.Printf("Data: %d\n", sliceData1)
	fmt.Printf("Len: %d\n",len(sliceData1))
	fmt.Printf("Cap: %d\n",cap(sliceData1))
	fmt.Println()

	// Membuat slice baru dari sliceData
	sliceData2 := sliceData[3:]

	fmt.Printf("===[Slice1]===\n")
	fmt.Printf("Data: %d\n", sliceData2)
	fmt.Printf("Len: %d\n",len(sliceData2))
	fmt.Printf("Cap: %d\n",cap(sliceData2))
	fmt.Println()

	// Menambahkan data ke sliceData1
	sliceData1 = append(sliceData1, 60)

	fmt.Printf("===[Setelah Append ke Slice1]===\n")
	fmt.Printf("Data Slice1: %d\n", sliceData1)
	fmt.Printf("Len Slice1: %d\n",len(sliceData1))
	fmt.Printf("Cap Slice1: %d\n",cap(sliceData1))
	fmt.Printf("Data Slice2: %d\n", sliceData2)
	fmt.Printf("Len Slice2: %d\n",len(sliceData2))
	fmt.Printf("Cap Slice2: %d\n",cap(sliceData2))
	fmt.Println()

	// Merubah data index 0 sliceData2
	sliceData2[0] = 60

	// Menambahkan data ke sliceData2
	sliceData2 = append(sliceData2, 70)


	fmt.Printf("===[Setelah Append ke Slice2]===\n")
	fmt.Printf("Data Slice1: %d\n", sliceData1)
	fmt.Printf("Len Slice1: %d\n",len(sliceData1))
	fmt.Printf("Cap Slice1: %d\n",cap(sliceData1))
	fmt.Printf("Data Slice2: %d\n", sliceData2)
	fmt.Printf("Len Slice2: %d\n",len(sliceData2))
	fmt.Printf("Cap Slice2: %d\n",cap(sliceData2))
}

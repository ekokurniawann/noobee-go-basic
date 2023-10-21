package main

import (
	"fmt"
	"sort"
)

type Examination struct {
	Name  string
	Score int
}

func main() {
	examResult := []Examination{
		{Name: "Budi", Score: 90},
		{Name: "Andi", Score: 85},
		{Name: "Nayla", Score: 87},
		{Name: "Danu", Score: 80},
		{Name: "Rahmat", Score: 75},
		{Name: "Erika", Score: 83},
		{Name: "Siska", Score: 87},
		{Name: "Mita", Score: 94},
		{Name: "Shinta", Score: 82},
		{Name: "Jojo", Score: 83},
	}

	// Urutkan data berdasarkan nilai (Score)
	sort.Slice(examResult, func(i, j int) bool {
		return examResult[i].Score < examResult[j].Score
	})

	// Median
	median := findMedian(examResult)

	// Max
	maximums := findMaximums(examResult)

	// Min
	minimums := findMinimums(examResult)

	// Average
	average := findAverage(examResult)

	// Nilai di atas 80 (menggunakan Binary Search)
	above80 := findAbove80(examResult)

	// Cetak hasil
	fmt.Printf("Median: %d\n", median)
	fmt.Printf("Max: %v\n", maximums)
	fmt.Printf("Min: %v\n", minimums)
	fmt.Printf("Average: %.2f\n", average)
	fmt.Printf("Nilai di atas 80: %v\n", above80)
}

// Fungsi untuk mencari median
func findMedian(data []Examination) int {
	n := len(data)
	if n%2 == 1 {
		return data[n/2].Score
	}
	return (data[n/2-1].Score + data[n/2].Score) / 2
}

// Fungsi untuk mencari nilai maksimum
func findMaximums(data []Examination) []Examination {
	maximums := make([]Examination, 0)
	maxScore := data[len(data)-1].Score
	for i := len(data) - 1; i >= 0; i-- {
		if data[i].Score == maxScore {
			maximums = append(maximums, data[i])
		} else {
			break
		}
	}
	return maximums
}

// Fungsi untuk mencari nilai minimum
func findMinimums(data []Examination) []Examination {
	minimums := make([]Examination, 0)
	minScore := data[0].Score
	for i := 0; i < len(data); i++ {
		if data[i].Score == minScore {
			minimums = append(minimums, data[i])
		} else {
			break
		}
	}
	return minimums
}

// Fungsi untuk menghitung rata-rata
func findAverage(data []Examination) float64 {
	total := 0
	for _, exam := range data {
		total += exam.Score
	}
	return float64(total) / float64(len(data))
}

// Fungsi untuk mencari nilai di atas 80 dengan binary search
func findAbove80(data []Examination) []Examination {
	result := make([]Examination, 0)
	low, high := 0, len(data)-1

	for low <= high {
		mid := (low + high) / 2
		if data[mid].Score > 80 {
			result = append(result, data[mid])
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return result
}


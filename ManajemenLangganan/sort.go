package main

import "fmt"

// SortBiaya menggunakan Selection Sort untuk mengurutkan berdasarkan biaya.
func SortBiaya(data *[100]Langganan, count int, ascending bool) {
    for i := 0; i < count-1; i++ {
        idx := i
        for j := i + 1; j < count; j++ {
            if (ascending && data[j].BiayaBulanan < data[idx].BiayaBulanan) ||
                (!ascending && data[j].BiayaBulanan > data[idx].BiayaBulanan) {
                idx = j
            }
        }
        data[i], data[idx] = data[idx], data[i]
    }
}

// SortNama menggunakan Insertion Sort untuk mengurutkan berdasarkan nama.
func SortNama(data *[100]Langganan, count int, ascending bool) {
    for i := 1; i < count; i++ {
        temp := data[i]
        j := i - 1
        for j >= 0 && ((ascending && data[j].NamaLayanan > temp.NamaLayanan) ||
            (!ascending && data[j].NamaLayanan < temp.NamaLayanan)) {
            data[j+1] = data[j]
            j--
        }
        data[j+1] = temp
    }
}

package langganan

import (
    "fmt"
    "strings"
)

func MenuPengurutan() {
    fmt.Println("----- Pengurutan Layanan -----")
    fmt.Println("1. Selection Sort (ASC)")
    fmt.Println("2. Insertion Sort (DESC)")
    fmt.Print("Pilih metode: ")
    var metode int
    fmt.Scanln(&metode)

    switch metode {
    case 1:
        SelectionSortBiaya()
    case 2:
        InsertionSortNama()
    default:
        fmt.Println("Metode tidak valid.")
    }

    fmt.Println("Data setelah diurutkan:")
    TampilkanLangganan()
}

func SelectionSortBiaya() {
    for i := 0; i < JumlahLangganan-1; i++ {
        minIdx := i
        for j := i + 1; j < JumlahLangganan; j++ {
            if DaftarLangganan[j].BiayaBulanan < DaftarLangganan[minIdx].BiayaBulanan {
                minIdx = j
            }
        }
        DaftarLangganan[i], DaftarLangganan[minIdx] = DaftarLangganan[minIdx], DaftarLangganan[i]
    }
}

func InsertionSortNama() {
    for i := 1; i < JumlahLangganan; i++ {
        key := DaftarLangganan[i]
        j := i - 1
        for j >= 0 && strings.ToLower(DaftarLangganan[j].NamaLayanan) > strings.ToLower(key.NamaLayanan) {
            DaftarLangganan[j+1] = DaftarLangganan[j]
            j--
        }
        DaftarLangganan[j+1] = key
    }
}
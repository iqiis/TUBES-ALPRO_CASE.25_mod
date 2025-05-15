package langganan

import (
    "fmt"
    "strings"
)

func MenuPencarian() {
    fmt.Println("----- Pencarian Layanan -----")
    fmt.Println("1. Sequential Search")
    fmt.Println("2. Binary Search (Data Harus Terurut)")
    fmt.Print("Pilih metode: ")
    var metode int
    fmt.Scanln(&metode)

    fmt.Print("Masukkan nama layanan: ")
    var keyword string
    fmt.Scanln(&keyword)

    switch metode {
    case 1:
        SequentialSearch(keyword)
    case 2:
        BinarySearch(keyword)
    default:
        fmt.Println("Metode tidak valid.")
    }
}

func SequentialSearch(keyword string) {
    found := false
    for i := 0; i < JumlahLangganan; i++ {
        if strings.Contains(strings.ToLower(DaftarLangganan[i].NamaLayanan), strings.ToLower(keyword)) {
            fmt.Printf("%+v\n", DaftarLangganan[i])
            found = true
        }
    }
    if !found {
        fmt.Println("Data tidak ditemukan.")
    }
}

func BinarySearch(keyword string) {
    left, right := 0, JumlahLangganan-1
    found := false
    for left <= right {
        mid := (left + right) / 2
        cmp := strings.Compare(strings.ToLower(DaftarLangganan[mid].NamaLayanan), strings.ToLower(keyword))
        if cmp == 0 {
            fmt.Printf("%+v\n", DaftarLangganan[mid])
            found = true
            break
        } else if cmp < 0 {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    if !found {
        fmt.Println("Data tidak ditemukan.")
    }
}
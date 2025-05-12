package main

import (
    "fmt"
    "strings"
)

// Pencarian menggunakan Sequential dan Binary Search
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
    for i := 0; i < jumlahLangganan; i++ {
        if strings.Contains(strings.ToLower(daftarLangganan[i].NamaLayanan), strings.ToLower(keyword)) {
            TampilkanSatuLangganan(daftarLangganan[i])
            found = true
        }
    }
    if !found {
        fmt.Println("Data tidak ditemukan.")
    }
}

func BinarySearch(keyword string) {
    left, right := 0, jumlahLangganan-1
    found := false
    for left <= right {
        mid := (left + right) / 2
        cmp := strings.Compare(strings.ToLower(daftarLangganan[mid].NamaLayanan), strings.ToLower(keyword))
        if cmp == 0 {
            TampilkanSatuLangganan(daftarLangganan[mid])
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

func TampilkanSatuLangganan(l Langganan) {
    aktifStr := "Tidak"
    if l.Aktif {
        aktifStr = "Ya"
    }
    fmt.Printf("Nama Layanan: %s\nKategori: %s\nBiaya: Rp %d\nJatuh Tempo: %s\nMetode Pembayaran: %s\nAktif: %s\n",
        l.NamaLayanan, l.Kategori, l.BiayaBulanan, l.TanggalJatuhTempo, l.MetodePembayaran, aktifStr)
}

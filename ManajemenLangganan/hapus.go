package main

import (
    "fmt"
    "strings"
)

// HapusLangganan menghapus data berdasarkan nama layanan.
func HapusLangganan(data *[100]Langganan, count *int) {
    var nama string
    fmt.Println("----- Hapus Layanan -----")
    fmt.Print("Masukkan nama layanan yang ingin dihapus: ")
    fmt.Scanln(&nama)

    idx := CariIndexByNama(data, *count, nama)
    if idx == -1 {
        fmt.Println("Layanan tidak ditemukan.")
        return
    }

    for i := idx; i < *count-1; i++ {
        data[i] = data[i+1]
    }
    *count--
    fmt.Println("Layanan berhasil dihapus!")
}

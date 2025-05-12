package main

import (
    "fmt"
    "strings"
)

// EditLangganan mengedit data berdasarkan nama layanan.
func EditLangganan(data *[100]Langganan, count int) {
    var nama string
    fmt.Println("----- Edit Layanan -----")
    fmt.Print("Masukkan nama layanan yang ingin diubah: ")
    fmt.Scanln(&nama)

    idx := CariIndexByNama(data, count, nama)
    if idx == -1 {
        fmt.Println("Layanan tidak ditemukan.")
        return
    }

    l := &data[idx]
    fmt.Printf("Data saat ini: %v
", *l)
    fmt.Print("Nama Layanan Baru (" + l.NamaLayanan + "): ")
    fmt.Scanln(&l.NamaLayanan)
    fmt.Print("Kategori Baru (" + l.Kategori + "): ")
    fmt.Scanln(&l.Kategori)
    fmt.Print("Biaya Bulanan Baru (" + fmt.Sprint(l.BiayaBulanan) + "): ")
    fmt.Scanln(&l.BiayaBulanan)
    fmt.Print("Tanggal Jatuh Tempo Baru (dd-mm-yyyy) (" + l.TanggalJatuhTempo + "): ")
    fmt.Scanln(&l.TanggalJatuhTempo)
    fmt.Print("Metode Pembayaran Baru (" + l.MetodePembayaran + "): ")
    fmt.Scanln(&l.MetodePembayaran)
    fmt.Print("Status Aktif Baru (y/n): ")
    var status string
    fmt.Scanln(&status)
    l.Aktif = strings.ToLower(status) == "y"

    fmt.Println("Data berhasil diperbarui!")
}

// CariIndexByNama mencari index data berdasarkan nama layanan
func CariIndexByNama(data *[100]Langganan, count int, nama string) int {
    for i := 0; i < count; i++ {
        if strings.EqualFold(data[i].NamaLayanan, nama) {
            return i
        }
    }
    return -1
}
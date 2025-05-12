package main

import (
    "fmt"
    "strings"
)

// TambahLangganan menambahkan data langganan baru ke array utama.
func TambahLangganan(data *[100]Langganan, count *int) {
    if *count >= len(data) {
        fmt.Println("Data langganan penuh.")
        return
    }

    var l Langganan
    fmt.Println("----- Tambah Layanan Baru -----")
    fmt.Print("Nama Layanan        : ")
    fmt.Scanln(&l.NamaLayanan)
    fmt.Print("Kategori            : ")
    fmt.Scanln(&l.Kategori)
    fmt.Print("Biaya Bulanan (Rp)  : ")
    fmt.Scanln(&l.BiayaBulanan)
    fmt.Print("Tanggal Jatuh Tempo (format: dd-mm-yyyy): ")
    fmt.Scanln(&l.TanggalJatuhTempo)
    fmt.Print("Metode Pembayaran   : ")
    fmt.Scanln(&l.MetodePembayaran)
    fmt.Print("Status Aktif (y/n)  : ")
    var status string
    fmt.Scanln(&status)
    l.Aktif = strings.ToLower(status) == "y"

    data[*count] = l
    *count++
    fmt.Println("Data berhasil ditambahkan!")
}

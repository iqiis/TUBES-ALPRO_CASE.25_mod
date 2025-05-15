package langganan

import (
    "fmt"
    "strings"
)

func TampilkanMenu() {
    fmt.Println("=======================================")
    fmt.Println(" MANAJEMEN SUBSKRIPSI DIGITAL")
    fmt.Println("=======================================")
    fmt.Println("1. Tambah Layanan Baru")
    fmt.Println("2. Tampilkan Semua Langganan")
    fmt.Println("3. Edit Layanan")
    fmt.Println("4. Hapus Layanan")
    fmt.Println("5. Cari Layanan")
    fmt.Println("6. Urutkan Langganan")
    fmt.Println("7. Cek Reminder Jatuh Tempo")
    fmt.Println("8. Total Pengeluaran Bulanan")
    fmt.Println("9. Rekomendasi Penghematan")
    fmt.Println("0. Keluar")
    fmt.Println("---------------------------------------")
}

func TambahLangganan() {
    if JumlahLangganan >= len(DaftarLangganan) {
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

    DaftarLangganan[JumlahLangganan] = l
    JumlahLangganan++
    fmt.Println("Data berhasil ditambahkan!")
}

func TampilkanLangganan() {
    fmt.Println("--------------------------- Daftar Langganan ---------------------------")
    fmt.Printf("| %-3s | %-12s | %-7s | %-12s | %-10s | %-6s |\n", "No", "Nama", "Biaya", "Jatuh Tempo", "Kategori", "Aktif")
    fmt.Println("----------------------------------------------------------------------")
    for i := 0; i < JumlahLangganan; i++ {
        aktifStr := "Tidak"
        if DaftarLangganan[i].Aktif {
            aktifStr = "Ya"
        }
        fmt.Printf("| %-3d | %-12s | %-7d | %-12s | %-10s | %-6s |\n",
            i+1,
            DaftarLangganan[i].NamaLayanan,
            DaftarLangganan[i].BiayaBulanan,
            DaftarLangganan[i].TanggalJatuhTempo,
            DaftarLangganan[i].Kategori,
            aktifStr)
    }
    fmt.Println("----------------------------------------------------------------------")
    fmt.Printf("Total: %d layanan\n", JumlahLangganan)
}

func EditLangganan() {
    var nama string
    fmt.Println("----- Edit Layanan -----")
    fmt.Print("Masukkan nama layanan yang ingin diubah: ")
    fmt.Scanln(&nama)

    index := -1
    for i := 0; i < JumlahLangganan; i++ {
        if strings.EqualFold(DaftarLangganan[i].NamaLayanan, nama) {
            index = i
            break
        }
    }

    if index == -1 {
        fmt.Println("Layanan tidak ditemukan.")
        return
    }

    var l *Langganan = &DaftarLangganan[index]
    fmt.Print("Nama Layanan Baru (" + l.NamaLayanan + "): ")
    fmt.Scanln(&l.NamaLayanan)
    fmt.Print("Kategori Baru (" + l.Kategori + "): ")
    fmt.Scanln(&l.Kategori)
    fmt.Print("Biaya Bulanan Baru (" + fmt.Sprint(l.BiayaBulanan) + "): ")
    fmt.Scanln(&l.BiayaBulanan)
    fmt.Print("Tanggal Jatuh Tempo Baru (dd-mm-yyyy): ")
    fmt.Scanln(&l.TanggalJatuhTempo)
    fmt.Print("Metode Pembayaran Baru (" + l.MetodePembayaran + "): ")
    fmt.Scanln(&l.MetodePembayaran)
    fmt.Print("Status Aktif Baru (y/n): ")
    var status string
    fmt.Scanln(&status)
    l.Aktif = strings.ToLower(status) == "y"

    fmt.Println("Data berhasil diperbarui!")
}

func HapusLangganan() {
    var nama string
    fmt.Println("----- Hapus Layanan -----")
    fmt.Print("Masukkan nama layanan yang ingin dihapus: ")
    fmt.Scanln(&nama)

    index := -1
    for i := 0; i < JumlahLangganan; i++ {
        if strings.EqualFold(DaftarLangganan[i].NamaLayanan, nama) {
            index = i
            break
        }
    }

    if index == -1 {
        fmt.Println("Layanan tidak ditemukan.")
        return
    }

    for i := index; i < JumlahLangganan-1; i++ {
        DaftarLangganan[i] = DaftarLangganan[i+1]
    }
    JumlahLangganan--

    fmt.Println("Layanan berhasil dihapus!")
}

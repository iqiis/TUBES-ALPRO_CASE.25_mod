package langganan

import (
    "fmt"
)

func TotalPengeluaran() {
    total := 0
    for i := 0; i < JumlahLangganan; i++ {
        if DaftarLangganan[i].Aktif {
            total += DaftarLangganan[i].BiayaBulanan
        }
    }
    fmt.Printf("Total pengeluaran bulanan aktif: Rp %d\n", total)
}

func RekomendasiPenghematan() {
    fmt.Println("----- Rekomendasi Penghematan -----")
    fmt.Println("Berikut layanan non-aktif atau biaya tinggi (di atas Rp 100.000):")
    rekomendasiDitemukan := false
    for i := 0; i < JumlahLangganan; i++ {
        l := DaftarLangganan[i]
        if !l.Aktif || l.BiayaBulanan > 100000 {
            fmt.Printf("- %s | Aktif: %v | Biaya: Rp %d\n", l.NamaLayanan, l.Aktif, l.BiayaBulanan)
            rekomendasiDitemukan = true
        }
    }

    if !rekomendasiDitemukan {
        fmt.Println("Semua langganan tampak optimal.")
    }
}
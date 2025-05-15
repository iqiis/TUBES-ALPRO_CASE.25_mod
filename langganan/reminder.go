package langganan

import (
    "fmt"
    "time"
)

func ReminderJatuhTempo() {
    fmt.Println("----- Reminder Jatuh Tempo dalam 7 Hari -----")
    sekarang := time.Now()
    batas := sekarang.AddDate(0, 0, 7)

    ditemukan := false
    for i := 0; i < JumlahLangganan; i++ {
        if AdaJatuhTempoDalamBatas(DaftarLangganan[i].TanggalJatuhTempo, sekarang, batas) {
            fmt.Printf("- %s (Jatuh tempo: %s)\n", DaftarLangganan[i].NamaLayanan, DaftarLangganan[i].TanggalJatuhTempo)
            ditemukan = true
        }
    }

    if !ditemukan {
        fmt.Println("Tidak ada layanan yang akan jatuh tempo dalam 7 hari.")
    }
}
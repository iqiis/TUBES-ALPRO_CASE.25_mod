package main

import (
    "fmt"
    "manajemen_subs/langganan"
)

func main() {
    for {
        langganan.TampilkanMenu()
        var pilihan int
        fmt.Print("Pilih menu: ")
        fmt.Scanln(&pilihan)

        switch pilihan {
        case 1:
            langganan.TambahLangganan()
        case 2:
            langganan.TampilkanLangganan()
        case 3:
            langganan.EditLangganan()
        case 4:
            langganan.HapusLangganan()
        case 5:
            langganan.MenuPencarian()
        case 6:
            langganan.MenuPengurutan()
        case 7:
            langganan.ReminderJatuhTempo()
        case 8:
            langganan.TotalPengeluaran()
        case 9:
            langganan.RekomendasiPenghematan()
        case 0:
            fmt.Println("Keluar dari program...")
            return
        default:
            fmt.Println("Pilihan tidak valid")
        }
        fmt.Println()
    }
}
package app
import "fmt"

func JalankanProgram() {
    for {
        TampilkanMenu()
        var pilihan int
        fmt.Print("Pilih menu: ")
        fmt.Scanln(&pilihan)

        if pilihan == 0 {
            fmt.Println("Keluar dari program...")
            return
        }

        ProsesPilihanMenu(pilihan)
        fmt.Println()
    }
}

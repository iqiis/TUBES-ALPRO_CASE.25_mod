// bantu tanggal jatuh tempo
package langganan

import (
    "time"
    "fmt"
)

func AdaJatuhTempoDalamBatas(tanggalStr string, dari, sampai time.Time) bool {
    layout := "02-01-2006"
    tgl, err := time.Parse(layout, tanggalStr)
    if err != nil {
        fmt.Println("Format tanggal salah:", tanggalStr)
        return false
    }

    tahunSekarang, bulanSekarang, _ := dari.Date()
    tgl = time.Date(tahunSekarang, bulanSekarang, tgl.Day(), 0, 0, 0, 0, dari.Location())

    return !tgl.Before(dari) && !tgl.After(sampai)
}
package langganan

type Langganan struct {
    NamaLayanan       string
    Kategori          string
    BiayaBulanan      int
    TanggalJatuhTempo string
    MetodePembayaran  string
    Aktif             bool
}

var DaftarLangganan [100]Langganan
var JumlahLangganan int

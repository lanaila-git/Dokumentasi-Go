// Sebuah aplikasi tiket bus ingin membantu penumpang memilih tujuan perjalanan,
// waktu keberangkatan, dan tipe kursi. Harga tiket berbeda tergantung rute dan
// waktu keberangkatan. Penumpang juga bisa mendapatkan promo tertentu jika
// memenuhi syarat tertentu. Setelah semua pilihan dibuat, sistem menampilkan
// total harga dan status tiket.
package main

import "fmt"

type bus struct {
  tujuan string
  waktu string
  tipe string
}
func main () {
  var pilih bus
  fmt.Println("--Pilih tiket yang anda butuhkan--")
  tujuan()
}

func tujuan() {
  var pilih int
  fmt.Println("-Pilih tujuan anda-")
  fmt.Println("1. Kota A")
  fmt.Println("2. Kota B")
  fmt.Println("3. Kota C")
  fmt.Println("4. Kota D")
  fmt.Println("5. Kota E")
  fmt.Print("\nPilih: ")
  fmt.Scan(&pilih)
  waktu()
}

func waktu() {
  var pilih bus
  fmt.Println("-kapan waktu keberangkatan anda-")
  fmt.Scan(&pilih.waktu)
}

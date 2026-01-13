// Seorang kasir di toko kecil harus menghitung total belanja pelanggan yang
// membeli berbagai barang dengan harga berbeda. Kadang pelanggan mendapat potongan
// harga berdasarkan total belanja atau jenis barang tertentu. Setelah transaksi selesai,
// kasir perlu menampilkan ringkasan pembelian beserta keterangan
// apakah pelanggan mendapatkan diskon atau tidak.

// rule
// variabel biasa X apabila diinisiasi sebuah nilai, ketika variabel di print tanpa '&' maka akan menghasilkan nilainya
// var X int
// cth: X = 7
// fmt.Println(X) --> 7
// namun ketika di print dengan '&' akan menghasilkan alamat memorinya
// fmt.Println(&X) --> 0x0000##
// variabel pointer Y apabila diinisiasi(menunjuk) sebuah nilai atau variabel biasa X, 
// ketika variabel di print tanpa '*' maka akan menghasilkan alamat memori dari X.
// cth: var Y *int
// *Y = X
// fmt.Println(Y) --> 0x0000##
// namun ketika di print dengan '*' akan menghasilkan nilai X
// fmt.Println(*Y) --> 7
// ini berarti Y menunjuk variabel X sehingga mampu mengakses nilainya

package main

import "fmt"

type dataBarang struct {
  nama, jenis string
  jumlah int
  harga float64
}

func main () {
  var i, n int
  var produk dataBarang

  fmt.Print("Jumlah barang yang ingin di input: ")
  fmt.Scan(&n)
  for i = 1; i <= n; i++ {
    fmt.Print("Nama barang: ")
    fmt.Scan(&produk.nama)
    fmt.Print("Harga: ")
    fmt.Scan(&produk.harga)
    fmt.Print("Jumlah: ")
    fmt.Scan(&produk.jumlah)
    fmt.Print("Jenis: ")
    fmt.Scan(&produk.jenis) //hanya jika "Promo" yang akan medapatkan diskon
    sistem(&produk.harga, &produk.jumlah, &produk.jenis, &produk.nama)
  }
}

func sistem(*h float64, *sum int, jns, nm *string) {
  var calculation float64
  if sum > 2 && jenis != "Promo"{
    calculation = discCountBySum(h, sum, jns)
    fmt.Printf("Harga total untuk %s\n", nm)
    fmt.Print("Rp.", sum * h)
    fmt.Print("Harga setelah diskon: Rp. ", calculation)
    fmt.Println()
    sumOfShopping(&calculation)
  } else if sum <= 2  && sum > 0 && jenis != "Promo"{
    fmt.Printf("Harga total untuk %s\n", nm)
    fmt.Print("Rp. ")
    calculation = normalCountBelanja(h, sum, jns)
    fmt.Println()
    sumOfShopping(&calculation)
  } else if jenis == "Promo" {
    calculation = disCountByJns(h, sum, jns)
    fmt.Printf("Harga total untuk %s\n", nm)
    fmt.Print("Rp.", sum * h)
    fmt.Print("Harga setelah diskon: Rp. ", calculation)
    fmt.Println()
    sumOfShopping(&calculation)
  }
}

func normalCountBelanja(harga float64, jumlah int, jenis string) float64 {
  return harga * jumlah
}

func dicCountBySum(harga float64, jumlah int, jenis string) float64 {
  var disc, total float64
  if jumlah == 3 {
    total = jumlah * harga
    disc = total * 0.25
    total = total - disc
  } else if jumlah > 3 {
    total = jumlah * harga
    disc = total * 0.6
    total = total - disc
  }

  return total
}

func dicCountByJns(harga float64, jumlah int, jenis string) float64 {
  var total, disc float64
  total = jumlah * harga
  disc = total * 0.45
  return total - disc
}

func sumOfShopping(calcu *float64) {
  var total float64
  total = total + calcu
  fmt.Println("Total belanja anda kali ini adalah")
  fmt.Println(calcu)
}

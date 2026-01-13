// Bayangkan kamu membuat sistem yang memantau kondisi cuaca harian
// selama seminggu. Setiap hari, sistem mencatat suhu, curah hujan,
// dan status cuaca (cerah, mendung, hujan). Di akhir minggu, sistem
// akan menunjukkan tren cuaca: apakah minggu ini cenderung panas, sejuk, atau hujan terus.

package main

import "fmt"

func main () {

  for i = 1; i <=7; i++ {
    fmt.Println("Bagaimana cuaca hari ini: ")
    fmt.Scan(&cuaca)
    if cuaca == "cerah" {
      crh = crh + 1
    } else if cuaca == "mendung" {
      mdg = mdg + 1
    } else if cuaca == "hujan" {
      hjn = hjn + 1
    }
  }
  if crh > mdg && crh > hjn {
    fmt.Println("predikasi cuaca akan cenderung panas")
  } else if mdg > crh && mdg > hjn {
    fmt.Println("predikasi cuaca akan cenderung sejuk")
  } else if hjn > mdg && hjn > crh {
    fmt.Println("predikasi cuaca akan cenderung hujan")
  } else {
    fmt.Println("tren cuaca tidak dapat di prediksi")
  }
}


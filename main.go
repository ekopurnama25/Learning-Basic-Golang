package main

import (
	"fmt"
	L "myproject/calculator"
	"time"
)

func Segitiga() {

}

func main() {
	L.Array()
	fmt.Println("Hello Dunia, Ini Merupakan Perkenalan Saya Sebagai Software Developer")
	//ini adalah komputer yang tidak akan ter eksekusi dalam bahasa program
	//fmt.Println("Ini adalah Komputer")

	/* Ini Merupakan

	fmt.Println("Komentar Multi Line")

	Komentar Multi Line */

	/*
		1. Integer (int)
		2. String ("")
		3. Float
		4. Bool
		5. Nil and Zero Value
	*/

	//Variabel
	//Cara Pertama
	var Nama string
	Nama = "Eko Purnama Azi"
	fmt.Println(Nama)
	//Cara Kedua
	var umur int = 20
	fmt.Println("umur saya adalah", umur)
	//Cara Ketiga
	alamat := "Serang"
	fmt.Println("Domisili Saya Di", alamat)
	//Aku mau ngubah Domisili
	alamat = "Cilegon"
	fmt.Println(alamat)

	//Kostanta = adalah Nilai Yang tidak bis diubah
	const Phi float32 = 3.14
	fmt.Println(Phi)

	var (
		Nama1 string
		Nama2 string = "Ayam"
	)

	Nama1 = "Sate"
	fmt.Println(Nama1, Nama2)
	fmt.Println(Nama2, Nama1)

	//Seleksi Kondisi atau Oprasi
	//Oprasi Aritmatika
	fmt.Println("-----Oprasi-----")
	var value = (((2+6)%3)*4 - 2) / 3

	fmt.Println(value)

	//Opraso Perbandingan
	var isEquel = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEquel)

	//Oprator Logika

	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)

	fmt.Println("-------Seleksi Kondisi-------")

	nilai_ujian := 77

	if nilai_ujian >= 80 {
		fmt.Println("Kamu Hebat")
	} else if nilai_ujian > 60 {
		fmt.Println("GGP Kamu udah berusaha")
	} else {
		fmt.Println("Remedial")
	}

	var point = 10

	switch {
	case point >= 8:
		fmt.Println("Kamu Hebat")
	case point > 7:
		fmt.Println("GGP Kamu udah berusaha")
	default:
		fmt.Println("Remedial")
	}

	fmt.Println("----- Perulangan For ---------")

	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	fmt.Println("------ Function -------")
	Persegi()
	fmt.Println("Luasnya Adalah", LuasSegitiga())

	Perkenalan("Sate Ayam", 20, "Serang")

}

func Perkenalan(nama string, umur int, domisili string) {
	fmt.Println("Hallo Nama Saya adalah:", nama, "Umur Saya adalah:", umur, "Domisili Saya Adalah:", domisili)
	fmt.Println("Waktu Perkenalan Saya adalah", time.Now())
}

func Persegi() {
	sisi := 3
	//menghitung Luas Persegi Panjang
	luas := sisi * sisi
	fmt.Println("Luas Dari Persegi Panjang Adalah", luas)

	keliling := sisi * 4
	fmt.Println("Keliling Persegi Panjang Adalah", keliling)
}

func LuasSegitiga() float32 {
	//Menghitung Luas Segitiga
	alas := 15.0
	tinggi := 10.0
	luas := 0.5 * alas * tinggi

	return float32(luas)
}

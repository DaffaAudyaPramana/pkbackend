package pkbackend

import (
	"fmt"
	"testing"
)

// func TestInsertUser(t *testing.T) {
// 	nm := "Daffa Audya Pramana"
// 	gnr := "1214015"
// 	eml := "Daffaaudyapramana03@gmail.com"
// 	nhp := "085722374028"
// 	hasil := InsertUser(nm, gnr, eml, nhp)
// 	fmt.Println(hasil)

// }

// func TestPembayaran(t *testing.T) {
// 	sts := "Dibayar"
// 	tbr := "200,000"
// 	hasil := InsertPembayaran(sts, tbr)
// 	fmt.Println(hasil)
// }

// func TestPendaftaran(t *testing.T) {
// 	nis := "92348927348"
// 	nik := "72619873972094293"
// 	nsa := "Daffa Audya Pramana"
// 	hasil := InsertPendaftaran(nis, nik, nsa)
// 	fmt.Println(hasil)
// }

// func TestInsertPengumuman(t *testing.T) {
// 	hsi := "LULUS"
// 	nli := "97"
// 	pgm := "Fighting"
// 	hasil := InsertPengumuman(hsi, nli, pgm)
// 	fmt.Println(hasil)
// }

// func TestInsertProgramKursus(t *testing.T) {
// 	nks := "FIGHTING GAME"
// 	jks := "Program 3"
// 	pgr := "Budi Setiawan"
// 	hasil := InsertProgramKursus(nks, jks, pgr)
// 	fmt.Println(hasil)
// }

func TestGetDataUser(t *testing.T) {
	stats := "Daffa Audya Pramana"
	data := GetDataUser(stats)
	fmt.Println(data)
}

// func TestGetDataAbout(t *testing.T) {
// 	stats := "Apakah kurikulum ini menyulitkan hidup?"
// 	data := GetDataAbout(stats)
// 	fmt.Println(data)
// }

// func TestGetDataContacus(t *testing.T) {
// 	stats := "0822126722"
// 	data := GetDataContacus(stats)
// 	fmt.Println(data)
// }

// func TestGetDataDashboard(t *testing.T) {
// 	stats := "Denmark"
// 	data := GetDataDashboard(stats)
// 	fmt.Println(data)
// }
// func TestGetDataMahasiswa(t *testing.T) {
// 	stats := "johndoe@flex.co"
// 	data := GetDataMahasiswa(stats)
// 	fmt.Println(data)
// }

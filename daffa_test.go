package pkbackend

import (
	"fmt"
	"testing"
)

// func TestInsertUser(t *testing.T) {
// 	dbname := "proyek-2"
// 	user := User{
// 		ID:     primitive.NewObjectID(),
// 		Nama:   "Arya",
// 		Gender: "Perempuan",
// 		Email:  "Arya@gmail.com",
// 		No_hp:  "085722374027"
// 	}
// 	insertedID := InsertUser(dbname, user)
// 	if insertedID == nil {
// 		t.Error("Failed to insert user")
// 	}
// }

// func TestInsertPendaftaran(t *testing.T) {
// 	dbname := "proyek-2"
// 	pendaftaran := Pendaftaran{
// 		ID:         primitive.NewObjectID(),
// 		Nama_siswa: "Arya",
// 		Nis:        "92348927348",
// 		Nik:        "72619873972094293",
// 	}
// 	insertedID := InsertPendaftaran(dbname, pendaftaran)
// 	if insertedID == nil {
// 		t.Error("Failed to insert pendaftaran")
// 	}
// }

// func TestInsertPembayaran(t *testing.T) {
// 	dbname := "proyek-2"
// 	pembayaran := Pembayaran{
// 		ID:          primitive.NewObjectID(),
// 		Status:      "92348927348",
// 		Total_bayar: "72619873972094293",
// 	}
// 	insertedID := InsertPembayaran(dbname, pembayaran)
// 	if insertedID == nil {
// 		t.Error("Failed to insert pembayaran")
// 	}
// }
// func TestInsertPengumuman(t *testing.T) {
// 	dbname := "proyek-2"
// 	pengumuman := Pengumuman{
// 		ID:            primitive.NewObjectID(),
// 		Hasil_seleksi: "LULUS",
// 		Nilai:         "97",
// 		Program:       "Fighting",
// 	}
// 	insertedID := InsertPengumuman(dbname, pengumuman)
// 	if insertedID == nil {
// 		t.Error("Failed to insert pengumuman")
// 	}
// }
// func TestInsertKursus(t *testing.T) {
// 	dbname := "proyek-2"
// 	kursus := Kursus{
// 		ID:             primitive.NewObjectID(),
// 		Nama_kursus:    "FIGHTING GAME",
// 		Jenjang_kursus: "Program 3",
// 		Pengajar:       "Jon Snow",
// 	}
// 	insertedID := InsertKursus(dbname, kursus)
// 	if insertedID == nil {
// 		t.Error("Failed to insert kursus")
// 	}
// }

func TestGetDataUser(t *testing.T) {
	stats := "Arya"
	data := GetDataUser(stats)
	fmt.Println(data)
}

func TestGetDataPendaftaran(t *testing.T) {
	stats := "92348927348"
	data := GetDataPendaftaran(stats)
	fmt.Println(data)
}

func TestGetDataPembayaran(t *testing.T) {
	stats := "Dibayar"
	data := GetDataPembayaran(stats)
	fmt.Println(data)
}

func TestGetDataPengumuman(t *testing.T) {
	stats := "Fighting"
	data := GetDataPengumuman(stats)
	fmt.Println(data)
}
func TestGetDataKursus(t *testing.T) {
	stats := "Jon Snow"
	data := GetDataKursus(stats)
	fmt.Println(data)
}

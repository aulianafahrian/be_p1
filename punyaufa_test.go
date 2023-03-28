package namapackage

import (
	"fmt"
	"testing"

	model "github.com/aulianafahrian/be_p1/model"
	module "github.com/aulianafahrian/be_p1/module"
)

func TestInsertPresensi(t *testing.T) {
	long := 98.345345
	lat := 123.561651
	lokasi := "Rumah"
	phonenumber := "685754189853"
	checkin := "masuk"
	biodata := model.Karyawan{
		Nama:         "Aufa",
		Phone_number: "628456456211",
		Jabatan:      "Bukan Rakyat Biasa",
	}
	hasil := module.InsertPresensi(long, lat, lokasi, phonenumber, checkin, biodata)
	fmt.Println(hasil)
}

func TestGetKaryawanFromPhoneNumber(t *testing.T) {
	phonenumber := "685754189853"
	biodata := module.GetKaryawanFromPhoneNumber(phonenumber)
	fmt.Println(biodata)
}

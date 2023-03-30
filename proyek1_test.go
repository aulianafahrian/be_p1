package proyek123

import (
	"fmt"
	"testing"

	model "github.com/aulianafahrian/be_p1/model"
	module "github.com/aulianafahrian/be_p1/module"
)

func TestInsertProyek1(t *testing.T) {
	biodata_mahasiswa := model.Mahasiswa{
		NPM:   "1214049",
		Nama:  "Auliana Fahrian",
		Kelas: "2B",
		Jurusan: model.Jurusan{
			Nama: "Teknik Informatika",
		},
		Prodi: model.Prodi{
			Nama: "D4 Teknik Informatika",
		},
	}
	dosen_pembimbing := model.Dosen{
		Nama: "Dimas",
		Prodi: model.Prodi{
			Nama: "D4 Teknik Informatika",
		},
	}

	dosen_penguji := model.Dosen{
		Nama: "Dani",
		Prodi: model.Prodi{
			Nama: "D4 Teknik Informatika",
		},
	}
	Judul := "Perangkat Keras"
	tanggal_sidang := "2023-12-12"
	hasil := module.InsertProyek1(module.MongoConn, "proyek1", biodata_mahasiswa, dosen_pembimbing, dosen_penguji, Judul, tanggal_sidang)
	fmt.Println(hasil)
}

func TestInsertDataMahasiswa(t *testing.T) {
	npm := "1214049"
	nama := "Auliana Fahrian"
	kelas := "2B"
	jurusan := model.Jurusan{
		Nama: "Teknik Informatika",
	}
	prodi := model.Prodi{
		Nama: "D4 Teknik Informatika",
	}
	hasil := module.InsertDataMahasiswa(module.MongoConn, "mahasiswa", npm, nama, kelas, jurusan, prodi)
	fmt.Println(hasil)
}

func TestInsertDataDosen(t *testing.T) {
	nid := "123456"
	nama := "Dani Ferdinan"
	prodi := model.Prodi{
		Nama: "D4 Teknik Informatika",
	}
	hasil := module.InsertDataDosen(module.MongoConn, "dosen", nid, nama, prodi)
	fmt.Println(hasil)
}

func TestInsertDataProdi(t *testing.T) {
	kode_prodi := "11111"
	nama := "D4 Teknik Informatika"
	hasil := module.InsertDataProdi(module.MongoConn, "prodi", kode_prodi, nama)
	fmt.Println(hasil)
}

func TestInsertDataJurusan(t *testing.T) {
	kode_jurusan := "22222"
	nama := "Teknik Informatika"
	hasil := module.InsertDataJurusan(module.MongoConn, "jurusan", kode_jurusan, nama)
	fmt.Println(hasil)
}

func TestGetProyek1FromNPM(t *testing.T) {
	npm := "1214049"
	hasil := module.GetProyek1FromNPM(npm, module.MongoConn, "proyek1")
	fmt.Println(hasil)
}
func TestGetDataMahasiswaFromNPM(t *testing.T) {
	npm := "1214049"
	mahasiswa := module.GetDataMahasiswaFromNPM(npm, module.MongoConn, "mahasiswa")
	fmt.Println(mahasiswa)
}

func TestGetDataDosenFromNID(t *testing.T) {
	nid := "123456"
	dosen := module.GetDataDosenFromNID(nid, module.MongoConn, "dosen")
	fmt.Println(dosen)
}

func TestGetDataProdiFromKodeProdi(t *testing.T) {
	kode_prodi := "11111"
	prodi := module.GetDataProdiFromKodeProdi(kode_prodi, module.MongoConn, "prodi")
	fmt.Println(prodi)
}

func TestGetDatajurusanFromKodeJurusan(t *testing.T) {
	kode_jurusan := "22222"
	jurusan := module.GetDataJurusanFromKodeJurusan(kode_jurusan, module.MongoConn, "jurusan")
	fmt.Println(jurusan)
}

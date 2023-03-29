package namapackage

import (
	"fmt"
	"testing"

	model "github.com/aulianafahrian/be_p1/model"
	module "github.com/aulianafahrian/be_p1/module"
)

func TestInsertProyek1(t *testing.T) {
	biodata_mahasiswa := model.Mahasiswa{
		NPM:   1214049,
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
	hasil := module.InsertProyek1(biodata_mahasiswa, dosen_pembimbing, dosen_penguji, Judul, tanggal_sidang)
	fmt.Println(hasil)
}

func TestGetProyek1FromNPM(t *testing.T) {
	npm := 1214049
	hasil := module.GetProyek1FromNPM(npm)
	fmt.Println(hasil)
}

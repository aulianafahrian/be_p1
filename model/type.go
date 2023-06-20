package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Proyek1 struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Biodata_mahasiswa Mahasiswa          `bson:"biodata_mahasiswa,omitempty" json:"biodata_mahasiswa,omitempty"`
	Dosen_pembimbing  Dosen              `bson:"dosen_pembimbing,omitempty" json:"dosen_pembimbing,omitempty"`
	Dosen_penguji     Dosen              `bson:"dosen_penguji,omitempty" json:"dosen_penguji,omitempty"`
	Judul             string             `bson:"judul,omitempty" json:"judul,omitempty"`
	Tanggal_sidang    string             `bson:"tanggal_sidang,omitempty" json:"tanggal_sidang,omitempty"`
}

type Proyek struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Tipe_proyek       string             `bson:"tipe_proyek,omitempty" json:"tipe_proyek,omitempty"`
	Biodata_mahasiswa Mahasiswa          `bson:"biodata_mahasiswa,omitempty" json:"biodata_mahasiswa,omitempty"`
	Partner_mahasiswa Mahasiswa          `bson:"partner_mahasiswa,omitempty" json:"partner_mahasiswa,omitempty"`
	Dosen_pembimbing  Dosen              `bson:"dosen_pembimbing,omitempty" json:"dosen_pembimbing,omitempty"`
	Dosen_penguji     Dosen              `bson:"dosen_penguji,omitempty" json:"dosen_penguji,omitempty"`
	Judul             string             `bson:"judul,omitempty" json:"judul,omitempty"`
	Tanggal_sidang    string             `bson:"tanggal_sidang,omitempty" json:"tanggal_sidang,omitempty"`
	Ruang_sidang      string             `bson:"ruang_sidang,omitempty" json:"ruang_sidang,omitempty"`
}

type Mahasiswa struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NPM     string             `bson:"npm,omitempty" json:"npm,omitempty"`
	Nama    string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Kelas   string             `bson:"kelas,omitempty" json:"kelas,omitempty"`
	Jurusan Jurusan            `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
	Prodi   Prodi              `bson:"prodi,omitempty" json:"prodi,omitempty"`
}

type Dosen struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NID   string             `bson:"nid,omitempty" json:"nid,omitempty"`
	Nama  string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Prodi Prodi              `bson:"prodi,omitempty" json:"prodi,omitempty"`
}

type Prodi struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode_Prodi string             `bson:"kode_prodi,omitempty" json:"kode_prodi,omitempty"`
	Nama       string             `bson:"nama,omitempty" json:"nama,omitempty"`
}

type Jurusan struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode_jurusan string             `bson:"kode_jurusan,omitempty" json:"kode_jurusan,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
}

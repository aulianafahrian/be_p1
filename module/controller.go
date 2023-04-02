package module

import (
	"context"
	"fmt"
	"os"

	"github.com/aiteung/atdb"
	model "github.com/aulianafahrian/be_p1/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "tes_db",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertProyek1(db *mongo.Database, col string, mahasiswa model.Mahasiswa, dosen_pembimbing model.Dosen, dosen_penguji model.Dosen, judul string, tanggal_sidang string) (InsertedID interface{}) {
	var proyek model.Proyek1
	proyek.Biodata_mahasiswa = mahasiswa
	proyek.Dosen_pembimbing = dosen_pembimbing
	proyek.Dosen_penguji = dosen_penguji
	proyek.Judul = judul
	proyek.Tanggal_sidang = tanggal_sidang
	return InsertOneDoc(db, col, proyek)
}

func InsertDataMahasiswa(db *mongo.Database, col string, npm string, nama string, kelas string, jurusan model.Jurusan, prodi model.Prodi) (InsertedID interface{}) {
	var mahasiswa model.Mahasiswa
	mahasiswa.NPM = npm
	mahasiswa.Nama = nama
	mahasiswa.Kelas = kelas
	mahasiswa.Jurusan = jurusan
	mahasiswa.Prodi = prodi
	return InsertOneDoc(db, col, mahasiswa)
}

func InsertDataDosen(db *mongo.Database, col string, nid string, nama string, prodi model.Prodi) (InsertedID interface{}) {
	var dosen model.Dosen
	dosen.NID = nid
	dosen.Nama = nama
	dosen.Prodi = prodi
	return InsertOneDoc(db, col, dosen)
}

func InsertDataProdi(db *mongo.Database, col string, kode_prodi string, nama string) (InsertedID interface{}) {
	var prodi model.Prodi
	prodi.Kode_Prodi = kode_prodi
	prodi.Nama = nama
	return InsertOneDoc(db, col, prodi)
}

func InsertDataJurusan(db *mongo.Database, col string, kode_jurusan string, nama string) (InsertedID interface{}) {
	var jurusan model.Jurusan
	jurusan.Kode_jurusan = kode_jurusan
	jurusan.Nama = nama
	return InsertOneDoc(db, col, jurusan)
}

func GetProyek1FromNPM(npm string, db *mongo.Database, col string) (data model.Proyek1) {
	proyek1 := db.Collection(col)
	filter := bson.M{"biodata_mahasiswa.npm": npm}
	fmt.Print(npm)
	err := proyek1.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("GetProyek1FromNPM: %v\n", err)
	}
	return data
}
func GetDataMahasiswaFromNPM(npm string, db *mongo.Database, col string) (data model.Mahasiswa) {
	mahasiswa := db.Collection(col)
	filter := bson.M{"npm": npm}
	fmt.Print(npm)
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("GetDataMahasiswaFromNPM: %v\n", err)
	}
	return data
}

func GetDataDosenFromNID(nid string, db *mongo.Database, col string) (data model.Dosen) {
	dosen := db.Collection(col)
	filter := bson.M{"nid": nid}
	fmt.Print(nid)
	err := dosen.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("GetDataDosenMahasiswaFromNPM: %v\n", err)
	}
	return data
}

func GetDataProdiFromKodeProdi(kode_prodi string, db *mongo.Database, col string) (data model.Prodi) {
	prodi := db.Collection(col)
	filter := bson.M{"kode_prodi": kode_prodi}
	fmt.Print(kode_prodi)
	err := prodi.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("GetDataProdiFromKodeProdi: %v\n", err)
	}
	return data
}

func GetDataJurusanFromKodeJurusan(kode_jurusan string, db *mongo.Database, col string) (data model.Jurusan) {
	jurusan := db.Collection(col)
	filter := bson.M{"kode_jurusan": kode_jurusan}
	fmt.Print(kode_jurusan)
	err := jurusan.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("GetDatajurusanFromKodeJurusan: %v\n", err)
	}
	return data
}

func GetAllProyek1(db *mongo.Database, col string) (proyek1 []model.Proyek1) {
	data_proyek1 := db.Collection(col)
	filter := bson.M{}
	cursor, err := data_proyek1.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &proyek1)
	if err != nil {
		fmt.Println(err)
	}
	return proyek1
}

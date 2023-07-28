package main

import (
	"fmt"

	_ "github.com/lib/pq"
	config "github.com/septian03yogi/enigmalaundryinc/config/database"
	"github.com/septian03yogi/enigmalaundryinc/model"
	"github.com/septian03yogi/enigmalaundryinc/repository"
	"github.com/septian03yogi/enigmalaundryinc/usecase"
)

// func connectDb() *sql.DB {
// 	//Koneksi
// 	// 1. URL -> driver://user:password@localhost:5432/database?sslmode=disable
// 	// 2. Default -> host=localhost port=5432 user=postgres password=password dbname=enigmalaundryinc sslmode=disable

// 	host := "localhost"
// 	port := 5432
// 	user := "postgres"
// 	password := "@03Yehezkiel"
// 	database := "enigmalaundryinc"
// 	driver := "postgres"
// 	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
// 	db, err := sql.Open(driver, dsn)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return db
// }

// func createUom(db *sql.DB, uom Uom) error {
// 	var err error
// 	_, err = db.Exec("INSERT INTO uom(id, name) VALUES ($1, $2)", uom.Id, uom.Name)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("UOM created succesfully")
// 	return nil
// }

// func mainMenuForm() {
// 	fmt.Println(`
// 	|++++ Enigma Lundry Menu ++++|
// 	| 1. Master UOM               |
// 	| 2. Master Product           |
// 	| 3. Master Customer          |
// 	| 4. Master Employee          |
// 	| 5. Transaksi               |
// 	| 6. Keluar                  |
// 	`)
// 	fmt.Print(" Pilih Menu (1-6): ")
// }

// // UOM Menu Form
// func uomMenuForm() {
// 	fmt.Println(`
// 	|++++ Enigma Lundry Menu ++++|
// 	| 1. Tambah Data             |
// 	| 2. Lihat Data              |
// 	| 3. Update Data             |
// 	| 4. Hapus Data              |
// 	| 5. Kembali Ke Menu Utama   |

// 	`)
// 	fmt.Print(" Pilih Menu (1-4): ")
// 	db := connectDb()
// 	defer db.Close()
// 	for {
// 		var selectedMenu string
// 		fmt.Scanln(&selectedMenu)
// 		switch selectedMenu {
// 		case "1":
// 			uom := uomCreateForm()
// 			err := createUom(db, uom)
// 			checkErr(err)
// 			return
// 		case "2":
// 			fmt.Println("Lihat data")
// 		case "3":
// 			fmt.Println("Update data")
// 		case "4":
// 			fmt.Println("Hapus data")
// 		case "5":
// 			return
// 		default:
// 			fmt.Println("Menu Tidak ditemukan")
// 		}
// 	}
// }

// func uomCreateForm() Uom {
// 	var (
// 		uomId, uomName, saveConfirmation string
// 	)
// 	fmt.Print("UOM ID: ")
// 	fmt.Scanln(&uomId)

// 	fmt.Print("UOM Name: ")
// 	fmt.Scanln(&uomName)

// 	fmt.Printf("UOM Id: %s, Name: %s akan disimpan (y/t)", uomId, uomName)
// 	fmt.Scanln(&saveConfirmation)

// 	if saveConfirmation == "y" {
// 		uom := Uom{
// 			Id:   uomId,
// 			Name: uomName,
// 		}
// 		return uom
// 	}
// 	return Uom{}
// }
// func runConsole() {
// 	for {
// 		mainMenuForm()
// 		var selectedMenu string
// 		fmt.Scanln(&selectedMenu)
// 		switch selectedMenu {
// 		case "1":
// 			uomMenuForm()
// 			var selectedMenu string
// 			fmt.Scanln(&selectedMenu)
// 			switch selectedMenu {
// 			case "1":
// 				uomMenuForm()
// 			case "6":
// 				os.Exit(0)
// 			default:
// 				fmt.Println("Menu tidak ditemukan")
// 			}
// 		}
// 	}
// }

// func checkErr(err error) {
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

func main() {

	// db := connectDb()
	// var (
	// 	uomId, uomName string
	// )

	// fmt.Print("UOM ID:")
	// fmt.Scanln(&uomId)

	// fmt.Print("UOM NAME:")
	// fmt.Scanln(&uomName)
	// uom := Uom{
	// 	Id:   uomId,
	// 	Name: uomName,
	// }

	// err := createUom(db, uom)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// runConsole()

	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	dbConn, _ := config.NewDbConnection(cfg)
	db := dbConn.Conn()
	uomRepo := repository.NewUomRepository(db)
	uomUseCase := usecase.NewUomUseCase(uomRepo)

	//repository
	uom := model.Uom{
		Id:   "7",
		Name: "Kotak",
	}

	err = uomUseCase.RegisterNewUom(uom)
	if err != nil {
		fmt.Print(err)
	}

	// db.Exec("INSERT INTO uom (id, name) VALUES ($1, $2)", uom.Id, uom.Name)

}

package delivery

import (
	"fmt"
	"os"

	"github.com/septian03yogi/enigmalaundryinc/config"
	"github.com/septian03yogi/enigmalaundryinc/delivery/controller"
	"github.com/septian03yogi/enigmalaundryinc/repository"
	"github.com/septian03yogi/enigmalaundryinc/usecase"
)

type Console struct {
	//semua usecase disimpan disini
	uomUC     usecase.UomUseCase
	productUC usecase.ProductUseCase
}

// func (c *Console) uomCreateForm() model.Uom {
// 	var (
// 		uomId, uomName, saveConfirmation string
// 	)
// 	fmt.Print("UOM Name: ")
// 	fmt.Scanln(&uomName)

// 	fmt.Printf("UOM Id: %s, Name: %s akan disimpan (y/t)", uomId, uomName)
// 	fmt.Scanln(&saveConfirmation)

// 	if saveConfirmation == "y" {
// 		uom := model.Uom{
// 			Id:   uuid.New().String(),
// 			Name: uomName,
// 		}
// 		return uom
// 	}
// 	return model.Uom{}
// }

func (c *Console) mainMenuForm() {
	fmt.Println(`
	|++++ Enigma Laundry Menu++++|
	| 1. Master UOM              |
	| 2. Master Product          |
	| 3. Master Customer         |
	| 4. Master Employee         |
	| 5. Transaksi               |
	| 6. Keluar 				 |

	`)
	fmt.Print("Pilih menu 1-6: ")
}

func (c *Console) Run() {
	for {
		c.mainMenuForm()
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			controller.NewUomController(c.uomUC).UomMenuForm()
		case "2":
			controller.NewProductController(c.productUC).HandlerMainForm()
		case "6":
			os.Exit(0)
		default:
			fmt.Println("Menu tidak ditemukan")
		}
	}
}

func NewConsole() *Console {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	dbConn, _ := config.NewDbConnection(cfg)
	db := dbConn.Conn()
	uomRepo := repository.NewUomRepository(db)
	productRepo := repository.NewProductRepository(db)
	uomUseCase := usecase.NewUomUseCase(uomRepo)
	productUseCase := usecase.NewProductUseCase(productRepo, uomUseCase)
	return &Console{
		uomUC:     uomUseCase,
		productUC: productUseCase,
	}

	// //repository
	// uom := model.Uom{
	// 	Id:   "7",
	// 	Name: "Kotak",
	// }

	// err = uomUseCase.RegisterNewUom(uom)
	// if err != nil {
	// 	fmt.Print(err)
	// }
}

package controller

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/septian03yogi/enigmalaundryinc/model"
	"github.com/septian03yogi/enigmalaundryinc/model/dto"
	"github.com/septian03yogi/enigmalaundryinc/usecase"
	"github.com/septian03yogi/enigmalaundryinc/utils/exceptions"
)

type ProductController struct {
	productUC usecase.ProductUseCase
}

func (p *ProductController) HandlerMainForm() {
	fmt.Println(`
	|++++ Master PRODUCT ++++|
	| 1. Tambah Data             |
	| 2. Lihat Data              |
	| 3. Detail Data             |
	| 4. Update Data             |
	| 5. Hapus Data              |
	| 6. Kembali Ke Menu Utama   |

	`)
	fmt.Print(" Pilih Menu (1-6): ")
	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			uom := p.createHandleForm()
			err := p.productUC.RegisterNewProduct(uom)
			exceptions.CheckErr(err)
			return
		case "2":
			requestPaging := dto.PaginationParam{
				Page: 1,
			}
			products, paging, err := p.productUC.FindAllProduct(requestPaging)
			exceptions.CheckErr(err)
			p.findAllHandlerForm(products, paging)
			return
		// case "3":
		// 	u.uomGetForm()
		// 	return
		// case "4":
		// 	uom := u.uomUpdateForm()
		// 	err := u.uomUC.UpdateUom(uom)
		// 	exceptions.CheckErr(err)
		// 	return
		// case "5":
		// 	fmt.Println("Hapus data")
		// 	id := u.uomDeleteForm()
		// 	err := u.uomUC.DeleteUom(id)
		// 	exceptions.CheckErr(err)
		// 	return
		case "6":
			return
		default:
			fmt.Println("Menu Tidak ditemukan")
		}
	}
}

func (p *ProductController) createHandleForm() model.Product {
	var (
		id, name, price, uomId, saveConfirmation string
	)
	fmt.Print("PRODUCT Name: ")
	fmt.Scanln(&name)

	fmt.Print("PRODUCT Price: ")
	fmt.Scanln(&price)

	fmt.Print("PRODUCT UOM ID: ")
	fmt.Scanln(&uomId)
	id = uuid.New().String()
	priceConv, _ := strconv.Atoi(price)
	fmt.Printf("Product Id: %s, Name: %s, Price:%d , UomId: %s akan disimpan (y/t)", id, name, priceConv, uomId)
	fmt.Scanln(&saveConfirmation)

	if saveConfirmation == "y" {
		product := model.Product{
			Id:    id,
			Name:  name,
			Price: priceConv,
			Uom:   model.Uom{Id: uomId},
		}
		return product
	}
	return model.Product{}
}

func (p *ProductController) findAllHandlerForm(products []model.Product, paging dto.Paging) {
	for _, product := range products {
		fmt.Println("Product list")
		fmt.Printf("ID: %s \n", product.Id)
		fmt.Printf("Name: %s \n", product.Name)
		fmt.Printf("Price: %d \n", product.Price)
		fmt.Printf("UOM Id: %s \n", product.Uom.Id)
		fmt.Printf("UOM Name: %s \n", product.Uom.Name)
		fmt.Println()
		fmt.Printf("Paging: ")
		fmt.Printf("Page:%d \n", paging.Page)
		fmt.Printf("RowsPerPage: %d \n", paging.RowsPerPage)
		fmt.Printf("TotalPages: %d \n", paging.TotalPages)
		fmt.Printf("TotalRows: %d \n", paging.TotalRows)
	}
}

// func (p *ProductController) uomUpdateForm() model.Uom {
// 	var (
// 		uomId, uomName, saveConfirmation string
// 	)
// 	fmt.Print("UOM ID: ")
// 	fmt.Scanln(&uomId)
// 	fmt.Println("UOM Name: ")
// 	fmt.Scanln(&uomName)
// 	fmt.Printf("UOM Id: %s, Name: %s akan disimpan? (y/t)", uomId, uomName)
// 	fmt.Scanln(&saveConfirmation)
// 	if saveConfirmation == "y" {
// 		uom := model.Uom{
// 			Id:   uomId,
// 			Name: uomName,
// 		}
// 		return uom
// 	}
// 	return model.Uom{}
// }

// func (p *ProductController) uomDeleteForm() string {
// 	var id string
// 	fmt.Println("UOM ID: ")
// 	fmt.Scanln(&id)
// 	return id
// }

// func (p *ProductController) uomGetForm() {
// 	var id string
// 	fmt.Print("UOM ID: ")
// 	fmt.Scanln(&id)
// 	uom, err := u.uomUC.FindByIdUom(id)
// 	exceptions.CheckErr(err)
// 	fmt.Printf("UOM ID %s \n", id)
// 	fmt.Println(strings.Repeat("=", 15))
// 	fmt.Printf("UOM ID: %s \n", uom.Id)
// 	fmt.Printf("UOM Name: %s \n", uom.Name)
// }

func NewProductController(usecase usecase.ProductUseCase) *ProductController {
	return &ProductController{productUC: usecase}
}

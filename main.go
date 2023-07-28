package main

import (
	"github.com/septian03yogi/enigmalaundryinc/delivery"
)

// func checkErr(err error) {
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

func main() {
	delivery.NewConsole().Run()

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

	// db.Exec("INSERT INTO uom (id, name) VALUES ($1, $2)", uom.Id, uom.Name)

}

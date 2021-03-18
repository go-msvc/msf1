package main

import (
	"fmt"

	"github.com/go-msvc/msf/config"
	"github.com/go-msvc/msf/crud"
	"github.com/go-msvc/msf/db"
	_ "github.com/go-msvc/msf/db/mysql"
	log "github.com/go-msvc/msf/logger"
	"github.com/go-msvc/msf/model"
	"github.com/go-msvc/msf/service"
	stockmodel "github.com/go-msvc/msf1/model"
	"github.com/go-msvc/msf1/service/stock/api"
)

var (
	stockDb db.IDatabase
)

func main() {
	log.Infof("Stock Service Main")
	config.MustLoadFile("./conf/stock.json")

	//start a new model
	dbModel := model.New()

	//open the database
	stockDb = db.MustOpen("stock")

	locationModel := dbModel.MustAdd(stockmodel.Location{})
	locationTable := db.MustAddTable(stockDb, locationModel)

	companyModel := dbModel.MustAdd(stockmodel.Company{})
	companyTable := db.MustAddTable(stockDb, companyModel)

	ownerModel := dbModel.MustAdd(stockmodel.Owner{})
	ownerTable := db.MustAddTable(stockDb, ownerModel)

	inventoryModel := dbModel.MustAdd(stockmodel.Inventory{})
	inventoryTable := db.MustAddTable(stockDb, inventoryModel)

	stockStatusModel := dbModel.MustAdd(stockmodel.StockStatus{})
	stockStatusTable := db.MustAddTable(stockDb, stockStatusModel)

	stockModel := dbModel.MustAdd(stockmodel.Stock{})
	stockTable := db.MustAddTable(stockDb, stockModel)

	stockLevelModel := dbModel.MustAdd(stockmodel.StockLevel{})
	stockLevelTable := db.MustAddTable(stockDb, stockLevelModel)

	service.NewService("stock").
		HandleMux("location", crud.New(locationTable)).
		HandleMux("company", crud.New(companyTable)).
		HandleMux("owner", crud.New(ownerTable)).
		HandleMux("inventory", crud.New(inventoryTable)).
		HandleMux("stock_status", crud.New(stockStatusTable)).
		HandleMux("stock", crud.New(stockTable)).
		HandleMux("stock_level", crud.New(stockLevelTable)).
		Handle("levels", getStockLevels).
		Run()
}

// func getLocation(ctx service.IContext, req api.GetLocationRequest) (res api.GetLocationResponse, err error) {
// 	location, err := locationModel.GetById(req.LocationId)
// 	if err != nil {
// 		err = fmt.Errorf("failed to get location.id=%d: %v", req.LocationId, err)
// 		return
// 	}
// 	return api.GetLocationResponse{
// 		Name: location.(stockModel.Location).Name,
// 	}, nil
// }

func getStockLevels(ctx service.IContext, req api.LevelsRequest) (res api.LevelsResponse, err error) {
	// stock, err := stockDb.Get(
	// 	ctx,
	// 	//model.Stock{},
	// 	db.Key{
	// 		"owner":    req.Owner,
	// 		"product":  req.ProductReference,
	// 		"location": req.Location,
	// 	},
	// )
	// if err != nil {
	// 	return api.LevelsResponse{}, fmt.Errorf("failed to read stock: %v", err)
	// }
	// ctx.Debugf("got stock: %+v", stock)
	return api.LevelsResponse{}, fmt.Errorf("NYI")
}

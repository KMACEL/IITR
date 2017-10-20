package operations

import "github.com/KMACEL/IITR/databasecenter"

type OtaBase struct{}

func (o OtaBase) Start() {
	var dataBase databasecenter.DB
	db := dataBase.Open("iTaksi")
	dataBase.CreateTable(db, "otaOperations",
		"imei INT PRIMARY KEY ,"+
			"pushApplication INT,"+
			"pushTime datetime,"+
			"startApplication INT,"+
			"startTime datetime,"+
			"controlOta int,"+
			"controlTime datetime")
	dataBase.InsertInto(db,
		"otaOPerations",
		"imei,pushApplication",
		"5555555", "aaaaaaaaaaa")

		dataBase.Close(db)
}

func readDataBase() {}

func insertDataBase() {}

func updateDataBase() {}

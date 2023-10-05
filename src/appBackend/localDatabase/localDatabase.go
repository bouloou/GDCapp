package localDatabase

import (
	"fmt"
	"io/fs"
	"os"
)

type Database struct {
	MainDirPath      string
	ClientsPath      string
	InternalDataPath string
}

func FirstInit(db Database) {
	err := os.Mkdir(db.MainDirPath, fs.ModeDir)
	if err != nil {
		fmt.Print(fmt.Errorf("%v", err))
	}

	// Init of client folder
	err = os.Mkdir(db.ClientsPath, fs.ModeDir)
	if err != nil {
		fmt.Print(fmt.Errorf("%v", err))
	}

	file, err := os.Create(db.ClientsPath + "/" + "clientsList.csv")
	if err != nil {
		fmt.Print(fmt.Errorf("%v", err))
	}
	file.Close()

	// Init of InternalData folder
	err = os.Mkdir(db.InternalDataPath, fs.ModeDir)
	if err != nil {
		fmt.Print(fmt.Errorf("%v", err))
	}
}

func GetClientFromFileNumber(db Database, fileNumber string) Client {

	var cli Client
	/*
		pathClientFolder := db.DirPath + "/clients/" + strconv.Itoa(cli.NumeroDossier)
		err := os.Mkdir(pathClientFolder, fs.ModeDir)
		if err != nil {
			panic(err)
		}

		file, err := os.Create(pathClientFolder + "/" + "infos.csv")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		reader := csv.NewReader(file)
		reader.FieldsPerRecord = -1 // Allow variable number of fields
		data, err := reader.ReadAll()
		if err != nil {
			panic(err)
		}

		// Print the CSV data
		for _, row := range data {
			for _, col := range row {
				fmt.Printf("%s,", col)
			}
			fmt.Println()
		}
	*/
	return cli
}

func AddClient(db Database, cli Client) error {

	/*
		pathClientFolder := db.DirPath + "/clients/" + cli.NumeroDossier

		file, err := os.Create(pathClientFolder + "/" + "infos.csv")
		if err != nil {
			fmt.Print(fmt.Errorf("%v", err))
		}
		defer file.Close()
	*/
	return nil

}

func DeleteClient(db Database, fileNumber string) error {
	/*
		pathClientFolder := db.DirPath + "/clients"
		err := os.Remove(pathClientFolder + "/infos.csv")
		if err != nil {
			fmt.Print(fmt.Errorf("%v", err))
		}

		err = os.Remove(pathClientFolder)
		if err != nil {
			fmt.Print(fmt.Errorf("%v", err))
		}
	*/
	return nil
}

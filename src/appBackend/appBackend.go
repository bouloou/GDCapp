package appBackend

import (
	"errors"
	"fmt"
	"io"
	"localDatabase"
	"net/http"
	"os"
	"strings"
)

var currentDatabase localDatabase.Database

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getClient(w http.ResponseWriter, r *http.Request) {
	splitedString := strings.Split(r.URL.Path, "/")
	io.WriteString(w, fmt.Sprintf("Get: %s\n", splitedString[3]))

	if localDatabase.IsValidFileNumberFormat(splitedString[3]) {

	} else {
		io.WriteString(w, fmt.Sprintf("Numero de dossier invalide\n"))
	}
}

func addClient(w http.ResponseWriter, r *http.Request) {

	splitedString := strings.Split(r.URL.Path, "/")
	io.WriteString(w, fmt.Sprintf("Add: %s\n", splitedString[3]))

	if localDatabase.IsValidFileNumberFormat(splitedString[3]) {

		var Marie localDatabase.Client
		Marie.Nom = "Marie"
		Marie.Concepteur = "Clementine"
		Marie.NumeroDossier = splitedString[3]
		localDatabase.AddClient(currentDatabase, Marie)

	} else {
		io.WriteString(w, fmt.Sprintf("Numero de dossier invalide\n"))
	}
}

func deleteClient(w http.ResponseWriter, r *http.Request) {

	splitedString := strings.Split(r.URL.Path, "/")
	io.WriteString(w, fmt.Sprintf("Delete: %s\n", splitedString[3]))

	if localDatabase.IsValidFileNumberFormat(splitedString[3]) {
		localDatabase.DeleteClient(currentDatabase, splitedString[3])
	} else {
		io.WriteString(w, fmt.Sprintf("Numero de dossier invalide\n"))
	}
}

func tests() {
	fmt.Printf("resultat %d : %v\n", 0, localDatabase.IsValidFileNumberFormat("0"))
	fmt.Printf("resultat %d : %v\n", 100000000, localDatabase.IsValidFileNumberFormat("100000000"))
	fmt.Printf("resultat %d : %v\n", 1, localDatabase.IsValidFileNumberFormat("1"))
	fmt.Printf("resultat %d : %v\n", 51515151, localDatabase.IsValidFileNumberFormat("51515151"))
	fmt.Printf("resultat %d : %v\n", 99999999, localDatabase.IsValidFileNumberFormat("99999999"))
}

func app() {
	// Add new data
	/*
		var Marie localDatabase.Client
		Marie.Nom = "Marie"
		Marie.Concepteur = "Clementine"
		Marie.NumeroDossier = localDatabase.ConvertIntToFileNumber(110116)
		localDatabase.AddNewClient(currentDatabase, Marie)
	*/

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/client/get/", getClient)
	http.HandleFunc("/client/add/", addClient)
	http.HandleFunc("/client/delete/", deleteClient)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
func StartBackend() {

	//Init database infos
	currentPath, _ := os.Getwd()
	currentDatabase.MainDirPath = currentPath + "/data"
	currentDatabase.ClientsPath = currentPath + "/data/clients"
	currentDatabase.InternalDataPath = currentPath + "/data/internaldata"

	//tests()
	localDatabase.FirstInit(currentDatabase)
	app()
}

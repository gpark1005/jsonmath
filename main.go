package main

import (
	"jsonmath/handlers"
	"log"
)

func main() {

	server := handlers.NewServer()

	log.Fatal(server.ListenAndServe())

	//fc := fileconverter.NewFileConverter("/home/gp/GolandProjects/jsonmath/test.json")
	//
	//od, err := fc.ConvertToGo()
	//if err != nil {
	//	os.Exit(1)
	//}
	//
	//od.DoConversion()
	//
	//err = fc.ConvertToFile(od)
	//if err != nil {
	//	os.Exit(1)
	//}
}

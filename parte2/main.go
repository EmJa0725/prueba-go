package main

import (
	"api_proyect/parte2/helpers"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	csvFile, err := os.Open("./parte2/example.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println("error leyendo CSV:" + err.Error())
	}

	sortedCsvLines := helpers.SortCsvData(csvLines)

	jsonData := helpers.ProcessData(sortedCsvLines)

	//Pretty print json json
	finalJson := helpers.PrettyJson(jsonData)

	//Final result
	log.Print(finalJson)
}

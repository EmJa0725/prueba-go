package api

import (
	"api_proyect/parte1/models"
	"encoding/json"
	"log"
)

func CalculateSingleDayInfo(data []byte) models.Report {
	var modelData []models.VentaData
	var total float32 = 0.0
	tdc := make(map[string]float32)
	var noCompraron int = 0
	var compraMasAlta float32 = 0.0

	err := json.Unmarshal(data, &modelData)
	if err != nil {
		log.Print(err)
	}

	for _, element := range modelData {
		// Monto total
		total += element.Monto
		// TDC
		_, exist := tdc[element.Tdc]
		if exist {
			tdc[element.Tdc] += element.Monto
		} else {
			tdc[element.Tdc] = element.Monto
		}
		// No compraron
		if !element.Compro {
			noCompraron += 1
		}
		// Compra mas alta
		if element.Monto > compraMasAlta {
			compraMasAlta = element.Monto
		}
	}
	// Remove empty default values from tdc map
	delete(tdc, "")

	// Set final values for single day
	return models.Report{
		Total:         total,
		ComprasPorTDC: tdc,
		NoCompraron:   noCompraron,
		CompraMasAlta: compraMasAlta,
	}
}

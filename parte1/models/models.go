package models

type VentaData struct {
	ClientId int     `json:"clientId"`
	Phone    string  `json:"phone"`
	Nombre   string  `json:"nombre"`
	Compro   bool    `json:"compro"`
	Tdc      string  `json:"tdc"`
	Monto    float32 `json:"monto"`
	Date     string  `json:"date"`
}

type DefaultResponse struct {
	StartDate string   `json:"startDate"`
	EndDate   string   `json:"endDate"`
	Success   bool     `json:"success"`
	Data      Report   `json:"data"`
	Errors    []string `json:"errors"`
}

type Report struct {
	Total         float32            `json:"total"`
	ComprasPorTDC map[string]float32 `json:"comprasPorTdc,omitempty"`
	NoCompraron   int                `json:"nocompraron"`
	CompraMasAlta float32            `json:"compraMasAlta"`
}

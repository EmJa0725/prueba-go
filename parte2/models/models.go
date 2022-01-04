package models

type CsvData struct {
	Organizacion string
	Usuario      string
	Rol          string
}

type JsonCsvData struct {
	Organizacion string  `json:"organizacion"`
	Users        []Users `json:"users"`
}

type Users struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

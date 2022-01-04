package helpers

import (
	"api_proyect/parte2/models"
	"bytes"
	"encoding/json"
	"log"
	"sort"
	"strings"
)

func ProcessData(data [][]string) []models.JsonCsvData {
	var finalData []models.JsonCsvData
	// Aux vars
	var organizations []string
	var usernames []string
	roles := make(map[string][]string)

	// Search unique organizations
	for i, line := range data {
		if i > 0 { // omit header line
			if !contains(organizations, line[0]) {
				organizations = append(organizations, line[0])
				finalData = append(finalData, models.JsonCsvData{
					Organizacion: line[0],
				})
			}
		}
	}

	//Search unique usernames for each organization
	for i, org := range organizations {
		for j, line := range data {
			if j > 0 { // omit header line
				// save roles on firts iteration
				if i == 1 {
					roles[line[0]+"-"+line[1]] = append(roles[line[0]+"-"+line[1]], line[2])
				}
				// save usernames
				if line[0] == org && !contains(usernames, line[1]) {
					usernames = append(usernames, line[1])
					finalData[i].Users = append(finalData[i].Users, models.Users{
						Username: line[1],
					})
				}
			}
		}
		usernames = nil
	}

	//Set roles list for each organization - username pair
	for key, value := range roles {
		res1 := strings.Split(key, "-")
		org, username := res1[0], res1[1]
		for i, element := range finalData {
			if element.Organizacion == org {
				for j, user := range finalData[i].Users {
					if user.Username == username {
						finalData[i].Users[j].Roles = value
					}
				}
			}
		}
	}

	return finalData
}

// verify if value exists in slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func SortCsvData(data [][]string) [][]string {
	sort.Slice(data[1:], func(i, j int) bool {
		return data[1:][i][0] < data[1:][j][0]
	})

	sort.Slice(data[1:], func(i, j int) bool {
		return data[1:][i][1] < data[1:][j][1]
	})

	return data
}

func PrettyJson(jsonData []models.JsonCsvData) *bytes.Buffer {
	byteData, err := json.Marshal(jsonData)
	if err != nil {
		log.Print("error:", err)
	}

	dst := &bytes.Buffer{}
	if err := json.Indent(dst, byteData, "", "  "); err != nil {
		log.Print("error:", err)
	}

	return dst
}

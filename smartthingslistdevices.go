package smartthings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"strings"
)

func (s SmartThingsConfig) ListSmartthingsDevices() (*DeviceData, error) {

	payload := strings.NewReader(``)

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.smartthings.com/v1/devices", payload)
	if err != nil {
		fmt.Println("Error on line 41 ListSmartthingsDevices:", err)
	}
	req.Header.Add("Authorization", "Bearer "+s.Token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error on line 48 ListSmartthingsDevices:", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error on line 53 ListSmartthingsDevices:", err)
	}
	// fmt.Println(string(body))
	var result DeviceData
	json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error on line 59 ListSmartthingsDevices:", err)
	}
	return &result, nil
}

// func ListDevices(w http.ResponseWriter, r *http.Request) {

// 	result, err := ListSmartthingsDevices()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	sort.Slice(result.Items, func(i, j int) bool {
// 		return result.Items[i].LocationID < result.Items[j].LocationID
// 	})
// 	for _, element := range result.Items {
// 		deviceid := element.DeviceID
// 		name := element.Label
// 		location := ""
// 		fmt.Println("element.LocationID - ", element.LocationID)
// 		if element.LocationID == "05cbd78c-11d8-4ea8-b2c5-50823a9aa470" {
// 			location = "South Pasadena"
// 		} else if element.LocationID == "1518f8a2-4ff4-4894-a06e-b3c28edfec69" {
// 			location = "Summerland"
// 		} else if element.LocationID == "b292c095-eade-43ee-acbd-f1a4236c7e07" {
// 			location = "i3d"
// 		}
// 		fmt.Println("Name: ", name, " ID: ", deviceid)
// 		fmt.Fprint(w, "Location: ", location, " Name: ", name, "ID: ", deviceid, "\n")
// 	}
// }

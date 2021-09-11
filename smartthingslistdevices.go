package smartthings

import (
	"encoding/json"
	"fmt"
	"homeServer/config"
	"io/ioutil"
	"net/http"

	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
)

type DeviceData struct {
	Items []struct {
		DeviceID               string `json:"deviceId"`
		Name                   string `json:"name"`
		Label                  string `json:"label"`
		ManufacturerName       string `json:"manufacturerName,omitempty"`
		PresentationID         string `json:"presentationId,omitempty"`
		DeviceManufacturerCode string `json:"deviceManufacturerCode,omitempty"`
		LocationID             string `json:"locationId"`
		RoomID                 string `json:"roomId,omitempty"`
		DeviceTypeID           string `json:"deviceTypeId,omitempty"`
		DeviceTypeName         string `json:"deviceTypeName,omitempty"`
		DeviceNetworkType      string `json:"deviceNetworkType,omitempty"`
	}
}

func (s SmartThingsConfig) ListSmartthingsDevices() (*DeviceData, error) {

	payload := strings.NewReader(``)

	client := &http.Client{}

	req, err := http.NewRequest("GET", config.Global.SmartThingsDeviceInfoURL, payload)
	if err != nil {
		log.Debug("Error on line 41 ListSmartthingsDevices:", err)
	}
	req.Header.Add("Authorization", "Bearer "+config.Global.SmarthingsToken)

	res, err := client.Do(req)
	if err != nil {
		log.Debug("Error on line 48 ListSmartthingsDevices:", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Debug("Error on line 53 ListSmartthingsDevices:", err)
	}
	// log.Debug(string(body))
	var result DeviceData
	json.Unmarshal(body, &result)
	if err != nil {
		log.Debug("Error on line 59 ListSmartthingsDevices:", err)
	}
	return &result, nil
}
func (s SmartThingsConfig) ListDevices(w http.ResponseWriter, r *http.Request) {

	result, err := ListSmartthingsDevices()
	if err != nil {
		log.Debug(err)
	}
	sort.Slice(result.Items, func(i, j int) bool {
		return result.Items[i].LocationID < result.Items[j].LocationID
	})
	for _, element := range result.Items {
		deviceid := element.DeviceID
		name := element.Label
		location := ""
		log.Debug("element.LocationID - ", element.LocationID)
		if element.LocationID == "05cbd78c-11d8-4ea8-b2c5-50823a9aa470" {
			location = "South Pasadena"
		} else if element.LocationID == "1518f8a2-4ff4-4894-a06e-b3c28edfec69" {
			location = "Summerland"
		} else if element.LocationID == "b292c095-eade-43ee-acbd-f1a4236c7e07" {
			location = "i3d"
		}
		log.Debug("Name: ", name, " ID: ", deviceid)
		fmt.Fprint(w, "Location: ", location, " Name: ", name, "ID: ", deviceid, "\n")
	}
}

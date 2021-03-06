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

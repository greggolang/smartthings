package smartthings

import (
	"encoding/json"
	"homeServer/config"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func (s SmartThingsConfig) SwitchDeviceStatus(deviceid string) (*SwitchDeviceStatusData, error) {

	url := "https://api.smartthings.com/v1/devices/" + deviceid + "/status"

	payload := strings.NewReader(``)

	req, err := http.NewRequest("GET", url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+config.Global.SmarthingsToken)

	res, err := config.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result SwitchDeviceStatusData
	json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (s SmartThingsConfig) MotionDeviceStatus(deviceid string) (*MotionDeviceStatusData, error) {

	url := "https://api.smartthings.com/v1/devices/" + deviceid + "/status"
	method := "GET"

	payload := strings.NewReader(``)

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+config.Global.SmarthingsToken)

	res, err := config.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	log.Info(string(body))
	var result MotionDeviceStatusData
	json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (s SmartThingsConfig) ThermostatDeviceStatus(deviceid string) (*ThermostatDeviceStatusData, error) {

	url := "https://api.smartthings.com/v1/devices/" + deviceid + "/status"
	method := "GET"

	payload := strings.NewReader(``)

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+config.Global.SmarthingsToken)

	res, err := config.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	log.Debug(string(body))
	var result ThermostatDeviceStatusData
	json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (s SmartThingsConfig) DoorDeviceStatus(deviceid string) (*DoorDeviceStatusData, error) {

	url := "https://api.smartthings.com/v1/devices/" + deviceid + "/status"
	method := "GET"

	payload := strings.NewReader(``)

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+config.Global.SmarthingsToken)

	res, err := config.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result DoorDeviceStatusData
	json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (s SmartThingsConfig) GarageDeviceStatus(deviceid string) (*GarageDeviceStatusData, error) {

	url := "https://api.smartthings.com/v1/devices/" + deviceid + "/status"
	method := "GET"

	payload := strings.NewReader(``)

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+config.Global.SmarthingsToken)

	res, err := config.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result GarageDeviceStatusData
	json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

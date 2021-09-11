package smartthings

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type SmartThingsConfig struct {
	Token string
}

func (s SmartThingsConfig) SendCommandOff(deviceid string) error {

	url := "https://api.smartthings.com/v1/devices/" + deviceid + "/commands"
	method := "POST"

	payload := strings.NewReader("{\"commands\": [{\"component\": \"main\",\"capability\": \"switch\",\"command\": \"off\"}]}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+s.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	return nil
}
func (s SmartThingsConfig) SendCommandOn(deviceid string) error {

	url := "https://api.smartthings.com/v1/devices/" + deviceid + "/commands"
	method := "POST"

	payload := strings.NewReader("{\"commands\": [{\"component\": \"main\",\"capability\": \"switch\",\"command\": \"on\"}]}")

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+s.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	return nil
}
func (s SmartThingsConfig) ThermostatSendCommandOff(deviceid string) error {

	url := "https://api.smartthings.com/v1/devices/" + deviceid + "/commands"
	method := "POST"

	payload := strings.NewReader("{\"commands\": [{\"component\": \"main\",\"capability\": \"thermostatMode\",\"command\": \"off\"}]}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+s.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	return nil
}
func (s SmartThingsConfig) ThermostatSendCommandOn(deviceid string) error {

	url := "https://api.smartthings.com/v1/devices/" + deviceid + "/commands"
	method := "POST"

	payload := strings.NewReader("{\"commands\": [{\"component\": \"main\",\"capability\": \"switch\",\"command\": \"on\"}]}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+s.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	return nil
}

package backend

import (
	"homeServer/config"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func SendCommandOff(deviceid string) {

	url := "https://api.smartthings.com/v1/devices/" + deviceid + "/commands"
	method := "POST"

	payload := strings.NewReader("{\"commands\": [{\"component\": \"main\",\"capability\": \"switch\",\"command\": \"off\"}]}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Debug(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+config.Global.SmarthingsToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Debug(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Debug(err)
		return
	}
	log.Debug(string(body))
}
func SendCommandOn(deviceid string) {

	url := "https://api.smartthings.com/v1/devices/" + deviceid + "/commands"
	method := "POST"

	payload := strings.NewReader("{\"commands\": [{\"component\": \"main\",\"capability\": \"switch\",\"command\": \"on\"}]}")

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Debug(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+config.Global.SmarthingsToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Debug(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Debug(err)
		return
	}
	log.Debug(string(body))
}
func ThermostatSendCommandOff(deviceid string) {

	url := "https://api.smartthings.com/v1/devices/" + deviceid + "/commands"
	method := "POST"

	payload := strings.NewReader("{\"commands\": [{\"component\": \"main\",\"capability\": \"thermostatMode\",\"command\": \"off\"}]}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Debug(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+config.Global.SmarthingsToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Debug(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Debug(err)
		return
	}
	log.Debug(string(body))

}
func ThermostatSendCommandOn(deviceid string) {

	url := "https://api.smartthings.com/v1/devices/" + deviceid + "/commands"
	method := "POST"

	payload := strings.NewReader("{\"commands\": [{\"component\": \"main\",\"capability\": \"switch\",\"command\": \"on\"}]}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Debug(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+config.Global.SmarthingsToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Debug(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Debug(err)
		return
	}
	log.Debug(string(body))

}

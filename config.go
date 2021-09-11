package smartthings

import (
	"time"
)

type SwitchDeviceStatusData struct {
	Components struct {
		Main struct {
			Switch struct {
				Switch struct {
					Value string `json:"value"`
				} `json:"switch"`
			} `json:"switch"`
			SwitchLevel struct {
				Level struct {
					Value int `json:"value"`
				} `json:"level"`
			} `json:"switchLevel"`
		} `json:"main"`
	} `json:"components"`
	HealthState struct {
		State           string    `json:"state"`
		LastUpdatedDate time.Time `json:"lastUpdatedDate"`
	} `json:"healthState"`
}
type MotionDeviceStatusData struct {
	Components struct {
		Main struct {
			Configuration struct {
			} `json:"configuration"`
			HealthCheck struct {
				CheckInterval struct {
					Value int    `json:"value"`
					Unit  string `json:"unit"`
					Data  struct {
						Protocol        string `json:"protocol"`
						HubHardwareID   string `json:"hubHardwareId"`
						OfflinePingable string `json:"offlinePingable"`
						DeviceScheme    string `json:"deviceScheme"`
					} `json:"data"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"checkInterval"`
				HealthStatus struct {
					Value interface{} `json:"value"`
					Data  struct {
					} `json:"data"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"healthStatus"`
				DeviceWatchEnroll struct {
					Value     interface{} `json:"value"`
					Timestamp time.Time   `json:"timestamp"`
				} `json:"DeviceWatch-Enroll"`
				DeviceWatchDeviceStatus struct {
					Value interface{} `json:"value"`
					Data  struct {
					} `json:"data"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"DeviceWatch-DeviceStatus"`
			} `json:"healthCheck"`
			TemperatureMeasurement struct {
				Temperature struct {
					Value     float64   `json:"value"`
					Unit      string    `json:"unit"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"temperature"`
			} `json:"temperatureMeasurement"`
			Refresh struct {
			} `json:"refresh"`
			MotionSensor struct {
				Motion struct {
					Value     string    `json:"value"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"motion"`
			} `json:"motionSensor"`
			Sensor struct {
			} `json:"sensor"`
			Battery struct {
				Battery struct {
					Value     int       `json:"value"`
					Unit      string    `json:"unit"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"battery"`
			} `json:"battery"`
			FirmwareUpdate struct {
				LastUpdateStatusReason struct {
					Value     interface{} `json:"value"`
					Timestamp time.Time   `json:"timestamp"`
				} `json:"lastUpdateStatusReason"`
				AvailableVersion struct {
					Value     string    `json:"value"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"availableVersion"`
				LastUpdateStatus struct {
					Value     interface{} `json:"value"`
					Timestamp time.Time   `json:"timestamp"`
				} `json:"lastUpdateStatus"`
				State struct {
					Value     string    `json:"value"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"state"`
				CurrentVersion struct {
					Value     string    `json:"value"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"currentVersion"`
				LastUpdateTime struct {
					Value     interface{} `json:"value"`
					Timestamp time.Time   `json:"timestamp"`
				} `json:"lastUpdateTime"`
			} `json:"firmwareUpdate"`
		} `json:"main"`
	} `json:"components"`
}
type ThermostatDeviceStatusData struct {
	Components struct {
		Main struct {
			RelativeHumidityMeasurement struct {
				Humidity struct {
					Value     int       `json:"value"`
					Unit      string    `json:"unit"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"humidity"`
			} `json:"relativeHumidityMeasurement"`
			Refresh struct {
			} `json:"refresh"`
			Battery struct {
				Battery struct {
					Value     int       `json:"value"`
					Unit      string    `json:"unit"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"battery"`
			} `json:"battery"`
			Actuator struct {
			} `json:"actuator"`
			Thermostat struct {
				ThermostatSetpointRange struct {
					Value     interface{} `json:"value"`
					Timestamp time.Time   `json:"timestamp"`
				} `json:"thermostatSetpointRange"`
				HeatingSetpointRange struct {
					Value     interface{} `json:"value"`
					Timestamp time.Time   `json:"timestamp"`
				} `json:"heatingSetpointRange"`
				ThermostatSetpoint struct {
					Value     int       `json:"value"`
					Unit      string    `json:"unit"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"thermostatSetpoint"`
				SupportedThermostatFanModes struct {
					Value     []string  `json:"value"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"supportedThermostatFanModes"`
				Schedule struct {
					Value     interface{} `json:"value"`
					Timestamp time.Time   `json:"timestamp"`
				} `json:"schedule"`
				CoolingSetpointRange struct {
					Value     interface{} `json:"value"`
					Timestamp time.Time   `json:"timestamp"`
				} `json:"coolingSetpointRange"`
				HeatingSetpoint struct {
					Value     int       `json:"value"`
					Unit      string    `json:"unit"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"heatingSetpoint"`
				CoolingSetpoint struct {
					Value     int       `json:"value"`
					Unit      string    `json:"unit"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"coolingSetpoint"`
				ThermostatOperatingState struct {
					Value     string    `json:"value"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"thermostatOperatingState"`
				Temperature struct {
					Value     int       `json:"value"`
					Unit      string    `json:"unit"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"temperature"`
				ThermostatFanMode struct {
					Value string `json:"value"`
					Data  struct {
						SupportedThermostatFanModes []string `json:"supportedThermostatFanModes"`
					} `json:"data"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"thermostatFanMode"`
				ThermostatMode struct {
					Value string `json:"value"`
					Data  struct {
						SupportedThermostatModes []string `json:"supportedThermostatModes"`
					} `json:"data"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"thermostatMode"`
				SupportedThermostatModes struct {
					Value     []string  `json:"value"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"supportedThermostatModes"`
			} `json:"thermostat"`
			ThermostatOperatingState struct {
				ThermostatOperatingState struct {
					Value     string    `json:"value"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"thermostatOperatingState"`
			} `json:"thermostatOperatingState"`
			HealthCheck struct {
				CheckInterval struct {
					Value int    `json:"value"`
					Unit  string `json:"unit"`
					Data  struct {
						Protocol      string `json:"protocol"`
						HubHardwareID string `json:"hubHardwareId"`
						DeviceScheme  string `json:"deviceScheme"`
					} `json:"data"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"checkInterval"`
				HealthStatus struct {
					Value interface{} `json:"value"`
					Data  struct {
					} `json:"data"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"healthStatus"`
				DeviceWatchEnroll struct {
					Value     interface{} `json:"value"`
					Timestamp time.Time   `json:"timestamp"`
				} `json:"DeviceWatch-Enroll"`
				DeviceWatchDeviceStatus struct {
					Value interface{} `json:"value"`
					Data  struct {
					} `json:"data"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"DeviceWatch-DeviceStatus"`
			} `json:"healthCheck"`
			TemperatureMeasurement struct {
				Temperature struct {
					Value     int       `json:"value"`
					Unit      string    `json:"unit"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"temperature"`
			} `json:"temperatureMeasurement"`
			ThermostatFanMode struct {
				ThermostatFanMode struct {
					Value string `json:"value"`
					Data  struct {
						SupportedThermostatFanModes []string `json:"supportedThermostatFanModes"`
					} `json:"data"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"thermostatFanMode"`
				SupportedThermostatFanModes struct {
					Value     []string  `json:"value"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"supportedThermostatFanModes"`
			} `json:"thermostatFanMode"`
			ThermostatHeatingSetpoint struct {
				HeatingSetpoint struct {
					Value     int       `json:"value"`
					Unit      string    `json:"unit"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"heatingSetpoint"`
			} `json:"thermostatHeatingSetpoint"`
			Sensor struct {
			} `json:"sensor"`
			ThermostatMode struct {
				ThermostatMode struct {
					Value string `json:"value"`
					Data  struct {
						SupportedThermostatModes []string `json:"supportedThermostatModes"`
					} `json:"data"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"thermostatMode"`
				SupportedThermostatModes struct {
					Value     []string  `json:"value"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"supportedThermostatModes"`
			} `json:"thermostatMode"`
			ThermostatCoolingSetpoint struct {
				CoolingSetpoint struct {
					Value     int       `json:"value"`
					Unit      string    `json:"unit"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"coolingSetpoint"`
			} `json:"thermostatCoolingSetpoint"`
		} `json:"main"`
	} `json:"components"`
}
type DoorDeviceStatusData struct {
	Components struct {
		Main struct {
			Battery struct {
				Battery struct {
					Timestamp string `json:"timestamp"`
					Unit      string `json:"unit"`
					Value     int64  `json:"value"`
				} `json:"battery"`
			} `json:"battery"`
			Configuration struct{} `json:"configuration"`
			ContactSensor struct {
				Contact struct {
					Timestamp string `json:"timestamp"`
					Value     string `json:"value"`
				} `json:"contact"`
			} `json:"contactSensor"`
			HealthCheck struct {
				DeviceWatch_DeviceStatus struct {
					Data      struct{}    `json:"data"`
					Timestamp string      `json:"timestamp"`
					Value     interface{} `json:"value"`
				} `json:"DeviceWatch-DeviceStatus"`
				DeviceWatch_Enroll struct {
					Timestamp string      `json:"timestamp"`
					Value     interface{} `json:"value"`
				} `json:"DeviceWatch-Enroll"`
				CheckInterval struct {
					Data struct {
						DeviceScheme  string `json:"deviceScheme"`
						HubHardwareID string `json:"hubHardwareId"`
						Protocol      string `json:"protocol"`
					} `json:"data"`
					Timestamp string `json:"timestamp"`
					Unit      string `json:"unit"`
					Value     int64  `json:"value"`
				} `json:"checkInterval"`
				HealthStatus struct {
					Data      struct{}    `json:"data"`
					Timestamp string      `json:"timestamp"`
					Value     interface{} `json:"value"`
				} `json:"healthStatus"`
			} `json:"healthCheck"`
			Sensor      struct{} `json:"sensor"`
			TamperAlert struct {
				Tamper struct {
					Timestamp string `json:"timestamp"`
					Value     string `json:"value"`
				} `json:"tamper"`
			} `json:"tamperAlert"`
		} `json:"main"`
	} `json:"components"`
}
type GarageDeviceStatusData struct {
	Components struct {
		Main struct {
			Battery struct {
				Battery struct {
					Timestamp string `json:"timestamp"`
					Unit      string `json:"unit"`
					Value     int64  `json:"value"`
				} `json:"battery"`
			} `json:"battery"`
			Configuration struct{} `json:"configuration"`
			ContactSensor struct {
				Contact struct {
					Timestamp string `json:"timestamp"`
					Value     string `json:"value"`
				} `json:"contact"`
			} `json:"contactSensor"`
			HealthCheck struct {
				DeviceWatch_DeviceStatus struct {
					Data      struct{}    `json:"data"`
					Timestamp string      `json:"timestamp"`
					Value     interface{} `json:"value"`
				} `json:"DeviceWatch-DeviceStatus"`
				DeviceWatch_Enroll struct {
					Timestamp string      `json:"timestamp"`
					Value     interface{} `json:"value"`
				} `json:"DeviceWatch-Enroll"`
				CheckInterval struct {
					Data struct {
						DeviceScheme  string `json:"deviceScheme"`
						HubHardwareID string `json:"hubHardwareId"`
						Protocol      string `json:"protocol"`
					} `json:"data"`
					Timestamp string `json:"timestamp"`
					Unit      string `json:"unit"`
					Value     int64  `json:"value"`
				} `json:"checkInterval"`
				HealthStatus struct {
					Data      struct{}    `json:"data"`
					Timestamp string      `json:"timestamp"`
					Value     interface{} `json:"value"`
				} `json:"healthStatus"`
			} `json:"healthCheck"`
			Sensor      struct{} `json:"sensor"`
			TamperAlert struct {
				Tamper struct {
					Timestamp string `json:"timestamp"`
					Value     string `json:"value"`
				} `json:"tamper"`
			} `json:"tamperAlert"`
		} `json:"main"`
	} `json:"components"`
}
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

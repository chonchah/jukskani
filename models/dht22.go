package models

import (
	"errors"
	"fmt"

	"github.com/MichaelS11/go-dht"
)

type SensorDHT struct {
	PinName     string `json:"PinName"`
	Temperature float64
	Humidity    float64
	Sensor      *dht.DHT
}

func (DHT *SensorDHT) Init() {
	err := dht.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
		return
	}
	if DHT.PinName == "" {
		panic(errors.New("Pin for DHT is nil"))
	}
	dht, err := dht.NewDHT(DHT.PinName, dht.Celsius, "")
	if err != nil {
		fmt.Println("NewDHT error:", err, DHT)
		return
	}
	DHT.Sensor = dht

}
func (DHT *SensorDHT) Read() error {
	humidity, temperature, err := DHT.Sensor.ReadRetry(3)

	DHT.Humidity = humidity
	DHT.Temperature = temperature
	return err
}

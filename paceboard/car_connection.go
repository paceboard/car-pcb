package paceboard

import "backend/model"

type CarConnection struct {
	Data          *model.Car
	Name          string
	TrafficIn     uint
	TrafficOut    uint
	Ping          uint
	ArbitraryData interface{}
}

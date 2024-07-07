package model

import "time"

type ParkingLot struct {
	Id          string
	TotalFloors int64
	Floors      []ParkingLotFloor
}

type ParkingLotFloor struct {
	Id         int64
	TotalSlots int64
	Slots      []FloorSlots
}

type FloorSlots struct {
	Id           int64
	Type         string
	Availability bool
}

type Vehicle struct {
	RegNo  string
	Type   string
	Color  string
	Parked bool
}

type ParkingTicket struct {
	Id        string
	Vehicle   Vehicle
	CreatedAt time.Time
}

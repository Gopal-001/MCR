package core

import (
	"fmt"
	"parkingLot/globals"
	"parkingLot/model"
	"strconv"
	"strings"
	"time"
)

func CreateParkingLot(lot *model.ParkingLot, lotId string, noOfFloors int64, noOfSlots int64) {
	floors := []model.ParkingLotFloor{}
	for i := int64(0); i < noOfFloors; i++ {
		slots := []model.FloorSlots{}
		for i := int64(0); i < noOfSlots; i++ {
			if i == 0 {
				slots = append(slots, model.FloorSlots{
					Id:           i + 1,
					Type:         globals.Truck,
					Availability: true,
				})
			} else if i < 3 {
				slots = append(slots, model.FloorSlots{
					Id:           i + 1,
					Type:         globals.Bike,
					Availability: true,
				})
			} else {
				slots = append(slots, model.FloorSlots{
					Id:           i + 1,
					Type:         globals.Car,
					Availability: true,
				})
			}
		}
		floors = append(floors, model.ParkingLotFloor{
			Id:         i + 1,
			TotalSlots: noOfSlots,
			Slots:      slots,
		})
	}
	lot.Id = lotId
	lot.TotalFloors = noOfFloors
	lot.Floors = floors
}

func ParkVehicle(lot *model.ParkingLot, tickets *[]model.ParkingTicket, vehicle *[]model.Vehicle, vehicleType string, regNo string, color string) (string, bool) {
	reqVehicle := model.Vehicle{}
	for i := range *vehicle {
		if (*vehicle)[i].RegNo == regNo {
			reqVehicle = (*vehicle)[i]
			break
		}
	}
	if reqVehicle.Parked {
		return globals.InvalidTicket, false
	}
	if reqVehicle.RegNo != regNo {
		*vehicle = append(*vehicle, model.Vehicle{
			RegNo:  regNo,
			Color:  color,
			Type:   vehicleType,
			Parked: true,
		})
		reqVehicle = (*vehicle)[len(*vehicle)-1]
	}

	floor, slot, isAvail := findAvailSlot(lot, reqVehicle.Type)
	if !isAvail {
		return globals.ParkingLotFull, false
	}

	ticketId := generateTicketId(lot.Id, floor, slot)
	*tickets = append(*tickets, model.ParkingTicket{
		Id:        ticketId,
		Vehicle:   reqVehicle,
		CreatedAt: time.Now(),
	})
	lot.Floors[floor-1].Slots[slot-1].Availability = false

	return globals.ParkedVehicle + " " + ticketId, true
}

func UnParkVehicle(lot *model.ParkingLot, tickets *[]model.ParkingTicket, vehicles *[]model.Vehicle, ticketId string) string {
	parkedVehicle := model.Vehicle{}

	for _, ticket := range *tickets {
		if ticket.Id == ticketId {
			ticketId = ticket.Id
			parkedVehicle = ticket.Vehicle
		}
	}

	for i, vehicle := range *vehicles {
		if vehicle.RegNo == parkedVehicle.RegNo {
			(*vehicles)[i].Parked = false
		}
	}

	if parkedVehicle.RegNo == globals.EMPTY_STRING {
		return globals.InvalidTicket
	}

	idAttr := strings.Split(ticketId, globals.UNDER_SCORE)
	floor, _ := strconv.Atoi(idAttr[1])
	slot, _ := strconv.Atoi(idAttr[2])

	if lot.Floors[floor-1].Slots[slot-1].Availability {
		return globals.InvalidTicket
	} else {
		lot.Floors[floor-1].Slots[slot-1].Availability = true
	}

	return strings.Replace(strings.Replace(globals.UnParkedVehicle, "{regNo}", parkedVehicle.RegNo, 1), "{color}", parkedVehicle.Color, 1)
}

func DisplayFreeSlots(lot *model.ParkingLot, vehicleType string) map[int64][]int64 {
	freeSlots := make(map[int64][]int64)

	for _, floor := range lot.Floors {
		for _, slot := range floor.Slots {
			if slot.Type == vehicleType && slot.Availability {
				freeSlots[floor.Id] = append(freeSlots[floor.Id], slot.Id)
			}
		}
		if len(freeSlots[floor.Id]) == 0 {
			freeSlots[floor.Id] = []int64{}
		}
	}

	return freeSlots
}

func DisplayOccupiedSlots(lot *model.ParkingLot, vehicleType string) map[int64][]int64 {
	occupiedSlots := make(map[int64][]int64)

	for _, floor := range lot.Floors {
		for _, slot := range floor.Slots {
			if slot.Type == vehicleType && !slot.Availability {
				occupiedSlots[floor.Id] = append(occupiedSlots[floor.Id], slot.Id)
			}
		}
		if len(occupiedSlots[floor.Id]) == 0 {
			occupiedSlots[floor.Id] = []int64{}
		}
	}

	return occupiedSlots
}

func findAvailSlot(lot *model.ParkingLot, vehicleType string) (int64, int64, bool) {
	for _, floor := range lot.Floors {
		for _, slot := range floor.Slots {
			if slot.Type == vehicleType && slot.Availability {
				return floor.Id, slot.Id, true
			}
		}
	}
	return 0, 0, false
}

func generateTicketId(parkingId string, floor int64, slot int64) string {
	return parkingId + "_" + fmt.Sprintf("%d", floor) + "_" + fmt.Sprintf("%d", slot)
}

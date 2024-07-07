package main

import (
	"bufio"
	"fmt"
	"os"
	"parkingLot/core"
	"parkingLot/globals"
	"parkingLot/model"
	"strconv"
	"strings"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	var lot = model.ParkingLot{}
	var tickets = []model.ParkingTicket{}
	var vehicles = []model.Vehicle{}

	for {
		scanner.Scan()
		input = scanner.Text()
		if input == globals.EXIT {
			break
		}

		commands := strings.Split(input, " ")
		switch commands[0] {
		case globals.CreateParkingLot:
			floors, _ := strconv.Atoi(commands[2])
			slots, _ := strconv.Atoi(commands[3])
			core.CreateParkingLot(&lot, commands[1], int64(floors), int64(slots))
			fmt.Printf("Created parking lot with %d floors and %d slots per floor\n", floors, slots)
			break

		case globals.ParkVehicle:
			msg, _ := core.ParkVehicle(&lot, &tickets, &vehicles, commands[1], commands[2], commands[3])
			fmt.Printf("%s\n", msg)
			break
		case globals.UnparkVehicle:
			msg := core.UnParkVehicle(&lot, &tickets, &vehicles, commands[1])
			fmt.Printf("%s\n", msg)
			break
		case globals.Display:
			switch commands[1] {
			case globals.FreeCount:
				mp := core.DisplayFreeSlots(&lot, commands[2])
				for key, val := range mp {
					fmt.Printf("No. of free slots for %s on Floor %d: %d\n", commands[2], key, len(val))
				}
				break
			case globals.FreeSlots:
				mp := core.DisplayFreeSlots(&lot, commands[2])
				for key, val := range mp {
					slotIds := ""
					for i := range val {
						slotIds = slotIds + fmt.Sprintf("%d", val[i])
						if i != len(val)-1 {
							slotIds = slotIds + ", "
						}
					}
					fmt.Printf("Free slots for %s on Floor %d: %s\n", commands[2], key, slotIds)
				}
				break
			case globals.OccupiedSlots:
				mp := core.DisplayOccupiedSlots(&lot, commands[2])
				for key, val := range mp {
					slotIds := ""
					for i := range val {
						slotIds = slotIds + fmt.Sprintf("%d", val[i])
						if i != len(val)-1 {
							slotIds = slotIds + ", "
						}
					}
					fmt.Printf("Occupied slots for %s on Floor %d: %s\n", commands[2], key, slotIds)
				}
				break
			}
		}
	}
}

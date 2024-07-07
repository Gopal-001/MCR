package globals

const (
	UNDER_SCORE  = "_"
	EMPTY_STRING = ""
	ParkingLotId = "PR1234"
	EXIT         = "exit"
	Truck        = "TRUCK"
	Bike         = "BIKE"
	Car          = "CAR"

	// Commands
	CreateParkingLot = "create_parking_lot"
	ParkVehicle      = "park_vehicle"
	UnparkVehicle    = "unpark_vehicle"
	Display          = "display"
	FreeCount        = "free_count"
	FreeSlots        = "free_slots"
	OccupiedSlots    = "occupied_slots"

	// Output statements

	ParkedVehicle   = "Parked vehicle. Ticket ID:"
	ParkingLotFull  = "Parking Lot Full"
	UnParkedVehicle = "Unparked vehicle with Registration Number: {regNo} and Color: {color}"
	InvalidTicket   = "Invalid Ticket"
)

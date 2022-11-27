package nilan

import (
	"encoding/binary"
	"errors"
	"time"

	"github.com/goburrow/modbus"
)

// Controller is used for communicating with Nilan CTS700 heatpump over
// Modbus TCP.
type Controller struct {
	Config Config
}

func (c *Controller) getHandler() *modbus.RTUClientHandler {
	handler := modbus.NewRTUClientHandler(c.Config.NilanAddress)
	handler.BaudRate = 19200
	handler.DataBits = 8
	handler.Parity = "E"
	handler.StopBits = 1
	handler.SlaveId = DeviceSlaveID
	handler.Timeout = 10 * time.Second
	err := handler.Connect()

	if err != nil {
		panic(err)
	}

	return handler
}

// FetchValue from InputRegister
func (c *Controller) FetchInputValue(register Register) (uint16, error) {
	handler := c.getHandler()
	defer handler.Close()
	client := modbus.NewClient(handler)
	resultBytes, error := client.ReadInputRegisters(uint16(register), 1)
	if error != nil {
		return 0, error
	}
	if len(resultBytes) == 2 {
		return binary.BigEndian.Uint16(resultBytes), nil
	} else {
		return 0, errors.New("cannot read register value")
	}
}

// FetchValue from HoldingRegister
func (c *Controller) FetchHoldingValue(register Register) (uint16, error) {
	handler := c.getHandler()
	defer handler.Close()
	client := modbus.NewClient(handler)
	resultBytes, error := client.ReadHoldingRegisters(uint16(register), 1)
	if error != nil {
		return 0, error
	}
	if len(resultBytes) == 2 {
		return binary.BigEndian.Uint16(resultBytes), nil
	} else {
		return 0, errors.New("cannot read register value")
	}
}

// Fetch input register values from slave
func (c *Controller) FetchHoldingValues(registers []Register) (map[Register]uint16, error) {
	m := make(map[Register]uint16)

	handler := c.getHandler()
	defer handler.Close()
	client := modbus.NewClient(handler)

	for _, register := range registers {
		resultBytes, err := client.ReadHoldingRegisters(uint16(register), 1)
		if err != nil {
			return m, err
		}
		if len(resultBytes) == 2 {
			resultWord := binary.BigEndian.Uint16(resultBytes)
			m[register] = resultWord
		} else {
			return m, errors.New("no result bytes")
		}
	}

	return m, nil
}

// Fetch input register values from slave
func (c *Controller) FetchInputRegisterValues(registers []Register) (map[Register]uint16, error) {
	m := make(map[Register]uint16)

	handler := c.getHandler()
	defer handler.Close()
	client := modbus.NewClient(handler)

	for _, register := range registers {
		resultBytes, err := client.ReadInputRegisters(uint16(register), 1)
		if err != nil {
			return m, err
		}
		if len(resultBytes) == 2 {
			resultWord := binary.BigEndian.Uint16(resultBytes)
			m[register] = resultWord
		} else {
			return m, errors.New("no result bytes")
		}
	}

	return m, nil
}

// SetRegisterValues on slave
func (c *Controller) SetRegisterValues(values map[Register]uint16) error {
	handler := c.getHandler()
	defer handler.Close()
	client := modbus.NewClient(handler)

	for register, value := range values {
		_, error := client.WriteSingleRegister(uint16(register), value)
		if error != nil {
			return error
		}
	}
	return nil
}

// Register is address of register on client
type Register uint16

const (
	// Default modbus slave address of Nilan EC9
	DeviceSlaveID byte = 30

	//Input registers

	//Errors
	AirFilter     Register = 101
	DoorOpen      Register = 102
	FireSmoke     Register = 103
	FrostOverHeat Register = 105
	HighPressure  Register = 106
	BoilWater     Register = 107
	Defrost       Register = 112

	// T0 - T15 scale is 100
	T0Controller          Register = 200
	T1InTakeTemp          Register = 201
	T2InletTemp           Register = 202 //not in use EC9
	T3ExhaustTemp         Register = 203 //not in use EC9
	T4OutletTemp          Register = 204 //not in use EC9
	T5CondTemp            Register = 205
	T6EvapTemp            Register = 206
	T7InletTemp           Register = 207
	T8OutdoorTemp         Register = 208
	T9HeaterTemp          Register = 209
	T10ExtTemp            Register = 210
	T11HotWaterTopTemp    Register = 211
	T12HotWaterBottomTemp Register = 212
	T13ReturnTemp         Register = 213
	T14SupplyTemp         Register = 214
	T15RoomTemp           Register = 215
	RelativeHumidity      Register = 221
	CO2                   Register = 222
	AlarmStatus           Register = 400
	AlarmID1              Register = 401
	AlarmID2              Register = 404
	AlarmID3              Register = 407
	ControlRunAct         Register = 1000
	ControlModeAct        Register = 1001
	ControlState          Register = 1002
	ControlSecInState     Register = 1003
	VentSetAct            Register = 1100
	InletAct              Register = 1101
	ExhaustAct            Register = 1102
	SinceFiltDay          Register = 1103
	ToFiltDay             Register = 1104
	IsSummer              Register = 1200
	TemperatureSet        Register = 1201
	TempControl           Register = 1202
	TempRoom              Register = 1203
	Efficiency            Register = 1204
	RequestedCapacity     Register = 1205
	ActualCapacity        Register = 1206
	HotWaterType          Register = 1700
	HotWaterAnodeState    Register = 1701
	DisplayLed1           Register = 2000
	DisplayLed2           Register = 2001
	HeatExtSet            Register = 2100

	//Holding registers
	Compressor      Register = 109
	WatherHeat      Register = 116
	CenCircPump     Register = 118
	CenHeat1        Register = 119
	CenHeat2        Register = 120
	CenHeat3        Register = 121
	CenHeatExt      Register = 122
	Defrosting      Register = 125
	ExhaustFanSpeed Register = 200
	InletFanSpeed   Register = 201
	AirHeatCap      Register = 202
	CenHeatCap      Register = 203
	CompresorCap    Register = 204

	// ControlModeSet {0-1; off,on}
	ControlRunSet Register = 1001

	// ControlModeSet {0-4; off,heat,cool,auto,service}
	ControlModeSet Register = 1002
	ControlVentSet Register = 1003
	ControlTempSet Register = 1004

	CoolVent Register = 1101

	CoolSet         Register = 1200
	SummerTempMin   Register = 1201
	WinterTempMin   Register = 1202
	SummerTempMax   Register = 1203
	WinterTempMax   Register = 1204
	SummerTempLimit Register = 1205

	HotWaterTempElectricT11   Register = 1700
	HotWaterTempCompressorT12 Register = 1701
)

// FetchHoldingRegisters
func (c *Controller) FetchHoldingRegisters() (*ReadingHoldings, error) {
	holdingRegisters := []Register{
		Compressor,
		WatherHeat,
		CenCircPump,
		CenHeat1,
		CenHeat2,
		CenHeat3,
		CenHeatExt,
		Defrosting,
		ExhaustFanSpeed,
		InletFanSpeed,
		AirHeatCap,
		CenHeatCap,
		CompresorCap,
		ControlRunSet,
		ControlModeSet,
		ControlVentSet,
		ControlTempSet,
		CoolVent,
		CoolSet,
		SummerTempMin,
		WinterTempMin,
		SummerTempMax,
		WinterTempMax,
		SummerTempLimit,
		HotWaterTempElectricT11,
		HotWaterTempCompressorT12,
	}
	holdingRegistersValues, e1 := c.FetchHoldingValues(holdingRegisters)
	if e1 != nil {
		return nil, e1
	}

	readings := &ReadingHoldings{
		Compressor:                int(holdingRegistersValues[Compressor]),
		WatherHeat:                int(holdingRegistersValues[WatherHeat]),
		CenCircPump:               int(holdingRegistersValues[CenCircPump]),
		CenHeat1:                  int(holdingRegistersValues[CenHeat1]),
		CenHeat2:                  int(holdingRegistersValues[CenHeat2]),
		CenHeat3:                  int(holdingRegistersValues[CenHeat3]),
		CenHeatExt:                int(holdingRegistersValues[CenHeatExt]),
		Defrosting:                int(holdingRegistersValues[Defrosting]),
		ExhaustFanSpeed:           int(holdingRegistersValues[ExhaustFanSpeed]),
		InletFanSpeed:             int(holdingRegistersValues[InletFanSpeed]),
		AirHeatCap:                int(holdingRegistersValues[AirHeatCap]),
		CenHeatCap:                int(holdingRegistersValues[CenHeatCap]),
		CompresorCap:              int(holdingRegistersValues[CompresorCap]),
		ControlRunSet:             int(holdingRegistersValues[ControlRunSet]),
		ControlModeSet:            int(holdingRegistersValues[ControlModeSet]),
		ControlVentSet:            int(holdingRegistersValues[ControlVentSet]),
		ControlTempSet:            int16(holdingRegistersValues[ControlTempSet]),
		CoolVent:                  int(holdingRegistersValues[CoolVent]),
		CoolSet:                   int(holdingRegistersValues[CoolSet]),
		SummerTempMin:             int16(holdingRegistersValues[SummerTempMin]),
		WinterTempMin:             int16(holdingRegistersValues[WinterTempMin]),
		SummerTempMax:             int16(holdingRegistersValues[SummerTempMax]),
		WinterTempMax:             int16(holdingRegistersValues[WinterTempMax]),
		SummerTempLimit:           int16(holdingRegistersValues[SummerTempLimit]),
		HotWaterTempElectricT11:   int16(holdingRegistersValues[HotWaterTempElectricT11]),
		HotWaterTempCompressorT12: int16(holdingRegistersValues[HotWaterTempCompressorT12])}

	return readings, nil
}

// FetchReadings of Nilan sensors
func (c *Controller) FetchReadings() (*Readings, error) {

	inputRegisters := []Register{
		T0Controller,
		T1InTakeTemp,
		T2InletTemp,
		T3ExhaustTemp,
		T4OutletTemp,
		T5CondTemp,
		T6EvapTemp,
		T7InletTemp,
		T8OutdoorTemp,
		T9HeaterTemp,
		T10ExtTemp,
		T11HotWaterTopTemp,
		T12HotWaterBottomTemp,
		T13ReturnTemp,
		T14SupplyTemp,
		T15RoomTemp,
		RelativeHumidity,
		CO2,
		AlarmStatus,
		AlarmID1,
		AlarmID2,
		AlarmID3,
		ControlRunAct,
		ControlModeAct,
		ControlState,
		ControlSecInState,
		VentSetAct,
		InletAct,
		ExhaustAct,
		SinceFiltDay,
		ToFiltDay,
		IsSummer,
		TemperatureSet,
		TempControl,
		TempRoom,
		Efficiency,
		RequestedCapacity,
		ActualCapacity,
		HotWaterType,
		HotWaterAnodeState,
		DisplayLed1,
		DisplayLed2,
		HeatExtSet}

	inputRegistersReadingsRaw, e1 := c.FetchInputRegisterValues(inputRegisters)
	if e1 != nil {
		return nil, e1
	}

	readings := &Readings{
		T0Controller:          int16(inputRegistersReadingsRaw[T0Controller]),
		T1InTakeTemp:          int16(inputRegistersReadingsRaw[T1InTakeTemp]),
		T2InletTemp:           int16(inputRegistersReadingsRaw[T2InletTemp]),
		T3ExhaustTemp:         int16(inputRegistersReadingsRaw[T3ExhaustTemp]),
		T4OutletTemp:          int16(inputRegistersReadingsRaw[T4OutletTemp]),
		T5CondTemp:            int16(inputRegistersReadingsRaw[T5CondTemp]),
		T6EvapTemp:            int16(inputRegistersReadingsRaw[T6EvapTemp]),
		T7InletTemp:           int16(inputRegistersReadingsRaw[T7InletTemp]),
		T8OutdoorTemp:         int16(inputRegistersReadingsRaw[T8OutdoorTemp]),
		T9HeaterTemp:          int16(inputRegistersReadingsRaw[T9HeaterTemp]),
		T10ExtTemp:            int16(inputRegistersReadingsRaw[T10ExtTemp]),
		T11HotWaterTopTemp:    int16(inputRegistersReadingsRaw[T11HotWaterTopTemp]),
		T12HotWaterBottomTemp: int16(inputRegistersReadingsRaw[T12HotWaterBottomTemp]),
		T13ReturnTemp:         int16(inputRegistersReadingsRaw[T13ReturnTemp]),
		T14SupplyTemp:         int16(inputRegistersReadingsRaw[T14SupplyTemp]),
		T15RoomTemp:           int16(inputRegistersReadingsRaw[T15RoomTemp]),
		RelativeHumidity:      int16(inputRegistersReadingsRaw[RelativeHumidity]),
		CO2:                   int16(inputRegistersReadingsRaw[CO2]),
		AlarmStatus:           int(inputRegistersReadingsRaw[AlarmStatus]),
		AlarmID1:              int(inputRegistersReadingsRaw[AlarmID1]),
		AlarmID2:              int(inputRegistersReadingsRaw[AlarmID2]),
		AlarmID3:              int(inputRegistersReadingsRaw[AlarmID3]),
		ControlRunAct:         int(inputRegistersReadingsRaw[ControlRunAct]),
		ControlModeAct:        int(inputRegistersReadingsRaw[ControlModeAct]),
		ControlState:          int(inputRegistersReadingsRaw[ControlState]),
		ControlSecInState:     int(inputRegistersReadingsRaw[ControlSecInState]),
		VentSetAct:            int(inputRegistersReadingsRaw[VentSetAct]),
		InletAct:              int(inputRegistersReadingsRaw[InletAct]),
		ExhaustAct:            int(inputRegistersReadingsRaw[ExhaustAct]),
		SinceFiltDay:          int(inputRegistersReadingsRaw[SinceFiltDay]),
		ToFiltDay:             int(inputRegistersReadingsRaw[ToFiltDay]),
		IsSummer:              int(inputRegistersReadingsRaw[IsSummer]),
		TemperatureSet:        int16(inputRegistersReadingsRaw[TemperatureSet]),
		TempControl:           int16(inputRegistersReadingsRaw[TempControl]),
		TempRoom:              int16(inputRegistersReadingsRaw[TempRoom]),
		Efficiency:            int16(inputRegistersReadingsRaw[Efficiency]),
		RequestedCapacity:     int16(inputRegistersReadingsRaw[RequestedCapacity]),
		ActualCapacity:        int16(inputRegistersReadingsRaw[ActualCapacity]),
		HotWaterType:          int(inputRegistersReadingsRaw[HotWaterType]),
		HotWaterAnodeState:    int(inputRegistersReadingsRaw[HotWaterAnodeState]),
		DisplayLed1:           int(inputRegistersReadingsRaw[DisplayLed1]),
		DisplayLed2:           int(inputRegistersReadingsRaw[DisplayLed2]),
		HeatExtSet:            int16(inputRegistersReadingsRaw[HeatExtSet])}

	return readings, nil
}

func (c *Controller) FetchErrors() (*Errors, error) {
	registers := []Register{
		AirFilter,
		DoorOpen,
		FireSmoke,
		FrostOverHeat,
		HighPressure,
		BoilWater,
		Defrost}

	readings, err := c.FetchInputRegisterValues(registers)
	if err != nil {
		return nil, err
	}

	airFilter := int(readings[AirFilter]) == 1
	//doorOpen := int(readings[DoorOpen]) == 1
	fireSmoke := int(readings[FireSmoke]) == 1
	frostOverHeat := int(readings[FrostOverHeat]) == 1
	highPressure := int(readings[HighPressure]) == 1
	boilWater := int(readings[BoilWater]) == 1
	defrost := int(readings[Defrost]) == 1

	otherErrors := fireSmoke || frostOverHeat || highPressure || boilWater || defrost

	errors := Errors{
		OldFilterWarning: airFilter,
		OtherErrors:      otherErrors,
	}

	return &errors, nil
}

func (c *Controller) FetchSettings() (*Settings, error) {
	registers := []Register{
		ControlModeSet,
		ControlVentSet,
		ControlTempSet,
		CoolVent,
		CoolSet}

	readings, err := c.FetchInputRegisterValues(registers)
	if err != nil {
		return nil, err
	}

	settings := &Settings{
		ModeSet:              ModeSet(readings[ControlModeSet]),
		FanSpeed:             FanSpeed(readings[ControlVentSet]),
		RequestedTemperature: int16(readings[ControlTempSet]),
		CoolVent:             FanSpeed(readings[CoolVent]),
		CoolTemperature:      int16(readings[CoolSet])}

	return settings, nil
}

package nilan

// FanSpeed represents Nilan ventilation intensity value in range from 101 (lowest) to 104 (highest).
type FanSpeed uint16

const (
	FanSpeedOff FanSpeed = 0
	// FanSpeedLow represents lowest fan speed aka level 1
	FanSpeedLow FanSpeed = 1
	// FanSpeedNormal represents normal fan speed aka level 2
	FanSpeedNormal FanSpeed = 2
	// FanSpeedHigh represents high fan speed aka level 3
	FanSpeedHigh FanSpeed = 3
	// FanSpeedVeryHigh represents highest fan speed aka level 4
	FanSpeedVeryHigh FanSpeed = 4
)

type ModeSet uint16

const (
	ModeSetOff     ModeSet = 0
	ModeSetHeat    ModeSet = 1
	ModeSetCool    ModeSet = 2
	ModeSetAuto    ModeSet = 3
	ModeSetService ModeSet = 4
)

// Settings of Nilan system
type Settings struct {
	//main state on/off
	ControlRunSet int
	// Mode set Off, Heat, Cool, Auto, Service
	ModeSet ModeSet
	// FanSpeed of ventilation (VentSet 0 - 5)
	FanSpeed FanSpeed
	// Requested temperature in C (5-40) times 100
	RequestedTemperature int16
	//Cooling ventilation (0 - 5) similar to fanspeed
	CoolVent FanSpeed
	//Cooling temperature set point scale 100
	CoolTemperature int16
}

// ReadHoldingRegisters

type ReadingHoldings struct {
	Compressor                int
	WatherHeat                int
	CenCircPump               int
	CenHeat1                  int
	CenHeat2                  int
	CenHeat3                  int
	CenHeatExt                int
	Defrosting                int
	ExhaustFanSpeed           int
	InletFanSpeed             int
	AirHeatCap                int
	CenHeatCap                int
	CompresorCap              int
	ControlRunSet             int
	ControlModeSet            int
	ControlVentSet            int
	ControlTempSet            int16
	CoolVent                  int
	CoolSet                   int
	SummerTempMin             int16
	WinterTempMin             int16
	SummerTempMax             int16
	WinterTempMax             int16
	SummerTempLimit           int16
	HotWaterTempElectricT11   int16
	HotWaterTempCompressorT12 int16
}

// Readings from Nilan sensors
type Readings struct {
	//T0 - T15 temp values in C scale 100
	T0Controller          int16
	T1InTakeTemp          int16
	T2InletTemp           int16
	T3ExhaustTemp         int16
	T4OutletTemp          int16
	T5CondTemp            int16
	T6EvapTemp            int16
	T7InletTemp           int16
	T8OutdoorTemp         int16
	T9HeaterTemp          int16
	T10ExtTemp            int16
	T11HotWaterTopTemp    int16
	T12HotWaterBottomTemp int16
	T13ReturnTemp         int16
	T14SupplyTemp         int16
	T15RoomTemp           int16
	RelativeHumidity      int16
	CO2                   int16
	AlarmStatus           int
	AlarmID1              int
	AlarmID2              int
	AlarmID3              int
	// Control Run actual (0-1; off, on)
	ControlRunAct int
	// Control Mode (0-4; off, heat, cool, auto, service)
	ControlModeAct     int
	ControlState       int
	ControlSecInState  int
	VentSetAct         int
	InletAct           int
	ExhaustAct         int
	SinceFiltDay       int
	ToFiltDay          int
	IsSummer           int
	TemperatureSet     int16
	TempControl        int16
	TempRoom           int16
	Efficiency         int16
	RequestedCapacity  int16
	ActualCapacity     int16
	HotWaterType       int
	HotWaterAnodeState int
	DisplayLed1        int
	DisplayLed2        int
	HeatExtSet         int16
}

type Errors struct {
	// Indicates a need of new filters
	OldFilterWarning bool
	// Indicates other problems that must to be checked
	OtherErrors bool
}

package loadout

// Fuel contains the main and reserve fuel capacities.
type Fuel struct {
	Main    float64 `json:"Main"`
	Reserve float64 `json:"Reserve"`
}

// Module contains information about a ship module.
type Module struct {
	Slot        string      `json:"Slot"`
	Item        string      `json:"Item"`
	On          bool        `json:"On"`
	Priority    int64       `json:"Priority"`
	Health      int64       `json:"Health"`
	Engineering Engineering `json:"Engineering"`
}

// Modifier describes an engineering modifier.
type Modifier struct {
	Label         string  `json:"Label"`
	Value         float64 `json:"Value"`
	OriginalValue float64 `json:"OriginalValue"`
	LessIsGood    int64   `json:"LessIsGood"`
}

// Engineering represents the engineering modifications performed on a module.
type Engineering struct {
	Engineer           string     `json:"Engineer"`
	EngineerID         int64      `json:"EngineerID"`
	BlueprintID        int64      `json:"BlueprintID"`
	BlueprintName      string     `json:"BlueprintName"`
	Level              int64      `json:"Level"`
	Quality            float64    `json:"Quality"`
	ExperimentalEffect string     `json:"ExperimentalEffect"`
	Modifiers          []Modifier `json:"Modifiers"`
}

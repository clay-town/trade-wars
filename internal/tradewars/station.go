package tradewars

type Station struct {
    Designation string `json:"designation"`
    Location string `json:"location"`
    Cargos []Cargo `json:"cargos"`
}

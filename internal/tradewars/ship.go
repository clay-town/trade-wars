package tradewars

type Ship struct {
    Callsign string `json:"callsign"`
    Location string `json:"location"`
    Cargos []Cargo `json:"cargos"`
    Cubits int `json:"cubits"`
    Online string `json:"online"`

}

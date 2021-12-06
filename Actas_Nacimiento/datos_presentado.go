package datos_presentado

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    FechaNaci    uint      `json:"FechaNaci,omitempty"`
    Sexo    string      `json:"Sexo,omitempty"`
    HoraNaci    uint      `json:"HoraNaci,omitempty"`
    PaisNaci    string      `json:"PaisNaci,omitempty"`
    EdoNaci    string      `json:"EdoNaci,omitempty"`
    muniNaci    string      `json:"muniNaci,omitempty"`
    ParroquiaNaci    string      `json:"ParroquiaNaci,omitempty"`
    Direccion    string      `json:"Direccion,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
  }
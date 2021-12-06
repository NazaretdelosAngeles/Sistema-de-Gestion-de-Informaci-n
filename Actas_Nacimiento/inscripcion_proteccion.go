package inscripcion_proteccion

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    ConcejoProteNiños    string      `json:"ConcejoProteNiños,omitempty"`
    Medida    uint      `json:"Medida,omitempty"`
    Fecha    uint      `json:"Fecha,omitempty"`
    NombresApellidosConce    string      `json:"NombresApellidosConce,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
}
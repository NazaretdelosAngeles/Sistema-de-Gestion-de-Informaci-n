package inscripcion_judicial

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Tribunal    string      `json:"Tribunal,omitempty"`
    SentenciaN    uint      `json:"SentenciaN,omitempty"`
    NombresApellidosJuez    string      `json:"NombresApellidosJuez,omitempty"`
    Fecha    uint      `json:"Fecha,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
}
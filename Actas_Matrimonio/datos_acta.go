package datos_acta

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Numero    uint      `json:"Numero,omitempty"`
    Fecha    uint      `json:"Fecha,omitempty"`
    Autoridad    string      `json:"Autoridad,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
}
package inscripcion_extem

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    InfoConceProte    string      `json:"InfoConceProte,omitempty"`
    Autoridad    string      `json:"Autoridad,omitempty"`
    Fecha    uint      `json:"Fecha,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
}
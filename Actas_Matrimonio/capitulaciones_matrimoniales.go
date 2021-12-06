package capitulaciones_matrimoniales

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    DatosRegistro    string      `json:"DatosRegistro,omitempty"`
    Nro    uint      `json:"Nro,omitempty"`
    Fecha    uint      `json:"Fecha,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
    }
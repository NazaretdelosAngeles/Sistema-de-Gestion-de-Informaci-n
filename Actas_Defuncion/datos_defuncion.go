package datos_defuncion

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Fecha    uint      `json:"Fecha,omitempty"`
    Hora    uint      `json:"Hora,omitempty"`
    Pais    string      `json:"Pais,omitempty"`
    Edo    string      `json:"Edo,omitempty"`
    Mnpio    string      `json:"Mnpio,omitempty"`
    Parr    string      `json:"Parr,omitempty"`
    Causas    string      `json:"Causas,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
    }
package Lugar_Hr_Fecha

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Lugar    string      `json:"Lugar,omitempty"`
    Hora    uint      `json:"Hora,omitempty"`
    Fecha    uint      `json:"Fecha,omitempty"`
    DatosExpediente    uint      `json:"DatosExpediente,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
    }
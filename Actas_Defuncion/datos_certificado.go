package datos_certificado

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Certficado    uint      `json:"Certficado,omitempty"`
    FechaExpedicion    uint      `json:"FechaExpedicion,omitempty"`
    Autoridad    string      `json:"Autoridad,omitempty"`
    Sexo    string      `json:"Sexo,omitempty"`
    NMPPS    uint      `json:"NMPPS,omitempty"`
    CentroSalud    string      `json:"CentroSalud,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
    }
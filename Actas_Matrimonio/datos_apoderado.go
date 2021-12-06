package datos_apoderado

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    DatosNotaria    uint      `json:"DatosNotaria,omitempty"`
    Nro    uint      `json:"Nro,omitempty"`
    Fecha    uint      `json:"Fecha,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
    }
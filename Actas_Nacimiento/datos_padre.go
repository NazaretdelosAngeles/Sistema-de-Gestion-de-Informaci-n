package datos_padre

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    Edad    uint      `json:"Edad,omitempty"`
    Nacionalidad    string      `json:"Nacionalidad,omitempty"`
    Profesion    string      `json:"Profesion,omitempty"`
    Comunidad    string      `json:"Comunidad,omitempty"`
    Direccion    string      `json:"Direccion,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
}
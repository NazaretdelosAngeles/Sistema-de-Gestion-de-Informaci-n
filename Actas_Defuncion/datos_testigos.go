package datos_testigos

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
    Nombres2    string      `json:"Nombres2,omitempty"`
    Apellidos2    string      `json:"Apellidos2,omitempty"`
    DocuIde2    uint      `json:"DocuIde2,omitempty"`
    Edad2    uint      `json:"Edad2,omitempty"`
    Nacionalidad2    string      `json:"Nacionalidad2,omitempty"`
    Profesion2    string      `json:"Profesion2,omitempty"`
    Comunidad2    string      `json:"Comunidad2,omitempty"`
    Direccion2    string      `json:"Direccion2,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
    }
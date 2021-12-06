package datos_testigos

import "time"

// Post created by a user.
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
    Nombres2    string      `json:"Nombres,omitempty"`
    Apellidos2    string      `json:"Apellidos,omitempty"`
    DocuIde2    uint      `json:"DocuIde,omitempty"`
    Edad2    uint      `json:"Edad,omitempty"`
    Nacionalidad2    string      `json:"Nacionalidad,omitempty"`
    Profesion2    string      `json:"Profesion,omitempty"`
    Comunidad2    string      `json:"Comunidad,omitempty"`
    Direccion2    string      `json:"Direccion,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
  }
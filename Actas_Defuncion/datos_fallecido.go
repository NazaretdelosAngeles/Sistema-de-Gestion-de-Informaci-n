package datos_fallecido

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    FechaNaci    uint      `json:"FechaNaci,omitempty"`
    LugarNaci    string      `json:"PaisNaci,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    Edad    uint      `json:"Edad,omitempty"`
    Sexo    string      `json:"Sexo,omitempty"`
    EstadoCivil    string      `json:"EstadoCivil,omitempty"`
    Nacionalidad    string      `json:"Nacionalidad,omitempty"`
    Profesion    string      `json:"Profesion,omitempty"`
    Comunidad    string      `json:"Comunidad,omitempty"`
    Direccion    string      `json:"Direccion,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
    }
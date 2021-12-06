package datos_conyugem

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    FechaNaci    uint      `json:"FechaNaci,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    EdoNaci    string      `json:"EdoNaci,omitempty"`
    MnpioNaci    string      `json:"MnpioNaci,omitempty"`
    ParrNaci    string      `json:"ParrNaci,omitempty"`
    Nacionalidad    string      `json:"Nacionalidad,omitempty"`
    EdoResid    string      `json:"EdoResid,omitempty"`
    MnpioResid    string      `json:"MnpioResid,omitempty"`
    ParrResid    string      `json:"ParrResid,omitempty"`
    Profesion    string      `json:"Profesion,omitempty"`
    NombresApellidosPadre    string      `json:"NombresApellidosPadre,omitempty"`
    DocuIdePadre    uint      `json:"DocuIdePadre,omitempty"`
    ProfesionPadre    string      `json:"ProfesionPadre,omitempty"`
    ResidenciaPadre    string      `json:"ResidenciaPadre,omitempty"`
    NombresApellidosMadre    string      `json:"NombresApellidosMadre,omitempty"`
    DocuIdeMadre    uint      `json:"DocuIdeMadre,omitempty"`
    ProfesionMadre    string      `json:"ProfesionMadre,omitempty"`
    ResidenciaMadre    string      `json:"ResidenciaMadre,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
    }
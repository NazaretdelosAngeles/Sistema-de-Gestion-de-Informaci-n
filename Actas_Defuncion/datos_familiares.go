package datos_familiares

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    NombresApellidosPareja    string      `json:"NombresApellidosPareja,omitempty"`
    Vive    string      `json:"Vive,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    Profesion    string      `json:"Profesion,omitempty"`
    Nacionalidad    string      `json:"Nacionalidad,omitempty"`
    Residencia    string      `json:"Residencia,omitempty"`
    NombresApellidoshijo1    string      `json:"NombresApellidoshijo1,omitempty"`
    DocuIdehijo1    uint      `json:"DocuIdehijo1,omitempty"`
    ViveHijo1    string      `json:"ViveHijo1,omitempty"`
    NombresApellidoshijo2    string      `json:"NombresApellidoshijo2,omitempty"`
    DocuIdehijo2    uint      `json:"DocuIdehijo2,omitempty"`
    ViveHijo2    string      `json:"ViveHijo2,omitempty"`
    NombresApellidoshijo3    string      `json:"NombresApellidoshijo3,omitempty"`
    DocuIdehijo3    uint      `json:"DocuIdehijo3,omitempty"`
    ViveHijo3    string      `json:"ViveHijo3,omitempty"`
    NombresApellidoshijo4    string      `json:"NombresApellidoshijo4,omitempty"`
    DocuIdehijo4    uint      `json:"DocuIdehijo4,omitempty"`
    ViveHijo4    string      `json:"ViveHijo4,omitempty"`
    NombresApellidoshijo5    string      `json:"NombresApellidoshijo5,omitempty"`
    DocuIdehijo5    uint      `json:"DocuIdehijo5,omitempty"`
    ViveHijo5    string      `json:"ViveHijo5,omitempty"`
    NombresApellidoshijo6    string      `json:"NombresApellidoshijo6,omitempty"`
    DocuIdehijo6    uint      `json:"DocuIdehijo6,omitempty"`
    ViveHijo6    string      `json:"ViveHijo6,omitempty"`
    NombresApellidoshijo7    string      `json:"NombresApellidoshijo7,omitempty"`
    DocuIdehijo7    uint      `json:"DocuIdehijo7,omitempty"`
    ViveHijo7    string      `json:"ViveHijo7,omitempty"`
    NombresApellidosMadre    string      `json:"NombresApellidosMadre,omitempty"`
    DocuIdeMadre    uint      `json:"DocuIdeMadre,omitempty"`
    ViveMadre    string      `json:"ViveMadre,omitempty"`
    NombresApellidosPadre    string      `json:"NombresApellidosPadre,omitempty"`
    DocuIdePadre    uint      `json:"DocuIdePadre,omitempty"`
    VivePadre    string      `json:"VivePadre,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
    }
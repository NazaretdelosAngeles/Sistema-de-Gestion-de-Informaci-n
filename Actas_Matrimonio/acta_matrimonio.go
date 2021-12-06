package acta_matrimonio

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    DatosRegistrador    uint      `json:"DatosRegistrador,omitempty"`
    LugarHrFecha    uint      `json:"LugarHrFecha,omitempty"`
    DatosConyugeM    uint      `json:"DatosConyugeM,omitempty"`
    DatosConyugeF    uint      `json:"DatosConyugeF,omitempty"`
    DDConyuges    uint      `json:"DDConyuges,omitempty"`
    AceptacionExpresa    uint      `json:"AceptacionExpresa,omitempty"`
    CapitulacionesMatrimoniales    uint      `json:"CapitulacionesMatrimoniales,omitempty"`
    DatosApoderado    uint      `json:"DatosApoderado,omitempty"`
    DatosHijos    uint      `json:"DatosHijos,omitempty"`
    DatosTestigo    uint      `json:"DatosTestigo,omitempty"`
    DatosActa    uint      `json:"DatosActa,omitempty"`
    CircunstanciasEspe    uint      `json:"CircunstanciasEspe,omitempty"`
    NotaMarginal    uint      `json:"NotaMarginal,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
    }
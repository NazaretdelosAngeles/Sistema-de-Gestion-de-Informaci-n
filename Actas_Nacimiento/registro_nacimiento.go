package registro_nacimiento

import "time"


type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    DatosRegistrador    uint      `json:"DatosRegistrador,omitempty"`
    DatosPresentado    uint      `json:"DatosPresentado,omitempty"`
    DatosCertificado    uint      `json:"DatosCertificado,omitempty"`
    DatosMadre    uint      `json:"DatosMadre,omitempty"`
    DatosPadre    uint      `json:"DatosPadre,omitempty"`
    DatosDeclarante    uint      `json:"DatosDeclarante,omitempty"`
    DatosTestigos    uint      `json:"DatosTestigos,omitempty"`
    DatosActa    uint      `json:"DatosActa,omitempty"`
    InscripcionProteccion    uint      `json:"InscripcionProteccion,omitempty"`
    InscripcionJudicial    uint      `json:"InscripcionJudicial,omitempty"`
    InscripcionExtem    uint      `json:"InscripcionExtem,omitempty"`
    CircunstanciasEspe    uint      `json:"CircunstanciasEspe,omitempty"`
    DocuPresentados    uint      `json:"DocuPresentados,omitempty"`
    NotaMarginal    uint      `json:"NotaMarginal,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
    }
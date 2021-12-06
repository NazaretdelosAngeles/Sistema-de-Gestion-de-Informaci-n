package registro_defuncion

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    DatosRegistrador    string      `json:"DatosRegistrador,omitempty"`
    DatosFallecido    string      `json:"DatosFallecido,omitempty"`
    DatosDefuncion    string      `json:"DatosDefuncion,omitempty"`
    DatosCertificado    string      `json:"DatosCertificado,omitempty"`
    DatosFamiliares    string      `json:"DatosFamiliares,omitempty"`
    DatosDeclarante    string      `json:"DatosDeclarante,omitempty"`
    InscripcionJudicial    string      `json:"InscripcionJudicial,omitempty"`
    DatosTestigos    string      `json:"DatosTestigos,omitempty"`
    CircunstanciasEspe    string      `json:"CircunstanciasEspe,omitempty"`
    DocuPresentados    string      `json:"DocuPresentados,omitempty"`
    NotaMarginal    uint      `json:"NotaMarginal,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
}
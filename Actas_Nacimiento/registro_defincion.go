// pkg/acta/acta.go
package registro_defuncion.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    DatosRegistrador    uint      `json:"user_id,omitempty"`
    DatosFallecido    uint      `json:"user_id,omitempty"`
    DatosDefuncion    uint      `json:"user_id,omitempty"`
    DatosCertificado    uint      `json:"user_id,omitempty"`
    DatosFamiliares    uint      `json:"user_id,omitempty"`
    DatosDeclarante    uint      `json:"user_id,omitempty"`
    InscripcionJudicial    uint      `json:"InscripcionJudicial,omitempty"`
    DatosTestigos    uint      `json:"user_id,omitempty"`
    CircunstanciasEspe    uint      `json:"user_id,omitempty"`
    DocuPresentados    uint      `json:"user_id,omitempty"`
    NotaMarginal    uint      `json:"user_id,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_Registrador.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    Oficina    string      `json:"Oficina,omitempty"`
    Resolucion    uint      `json:"Resolucion,omitempty"`
    Fecha    string      `json:"Fecha,omitempty"`
    Resolucion    uint      `json:"Resolucion,omitempty"`
    Fecha    string      `json:"Fecha,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_Fallecido.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    FechaNaci    string      `json:"FechaNaci,omitempty"`
    LugarNaci    string      `json:"PaisNaci,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    Edad    string      `json:"HoraNaci,omitempty"`
    Sexo    string      `json:"Sexo,omitempty"`
    EstadoCivil    string      `json:"HoraNaci,omitempty"`
    Nacionalidad    string      `json:"PaisNaci,omitempty"`
    Profesion    string      `json:"EdoNaci,omitempty"`
    Comunidad    string      `json:"muniNaci,omitempty"`
    Direccion    string      `json:"Direccion,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_Defuncion.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Fecha    string      `json:"Fecha,omitempty"`
    Hora    string      `json:"Hora,omitempty"`
    Pais    string      `json:"Pais,omitempty"`
    Edo    string      `json:"Edo,omitempty"`
    Mnpio    string      `json:"Mnpio,omitempty"`
    Parr    string      `json:"Parr,omitempty"`
    Causas    string      `json:"Causas,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_Certificado.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Certficado    string      `json:"Certficado,omitempty"`
    FechaExpedicion    string      `json:"FechaExpedicion,omitempty"`
    Autoridad    string      `json:"Autoridad,omitempty"`
    Sexo    string      `json:"Sexo,omitempty"`
    NMPPS    string      `json:"NMPPS,omitempty"`
    CentroSalud    string      `json:"CentroSalud,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_Familiares.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    NombresApellidosPareja    string      `json:"Nombres,omitempty"`
    Vive    string      `json:"Apellidos,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    Profesion    string      `json:"Profesion,omitempty"`
    Nacionalidad    uint      `json:"Nacionalidad,omitempty"`
    Residencia    uint      `json:"Comunidad,omitempty"`
    NombresApellidoshijo1    string      `json:"Direccion,omitempty"`
    DocuIdehijo1    uint      `json:"DocuIde,omitempty"`
    ViveHijo1    string      `json:"Apellidos,omitempty"`
    NombresApellidoshijo2    string      `json:"Direccion,omitempty"`
    DocuIdehijo2    uint      `json:"DocuIde,omitempty"`
    ViveHijo2    string      `json:"Apellidos,omitempty"`
    NombresApellidoshijo3    string      `json:"Direccion,omitempty"`
    DocuIdehijo3    uint      `json:"DocuIde,omitempty"`
    ViveHijo3    string      `json:"Apellidos,omitempty"`
    NombresApellidoshijo4    string      `json:"Direccion,omitempty"`
    DocuIdehijo4    uint      `json:"DocuIde,omitempty"`
    ViveHijo4    string      `json:"Apellidos,omitempty"`
    NombresApellidoshijo5    string      `json:"Direccion,omitempty"`
    DocuIdehijo5    uint      `json:"DocuIde,omitempty"`
    ViveHijo5    string      `json:"Apellidos,omitempty"`
    NombresApellidoshijo6    string      `json:"Direccion,omitempty"`
    DocuIdehijo6    uint      `json:"DocuIde,omitempty"`
    ViveHijo6    string      `json:"Apellidos,omitempty"`
    NombresApellidoshijo7    string      `json:"Direccion,omitempty"`
    DocuIdehijo7    uint      `json:"DocuIde,omitempty"`
    ViveHijo7    string      `json:"Apellidos,omitempty"`
    NombresApellidosMadre    string      `json:"Nombres,omitempty"`
    DocuIdeMadre    uint      `json:"DocuIde,omitempty"`
    ViveMadre    string      `json:"Apellidos,omitempty"`
    NombresApellidosPadre    string      `json:"Nombres,omitempty"`
    DocuIdePadre    uint      `json:"DocuIde,omitempty"`
    VivePadre    string      `json:"Apellidos,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_Declarante.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    Edad    string      `json:"Edad,omitempty"`
    Nacionalidad    uint      `json:"Nacionalidad,omitempty"`
    Profesion    string      `json:"Profesion,omitempty"`
    Comunidad    uint      `json:"Comunidad,omitempty"`
    Direccion    string      `json:"Direccion,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Inscripcion_Judicial.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Tribunal    string      `json:"Tribunal,omitempty"`
    SentenciaN    string      `json:"SentenciaN,omitempty"`
    NombresApellidosJuez    uint      `json:"NombresApellidosJuez,omitempty"`
    ExtractoSentencia    uint      `json:"Fecha,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_Testigos.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    Edad    string      `json:"Edad,omitempty"`
    Nacionalidad    uint      `json:"Nacionalidad,omitempty"`
    Profesion    string      `json:"Profesion,omitempty"`
    Comunidad    uint      `json:"Comunidad,omitempty"`
    Direccion    string      `json:"Direccion,omitempty"`
    Nombres2    string      `json:"Nombres,omitempty"`
    Apellidos2    string      `json:"Apellidos,omitempty"`
    DocuIde2    uint      `json:"DocuIde,omitempty"`
    Edad2    string      `json:"Edad,omitempty"`
    Nacionalidad2    uint      `json:"Nacionalidad,omitempty"`
    Profesion2    string      `json:"Profesion,omitempty"`
    Comunidad2    uint      `json:"Comunidad,omitempty"`
    Direccion2    string      `json:"Direccion,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Circunstancias_Espe.go

import "time"

// Post created by a note.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nota    string      `json:"InfoConceProte,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Docu_Presentados.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nota    string      `json:"InfoConceProte,omitempty"`    
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

    // pkg/acta/acta.go
package Nota_Marginal.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nota    string      `json:"InfoConceProte,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;


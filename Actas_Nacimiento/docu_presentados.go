package docu_presentados

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nota    string      `json:"Nota,omitempty"`    
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
}
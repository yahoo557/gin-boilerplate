package common

import "time"

type BaseEntity struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Product) BeforeCreate() {
	// Set CreatedAt and UpdatedAt before creating the entity
	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now
}

func (p *Product) BeforeUpdate() {
	// Update the UpdatedAt timestamp before updating the entity
	p.UpdatedAt = time.Now()
}

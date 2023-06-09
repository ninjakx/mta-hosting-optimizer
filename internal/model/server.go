package model

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

type Server struct {
	gorm.Model        // Includes fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	IP         string `json:"IP" validate:"ip_address"`
	Hostname   string
	Active     bool
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	// Auto migrate the models
	db.AutoMigrate(&Server{})
	db.AutoMigrate(&User{})

	return db
}

func (s *Server) Disable() {
	s.Active = false
}

func (s *Server) Enable() {
	s.Active = true
}

package api

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Config struct {
	Db     *gorm.DB
	Router *mux.Router
}

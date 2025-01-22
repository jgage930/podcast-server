package common

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
)

func GetById[T any](buf *T, id string, db *gorm.DB, w http.ResponseWriter) {
	err := db.First(&buf, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "Record Not Found", http.StatusNotFound)
	}
}

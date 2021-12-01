package handlers

import (
	"log"
)

type Handler struct {
	l *log.Logger
}

func NewHandler(l *log.Logger) *Handler {
	return &Handler{l}
}

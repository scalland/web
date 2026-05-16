package handlers

import (
	"/internal/worker"
	"/pkg/utils"
)

// WebHandlers holds dependencies for all HTTP handlers.
type WebHandlers struct {
	Utils  *utils.Utils
	Worker *worker.Worker
}

// NewWebHandlers creates a new WebHandlers instance.
func NewWebHandlers(u *utils.Utils, w *worker.Worker) *WebHandlers {
	return &WebHandlers{Utils: u, Worker: w}
}

package handlers

/*
 * Copyright (c) 2022 Norwegian University of Science and Technology
 */

import (
	"go.uber.org/zap"
	"net/http"
)

// NotFound is a convenience handler for catching unknown paths and log the path from the request.
func NotFound(log *zap.SugaredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infof("path not found %s", r.URL.Path)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

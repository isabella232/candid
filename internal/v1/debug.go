// Copyright 2014 Canonical Ltd.

package v1

import (
	"net/http"

	"github.com/juju/utils/debugstatus"
)

// GET /debug
func (h *Handler) serveDebug(http.ResponseWriter, *http.Request) error {
	return errNotImplemented
}

// GET /debug/status
func (h *Handler) serveDebugStatus(_ http.Header, req *http.Request) (interface{}, error) {
	return debugstatus.Check(
		debugstatus.StartTime,
		debugstatus.Connection(h.store.DB.Session),
		debugstatus.MongoCollections(h.store.DB),
	), nil
}
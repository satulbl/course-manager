package v1

import "github.com/google/uuid"

type simpleIdResponse struct {
	ID uuid.UUID `json:"id"`
}

func newSimpleIdResponse(id uuid.UUID) simpleIdResponse {
	return simpleIdResponse{ID: id}
}
package fetch

import (
	"time"

	"github.com/gofrs/uuid"
)

type (
	Status string
	Kind   *int

	SimpleRequest struct {
		UserID string `json:"userID"`
	}

	MoreComplexRequest struct {
		ID     uuid.UUID `json:"id"`
		UserID string    `json:"userID"`
		Age    int       `json:"age"`
		Now    time.Time `json:"now"`
	}

	PointerRequest struct {
		Age *int `json:"age"`
	}

	ArrayRequest struct {
		IDs []string `json:"ids"`
	}

	TypeDefRequest struct {
		RequestStatus Status `json:"requestStatus"`
		RequestKind   Kind   `json:"requestKind"`
	}
)

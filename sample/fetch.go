package fetch

import (
	"time"
)

type (
	String  string
	Int     int
	Times   []*time.Time
	Numbers []int

	FetchRequest struct {
		ID string `json:"ID"`
	}

	FetchResponse struct {
		ID                 string     `json:"ID"`
		GroupID            string     `json:"groupID"`
		Name               string     `json:"name"`
		ShortName          *string    `json:"shortName"`
		Age                int        `json:"age"`
		StringSample       String     `json:"stringSample"`
		IntSample          Int        `json:"intSample"`
		Strings            []string   `json:"strings"`
		TimeSample         time.Time  `json:"timeSample"`
		TimeNullableSample *time.Time `json:"timeNullableSample"`
		TimesSample        Times      `json:"timesSample"`
		NumbersSample      Numbers    `json:"numbersSample"`
	}
)

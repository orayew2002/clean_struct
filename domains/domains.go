package domains

type TaggedStruct struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Score    int64  `db:"score"`
	Verified bool   `json:"verified"`
}

type MixedStruct struct {
	Name   string
	ID     int
	Active bool
	Flag   bool
}

type PointerStruct struct {
	Name *string
	Age  *int
	Data *[]byte
}

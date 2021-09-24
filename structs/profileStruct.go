package structs

type Profile struct {
	ID    *int64  `json:"id"`
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Age   *int64  `json:"age"`
}

type DelResult struct {
	RowNum int64 `json:"rowNum"`
}

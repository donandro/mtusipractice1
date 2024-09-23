package models

type GetPillsResponse struct {
	Pills []Pill `json:"pills"`
}

type Pill struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Dosage      string `json:"dosage"`
	Instruction string `json:"instruction"`
}

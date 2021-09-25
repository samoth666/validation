package models

type Dna struct {
	Id           string   `json:"id"`
	DnaSecuences []string `json:"dna"`
	IsMutant     bool     `json:"is_mutant"`
}

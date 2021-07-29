package internal

import "meteor/enum"

type FamilyMember struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	Gender enum.Gender `json:"gender"`
	MaritalStatus enum.MaritalStatus `json:"marital_status"`
	Spouse string `json:"spouse"`
	OccupationType enum.OccupationType `json:"occupation_type"`
	AnnualIncome float64 `json:"annual_income"`
	DOB string `json:"dob"`
}

func (member *FamilyMember) AddMember() error {
	return nil
}



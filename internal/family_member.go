package internal

import "meteor/enum"

type FamilyMember struct {
	Name string
	Gender enum.Gender
	MaritalStatus enum.MaritalStatus
	Spouse string
	OccupationType enum.OccupationType
	AnnualIncome float64
	DOB string
}

func (member *FamilyMember) AddMember() error {
	return nil
}



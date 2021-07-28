package internal

import "meteor/enum"

type Household struct {
	HousingType enum.HouseholdType
	FamilyMembers []*FamilyMember
}

func (Household *Household) Create(housingType enum.HouseholdType) error {
	// Create Household here
	return nil
}

func QueryHouseholds() []*Household {
	// Query all households
	return []*Household{}
}

func QueryUniqueHousehold(MemberName string, HousingType enum.HouseholdType) *Household{
	return &Household{}
}


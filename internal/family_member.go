package internal

import (
	"fmt"
	"gorm.io/gorm"
	"meteor/enum"
)

type FamilyMember struct {
	ID uint64 `json:"id"`
	HouseholdID uint64 `json:"household_id"`
	Name string `json:"name"`
	Gender enum.Gender `json:"gender"`
	MaritalStatus enum.MaritalStatus `json:"marital_status"`
	Spouse string `json:"spouse"`
	OccupationType enum.OccupationType `json:"occupation_type"`
	AnnualIncome float64 `json:"annual_income"`
	DOB string `json:"dob"`
}

// TODO: Check if marital status must not have spouse
func AddFamilyMember(db *gorm.DB, newMember *FamilyMember) error {
	// Check if household exists
	err := QueryHouseholdById(db, fmt.Sprintf("%+v",newMember.HouseholdID))
	if err != nil {
		// If household does not exists
		// TODO: Classify errors
		return fmt.Errorf("household does not exist")
	}
	// If a household exists, add member
	err = db.Create(newMember).Error
	if err != nil {
		return err
	}
	return nil
}



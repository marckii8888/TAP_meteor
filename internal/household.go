package internal

import (
	"gorm.io/gorm"
	"meteor/enum"
)

type Household struct {
	ID uint64 `json:"id"`
	HousingType enum.HouseholdType `json:"housing_type"`
	FamilyMembers []FamilyMember `json:"family_members,omitempty" gorm:"-"`
}

type HouseholdReq struct {
	Households []Household `json:"households"`
}

// Create
// @Summary Function to add a new household to the database
func Create(db *gorm.DB, housingType enum.HouseholdType) error {
	newHousehold := &Household{
		HousingType: housingType,
	}
	err := db.Create(newHousehold).Error
	if err != nil{
		return err
	}
	return nil
}

// QueryHouseholds
// @Summary Get all households from the database
func QueryHouseholds(db *gorm.DB, households *[]Household) error {
	err := db.Find(households).Error
	if err != nil {
		return err
	}
	return nil
}

// QueryUniqueHousehold
// @Summary Retrieve a unique household given the household id
func QueryUniqueHousehold(db *gorm.DB, household *Household, id string) error{
	err := db.Where("id = ?", id).First(household).Error
	if err != nil{
		return err
	}
	return nil
}

// QueryHouseholdById
// @Summary Checks if a household exists in the database
func QueryHouseholdById(db *gorm.DB, id string) error {
	var ret *Household
	err := db.Where("id = ?", id).First(&ret).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteHousehold
// @Summary Delete a household from the database given the household id
func DeleteHousehold(db *gorm.DB, household *Household, id string) error {
	err := db.Where("id = ?", id).Delete(household).Error
	if err != nil {
		return err
	}
	return nil
}

// IsReqValid
// @Summary Checks if the housing_type field in HouseholdReq struct is valid
func IsReqValid(req HouseholdReq) bool {
	// Check if household type is valid enum
	for _, household := range req.Households{
		if !enum.HouseholdTypeMap[household.HousingType]{
			// If enum type does not exists
			return false
		}
	}
	return true
}


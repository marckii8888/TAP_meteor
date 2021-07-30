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

func Create(db *gorm.DB, housingType enum.HouseholdType) error {
	// Create Household here
	newHousehold := &Household{
		HousingType: housingType,
	}
	err := db.Create(newHousehold).Error
	if err != nil{
		return err
	}
	return nil
}

func QueryHouseholds(db *gorm.DB, households *[]Household) error {
	err := db.Find(households).Error
	if err != nil {
		return err
	}
	return nil
}

func QueryUniqueHousehold(db *gorm.DB, household *Household, id string) error{
	err := db.Where("id = ?", id).First(household).Error
	if err != nil{
		return err
	}
	return nil
}

func QueryHouseholdById(db *gorm.DB, id string) error {
	var ret *Household
	err := db.Where("id = ?", id).First(&ret).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteHousehold(db *gorm.DB, household *Household, id string) error {
	err := db.Where("id = ?", id).Delete(household).Error
	if err != nil {
		return err
	}
	return nil
}


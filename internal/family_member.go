package internal

import (
	"fmt"
	"gorm.io/gorm"
	"meteor/enum"
	"strconv"
	"strings"
	"time"
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
// TODO: Check DOB is correct
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

func QueryFamilyMembers(db *gorm.DB, members *[]FamilyMember, householdId uint64) error {
	ids := []uint64{householdId}
	err := db.Where("household_id = ?", ids).Find(members).Error
	//err := db.First(&members, householdId).Error
	if err != nil {
		return err
	}
	return nil
}

func CalculateAge(dob string) uint64{
	dobSlice := strings.Split(dob, "-")
	birthYear := dobSlice[2]
	birthYearInt, _ := strconv.Atoi(birthYear)
	year, _, _ := time.Now().Date()
	return uint64(year - birthYearInt)
}



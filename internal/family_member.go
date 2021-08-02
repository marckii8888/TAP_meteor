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

func isDOBValid(dob string) bool {
	_, err := time.Parse("02-01-2006", dob)
	if err != nil { return false }
	return true
}

func isMaritalStatusValid(member *FamilyMember) bool {
	if member.MaritalStatus == enum.Married {
		if member.Spouse == "" { return false }
		return true
	}
	// If not married, should not have spouse
	if member.Spouse != "" { return false }

	return true
}

func AddFamilyMember(db *gorm.DB, newMember *FamilyMember) error {
	// Check if DOB is formatted correctly
	if !isDOBValid(newMember.DOB) { return fmt.Errorf("Invalid date of birth format. Should be DD-MM-YYYY")}

	// Check if martial status does not contradicts
	if !isMaritalStatusValid(newMember) { return fmt.Errorf("Contradiction in marital status.")}

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

func DeleteFamilyMember(db *gorm.DB, member *FamilyMember, id string) error {
	err := db.Where("id = ?", id).Delete(member).Error
	if err != nil {
		return err
	}
	return nil
}


func DeleteFamilyMemberFromHousehold(db *gorm.DB, member *FamilyMember, id string) error {
	err := db.Where("household_id = ?", id).Delete(member).Error
	if err != nil {
		return err
	}
	return nil
}

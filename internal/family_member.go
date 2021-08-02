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

// isDOBValid
// @Summary Check if the given string follows the format
// DD-MM-YYYY
func isDOBValid(dob string) bool {
	_, err := time.Parse("02-01-2006", dob)
	if err != nil { return false }
	return true
}

// isMaritalStatusValid
// @Summary Checks if there are any contradictions in the marital status
func isMaritalStatusValid(member *FamilyMember) bool {
	// Checks if a member is married and have a spouse
	if member.MaritalStatus == enum.Married {
		if member.Spouse == "" { return false }
		return true
	}
	// Checks if a member is not married and have no spouse
	if member.Spouse != "" { return false }

	return true
}

// AddFamilyMember
// @Summary Add a new member to the database
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

// QueryFamilyMembers
// @Summary Query the database for all family members in a given household id
func QueryFamilyMembers(db *gorm.DB, members *[]FamilyMember, householdId uint64) error {
	ids := []uint64{householdId}
	err := db.Where("household_id = ?", ids).Find(members).Error
	if err != nil {
		return err
	}
	return nil
}

// CalculateAge
// @Summary Calculate the age of a member
func CalculateAge(dob string) uint64{
	dobSlice := strings.Split(dob, "-")
	birthYear := dobSlice[2]
	birthYearInt, _ := strconv.Atoi(birthYear)
	year, _, _ := time.Now().Date()
	return uint64(year - birthYearInt)
}

// DeleteFamilyMember
// @Summary Delete a family member from the database given the member id
func DeleteFamilyMember(db *gorm.DB, member *FamilyMember, id string) error {
	err := db.Where("id = ?", id).Delete(member).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteFamilyMemberFromHousehold
// @Summary Remove the family member from the household
func DeleteFamilyMemberFromHousehold(db *gorm.DB, member *FamilyMember, id string) error {
	err := db.Where("household_id = ?", id).Delete(member).Error
	if err != nil {
		return err
	}
	return nil
}

// IsFamilyMemberValid
// @Summary Check if the fields in the POST requests are valid
func IsFamilyMemberValid(req *FamilyMember) bool {
	// Check if the Gender field is either MALE or FEMALE
	if !enum.FamilyMemberGenderMap[req.Gender] { return false }
	// Check if the MaritalStatus field is either MARRIED or SINGLE
	if !enum.FamilyMemberMaritalStatusMap[req.MaritalStatus] { return false }
	// Check if the OccupationType field is either EMPLOYED, UNEMPLOYED or STUDENT
	if !enum.FamilyMemberOccupationTypeMap[req.OccupationType] { return false }
	return true
}
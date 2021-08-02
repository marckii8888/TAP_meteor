package internal

import "github.com/marckii8888/TAP_meteor/enum"

type TotalGrantResp struct {
	StudentEncouragementBonus []GrantResp `json:"student_encouragement_bonus"`
	FamilyTogethernessScheme []GrantResp `json:"family_togetherness_scheme"`
	ElderBonus []GrantResp `json:"elder_bonus"`
	BabySunshineGrant []GrantResp `json:"baby_sunshine_grant"`
	YoloGstGrant 	[]GrantResp `json:"yolo_gst_grant"`
}

type GrantResp struct {
	HouseHold Household `json:"house_hold"`
	EligibleFamilyMembers []FamilyMember `json:"eligible_family_members"`
}

// IsEligible
// @Summary Check which grant a given household is eligible for
func IsEligible(resp *TotalGrantResp, household Household) {
	ok, res := isEligibleForStudentEncouragement(household)
	// If household / members eligible for student encouragement bonus
	if ok {
		resp.StudentEncouragementBonus = append(resp.StudentEncouragementBonus, res)
	}

	ok, res = isEligibleForFamilyScheme(household)
	// If household / members eligible for family togetherness scheme
	if ok {
		resp.FamilyTogethernessScheme = append(resp.FamilyTogethernessScheme, res)
	}

	ok, res = isEligibleForElderBonus(household)
	// If household / members eligible for Elderly Bonus
	if ok {
		resp.ElderBonus = append(resp.ElderBonus, res)
	}

	ok, res = isEligibleForBabySunshine(household)
	// If household / members eligible for Baby Sunshine Scheme
	if ok {
		resp.BabySunshineGrant = append(resp.BabySunshineGrant, res)
	}

	ok, res = isEligibleForYolo(household)
	// If household / members eligible for family togetherness scheme
	if ok {
		resp.YoloGstGrant = append(resp.YoloGstGrant, res)
	}
}

// isEligibleForStudentEncouragement
// @Summary Checks if a household is eligible for Student Encouragement Bonus
// Criteria:
// 1. Household must have children < 16 y.o
// 2. Household income < 150 000
func isEligibleForStudentEncouragement(household Household) (bool, GrantResp){
	var resp GrantResp
	totalIncome := 0.0
	eligible := false

	for _, member := range household.FamilyMembers{
		age := CalculateAge(member.DOB)
		// Check if children < 16 years old
		if age < 16 {
			eligible = true
			resp.EligibleFamilyMembers = append(resp.EligibleFamilyMembers, member)
		}
		totalIncome += member.AnnualIncome
	}

	// Check if total income < 150 000
	if totalIncome >= 150000 {
		return false, GrantResp{}
	}
	resp.HouseHold = household

	return eligible, resp
}

// isEligibleForFamilyScheme
// @Summary Checks if a household is eligible for Family Togetherness Scheme
// Criteria:
// 1. Household must have husband and wife
// 2. Household must have child(ren) > 18 y.o
func isEligibleForFamilyScheme(household Household) (bool, GrantResp){
	haveSpouse := false
	haveChild := false
	spouseMap := make(map[string]FamilyMember)
	var resp GrantResp
	for _, member := range household.FamilyMembers{
		// Check if spouse in map
		if _, ok := spouseMap[member.Name]; ok {
			// Spouse exists
			haveSpouse = true
			// Add both husband and wife to the list of eligible users
			resp.EligibleFamilyMembers = append(resp.EligibleFamilyMembers, member)
			resp.EligibleFamilyMembers = append(resp.EligibleFamilyMembers, spouseMap[member.Name])
		}

		// Check if family member is MARRIED and has a spouse
		if member.MaritalStatus == enum.Married && member.Spouse != "" {
			// Add to map
			spouseMap[member.Spouse] = member
		}

		// Check for child age
		age := CalculateAge(member.DOB)
		if age < 18 {
			haveChild = true
			resp.EligibleFamilyMembers = append(resp.EligibleFamilyMembers, member)
		}
	}

	if haveSpouse && haveChild {
		resp.HouseHold = household
		return true, resp
	}

	return false, GrantResp{}
}

// isEligibleForElderBonus
// @Summary Checks if a household is eligible for Elder Bonus
// Criteria:
// 1. Household must have member > 50 y.o
func isEligibleForElderBonus(household Household) (bool, GrantResp){
	var resp GrantResp
	eligible := false
	for _, member := range household.FamilyMembers{
		age := CalculateAge(member.DOB)
		// Check if member age is > 50
		if age > 50 {
			eligible = true
			resp.EligibleFamilyMembers = append(resp.EligibleFamilyMembers, member)
		}
	}
	if eligible {
		resp.HouseHold = household
		return true, resp
	}
	return false, GrantResp{}
}

// isEligibleForBabySunshine
// @Summary Checks if a household is eligible for Baby Sunshine Grant
// Criteria
// 1. Have children < 5 y.o
func isEligibleForBabySunshine(household Household) (bool, GrantResp){
	// Children younger than 5
	var resp GrantResp
	eligible := false

	for _, member := range household.FamilyMembers{
		age := CalculateAge(member.DOB)
		if age < 5 {
			eligible = true
			resp.EligibleFamilyMembers = append(resp.EligibleFamilyMembers, member)
		}
	}
	if eligible {
		resp.HouseHold = household
		return true, resp
	}
	return false, GrantResp{}
	return true, GrantResp{}
}

// isEligibleForYolo
// @Summary Checks if household is eligible for YOLO GST Grant
// Criteria:
// 1. Household with annual income < 100 000
func isEligibleForYolo(household Household) (bool, GrantResp){
	// Annual income less than $100 000
	var resp GrantResp
	eligible := false
	totalIncome := 0.0
	for _, member := range household.FamilyMembers{
		totalIncome += member.AnnualIncome
	}
	// Check if total household income < 100,000
	if totalIncome < 100000.0{
		eligible = true
		resp.HouseHold = household
	}
	return eligible, resp
}

// CheckTotalIncome
// @Summary check if household total income equals to query parameters.
// This is to filter households with annual income != query income
func CheckTotalIncome(household Household, query float64) bool {
	totalIncome := 0.0
	for _, member := range household.FamilyMembers {
		totalIncome += member.AnnualIncome
	}
	if totalIncome != query {
		return false
	}
	return true
}
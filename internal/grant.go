package internal

import "meteor/enum"

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

func isEligibleForStudentEncouragement(household Household) (bool, GrantResp){
	// List households and family members for student encouragement bonuses
	// Household must have children < 16 y.o
	// House hold income < 150 000
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

func isEligibleForFamilyScheme(household Household) (bool, GrantResp){
	// Households with husband & wife
	// Has child(ren) younger than 18 years old.
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
		return true, resp
	}

	return false, GrantResp{}
}

func isEligibleForElderBonus(household Household) (bool, GrantResp){
	// HDB members over 50 y.o
	var resp GrantResp
	eligible := false
	for _, member := range household.FamilyMembers{
		age := CalculateAge(member.DOB)
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

func isEligibleForYolo(household Household) (bool, GrantResp){
	// Annual income less than $100 000
	var resp GrantResp
	eligible := false
	totalIncome := 0.0
	for _, member := range household.FamilyMembers{
		totalIncome += member.AnnualIncome
	}
	if totalIncome < 100000.0{
		eligible = true
		resp.HouseHold = household
	}
	return eligible, resp
}

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
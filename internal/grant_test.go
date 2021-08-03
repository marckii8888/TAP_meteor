package internal_test

import (
	"github.com/marckii8888/TAP_meteor/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEligible(t *testing.T) {
	var resp internal.TotalGrantResp

	testHousehold := internal.Household{
		ID: 1,
		HousingType: "HDB",
		FamilyMembers: []internal.FamilyMember{
			{
				HouseholdID: 1,
				Name: "Mark Tan",
				Gender: "MALE",
				MaritalStatus: "SINGLE",
				Spouse: "",
				OccupationType: "STUDENT",
				AnnualIncome: 0.0,
				DOB: "24-12-2011",
			},
			{
				HouseholdID: 1,
				Name: "Mavis Tan",
				Gender: "FEMALE",
				MaritalStatus: "SINGLE",
				Spouse: "",
				OccupationType: "STUDENT",
				AnnualIncome: 0.0,
				DOB: "24-12-2017",
			},
			{
				HouseholdID: 1,
				Name: "Tan Jun Han",
				Gender: "MALE",
				MaritalStatus: "MARRIED",
				Spouse: "Rebecca Chew",
				OccupationType: "EMPLOYED",
				AnnualIncome: 75000.0,
				DOB: "24-12-1963",
			},
			{
				HouseholdID: 1,
				Name: "Rebecca Chew",
				Gender: "FEMALE",
				MaritalStatus: "MARRIED",
				Spouse: "Tan Jun Han",
				OccupationType: "EMPLOYED",
				AnnualIncome: 82000.0,
				DOB: "24-12-1964",
			},
		},
	}

	internal.IsEligible(&resp, testHousehold)
	assert.Nil(t, resp.StudentEncouragementBonus)
	assert.EqualValues(t, len(testHousehold.FamilyMembers), len(resp.FamilyTogethernessScheme[0].EligibleFamilyMembers))
	assert.EqualValues(t, 2, len(resp.ElderBonus[0].EligibleFamilyMembers)) // Tan Jun Han and Rebecca Chew
	assert.EqualValues(t, 1, len(resp.BabySunshineGrant[0].EligibleFamilyMembers)) // Mavis Tan
	assert.Nil(t, resp.YoloGstGrant)
}

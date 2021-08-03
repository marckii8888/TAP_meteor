package internal_test

import (
	"github.com/marckii8888/TAP_meteor/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHouseholdCRUD(t *testing.T){
	db, err := InitTestDB()
	assert.NoError(t, err, "Failed to init testdb")

	db.AutoMigrate(&internal.Household{})

	// Test Create household
	assert.NoError(t, internal.Create(db, "HDB"))

	// Test QueryHouseholds
	var resp []internal.Household
	assert.NoError(t, internal.QueryHouseholds(db, &resp), "Bad QueryHouseholds")
	assert.Equal(t, 1, len(resp))

	// Test QueryUniqueHousehold
	var resp2 internal.Household
	assert.NoError(t, internal.QueryUniqueHousehold(db, &resp2, "1"))

	// Test DeleteHousehold
	assert.NoError(t, internal.DeleteHousehold(db, &resp2, "1"))

	db.Delete(&internal.Household{})
}


func TestIsReqValid(t *testing.T) {
	assert.True(t, internal.IsReqValid(internal.HouseholdReq{[]internal.Household{{ID: 1, HousingType: "HDB"}}}))
	assert.True(t, internal.IsReqValid(internal.HouseholdReq{[]internal.Household{{ID: 1, HousingType: "CONDOMINIUM"}}}))
	assert.True(t, internal.IsReqValid(internal.HouseholdReq{[]internal.Household{{ID: 1, HousingType: "LANDED"}}}))
	assert.False(t, internal.IsReqValid(internal.HouseholdReq{[]internal.Household{{ID: 1, HousingType: "!@#"}}}))
}
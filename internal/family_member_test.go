package internal_test

import (
	"github.com/marckii8888/TAP_meteor/config"
	"github.com/marckii8888/TAP_meteor/internal"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"testing"
)

func InitTestDB() (*gorm.DB, error){
	os.Chdir("../")
	config.InitConf()
	dsn := config.Conf.Database.User +":"+ config.Conf.Database.Password +"@tcp"+ "(" + config.Conf.Database.Host + ":" + config.Conf.Database.Port +")/testdb?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestCRUDFamilyMembers(t *testing.T){
	db, err := InitTestDB()
	assert.NoError(t, err, "Failed to init testdb")
	db.AutoMigrate(&internal.FamilyMember{})
	db.AutoMigrate(&internal.Household{})

	testHousehold := &internal.Household{
		ID: 1,
		HousingType: "HDB",
	}
	db.Create(testHousehold)

	testMember1 := &internal.FamilyMember{
		HouseholdID: 1,
		Name: "Mark Tan",
		Gender: "MALE",
		MaritalStatus: "SINGLE",
		Spouse: "",
		OccupationType: "STUDENT",
		AnnualIncome: 0.0,
		DOB: "24-12-2011",
	}
	// Test Add new family member
	assert.NoError(t, internal.AddFamilyMember(db, testMember1), "Bad AddFamilyMember")

	// Test Query family member
	var req []internal.FamilyMember
	err = internal.QueryFamilyMembers(db, &req, 1)
	assert.NoError(t, internal.QueryFamilyMembers(db, &req, 1), "Bad queryFamilyMembers")

	// Test Delete Family Member
	assert.NoError(t, internal.DeleteFamilyMember(db, testMember1, "1"), "Bad DeleteFamilyMember")

	db.Delete(testHousehold)
	db.Delete(&internal.FamilyMember{})
	db.Delete(&internal.Household{})
}

func TestCalculateAge(t *testing.T) {
	assert.Equal(t, uint64(20), internal.CalculateAge("09-02-2001"))
}

func TestFamilyMemberValid(t *testing.T){
	// Test if isDOBValid
	assert.True(t, internal.IsFamilyMemberValid(&internal.FamilyMember{HouseholdID: 1, Name: "Test", Gender: "MALE", OccupationType: "STUDENT", MaritalStatus: "SINGLE", AnnualIncome: 0.0, DOB: "08-08-1997"}))
	assert.False(t, internal.IsFamilyMemberValid(&internal.FamilyMember{HouseholdID: 1, Name: "Test", Gender: "hello", OccupationType: "STUDENT", MaritalStatus: "SINGLE", AnnualIncome: 0.0, DOB: "08-08-1997"}))
	assert.False(t, internal.IsFamilyMemberValid(&internal.FamilyMember{HouseholdID: 1, Name: "Test", Gender: "MALE", OccupationType: "testing", MaritalStatus: "SINGLE", AnnualIncome: 0.0, DOB: "08-08-1997"}))
	assert.False(t, internal.IsFamilyMemberValid(&internal.FamilyMember{HouseholdID: 1, Name: "Test", Gender: "MALE", OccupationType: "STUDENT", MaritalStatus: "@!#!", AnnualIncome: 0.0, DOB: "08-08-1997"}))
}

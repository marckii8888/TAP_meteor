package internal

import (
	"gorm.io/gorm"
	"meteor/enum"
)

type Household struct {
	ID uint64 `json:"id"`
	HousingType enum.HouseholdType `json:"housing_type"`
	FamilyMembers []*FamilyMember `json:"family_members,omitempty" gorm:"-"`
}

type CreateHouseholdReq struct {
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

func QueryHouseholds(db *gorm.DB) error {
	var ret []*Household
	err := db.Find(ret).Error
	if err != nil {
		return err
	}
	return nil
}

func QueryUniqueHousehold(MemberName string, HousingType enum.HouseholdType) *Household{
	//get user by id
	//func GetUser(db *gorm.DB, User *User, id string) (err error) {
	//	err = db.Where("id = ?", id).First(User).Error
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//}
	return &Household{}
}


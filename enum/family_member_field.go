package enum

type Gender string
type MaritalStatus string
type OccupationType string

const(
	Male Gender = "MALE"
	Female Gender = "FEMALE"
	Married MaritalStatus = "MARRIED"
	Single MaritalStatus = "SINGLE"
	Unemp OccupationType = "UNEMPLOYED"
	Student OccupationType = "STUDENT"
	Emp OccupationType = "EMPLOYED"
	)
var FamilyMemberGenderMap = map[Gender]bool{
	Male: true,
	Female: true,
}

var FamilyMemberMaritalStatusMap = map[MaritalStatus]bool{
	Single: true,
	Married: true,
}

var FamilyMemberOccupationTypeMap = map[OccupationType]bool{
	Unemp: true,
	Student: true,
	Emp: true,
}

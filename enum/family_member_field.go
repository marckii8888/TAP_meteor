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

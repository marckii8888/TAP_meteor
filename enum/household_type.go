package enum

type HouseholdType string

const(
	Landed	HouseholdType = "LANDED"
	Condo HouseholdType = "CONDOMINIUM"
	HDB HouseholdType = "HDB"
)
var HouseholdTypeMap = map[HouseholdType]bool{
	Landed: true,
	Condo: true,
	HDB: true,
}


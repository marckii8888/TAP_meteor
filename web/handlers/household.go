package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
	"meteor/internal"
	"net/http"
	"strconv"
)

type Helper struct{
	db *gorm.DB
}
func New() *Helper{
	db:= internal.InitDB()
	db.AutoMigrate(&internal.Household{})
	db.AutoMigrate(&internal.FamilyMember{})
	return &Helper{db : db}
}

func (helper *Helper)CreateHousehold(c *gin.Context){
	/*
	{
	    "households" : [
	        {
	            "housing_type" : "HDB"
	        }
	    ]
	}
	*/
	var req internal.HouseholdReq
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		c.JSON(404, gin.H{
			"message" : fmt.Sprintf("Error: %+v... req = %+v", err, req),
		})
		return
	}

	// Check if req is valid
	if !internal.IsReqValid(req) {
		c.JSON(404, gin.H{
			"error" : "Invalid household type",
		})
		return
	}
	for _, household := range req.Households{
		err := internal.Create(helper.db, household.HousingType)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error" : fmt.Sprintf("Error - %+v", err),
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"message" : "Household created",
	})
}

// AddFamilyMember
// url http://localhost:8081/household/add_family_member
func (helper *Helper)AddFamilyMember(c *gin.Context){
	/*
	{
	     "household_id" : 3,
	     "name" : "Bobby",
	     "gender" : "MALE",
	     "marital_status" : "SINGLE",
	     "spouse" : "Alice",
	     "occupation_type" : "STUDENT",
	     "annual_income" : 1000.0,
	     "dob" : "08-08-1997"
	}
	*/
	var req *internal.FamilyMember
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error" : fmt.Sprintf("Error - %+v", err),
		})
		return
	}

	// Check if req is valid
	if !internal.IsFamilyMemberValid(req) {
		c.JSON(404, gin.H{
			"error" : "Invalid fields",
		})
		return
	}

	err := internal.AddFamilyMember(helper.db, req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error" : fmt.Sprintf("Error - %+v", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Added member",
	})
}

// ListAllHouseholds
// @Summary List all the households in the database
func (helper *Helper) ListAllHouseholds(c *gin.Context){
	var ret []internal.Household
	err := internal.QueryHouseholds(helper.db, &ret)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error" : fmt.Sprintf("Error - %+v", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": ret,
	})
}

// QueryHousehold
// @Summary List details of a specific household
// url: GET http://localhost:8081/household/query_household?id=1
func (helper *Helper) QueryHousehold(c *gin.Context){
	id := c.Query("id")
	var ret internal.Household
	err := internal.QueryUniqueHousehold(helper.db, &ret, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If record does not exists
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error" : fmt.Sprintf("Error - %+v", err),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error" : fmt.Sprintf("Error - %+v", err),
		})
	}

	// Query family members base on household
	var familyMembers []internal.FamilyMember
	err = internal.QueryFamilyMembers(helper.db, &familyMembers, ret.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If record does not exists
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error" : fmt.Sprintf("Error - %+v", err),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error" : fmt.Sprintf("Error - %+v", err),
		})
	}

	ret.FamilyMembers = familyMembers

	c.JSON(200, gin.H{
		"message": ret,
	})
}

// QueryHouseholdsGrantEligibility
// @Summary
// url: http://localhost:8081/grants/list_eligible_households?household_size=2&total_income=1000
func (helper *Helper) QueryHouseholdsGrantEligibility(c *gin.Context){
	householdSize := c.Query("household_size")
	totalIncome := c.Query("total_income")

	// Convert householdSize and totalIncome to int and float64 respectively
	householdSizeInt, _ := strconv.Atoi(householdSize)
	totalIncomeFloat, _ :=  strconv.ParseFloat(totalIncome, 64)

	// First find all the household
	var households []internal.Household
	err := internal.QueryHouseholds(helper.db, &households)

	if err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error" : fmt.Sprintf("Error - %+v", err),
		})
		return
	}

	var resp internal.TotalGrantResp
	for _, household := range households{
		err = internal.QueryFamilyMembers(helper.db, &household.FamilyMembers, household.ID)
		// Filter based on family members
		if len(household.FamilyMembers) != householdSizeInt && householdSizeInt != 0{
			continue
		}

		// Filter base on household income
		if !internal.CheckTotalIncome(household, totalIncomeFloat) && totalIncomeFloat != 0.0 {
			continue
		}

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// If record does not exists
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error" : fmt.Sprintf("Error - %+v", err),
				})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error" : fmt.Sprintf("Error - %+v", err),
			})
		}
		// Check eligibility here
		internal.IsEligible(&resp, household)
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : resp,
	})
}

func (helper *Helper) DeleteHousehold(c *gin.Context){
	//id := c.Query("id")
	id, _ := c.Params.Get("house_id")
	fmt.Println(id)
	var ret internal.Household
	var member internal.FamilyMember

	// Delete all family members
	err := internal.DeleteFamilyMemberFromHousehold(helper.db, &member, fmt.Sprintf("%+v", id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error" : fmt.Sprintf("Error - %+v", err),
		})
		return
	}

	// Delete household
	err = internal.DeleteHousehold(helper.db, &ret, fmt.Sprintf("%+v", id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error" : fmt.Sprintf("Error - %+v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Household deleted",
	})
}

func (helper *Helper) DeleteMember(c *gin.Context){
	id, _ := c.Params.Get("member_id")
	var member internal.FamilyMember

	err := internal.DeleteFamilyMember(helper.db, &member, fmt.Sprintf("%+v", id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error" : fmt.Sprintf("Error - %+v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Family member deleted",
	})
}
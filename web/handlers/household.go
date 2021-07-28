package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
	"meteor/internal"
	"net/http"
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
	var req internal.CreateHouseholdReq
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		c.JSON(404, gin.H{
			"message" : fmt.Sprintf("Error: %+v", err),
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
		"message" : "Done",
	})
}

func (helper *Helper)AddFamilyMember(c *gin.Context){
	var req *internal.FamilyMember
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error" : fmt.Sprintf("Error - %+v", err),
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

func ListAllHouseholds(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Listing all household",
	})
}

func QueryHouseHold(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Querying a household",
	})
}

package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateHousehold(c *gin.Context){
	fmt.Println("Creating household")
}

func AddFamilyMember(c *gin.Context){
	fmt.Println("Adding member")
}

func ListAllHouseholds(c *gin.Context){
	fmt.Println("Listing all households")
}

func QueryHouseHold(c *gin.Context){
	fmt.Println("Querying a household")
}

package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"meteor/web/handlers"
)

type Router struct{
	*gin.Engine
}
func NewRouter() *Router{
	router := gin.Default()

	helper := handlers.New()
	householdAPI := router.Group("/household")
	// 1. Create Household
	householdAPI.POST("/create", helper.CreateHousehold)
	// 2. Add family member
	householdAPI.POST("/add_family_member", helper.AddFamilyMember)
	// 3. List house hold
	householdAPI.GET("/list_households", helper.ListAllHouseholds)
	// 4. Show selected house hold
	householdAPI.GET("/query_household", helper.QueryHousehold)

	// 6. [OPTIONAL] Delete Household
	householdAPI.DELETE("/delete_household", helper.DeleteHousehold)
	// 7. [OPTIONAL] Delete family member
	//householdAPI.DELETE("/delete_member", helper.DeleteFamilyMember)

	// 5. Search for households and recipients **
	grantAPI := router.Group("/grants")
	grantAPI.GET("/list_eligible_households", helper.QueryHouseholdsGrantEligibility)

	return &Router{
		router,
	}
}

func (r *Router) Run(){
	port := 8081
	err := r.Engine.Run(fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to start router")
	}
	log.Printf("Connected to port %+v", port)
}

func Run() {
	NewRouter().Run()
}
package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"meteor/config"
	"meteor/web/handlers"
)

type Router struct{
	*gin.Engine
}
func NewRouter() *Router{
	router := gin.Default()

	// Read config file
	config.InitConf()

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

	// 5. [OPTIONAL] Delete Household
	householdAPI.DELETE("/delete_household/:house_id", helper.DeleteHousehold)
	// 6. [OPTIONAL] Delete family member
	householdAPI.DELETE("/delete_member/:member_id", helper.DeleteMember)

	// 7. Search for households and recipients
	grantAPI := router.Group("/grants")
	grantAPI.GET("/list_eligible_households", helper.QueryHouseholdsGrantEligibility)

	return &Router{
		router,
	}
}

func (r *Router) Run(){
	err := r.Engine.Run(fmt.Sprintf(":%v", config.Conf.Server.Port))
	if err != nil {
		log.Fatalf("Failed to start router")
	}
	log.Printf("Connected to port %+v", config.Conf.Server.Port)
}

func Run() {
	NewRouter().Run()
}
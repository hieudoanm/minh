package vietnam_national_assembly_router

import (
	committees_controller "chatbot-functions/src/modules/vietnam/controller/national-assembly/committees"
	members_controller "chatbot-functions/src/modules/vietnam/controller/national-assembly/members"

	"github.com/julienschmidt/httprouter"
)

func VietnamNationalAssemblyRouter(router *httprouter.Router) {
	router.GET("/vietnam/national-assembly/committees", committees_controller.GetCommittees)
	router.GET("/vietnam/national-assembly/committees/:id", committees_controller.GetCommittee)
	router.GET("/vietnam/national-assembly/members", members_controller.GetMembers)
	router.GET("/vietnam/national-assembly/members/:id", members_controller.GetMember)

}

package vietnam_router

import (
	vietnam_government_router "chatbot-functions/src/modules/vietnam/router/government"
	vietnam_maps_router "chatbot-functions/src/modules/vietnam/router/maps"
	vietnam_national_assembly_router "chatbot-functions/src/modules/vietnam/router/national-assembly"
	vietnam_vnindex_router "chatbot-functions/src/modules/vietnam/router/vnindex"

	"github.com/julienschmidt/httprouter"
)

func VietnamRouter(router *httprouter.Router) {
	vietnam_government_router.VietnamGovernmentRouter(router)              // Government
	vietnam_maps_router.VietnamMapsRouter(router)                          // Maps
	vietnam_national_assembly_router.VietnamNationalAssemblyRouter(router) // National Assembly
	vietnam_vnindex_router.VietnamVnindexRouter(router)                    // VNINDEX

}

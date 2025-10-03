package serverctl
import "github.com/gravitl/netmaker/controllers"
func setACLRoutes(networks *mux.Router) {
	acl := networks.PathPrefix("/{netid}/acl").Subrouter()
	acl.HandleFunc("", controllers.GetNetworkACLs).Methods("GET")
	acl.HandleFunc("", controllers.PutNetworkACL).Methods("PUT","DELETE")
}
package controllers
import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gravitl/netmaker/logger"
	"github.com/gravitl/netmaker/logic"
	"github.com/gravitl/netmaker/models"
)
// GetNetworkACLs – lista usuários autorizados numa rede
func GetNetworkACLs(w http.ResponseWriter, r *http.Request) {
	netID := mux.Vars(r)["netid"]
	var acls []models.UserNetworkACL
	database.DB.Where("net_id = ?", netID).Find(&acls)
	logic.ReturnSuccessResponse(w, r, acls)
}
// PutNetworkACL – cria/remove ACL
func PutNetworkACL(w http.ResponseWriter, r *http.Request) {
	netID := mux.Vars(r)["netid"]
	var payload struct{ UserName string `json:"user_name"` }
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		logic.ReturnErrorResponse(w, r, logic.FormatError(err, "badrequest"))
		return
	}
	switch r.Method {
	case http.MethodPut:
		database.DB.Save(&models.UserNetworkACL{UserName: payload.UserName, NetID: netID})
	case http.MethodDelete:
		database.DB.Where("user_name = ? AND net_id = ?", payload.UserName, netID).Delete(&models.UserNetworkACL{})
	}
	logic.ReturnSuccessResponse(w, r, nil)
}
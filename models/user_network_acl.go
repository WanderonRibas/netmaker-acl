package models
type UserNetworkACL struct {
	UserName string `json:"user_name" gorm:"primaryKey"`
	NetID    string `json:"net_id"    gorm:"primaryKey"`
}
func (UserNetworkACL) TableName() string { return "user_network_acls" }
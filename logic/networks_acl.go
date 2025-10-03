package logic
import "gorm.io/gorm"
// NetworksAllowedToUser devolve []string com os netids que o usuário pode ver
func NetworksAllowedToUser(userName string) ([]string, error) {
	var ids []string
	err := database.DB.Model(&models.UserNetworkACL{}).
		Where("user_name = ?", userName).
		Pluck("net_id", &ids).Error
	return ids, err
}
// UserCanAccessNetwork retorna true se o usuário tem acesso explícito OU for admin
func UserCanAccessNetwork(userName, netID string) (bool, error) {
	var isAdmin bool
	if err := database.DB.Model(&models.User{}).
		Where("user_name = ? AND is_admin = 1", userName).
		Select("1").Scan(&isAdmin).Error; err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if isAdmin { return true, nil }
	var count int64
	database.DB.Model(&models.UserNetworkACL{}).
		Where("user_name = ? AND net_id = ?", userName, netID).
		Count(&count)
	return count > 0, nil
}
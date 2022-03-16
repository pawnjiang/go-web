package dao

func IsPhoneExist(phone string) bool {
	var user model.user
	db.DB.Where("phone=?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

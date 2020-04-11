package models

type User struct {
	Id        Pk
	UserName  string
	RealName  string
	Password  string
	Telephone string
}

func FindCountByUserNameAndPassword(userName string, password string) int64 {

	rows, err := o.QueryTable(new(User)).Filter("user_name", userName).Filter("password", password).Count()
	if err != nil {
		log.Error("error :", err)
		return -1
	}
	return rows
}

func FindUserByUserName(username string) (User, bool) {

	var user User
	err := o.QueryTable(new(User)).Filter("user_name", username).One(&user)
	if err != nil {
		log.Error("error", err)
		return user, false
	}
	return user, true
}

func InsetUser(user *User) int64 {
	id, _ := o.Insert(user)
	return id
}

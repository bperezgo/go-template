package entities

type User struct {
	id       UserId
	email    UserEmail
	password UserPassword
}

type UserDto struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type UserId struct {
	id string
}

type UserPassword struct {
	password string
}

type UserEmail struct {
	email string
}

func NewUser(id string, email string, password string) User {
	userId := NewUserId(id)
	userEmail := NewUserEmail(email)
	userPassword := NewUserPassword(password)

	return User{
		id:       userId,
		email:    userEmail,
		password: userPassword,
	}
}

func (u User) ToDto() UserDto {
	return UserDto{
		Id:    u.id.String(),
		Email: u.email.String(),
	}
}

// VALUE OBJECTS
func NewUserId(id string) UserId {
	userId := id
	return UserId{
		id: userId,
	}
}

func (u UserId) String() string {
	return u.id
}

func NewUserPassword(password string) UserPassword {
	return UserPassword{
		password: password,
	}
}

func NewUserEmail(email string) UserEmail {
	return UserEmail{
		email: email,
	}
}

func (e UserEmail) String() string {
	return e.email
}

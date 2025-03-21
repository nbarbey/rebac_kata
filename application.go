package rebac_kata

type Application struct {
	ads   map[int]*Ad
	users map[string]*User
}

func NewApplication(users map[string]*User) *Application {
	a := &Application{ads: make(map[int]*Ad), users: users}
	for _, user := range a.users {
		user.a = a
	}
	return a
}

func (a *Application) Login(userName string) *User {
	return a.users[userName]
}

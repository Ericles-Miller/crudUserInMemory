package application

import "github.com/google/uuid"

type Application struct {
	data map[uuid.UUID]User
}

func New() *Application {
	return &Application{
		data: make(map[uuid.UUID]User),
	}
}

func (a *Application) FindAll() []User {
	users := make([]User, 0, len(a.data))
	for _, user := range a.data {
		users = append(users, user)
	}
	return users
}

func (a *Application) FindById(id uuid.UUID) (User, bool) {
	user, ok := a.data[id]
	return user, ok
}

func (a *Application) Insert(userBody UserBody) (User, error) {
	if err := ValidateUser(userBody); err != nil {
		return User{}, err
	}

	user := User{
		Id:        uuid.New(),
		FirstName: userBody.FirstName,
		LastName:  userBody.LastName,
		Biography: userBody.Biography,
	}

	a.data[user.Id] = user
	return user, nil
}

func (a *Application) Update(id uuid.UUID, userBody UserBody) (User, bool, error) {
	if _, ok := a.data[id]; !ok {
		return User{}, false, nil
	}

	if err := ValidateUser(userBody); err != nil {
		return User{}, true, err
	}

	user := User{
		Id:        id,
		FirstName: userBody.FirstName,
		LastName:  userBody.LastName,
		Biography: userBody.Biography,
	}

	a.data[id] = user
	return user, true, nil
}

func (a *Application) Delete(id uuid.UUID) (User, bool) {
	user, ok := a.data[id]
	if !ok {
		return User{}, false
	}
	delete(a.data, id)
	return user, true
}

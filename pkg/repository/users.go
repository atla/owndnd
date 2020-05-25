package repository

import (
	"github.com/atla/owndnd/pkg/db"
	e "github.com/atla/owndnd/pkg/entities"
)

//UsersRepository ...
type UsersRepository interface {
	FindByRefID(id string) (*e.User, error)
	Create(user *e.User) (*e.User, error)
	Update(id string, user *e.User) error
	Delete(id string) error
}

//--- Implementations

type usersRepo struct {
	*GenericRepo
}

//NewMongoDBUsersRepository creates a new mongodb partiesrep
func NewMongoDBUsersRepository(db *db.Client) UsersRepository {
	return &usersRepo{
		GenericRepo: &GenericRepo{
			db:         db,
			collection: "users",
			generator: func() interface{} {
				return &e.User{}
			},
		},
	}
}

func (pr *usersRepo) Create(user *e.User) (*e.User, error) {
	user.Entity = e.NewEntity()
	result, err := pr.GenericRepo.Store(user)
	return result.(*e.User), err
}

func (pr *usersRepo) Update(refID string, user *e.User) error {
	return pr.GenericRepo.UpdateByField(user, "refid", refID)
}

func (pr *usersRepo) FindByRefID(refID string) (*e.User, error) {
	result, err := pr.GenericRepo.FindByField("refid", refID)

	if user, ok := result.(*e.User); ok {
		return user, nil
	}
	return nil, err //ors.New("result is not a User")
}

func (pr *usersRepo) Delete(id string) error {
	return pr.GenericRepo.Delete(id)
}

package impl

import (
	"authapp/domain/user"
	"authapp/dto/request"
	"errors"
)

type service struct {
	repo user.Repository
}

func NewService(repo user.Repository) user.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) ConvertToRole(roleStr string) (user.Role, error) {
	roleMap := user.GetAllRoles()

	role := user.Role(roleStr)
	if _, ok := roleMap[role]; !ok {
		return "", errors.New("role does not exist")
	}

	return role, nil
}

func (s *service) CreateUserIfNotAny(req request.CreateUserRequest) (*user.User, error) {
	role, err := s.ConvertToRole(req.Role)
	if err != nil {
		return nil, err
	}

	u := &user.User{
		Phonenumber: req.Phonenumber,
		Name:        req.Name,
		// TODO Generate Password
		Password: "",
		Role:     role,
	}

	existingUser, err := s.repo.GetUserByPhonenumber(req.Phonenumber)
	if existingUser != nil {
		return existingUser, nil
	}
	if err != nil {
		return nil, err
	}

	u, err = s.repo.Persist(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *service) Login(phonenumber, password string) (*user.User, string, error) {

	u, err := s.repo.GetUserByUserPass(phonenumber, password)
	if u == nil {
		return nil, "", nil
	}
	if err != nil {
		return nil, "", err
	}

	// TODO Generate Token
	token := ""

	return u, token, nil
}

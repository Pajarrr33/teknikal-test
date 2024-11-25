package usecase

import (
	"fmt"
	"strings"
	"teknikal-test/entity"
	"teknikal-test/entity/request"
	"teknikal-test/entity/response"
	"teknikal-test/repository"
	"teknikal-test/service"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Login(request request.LoginRequest) (response.LoginResponse, error)
	Register(request request.RegisterRequest) (entity.Customer, error)
	Logout(token string,id string) (error)
}

type authUsecase struct {
	CustomerRepository repository.CustomerRepository
	ExpiredRepository  repository.ExpiredRepository
	JwtService         service.JWTService
}

// Login implements AuthUsecase.
func (a *authUsecase) Login(request request.LoginRequest) (response.LoginResponse, error) {
	customer, err := a.CustomerRepository.FindByEmail(request.Email)
	if err != nil {
		err = fmt.Errorf("customer not found")
		return response.LoginResponse{}, err
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(request.Password))
	if err != nil {
		fmt.Println(err)
		err = fmt.Errorf("wrong password")
		return response.LoginResponse{}, err
	}

	loginResponse , err := a.JwtService.GenerateToken(customer)
	if err != nil {
		err = fmt.Errorf("failed to generate token")
		return response.LoginResponse{}, err
	}

	service.AddLog(logrus.Fields{"customer" : customer.Name},"info","user has been logged in")
	service.SaveLog()
	return loginResponse, nil
}

func (a *authUsecase) Register(request request.RegisterRequest) (entity.Customer, error) {
	if request.Email == "" || request.Password == "" || request.Name == "" {
		err := fmt.Errorf("missing required fields")
		return entity.Customer{}, err
	}

	if (len(request.Password) < 8) {
		err := fmt.Errorf("password must be at least 8 characters")
		return entity.Customer{}, err
	}

	if !strings.Contains(request.Email, "@") {
		err := fmt.Errorf("invalid email format")
		return entity.Customer{}, err
	}

	isExist,_ := a.CustomerRepository.FindByEmail(request.Email)
	if isExist.Id != "" {
		err := fmt.Errorf("email already exist")
		return entity.Customer{}, err
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		err = fmt.Errorf("failed to hash password")
		return entity.Customer{}, err
	}

	request.Password = string(hashedPassword)

	customer , err := a.CustomerRepository.Create(request)
	if err != nil {
		err = fmt.Errorf("failed to create customer")
		return entity.Customer{}, err
	}

	service.AddLog(logrus.Fields{"customer" : customer.Name},"info","user has been registered")
	service.SaveLog()
	return customer, nil
}

func (a *authUsecase) Logout(token string, id string) error {
    // Check if the token is already expired
    isExist, err := a.ExpiredRepository.GetExpiredByToken(token)
    if err != nil {
        return fmt.Errorf("failed to get expired token: %w", err)
    }
    if isExist {
        return fmt.Errorf("token already exists")
    }

    // Insert the token into the expired_token table
    if err := a.ExpiredRepository.Insert(token); err != nil {
        return fmt.Errorf("failed to insert expired token: %w", err)
    }

    // Find the user by ID
    user, err := a.CustomerRepository.FindById(id)
    if err != nil {
        return fmt.Errorf("failed to find user: %w", err)
    }

    // Log the logout action
    a.logUserLogout(user.Name)
    return nil
}

// Helper function for logging
func (a *authUsecase) logUserLogout(userName string) {
    service.AddLog(logrus.Fields{"customer": userName}, "info", "user has been logged out")
    service.SaveLog()
}


func NewAuthUsecase(customerRepository repository.CustomerRepository, jwtService service.JWTService, expiredRepository repository.ExpiredRepository) AuthUsecase {
	return &authUsecase{customerRepository, expiredRepository, jwtService}
}

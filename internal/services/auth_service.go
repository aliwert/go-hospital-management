package services

import (
	"errors"
	"os"
	"time"

	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/repositories"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(req *models.RegisterRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
		Status:   true,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(req *models.LoginRequest) (*TokenDetails, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Update last login
	now := time.Now()
	user.LastLogin = &now
	s.userRepo.Update(user)

	// Generate tokens
	tokens, err := s.CreateTokens(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (s *AuthService) GetUserById(id uint) (*models.User, error) {
	return s.userRepo.FindById(id)
}

func (s *AuthService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *AuthService) CreateTokens(userID uint, role string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()   // Access token expires in 15 minutes
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix() // Refresh token expires in 7 days

	// Generate Access Token
	atClaims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     td.AtExpires,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	var err error
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	// Generate Refresh Token
	rtClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     td.RtExpires,
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func (s *AuthService) RefreshToken(refreshToken string) (*TokenDetails, error) {
	// Verify refresh token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil
	})

	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	// Get user from database
	userID := uint(claims["user_id"].(float64))
	user, err := s.userRepo.FindById(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Generate new tokens
	return s.CreateTokens(user.ID, user.Role)
}

func (s *AuthService) UpdateUser(id uint, req *models.UpdateRequest) (*models.User, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}
	if req.Role != "" {
		user.Role = req.Role
	}
	if req.Status != nil {
		user.Status = *req.Status
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

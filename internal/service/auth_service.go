package service

import (
	"github.com/ryo-y222/delivery-api/internal/model"
	"github.com/ryo-y222/delivery-api/internal/repository"
	"github.com/ryo-y222/delivery-api/internal/util"
)

type AuthService struct {
	repo      repository.IUserRepository
	jwtSecret string
}

func NewAuthService(repo repository.IUserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Login(email string, password string) (*model.User, string, error) {
	// サービス層の処理
	// 1．メールアドレスでユーザーを検索
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, "", err //DBエラー、repositoryに任せる。
	}
	if user == nil {
		return nil, "", ErrInvalidCredentials
	}

	//２．パスワードを検証
	if err = util.CheckPassword(password, user.PasswordHash); err != nil {
		return nil, "", ErrInvalidCredentials
	}

	//3.JWT生成
	token, err := util.GenerateToken(user.ID, user.Email, user.Role, s.jwtSecret)

	if err != nil {
		return nil, "", err
	}

	return user, token, nil

}

func (s *AuthService) Register(email, password, name, role, company, phone string) (*model.User, string, error) {

	//サービス層の処理
	// 1.メールアドレス重複チェック
	existingUser, err := s.repo.GetByEmail(email)

	if err != nil {
		return nil, "", err
	}
	if existingUser != nil {
		return nil, "", ErrEmailAlreadyExists
	}

	//2.パスワードをハッシュ化
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, "", err
	}

	//３.ユーザー構造体を作成
	user := &model.User{
		Email:        email,
		PasswordHash: hashedPassword,
		Name:         name,
		Role:         role,
		Company:      company,
		Phone:        phone,
	}

	//4.ユーザー情報を登録。
	if err := s.repo.Create(user); err != nil {
		return nil, "", err
	}

	//5.JWT生成
	token, err := util.GenerateToken(user.ID, user.Email, user.Role, s.jwtSecret)

	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}

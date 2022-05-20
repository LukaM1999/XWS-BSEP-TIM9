package domain

import (
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"unicode"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username" validate:"username"`
	Password string             `bson:"password" validate:"password"`
	Role     string             `bson:"role" validate:"required"`
}

func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func usernameValidator(fl validator.FieldLevel) bool {
	matchString, err := regexp.MatchString(`^[_a-zA-Z0-9]([._-]([._-]?)|[a-zA-Z0-9]){3,18}[_a-zA-Z0-9]$`, fl.Field().String())
	if err != nil {
		return false
	}
	return matchString
}

func passwordValidator(fl validator.FieldLevel) bool {
	return isValid(fl.Field().String())
}

func isValid(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) > 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func NewUserValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("username", usernameValidator)
	validate.RegisterValidation("password", passwordValidator)
	return validate
}

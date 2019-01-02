package common

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v8"
	"time"
)

// replace this with your own key, you can gen it with './marmalade gen_secret' command
const NBSecretPassword = "awE_PBAzcUQ2SCtzIveOQ-o-I70HHI6o7iWu1jcMOzs="


func GenToken(id int64) string {
	jwtToken := jwt.New(jwt.GetSigningMethod("HS256"))
	// Set some claims
	jwtToken.Claims = jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	// Sign and get the complete encoded token as a string
	token, _ := jwtToken.SignedString([]byte(NBSecretPassword))
	return token
}

func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

// To handle the error returned by c.Bind in gin framework
// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		// can translate each error one at a time.
		//fmt.Println("gg",v.NameNamespace)
		if v.Param != "" {
			res.Errors[v.Field] = fmt.Sprintf("{%v: %v}", v.Tag, v.Param)
		} else {
			res.Errors[v.Field] = fmt.Sprintf("{key: %v}", v.Tag)
		}

	}
	return res
}

// Warp the error info in a object
func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

func SetPassword(password string) string {
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return string(passwordHash)
}

func CheckPassword(password, storedPassword string) bool {
	bytePassword := []byte(password)
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), bytePassword)
	if err != nil {
		return false
	}
	return true
}

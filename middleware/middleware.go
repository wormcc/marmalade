package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	_accountRepo "github.com/wormcc/marmalade/account/repository"
	"github.com/wormcc/marmalade/common"
	"github.com/wormcc/marmalade/models"
	"net/http"
	"strings"
)

// Strips 'TOKEN ' prefix from token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
		return tok[6:], nil
	}
	return tok, nil
}

// Extract  token from Authorization header
// Uses PostExtractionFilter to strip "TOKEN " prefix from header
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromTokenString,
}

// Extractor for OAuth2 access tokens.  Looks in 'Authorization'
// header then 'access_token' argument for a token.
var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func UpdateContextAccountModel(c *gin.Context, accountId int64) error {
	var currentAccount *models.Account
	var err error
	if accountId != 0 {
		dbConn := common.GetDB()
		accountRepo := _accountRepo.NewMysqlAccountRepository(dbConn)
		currentAccount, err = accountRepo.GetById(c.Request.Context(), accountId)
		if err != nil {
			return err
		}
	}
	c.Set("current_account_id", accountId)
	c.Set("current_account", currentAccount)
	return nil
}


func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		UpdateContextAccountModel(c, 0)
		token, err := request.ParseFromRequest(c.Request, MyAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := []byte(common.NBSecretPassword)
			return b, nil
		})
		if err != nil {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			accountId := int64(claims["id"].(float64))
			err = UpdateContextAccountModel(c, accountId)
			if err != nil {
				c.AbortWithError(http.StatusUnauthorized, errors.New("ERROR TOKEN"))
			}
		}
	}
}

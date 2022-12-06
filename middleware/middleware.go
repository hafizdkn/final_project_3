package middleware

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"final_project_3/auth"
	"final_project_3/helper"
	"final_project_3/user"
)

// type resp struct {
// 	response helper.IResponse
// }

// type rr struct{}

var resp = helper.NewResponse

func AuthMiddleware(authService auth.IService, userService user.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		unauthorizedResponse := "Unauthorized"
		cantParseTokenReponse := "Can't parse token"
		tokenIvalidResponse := "Invalid Token"

		headerToken := c.Request.Header.Get("Authorization")
		bearer := strings.HasPrefix(headerToken, "Bearer")
		if !bearer {
			err := errors.New(unauthorizedResponse)
			response := resp.UnauthorizedResponse(err, unauthorizedResponse)
			resp.AbortJsonRespnse(c, response)
			return
		}

		stringToken := strings.Split(headerToken, " ")
		if len(stringToken) != 2 {
			err := errors.New(cantParseTokenReponse)
			response := resp.UnauthorizedResponse(err, "")
			resp.AbortJsonRespnse(c, response)
			return
		}

		token, err := authService.ValidateToken(stringToken[1])
		if err != nil {
			response := resp.UnauthorizedResponse(err, "")
			resp.AbortJsonRespnse(c, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			err := errors.New(tokenIvalidResponse)
			response := resp.UnauthorizedResponse(err, "")
			resp.AbortJsonRespnse(c, response)
			return
		}

		userId := int(claim["id"].(float64))
		user, err := userService.GetUserById(userId)
		if err != nil {
			response := resp.UnauthorizedResponse(err, "")
			resp.AbortJsonRespnse(c, response)
			return
		}

		c.Set("currentUser", user)
	}
}

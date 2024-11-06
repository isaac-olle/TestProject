package server

import (
	error2 "TestProject/internal/error"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"net/http"
)

const (
	authentication = "jwt"
	ContextString  = "AuthContext"
)

type AuthContext struct {
}

func ErrorHandlerMiddleware(c *gin.Context) {
	c.Next()

	err := c.Errors.Last()
	if err == nil {
		return
	}
	var httpError *error2.HttpError
	ok := errors.As(err.Err, &httpError)
	if !ok {
		c.JSON(http.StatusInternalServerError, "error in error handler")
	}
	c.AbortWithStatusJSON(httpError.HttpCode(), struct {
		cerror string `json:"error"`
	}{
		httpError.Error(),
	})
}

func validateToken(request *http.Request) (jwt.Token, error) {
	token, err := jwt.ParseHeader(request.Header, "Authorization", jwt.WithVerify(false))
	if err != nil {
		return nil, errors.New("invalid token: " + err.Error())
	}
	if token.Subject() == "" {
		return nil, errors.New("invalid token: Missing token subject")
	}
	return token, nil
}

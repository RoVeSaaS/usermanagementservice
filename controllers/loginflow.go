package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
)

// Login godoc
// @Summary Login Endpoint.
// @Description Login Endpoint for the service.
// @Tags LoginFlow
// @Produce json
// @Success 200
// @Router /login [get]
func GetAuthUrl(c *gin.Context) {
	url, err := usermanagement.GetAuthorizationURL(
		usermanagement.GetAuthorizationURLOpts{
			ClientID:    os.Getenv("WORKOS_CLIENT_ID"),
			Provider:    "authkit",
			RedirectURI: "http://localhost:8000/callback",
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusMovedPermanently, url.String())
}

func CallBack(c *gin.Context) {
	session := sessions.Default(c)
	LoginCode := c.Query("code")
	//LoginState := c.Param("state")
	fmt.Println(LoginCode)
	session.Set("code", LoginCode)
	session.Save()
	c.Redirect(http.StatusFound, "/userinfo")
}

func UserInfo(c *gin.Context) {
	session := sessions.Default(c)
	LoginCode := session.Get("code").(string)
	User, err := usermanagement.AuthenticateWithCode(
		c.Request.Context(),
		usermanagement.AuthenticateWithCodeOpts{
			ClientID: os.Getenv("WORKOS_CLIENT_ID"),
			Code:     LoginCode,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User": User})
}

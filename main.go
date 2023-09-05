package main

import (
	"gei"
	"net/http"
)

const (
	cookiename = "user-cookie"
	secretKey  = "your-secret-key"
)

func main() {
	r := gei.New()
	gei.Configinput()
	r.RunMiddware(gei.Logger())
	r.GET("/", func(c *gei.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.POST("/signup", gei.Signup())
	r.POST("/login", gei.Login())

	r.GET("/info", gei.AuthenticateMiddleware())

	r.GET("/logout", func(c *gei.Context) {
		cookie := http.Cookie{
			Name:   cookiename,
			Value:  "",
			MaxAge: -1,
		}
		http.SetCookie(c.Writer, &cookie)

		c.JSON(http.StatusOK, gei.H{"message": "Logged out successfully"})
	})

	r.Run(":9999")
}

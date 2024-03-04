package main

import (

	db "buyBuy-server/config"
	// "github.com/gorilla/sessions"

)

func main() {
	err := db.GetConnection()
	if err != nil {
		return nil, err
	}
// 	r := gin.Default()

// 	// Create a new session store.
// 	store := sessions.NewCookieStore([]byte("secret"))

// 	// Add the session middleware.
// 	r.Use(SessionMiddleware(store))

// 	r.GET("/login", controllers.LoginHandler)
// 	r.POST("/login", controllers.LoginHandler)

// 	r.GET("/logout", controllers.LogoutHandler)

// 	r.Run(":8080")
}


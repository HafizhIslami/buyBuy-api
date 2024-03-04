package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func SessionMiddleware(store *sessions.CookieStore) gin.HandlerFunc {
    return func(c *gin.Context) {
        session, err := store.Get(c.Request, "session")
        if err != nil {
            c.AbortWithStatus(http.StatusInternalServerError)
            return
        }

        c.Set("session", session)

        c.Next()
    }
}

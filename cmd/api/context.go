package main

import (
	"backend/internal/database"
	"github.com/gin-gonic/gin"
)

// this is a helper function to get the user from the context
func (app*application)GetuserFromContext(c *gin.Context) *database.User{
	contextUser, exists := c.Get("user")
	if !exists {
		return &database.User{}
	}

	user,ok := contextUser.(*database.User)
	if!ok{
		return &database.User{}
	}

	return user
}


/*In the Gin framework, a context (*gin.Context) is indeed used to store and pass request-scoped data (like authentication info, request metadata, etc.) throughout the lifecycle of a request.

What your statement gets right
“GetUserFromContext” is typically a helper function (not built into Gin itself).
It is used to retrieve user data stored in the context.
That data is usually placed there by an authentication middleware (e.g., after validating a JWT or session).
The context acts as a shared storage space for request-specific data.
Important clarification
GetUserFromContext is not a standard Gin method — it’s something developers usually define themselves.

Under the hood, it typically uses:

user, exists := c.Get("user")

where "user" is the key set earlier in middleware.*/
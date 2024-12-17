package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	// Login
	rt.router.POST("/session", rt.wrap(rt.doLogin))

	// User routes
	rt.router.PUT("/users/:id/name", rt.wrap(rt.setNameApi))
	rt.router.POST("/users/:id/conversations", rt.wrap(rt.newChatApi))

	return rt.router
}

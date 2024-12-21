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

	// Session
	rt.router.POST("/session", rt.wrap(rt.doLogin))

	// User routes
	rt.router.PUT("/users/:id/name", rt.wrap(rt.setName))
	rt.router.POST("/users/:id/conversations", rt.wrap(rt.newChat))
	rt.router.GET("/users/:id/conversations", rt.wrap(rt.getMyConversations))

	// Group routes
	rt.router.DELETE("/users/:id/groups/:group_id", rt.wrap(rt.leaveGroup))

	return rt.router
}

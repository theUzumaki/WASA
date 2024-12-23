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
	rt.router.POST("/session", rt.wrap(rt.doLogin, false))

	// User routes
	rt.router.PUT("/users/:id/name", rt.wrap(rt.setName, true))
	rt.router.POST("/users/:id/conversations", rt.wrap(rt.newChat, true))
	rt.router.GET("/users/:id/conversations", rt.wrap(rt.getMyConversations, true))
	rt.router.GET("/users/:id/conversations/:chat_id", rt.wrap(rt.getConversation, true))

	// Group routes
	rt.router.DELETE("/users/:id/groups/:group_id", rt.wrap(rt.leaveGroup, true))
	rt.router.PUT("/users/:id/groups/:group_id", rt.wrap(rt.addToGroup, true))

	return rt.router
}

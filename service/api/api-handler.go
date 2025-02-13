package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// Session
	rt.router.POST("/session", rt.wrap(rt.doLogin, false))

	// User routes
	rt.router.PUT("/users/:id/name", rt.wrap(rt.setName, true))
	rt.router.GET("/users/:id/search/:name", rt.wrap(rt.getUsers, true))
	rt.router.POST("/users/:id/conversations", rt.wrap(rt.newChat, true))
	rt.router.GET("/users/:id/conversations", rt.wrap(rt.getMyConversations, true))
	rt.router.GET("/users/:id/conversations/:chat_id", rt.wrap(rt.getConversation, true))
	rt.router.PUT("/users/:id/picture", rt.wrap(rt.setMyPhoto, true))

	// Group routes
	rt.router.DELETE("/users/:id/groups/:group_id", rt.wrap(rt.leaveGroup, true))
	rt.router.PUT("/users/:id/groups/:group_id", rt.wrap(rt.addToGroup, true))
	rt.router.PUT("/users/:id/groups/:group_id/name", rt.wrap(rt.setGroupName, true))
	rt.router.PUT("/users/:id/groups/:group_id/picture", rt.wrap(rt.setGroupPhoto, true))

	// Message routes
	rt.router.POST("/users/:id/conversations/:chat_id", rt.wrap(rt.sendMessage, true))
	rt.router.POST("/users/:id/conversations/:chat_id/message/:message_id", rt.wrap(rt.forwardMessage, true))
	rt.router.DELETE("/users/:id/conversations/:chat_id/message/:message_id", rt.wrap(rt.deleteMessage, true))
	rt.router.PUT("/users/:id/conversations/:chat_id/message/:message_id/comment", rt.wrap(rt.commentMessage, true))
	rt.router.DELETE("/users/:id/conversations/:chat_id/message/:message_id/comment", rt.wrap(rt.uncommentMessage, true))

	return rt.router
}

package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	//rt.router.GET("/user", rt.Getusername)
	rt.router.PUT("/user", rt.SetMyUserName)
	rt.router.GET("/user/:id", rt.Profile)
	rt.router.GET("/user", rt.Stream)
	rt.router.GET("/users", rt.Users_lst)
	rt.router.POST("/login", rt.Login)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))
	rt.router.GET("/Whoami", rt.Whoami)
	rt.router.POST("/user/:id/followers", rt.FollowUser)
	rt.router.DELETE("/user/:id/followers", rt.Remove_follower)
	rt.router.GET("/followers", rt.Follower_lst)
	rt.router.POST("/user/:id/ban", rt.BanUser)
	rt.router.DELETE("/user/:id/ban", rt.Remove_ban)
	rt.router.GET("/ban", rt.Baned_lst)
	rt.router.GET("/photos", rt.Photo_lst)
	rt.router.POST("/photos", rt.UploadPhoto)
	rt.router.DELETE("/photos/:id", rt.DeletePhoto)
	rt.router.POST("/photos/:id/likes", rt.Like)
	rt.router.GET("/photos/likes", rt.Like_lst)
	rt.router.DELETE("/photos/:id/likes", rt.DeleteLike)
	rt.router.POST("/photos/:id/comments", rt.Comment)
	rt.router.GET("/photos/comments/:id", rt.Comment_lst)
	rt.router.DELETE("/photos/:id/comments", rt.DeleteComment)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
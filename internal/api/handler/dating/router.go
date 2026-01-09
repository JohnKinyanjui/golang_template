package dating_handler

import "github.com/labstack/echo/v4"

type DatingRouter struct {
	private *echo.Group
	public  *echo.Group
}

func Handler(private, public *echo.Group) *DatingRouter {
	return &DatingRouter{public: public, private: private}
}

func (rt *DatingRouter) Routes() {
	rt.private.GET("/dating/feed", rt.getFeed)
	rt.private.GET("/dating/stories", rt.getStories)
	rt.private.GET("/dating/history", rt.getHistory)
	rt.private.POST("/dating/like", rt.likeUser)
	rt.private.POST("/dating/reject", rt.rejectUser)
}

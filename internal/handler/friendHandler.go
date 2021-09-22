package handler

import (
	"github.com/DarkReduX/social-network-server/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type FriendHandler struct {
	friendService *service.FriendService
}

func NewFriendHandler(friendService *service.FriendService) *FriendHandler {
	return &FriendHandler{friendService: friendService}
}

// AddFriendRequest godoc
// @Summary Send friend request to user with sent ID
// @Description User to add in friend list Identifier
// @Tags friend
// @Accept  json
// @Produce  json
// @Param friend_id query string true "User to add in friend identifier"
// @Security BearerToken
// @Success 200
// @Failure 400
// @Router /friend [post]
func (h FriendHandler) AddFriendRequest(c echo.Context) error {
	username := c.Get("username").(string)
	friendName := c.QueryParam("friend_id")
	if err := h.friendService.AddFriendRequest(c.Request().Context(), username, friendName); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}

// ProcessFriendRequest godoc
// @Summary Decline or Accept friend request from other user.
// @Tags friend
// @Accept  json
// @Produce  json
// @Security BearerToken
// @Param request_type query string true "DECLINE/ACCEPT"
// @Success 200
// @Failure 400
// @Router /friend [put]
func (h FriendHandler) ProcessFriendRequest(c echo.Context) error {
	username := c.Get("username").(string)
	friendName := c.QueryParam("friend_id")
	requestType := c.QueryParam("request_type")
	if err := h.friendService.ProcessFriendRequest(c.Request().Context(), username, friendName, requestType); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}

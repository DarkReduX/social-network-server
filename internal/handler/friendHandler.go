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

func (h FriendHandler) AddFriendRequest(c echo.Context) error {
	username := c.Get("username").(string)
	friendName := c.QueryParam("friend_id")
	if err := h.friendService.AddFriendRequest(c.Request().Context(), username, friendName); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}

func (h FriendHandler) ProcessFriendRequest(c echo.Context) error {
	username := c.Get("username").(string)
	friendName := c.QueryParam("friend_id")
	requestType := c.QueryParam("request_type")
	if err := h.friendService.ProcessFriendRequest(c.Request().Context(), username, friendName, requestType); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}

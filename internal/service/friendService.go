package service

import (
	"context"
	"github.com/DarkReduX/social-network-server/internal/repository"
)

type FriendService struct {
	friendRepository *repository.FriendRepository
}

func NewFriendService(friendRepository *repository.FriendRepository) *FriendService {
	return &FriendService{friendRepository: friendRepository}
}

func (s *FriendService) AddFriendRequest(ctx context.Context, userId string, friendId string) error {
	return s.friendRepository.AddFriendRequest(ctx, userId, friendId)
}

func (s *FriendService) ProcessFriendRequest(ctx context.Context, userId string, friendId string, requestType string) error {
	return s.friendRepository.ProcessFriendRequest(ctx, userId, friendId, requestType)
}

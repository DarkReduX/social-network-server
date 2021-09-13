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
	return s.friendRepository.Add(ctx, userId, friendId)
}

func (s *FriendService) DeleteFriend(ctx context.Context, userId string, friendId string) error {
	return s.friendRepository.Delete(ctx, userId, friendId)
}

func (s *FriendService) SubmitFriendRequest(ctx context.Context, userId string, friendId string) error {
	return s.friendRepository.SubmitFriend(ctx, userId, friendId)
}

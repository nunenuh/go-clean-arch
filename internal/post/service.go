package post

import (
	"context"
	"time"
)

type postService struct {
	postRepository PostRepository
}

func NewPostService(r PostRepository) PostService {
	return &postService{postRepository: r}
}

func (s *postService) FetchPosts(ctx context.Context) (*[]PostAndUser, error) {
	return s.postRepository.GetPosts(ctx)
}

func (s *postService) FetchPost(ctx context.Context, postID int) (*PostAndUser, error) {
	return s.postRepository.GetPost(ctx, postID)
}

func (s *postService) BuildPost(ctx context.Context, post *Post, userID int) error {
	post.Created = time.Now().Unix()
	post.Modified = time.Now().Unix()
	post.User = userID

	return s.postRepository.CreatePost(ctx, post)
}

func (s *postService) ModifyPost(ctx context.Context, postID int, post *Post, userID int) error {
	post.Modified = time.Now().Unix()
	post.User = userID

	return s.postRepository.UpdatePost(ctx, postID, post)
}

func (s *postService) DestroyPost(ctx context.Context, postID int) error {
	return s.postRepository.DeletePost(ctx, postID)
}

package usecase

import (
	"Test/domain"
	"errors"

	"github.com/google/uuid"
)

type postUseCase struct {
	postRepo domain.PostRepository
}

func NewPostUseCase(postRepo domain.PostRepository) *postUseCase {
	return &postUseCase{postRepo: postRepo}
}

func (uc *postUseCase) Create(post *domain.PostReq) (*domain.PostRes, error) {
	if post.Title == "" {
		return nil, errors.New("title is required")
	}

	id := uuid.New()

	posts := &domain.Post{
		ID:        id.String(),
		Title:     post.Title,
		Content:   post.Content,
		Published: false,
	}
	res, err := uc.postRepo.Create(posts)
	if err != nil {
		return nil, err
	}
	return &domain.PostRes{
		ID:        res.ID,
		Title:     res.Title,
		Content:   res.Content,
		Published: res.Published,
	}, nil

}

func (uc *postUseCase) GetAll(query *domain.PostAllReq, pagination *domain.Pagination) (*domain.PostResponse, error) {
	res, count, totalPage, err := uc.postRepo.FindAllField(query, pagination)
	if err != nil {
		return nil, err
	}

	response := &domain.PostResponse{
		Posts:     make([]domain.PostRes, len(res)),
		Count:     int(count),
		Limit:     pagination.PageSize,
		Page:      pagination.Page,
		TotalPage: int(totalPage),
	}

	for i, post := range res {
		response.Posts[i] = domain.PostRes{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			Published: post.Published,
		}
	}

	return response, nil
}

func (uc *postUseCase) GetByID(id string) (*domain.PostRes, error) {
	postID := &domain.Post{
		ID: string(id),
	}
	res, err := uc.postRepo.FindOne(postID)
	if err != nil {
		return nil, err
	}

	return &domain.PostRes{
		ID:        res.ID,
		Title:     res.Title,
		Content:   res.Content,
		Published: res.Published,
	}, nil

}

func (uc *postUseCase) UpdateByID(post *domain.PostUpdateReq) (*domain.PostRes, error) {
	posts := &domain.Post{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		Published: post.Published,
	}

	res, err := uc.postRepo.UpdateByID(posts)
	if err != nil {
		return nil, err
	}

	return &domain.PostRes{
		ID:        res.ID,
		Title:     res.Title,
		Content:   res.Content,
		Published: res.Published,
	}, nil
}

func (uc *postUseCase) DeleteByID(id string) error {
	post := &domain.Post{}
	post.ID = id
	err := uc.postRepo.DeleteByID(post)
	if err != nil {
		return err
	}
	return nil
}

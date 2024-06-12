package service

import (
	"github.com/devfullcycle/14-gRPC/internal/database"
	"github.com/devfullcycle/14-gRPC/internal/pb"
)

type CategoryService struct { 
	pb.UnimplementedCategoryServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.CreateCategory(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	CategoryResponse := &pb.Category{
		Id: category.Id,
		Name: category.Name,
		Description: category.Description,
	}
	return &pb.categoryResponse{
		Category: categoryResponse,
	}, nil
}

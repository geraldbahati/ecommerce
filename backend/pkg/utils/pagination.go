package utils

import (
	"context"
	"github.com/geraldbahati/ecommerce/pkg/config"
	"github.com/geraldbahati/ecommerce/pkg/model"
)

func Paginate(
	ctx context.Context,
	totalCount int64,
	page int32,
	pageSize int32,
	fetchData func(ctx context.Context, offset int32, limit int32) (interface{}, error),
) (*model.PaginationResult, error) {
	cfg := config.LoadConfig()

	if page < 1 {
		page = cfg.DefaultPage
	}

	if pageSize < 1 {
		pageSize = cfg.DefaultPageSize
	}

	offset := (page - 1) * pageSize
	data, err := fetchData(ctx, offset, pageSize)
	if err != nil {
		return nil, err
	}

	totalPages := totalCount / int64(pageSize)
	if totalCount%int64(pageSize) > 0 {
		totalPages++
	}

	return &model.PaginationResult{
		TotalCount: totalCount,
		TotalPages: int32(totalPages),
		Page:       page,
		PageSize:   pageSize,
		Data:       data,
	}, nil
}

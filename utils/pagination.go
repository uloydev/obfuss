package utils

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/models"
)

const (
	DEFAULT_PAGINATION_PAGE = 1
	DEFAULT_PAGINATION_SIZE = 10
	PAGINATION_MAX_SIZE     = 100
)

var (
	ErrPaginationMaxSize = errors.New("invalid pagination size, max 100")
	ErrFailedFetchData   = errors.New("failed to fetch data")
)

func Paginate[T any](
	params models.PaginationParams,
	query *gorm.DB,
	logger *zap.Logger,
) (
	meta *models.PaginationMeta,
	data []T,
	err error,
) {
	var total int64

	if err = query.Count(&total).Error; err != nil {
		logger.Error("failed to count data", zap.Error(err))
		return
	}

	if params.Page <= 0 {
		params.Page = DEFAULT_PAGINATION_PAGE
	}

	if params.Size <= 0 {
		params.Size = DEFAULT_PAGINATION_SIZE
	}

	if params.Size > PAGINATION_MAX_SIZE {
		err = ErrPaginationMaxSize
		return
	}

	meta = &models.PaginationMeta{
		Total:     total,
		TotalPage: total/int64(params.Size) + 1,
		Size:      params.Size,
		Page:      params.Page,
	}

	err = query.Offset((params.Page - 1) * params.Size).
		Limit(params.Size).
		Find(&data).Error

	if err != nil {
		logger.Error("failed to fetch data", zap.Error(err))
		err = ErrFailedFetchData
		return
	}

	return
}

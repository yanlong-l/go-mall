package app

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yanlong-l/go-mall/config"
)

type pagination struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewPagination(ctx *gin.Context) *pagination {
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page <= 0 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	if pageSize <= 0 {
		pageSize = config.App.Pagination.DefaultSize
	}
	if pageSize > config.App.Pagination.MaxSize {
		pageSize = config.App.Pagination.MaxSize
	}

	return &pagination{Page: page, PageSize: pageSize}
}

func (p *pagination) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *pagination) GetPageSize() int {
	return p.PageSize
}

func (p *pagination) GetPage() int {
	return p.Page
}

func (p *pagination) SetTotalRows(total int) {
	p.TotalRows = total
}

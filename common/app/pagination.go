package app

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yanlong-l/go-mall/config"
)

type Pagination struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewPagination(ctx *gin.Context) *Pagination {
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

	return &Pagination{Page: page, PageSize: pageSize}
}

func (p *Pagination) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *Pagination) GetPageSize() int {
	return p.PageSize
}

func (p *Pagination) GetPage() int {
	return p.Page
}

func (p *Pagination) SetTotalRows(total int) {
	p.TotalRows = total
}

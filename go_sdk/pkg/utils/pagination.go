package utils

import (
	"context"
	"fmt"

	"github.com/tiktok/tiktok-business-api-sdk/go_sdk/pkg/models"
)

// PaginationOptions holds pagination configuration
type PaginationOptions struct {
	PageSize int
	MaxPages int // 0 means no limit
}

// DefaultPaginationOptions returns sensible defaults for pagination
func DefaultPaginationOptions() *PaginationOptions {
	return &PaginationOptions{
		PageSize: 50,
		MaxPages: 0, // No limit
	}
}

// PaginatedResult represents a single page of results
type PaginatedResult[T any] struct {
	Data     []T                   `json:"data"`
	PageInfo models.PaginationInfo `json:"page_info"`
}

// PaginationIterator provides an iterator interface for paginated results
type PaginationIterator[T any] struct {
	fetchFunc   func(ctx context.Context, page int, pageSize int) (*PaginatedResult[T], error)
	options     *PaginationOptions
	currentPage int
	hasMore     bool
	lastResult  *PaginatedResult[T]
}

// NewPaginationIterator creates a new pagination iterator
func NewPaginationIterator[T any](
	fetchFunc func(ctx context.Context, page int, pageSize int) (*PaginatedResult[T], error),
	options *PaginationOptions,
) *PaginationIterator[T] {
	if options == nil {
		options = DefaultPaginationOptions()
	}

	return &PaginationIterator[T]{
		fetchFunc:   fetchFunc,
		options:     options,
		currentPage: 1,
		hasMore:     true,
	}
}

// Next fetches the next page of results
func (p *PaginationIterator[T]) Next(ctx context.Context) (*PaginatedResult[T], error) {
	if !p.hasMore {
		return nil, fmt.Errorf("no more pages available")
	}

	if p.options.MaxPages > 0 && p.currentPage > p.options.MaxPages {
		p.hasMore = false
		return nil, fmt.Errorf("maximum pages limit reached")
	}

	result, err := p.fetchFunc(ctx, p.currentPage, p.options.PageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch page %d: %w", p.currentPage, err)
	}

	p.lastResult = result
	p.currentPage++
	p.hasMore = result.PageInfo.HasMore

	return result, nil
}

// HasNext returns true if there are more pages available
func (p *PaginationIterator[T]) HasNext() bool {
	return p.hasMore
}

// CurrentPage returns the current page number (1-based)
func (p *PaginationIterator[T]) CurrentPage() int {
	return p.currentPage - 1 // Adjust for 0-based indexing after increment
}

// LastResult returns the last fetched result
func (p *PaginationIterator[T]) LastResult() *PaginatedResult[T] {
	return p.lastResult
}

// AllPages fetches all available pages and returns all results
func (p *PaginationIterator[T]) AllPages(ctx context.Context) ([]T, error) {
	var allResults []T

	for p.HasNext() {
		result, err := p.Next(ctx)
		if err != nil {
			return allResults, err
		}

		allResults = append(allResults, result.Data...)

		// Safety check to prevent infinite loops
		if len(allResults) > 100000 {
			return allResults, fmt.Errorf("too many results, stopping at 100,000 items")
		}
	}

	return allResults, nil
}

// ForEach iterates through all pages and calls the provided function for each item
func (p *PaginationIterator[T]) ForEach(ctx context.Context, fn func(item T) error) error {
	for p.HasNext() {
		result, err := p.Next(ctx)
		if err != nil {
			return err
		}

		for _, item := range result.Data {
			if err := fn(item); err != nil {
				return err
			}
		}
	}

	return nil
}

// ForEachPage iterates through all pages and calls the provided function for each page
func (p *PaginationIterator[T]) ForEachPage(ctx context.Context, fn func(page *PaginatedResult[T]) error) error {
	for p.HasNext() {
		result, err := p.Next(ctx)
		if err != nil {
			return err
		}

		if err := fn(result); err != nil {
			return err
		}
	}

	return nil
}

// PaginationHelper provides utility functions for pagination
type PaginationHelper struct{}

// NewPaginationHelper creates a new pagination helper
func NewPaginationHelper() *PaginationHelper {
	return &PaginationHelper{}
}

// CalculateOffset calculates the offset for a given page and page size
func (h *PaginationHelper) CalculateOffset(page, pageSize int) int {
	if page <= 0 {
		page = 1
	}
	return (page - 1) * pageSize
}

// CalculateTotalPages calculates the total number of pages for a given total count and page size
func (h *PaginationHelper) CalculateTotalPages(totalCount, pageSize int) int {
	if pageSize <= 0 {
		return 0
	}
	return (totalCount + pageSize - 1) / pageSize
}

// ValidatePaginationParams validates pagination parameters
func (h *PaginationHelper) ValidatePaginationParams(page, pageSize int) error {
	if page <= 0 {
		return models.NewValidationError("page", "page must be greater than 0")
	}

	if pageSize <= 0 {
		return models.NewValidationError("page_size", "page size must be greater than 0")
	}

	if pageSize > 1000 {
		return models.NewValidationError("page_size", "page size cannot exceed 1000")
	}

	return nil
}

// CreatePaginationInfo creates pagination info from parameters
func (h *PaginationHelper) CreatePaginationInfo(page, pageSize, totalCount int) models.PaginationInfo {
	totalPages := h.CalculateTotalPages(totalCount, pageSize)
	hasMore := page < totalPages

	return models.PaginationInfo{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: totalCount,
		TotalPage:  totalPages,
		HasMore:    hasMore,
	}
}

// BatchProcessor processes items in batches
type BatchProcessor[T any] struct {
	batchSize int
	processor func(batch []T) error
}

// NewBatchProcessor creates a new batch processor
func NewBatchProcessor[T any](batchSize int, processor func(batch []T) error) *BatchProcessor[T] {
	return &BatchProcessor[T]{
		batchSize: batchSize,
		processor: processor,
	}
}

// Process processes all items in batches
func (b *BatchProcessor[T]) Process(items []T) error {
	for i := 0; i < len(items); i += b.batchSize {
		end := i + b.batchSize
		if end > len(items) {
			end = len(items)
		}

		batch := items[i:end]
		if err := b.processor(batch); err != nil {
			return fmt.Errorf("failed to process batch starting at index %d: %w", i, err)
		}
	}

	return nil
}

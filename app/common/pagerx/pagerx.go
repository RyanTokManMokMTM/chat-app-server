package pagerx

import "math"

const (
	DEFAULT_MAX_LIMIT  uint = 20
	DEFAULT_MAX_OFFSET uint = 20
)

//// GetOffset to prevent over-range
//func GetOffset(offset uint) uint {
//	if offset > DEFAULT_MAX_OFFSET {
//		return DEFAULT_MAX_OFFSET
//	}
//	return offset
//}

// GetLimit to prevent over-range
func GetLimit(limit uint) uint {
	if limit > DEFAULT_MAX_LIMIT {
		return DEFAULT_MAX_LIMIT
	}

	return limit
}

// PageOffset to get current page
func PageOffset(pageSize, page uint) uint {
	return (pageSize) * (page - 1)
}

// GetTotalPageByPageSize to get total page
func GetTotalPageByPageSize(total, pageSize uint) uint {
	return uint(math.Ceil(float64(total) / float64(pageSize)))
}

package api

// ClampPageSize returns the page size clamped between the default and max values.
// If requested is <= 0, defaultSize is returned.
func ClampPageSize(requested, defaultSize, maxSize int) int {
	if requested <= 0 {
		return defaultSize
	}
	if requested > maxSize {
		return maxSize
	}
	return requested
}

// PaginateWith runs a fetch-and-accumulate loop over token-based paginated endpoints.
// fetch is called with a page token (empty string for first page).
// It returns the collected items, whether there are more pages, and the next page token.
func PaginateWith[T any](
	fetch func(pageToken string) (items []T, hasMore bool, nextToken string, err error),
) ([]T, error) {
	var all []T
	var pageToken string

	for {
		items, hasMore, nextToken, err := fetch(pageToken)
		if err != nil {
			return nil, err
		}
		all = append(all, items...)
		if !hasMore || nextToken == "" {
			break
		}
		pageToken = nextToken
	}

	return all, nil
}

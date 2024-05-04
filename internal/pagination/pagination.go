package pagination

var (
	defaultPageSize  = 20
	defaultPageIndex = 1
)

func GetDefaultPagination(pageIndex, pageSize int, order string) (int, int, string) {
	newPageSize := pageSize
	newPageIndex := pageIndex
	newOrder := order
	if newPageIndex == 0 {
		newPageIndex = defaultPageIndex
	}
	if newPageSize == 0 {
		newPageSize = defaultPageSize
	}

	return newPageIndex, newPageSize, newOrder
}

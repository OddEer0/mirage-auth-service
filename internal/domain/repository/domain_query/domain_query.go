package domainQuery

const (
	SortAsc  = "ASC"
	SortDesc = "DESC"
)

type (
	PaginationQuery struct {
		CurrentPage uint
		PageCount   uint
	}

	OrderQuery struct {
		OrderBy        string
		OrderDirection string
	}

	UserQueryRequest struct {
		PaginationQuery
		OrderQuery
	}
)

package domainQuery

const (
	OrderAsc  = "ASC"
	OrderDesc = "DESC"
)

type (
	PaginationQuery struct {
		CurrentPage uint `json:"current_page"`
		PageCount   uint `json:"page_count"`
	}

	OrderQuery struct {
		OrderBy        string `json:"order_by"`
		OrderDirection string `json:"order_direction"`
	}

	UserQueryRequest struct {
		PaginationQuery `json:"pagination_query"`
		OrderQuery      `json:"order_query"`
	}
)

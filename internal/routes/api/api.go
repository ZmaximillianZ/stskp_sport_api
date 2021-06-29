package api

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/repository"
)

type QueryParams struct {
	Criteria map[string][2]string `json:"criteria"`
	Order    map[string]string    `json:"order"`
	Limit    int                  `json:"limit"`
	Offset   int                  `json:"offset"`
}

func ParseQueryParams(
	allowedFields []string,
	qp *QueryParams,
) (criteria map[string][2]string, order map[string]string, limit, offset int, isComparisonOperator bool) {
	isComparisonOperator = true
	criteria = make(map[string][2]string)
	order = make(map[string]string)
	for _, value := range allowedFields {
		queryVal, okVal := qp.Criteria[value]
		orderVal, okOrder := qp.Order[value]
		if !okOrder {
			logging.Error("user doesn't have this field for order")
			break
		}
		order[value] = orderVal
		if !okVal {
			continue
		}
		isComparisonOperator = repository.IsComparisonOperator(queryVal[0])
		if !isComparisonOperator {
			logging.Error("last element must be comparison operator")
			break
		}
		criteria[value] = queryVal
	}

	return criteria, qp.Order, qp.Limit, qp.Offset, isComparisonOperator
}

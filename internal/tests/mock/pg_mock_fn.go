package mock

import (
	"errors"
	"time"
)

type (
	PgMockRow struct {
		Data []any
	}
)

func (m PgMockRow) Scan(dest ...any) error {
	if len(m.Data) != len(dest) {
		return errors.New("invalid dest count")
	}
	for i := 0; i < len(dest); i++ {
		switch v := dest[i].(type) {
		case *string:
			*v = m.Data[i].(string)
		case *time.Time:
			*v = m.Data[i].(time.Time)
		case *bool:
			*v = m.Data[i].(bool)
		case *int:
			*v = m.Data[i].(int)
		case *int64:
			*v = m.Data[i].(int64)
		case *float64:
			*v = m.Data[i].(float64)
		case **string:
			*v = m.Data[i].(*string)
		default:
			return errors.New("unsupported scan type")
		}
	}
	return nil
}

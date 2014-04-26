package models

import (
	"fmt"
	"github.com/coopernurse/gorp"
	"strings"
)

var (
	DbMap *gorp.DbMap
)

func SelectQuery(table string, attributes map[string]string) string {
	query := fmt.Sprintf("select * from %s", table)
	for key, value := range attributes {
		if strings.Index(query, "where") == -1 {
			query = fmt.Sprintf("%s where %s = '%s'", query, key, value)
		} else {
			query = fmt.Sprintf("%s and %s = '%s'", query, key, value)
		}
	}
	return query
}

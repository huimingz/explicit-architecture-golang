package dql

import "explicit/core/port/persistence"

type QueryBuilderInterface interface {
	persistence.QueryBuilderInterface

	SetParameter(k string, v interface{}) QueryBuilderInterface
	SetMaxResults(maxResults int) QueryBuilderInterface
	Distinct(flag bool) QueryBuilderInterface
	Select(fields ...string) QueryBuilderInterface
	Update(k string, v interface{}) QueryBuilderInterface
	OrderBy(k string, order string) QueryBuilderInterface
}

package db

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type QueryBuilder struct {
	Query strings.Builder
	Args  []interface{}
	Count int
}

func (qb *QueryBuilder) AddWhereCondition(condition string, value interface{}) {
	qb.Count++
	logThis := log.WithFields(logrus.Fields{
		"package":  "query_builder",
		"function": "QueryBuilder AddWhere",
	})
	logThis.WithFields(logrus.Fields{
		"condition":       condition,
		"value":           value,
		"argument number": qb.Count,
	}).Debug()
	format := " AND %s $%d"
	qb.Args = append(qb.Args, value)
	qb.Query.WriteString(fmt.Sprintf(format, condition, qb.Count))
}

func (qb *QueryBuilder) AddOrderBy(column string, order string) {
	condition := "ORDER BY"
	logThis := log.WithFields(logrus.Fields{
		"package":  "query_builder",
		"function": "QueryBuilder AddOrderBy",
	})
	logThis.WithFields(logrus.Fields{
		"condition":         condition,
		"column":            column,
		"order":             order,
		"argument number 1": qb.Count,
		"argument number 2": qb.Count + 1,
	}).Debug("Added order by filter")
	format := " %s %s %s"
	qb.Query.WriteString(fmt.Sprintf(format, condition, column, order))
}

func (qb *QueryBuilder) AddLimit(value interface{}) {
	qb.Count++
	condition := "LIMIT"
	logThis := log.WithFields(logrus.Fields{
		"package":  "query_builder",
		"function": "QueryBuilder AddLimit",
	})
	logThis.WithFields(logrus.Fields{
		"condition":       condition,
		"value":           value,
		"argument number": qb.Count,
	}).Debug()
	format := " %s $%d"
	qb.Args = append(qb.Args, value)
	qb.Query.WriteString(fmt.Sprintf(format, condition, qb.Count))
}

func (qb *QueryBuilder) AddOffset(value interface{}) {
	qb.Count++
	condition := "OFFSET"
	logThis := log.WithFields(logrus.Fields{
		"package":  "query_builder",
		"function": "QueryBuilder AddOffset",
	})
	logThis.WithFields(logrus.Fields{
		"condition":       condition,
		"value":           value,
		"argument number": qb.Count,
	}).Debug()
	format := " %s $%d"
	qb.Args = append(qb.Args, value)
	qb.Query.WriteString(fmt.Sprintf(format, condition, qb.Count))
}

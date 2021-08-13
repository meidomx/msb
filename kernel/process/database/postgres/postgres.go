package postgres

import (
	"context"
	"errors"

	"github.com/meidomx/msb/api/kern"
	"github.com/meidomx/msb/api/moduleapi"
	"github.com/meidomx/msb/module"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresDatabaseOperatorFactory struct {
}

func (p *PostgresDatabaseOperatorFactory) Name() string {
	return "builtin.database.operator.factory.postgres"
}

func (p *PostgresDatabaseOperatorFactory) LoadConfig(m map[string]interface{}, i interface{}) (kern.Service, error) {
	po := new(PostgresDatabaseOperator)
	name, ok := m["name"]
	if !ok {
		return nil, errors.New("postgres: no name specified")
	}
	po.InstName, ok = name.(string)
	if !ok {
		return nil, errors.New("postgres: name should be string")
	}
	ds, ok := i.(*pgxpool.Pool)
	if !ok {
		return nil, errors.New("postgres: datasource pool is required")
	}
	po.ds = ds
	return po, nil
}

type PostgresDatabaseOperator struct {
	kern.DefaultService

	ds *pgxpool.Pool
}

func (p *PostgresDatabaseOperator) Handle(i interface{}) (interface{}, error) {
	r, ok := i.(*moduleapi.DatabaseQueryRequest)
	if !ok {
		return nil, errors.New("postgres: request type is wrong")
	}

	switch r.Operation {
	case "insert":
		fallthrough
	case "delete":
		fallthrough
	case "update":
		for _, v := range r.Parameters {
			_, err := p.ds.Exec(context.Background(), r.SQL, v...)
			if err != nil {
				return nil, err
			}
		}
		return &moduleapi.DatabaseQueryResult{}, nil
	case "select":
		rs := new(moduleapi.DatabaseQueryResult)
		parameters := r.Parameters
		if len(parameters) == 0 {
			parameters = append(parameters, []interface{}{})
		}
		for _, v := range parameters {
			rows, err := p.ds.Query(context.Background(), r.SQL, v...)
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			var columnNames []string
			for _, fd := range rows.FieldDescriptions() {
				columnNames = append(columnNames, string(fd.Name))
			}
			var rrows [][]interface{}
			for rows.Next() {
				row, err := rows.Values()
				if err != nil {
					return nil, err
				}
				rrows = append(rrows, row)
			}
			rows.RawValues()
			rs.Results = append(rs.Results, struct {
				ColumnName []string
				Rows       [][]interface{}
			}{ColumnName: columnNames, Rows: rrows})
		}
		return rs, nil
	default:
		return nil, errors.New("postgres: operation not support:" + r.Operation)
	}
}

var _ kern.Service = new(PostgresDatabaseOperator)
var _ kern.ServiceFactory = new(PostgresDatabaseOperatorFactory)

func init() {
	module.RegisterKernFactory(new(PostgresDatabaseOperatorFactory))
}

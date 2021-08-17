package postgres

import (
	"context"
	"errors"

	"github.com/meidomx/msb/api"
	"github.com/meidomx/msb/api/kern"
	"github.com/meidomx/msb/module"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresDatasourceFactory struct {
}

func (p *PostgresDatasourceFactory) Name() string {
	return "builtin.database.provider.postgres"
}

func (p *PostgresDatasourceFactory) LoadConfig(m map[string]interface{}, i interface{}) (kern.Binding, error) {
	pd := new(PostgresDatasourceProvider)
	name, ok := m["name"]
	if !ok {
		return nil, errors.New("postgres: no name specified")
	}
	pd.InstName, ok = name.(string)
	if !ok {
		return nil, errors.New("postgres: name should be string")
	}
	s, ok := m["connect_string"]
	if !ok {
		return nil, errors.New("postgres: no connect string specified")
	}
	pd.ConnectString, ok = s.(string)
	if !ok {
		return nil, errors.New("postgres: connect string should be string")
	}
	return pd, nil
}

type PostgresDatasourceProvider struct {
	kern.DefaultBinding

	ConnectString string

	pool *pgxpool.Pool
}

func (p *PostgresDatasourceProvider) Bind(msbCtx api.MsbContext, parameter interface{}) (interface{}, error) {
	pool, err := pgxpool.Connect(context.Background(), p.ConnectString)
	if err != nil {
		return nil, err
	}
	p.pool = pool
	return pool, nil
}

var _ kern.BindingFactory = new(PostgresDatasourceFactory)
var _ kern.Binding = new(PostgresDatasourceProvider)

func init() {
	module.RegisterKernFactory(new(PostgresDatasourceFactory))
}

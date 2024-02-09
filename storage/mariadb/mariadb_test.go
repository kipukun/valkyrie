package mariadb_test

import (
	"context"
	"testing"

	radio "github.com/R-a-dio/valkyrie"
	"github.com/R-a-dio/valkyrie/config"
	mysqlmigrations "github.com/R-a-dio/valkyrie/migrations/mysql"
	"github.com/R-a-dio/valkyrie/storage"
	storagetest "github.com/R-a-dio/valkyrie/storage/test"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mariadb"
)

type MariaDBSetup struct {
	container *mariadb.MariaDBContainer
}

func (setup *MariaDBSetup) Setup(ctx context.Context) (radio.StorageService, error) {
	cfg, err := config.LoadFile()
	if err != nil {
		return nil, err
	}

	testcontainers.Logger = testcontainers.TestLogger(storagetest.CtxT(ctx))

	// setup a container to test in
	container, err := mariadb.RunContainer(ctx,
		testcontainers.WithImage("mariadb:latest"),
		mariadb.WithDatabase("test"),
		mariadb.WithUsername("root"),
		mariadb.WithPassword(""),
	)
	if err != nil {
		return nil, err
	}
	setup.container = container
	// then update our config to connect to the container
	bare := cfg.Conf()
	bare.Database.DSN, err = container.ConnectionString(ctx)
	if err != nil {
		return nil, err
	}
	cfg.StoreConf(bare)

	// run migrations
	err = setup.RunMigrations(ctx, cfg)
	if err != nil {
		return nil, err
	}

	// then open a storage instance
	s, err := storage.Open(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (setup *MariaDBSetup) RunMigrations(ctx context.Context, cfg config.Config) error {
	migr, err := mysqlmigrations.New(ctx, cfg)
	if err != nil {
		return err
	}

	err = migr.Up()
	if err != nil {
		return err
	}
	return nil
}

func (setup *MariaDBSetup) TearDown(ctx context.Context) error {
	err := setup.container.Terminate(ctx)
	if err != nil {
		return err
	}
	return nil
}

func TestMariaDBStorage(t *testing.T) {
	if !testing.Short() {
		storagetest.RunTests(t, new(MariaDBSetup))
	}
}

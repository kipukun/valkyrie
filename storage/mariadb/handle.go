package mariadb

import (
	"context"
	"database/sql"
	"strings"
	"time"

	radio "github.com/R-a-dio/valkyrie"
	"github.com/R-a-dio/valkyrie/config"
	"github.com/R-a-dio/valkyrie/storage"
	"github.com/go-sql-driver/mysql" // mariadb
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

func init() {
	storage.Register("mariadb", Connect)
}

var DatabaseConnectFunc = sqlx.ConnectContext

// specialCasedColumnNames is a map of Go <StructField> to SQL <ColumnName>
var specialCasedColumnNames = map[string]string{
	"CreatedAt":     "created_at",
	"DeletedAt":     "deleted_at",
	"UpdatedAt":     "updated_at",
	"RememberToken": "remember_token",
}

// mapperFunc implements the MapperFunc for sqlx to specialcase column names
// and lowercase them for scan matching
func mapperFunc(s string) string {
	n, ok := specialCasedColumnNames[s]
	if ok {
		s = n
	}
	return strings.ToLower(s)
}

// ConnectDB connects to the configured mariadb instance and returns the raw database
// object. Argument multistatement indicates if we should allow queries with multiple
// statements in them.
func ConnectDB(ctx context.Context, cfg config.Config, multistatement bool) (*sqlx.DB, error) {
	info := cfg.Conf().Database

	// we require some specific arguments in the DSN to have code work properly, so make
	// sure those are included
	dsn, err := mysql.ParseDSN(info.DSN)
	if err != nil {
		return nil, err
	}

	// enable multistatement queries if asked for
	if multistatement {
		dsn.MultiStatements = true
	}
	// UTC location to handle time.Time location
	dsn.Loc, err = time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	// parsetime to handle time.Time in the driver
	dsn.ParseTime = true
	// time_zone to have the database not try and interpret dates and times as the
	// locale of the system, but as UTC+0 instead
	if dsn.Params == nil {
		dsn.Params = map[string]string{}
	}
	dsn.Params["time_zone"] = "'+00:00'"
	conndsn := dsn.FormatDSN()

	// we want to print what we're connecting to, but not print our password
	if dsn.Passwd != "" {
		dsn.Passwd = "<redacted>"
	}

	zerolog.Ctx(ctx).Info().Str("address", dsn.FormatDSN()).Msg("trying to connect")

	db, err := DatabaseConnectFunc(ctx, "mysql", conndsn)
	if err != nil {
		return nil, err
	}

	db.MapperFunc(mapperFunc)

	return db, nil
}

// Connect connects to the database configured in cfg
func Connect(ctx context.Context, cfg config.Config) (radio.StorageService, error) {
	// TODO(wessie): check if we've applied all our SQL migrations before allowing usage
	db, err := ConnectDB(ctx, cfg, false)
	if err != nil {
		return nil, err
	}
	return &StorageService{db}, nil
}

// StorageService implements radio.StorageService with a sql database
type StorageService struct {
	db *sqlx.DB
}

// fakeTx is a *sqlx.Tx with the Commit method disabled
type fakeTx struct {
	*sqlx.Tx
}

// Commit does nothing
func (fakeTx) Commit() error {
	return nil
}

// tx either unwraps the tx given to a *sqlx.Tx, or creates a new transaction if tx is
// nil. Passing in a StorageTx not returned by this package will panic
func (s *StorageService) tx(ctx context.Context, tx radio.StorageTx) (*sqlx.Tx, radio.StorageTx, error) {
	if tx == nil {
		// new transaction
		tx, err := s.db.BeginTxx(ctx, nil)
		return tx, tx, err
	}

	// existing transaction, make sure it's one of ours and then use it
	switch txx := tx.(type) {
	case *sqlx.Tx:
		// if this is a real tx, we disable the commit so that the transaction can't
		// be committed earlier than expected by the creator
		return txx, fakeTx{txx}, nil
	case fakeTx:
		return txx.Tx, txx, nil
	default:
		panic("mariadb: invalid tx passed to StorageService")
	}
}

func (s *StorageService) Sessions(ctx context.Context) radio.SessionStorage {
	return SessionStorage{
		handle: handle{s.db, ctx, "sessions"},
	}
}

func (s *StorageService) SessionsTx(ctx context.Context, tx radio.StorageTx) (radio.SessionStorage, radio.StorageTx, error) {
	db, tx, err := s.tx(ctx, tx)
	if err != nil {
		return nil, nil, err
	}

	storage := SessionStorage{
		handle: handle{db, ctx, "sessions"},
	}
	return storage, tx, nil
}

func (s *StorageService) Queue(ctx context.Context) radio.QueueStorage {
	return QueueStorage{
		handle: handle{s.db, ctx, "queue"},
	}
}

func (s *StorageService) QueueTx(ctx context.Context, tx radio.StorageTx) (radio.QueueStorage, radio.StorageTx, error) {
	db, tx, err := s.tx(ctx, tx)
	if err != nil {
		return nil, nil, err
	}

	storage := QueueStorage{
		handle: handle{db, ctx, "queue"},
	}
	return storage, tx, nil
}

func (s *StorageService) Song(ctx context.Context) radio.SongStorage {
	return SongStorage{
		handle: handle{s.db, ctx, "song"},
	}
}

func (s *StorageService) SongTx(ctx context.Context, tx radio.StorageTx) (radio.SongStorage, radio.StorageTx, error) {
	db, tx, err := s.tx(ctx, tx)
	if err != nil {
		return nil, nil, err
	}

	storage := SongStorage{
		handle: handle{db, ctx, "song"},
	}
	return storage, tx, nil
}

func (s *StorageService) Track(ctx context.Context) radio.TrackStorage {
	return TrackStorage{
		handle: handle{s.db, ctx, "track"},
	}
}

func (s *StorageService) TrackTx(ctx context.Context, tx radio.StorageTx) (radio.TrackStorage, radio.StorageTx, error) {
	db, tx, err := s.tx(ctx, tx)
	if err != nil {
		return nil, nil, err
	}

	storage := TrackStorage{
		handle: handle{db, ctx, "track"},
	}
	return storage, tx, nil
}

func (s *StorageService) Request(ctx context.Context) radio.RequestStorage {
	return RequestStorage{
		handle: handle{s.db, ctx, "request"},
	}
}
func (s *StorageService) RequestTx(ctx context.Context, tx radio.StorageTx) (radio.RequestStorage, radio.StorageTx, error) {
	db, tx, err := s.tx(ctx, tx)
	if err != nil {
		return nil, nil, err
	}

	storage := RequestStorage{
		handle: handle{db, ctx, "request"},
	}
	return storage, tx, nil
}

func (s *StorageService) User(ctx context.Context) radio.UserStorage {
	return UserStorage{
		handle: handle{s.db, ctx, "user"},
	}
}

func (s *StorageService) UserTx(ctx context.Context, tx radio.StorageTx) (radio.UserStorage, radio.StorageTx, error) {
	db, tx, err := s.tx(ctx, tx)
	if err != nil {
		return nil, nil, err
	}

	storage := UserStorage{
		handle: handle{db, ctx, "user"},
	}
	return storage, tx, nil
}

func (s *StorageService) Submissions(ctx context.Context) radio.SubmissionStorage {
	return SubmissionStorage{
		handle: handle{s.db, ctx, "submissions"},
	}
}

func (s *StorageService) SubmissionsTx(ctx context.Context, tx radio.StorageTx) (radio.SubmissionStorage, radio.StorageTx, error) {
	db, tx, err := s.tx(ctx, tx)
	if err != nil {
		return nil, nil, err
	}

	storage := SubmissionStorage{
		handle: handle{db, ctx, "submissions"},
	}
	return storage, tx, nil
}

func (s *StorageService) News(ctx context.Context) radio.NewsStorage {
	return NewsStorage{
		handle: handle{s.db, ctx, "news"},
	}
}

func (s *StorageService) NewsTx(ctx context.Context, tx radio.StorageTx) (radio.NewsStorage, radio.StorageTx, error) {
	db, tx, err := s.tx(ctx, tx)
	if err != nil {
		return nil, nil, err
	}

	storage := NewsStorage{
		handle: handle{db, ctx, "news"},
	}
	return storage, tx, nil
}

func (s *StorageService) Status(ctx context.Context) radio.StatusStorage {
	return StatusStorage{
		handle: handle{s.db, ctx, "status"},
	}
}

func (s *StorageService) Relay(ctx context.Context) radio.RelayStorage {
	return RelayStorage{
		handle: handle{s.db, ctx, "relay"},
	}
}

func (s *StorageService) RelayTx(ctx context.Context, tx radio.StorageTx) (radio.RelayStorage, radio.StorageTx, error) {
	db, tx, err := s.tx(ctx, tx)
	if err != nil {
		return nil, nil, err
	}

	storage := RelayStorage{
		handle: handle{db, ctx, "relay"},
	}
	return storage, tx, nil
}

func (s *StorageService) Search() radio.SearchService {
	return SearchService{
		db: s.db,
	}
}

type extContext interface {
	sqlx.ExecerContext
	sqlx.QueryerContext
	// these are methods on sqlx.binder that is private, we need to implement these
	// to be a sqlx.Ext so that we can use all extensions added by sqlx
	DriverName() string
	Rebind(string) string
	BindNamed(string, interface{}) (string, []interface{}, error)
}

// requireTx returns a handle that uses a transaction, if the handle given already is
// one using a transaction it returns it as-is, otherwise makes a new transaction
func requireTx(h handle) (handle, radio.StorageTx, error) {
	if tx, ok := h.ext.(*sqlx.Tx); ok {
		return h, fakeTx{tx}, nil
	}

	db, ok := h.ext.(*sqlx.DB)
	if !ok {
		zerolog.Ctx(h.ctx).Panic().Any("ext", h.ext).Msg("unknown type")
	}

	tx, err := db.BeginTxx(h.ctx, nil)
	if err != nil {
		return h, nil, err
	}
	h.ext = tx
	return h, tx, nil
}

func namedExecLastInsertId(e sqlx.Ext, query string, arg any) (int64, error) {
	res, err := sqlx.NamedExec(e, query, arg)
	if err != nil {
		return 0, err
	}

	new, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return new, nil
}

// handle is an implementation of sqlx.Execer and sqlx.Queryer that can either use
// a *sqlx.DB directly, or a *sqlx.Tx. It implements these with the *Context equivalents
type handle struct {
	ext extContext
	ctx context.Context

	service string
}

func (h handle) telemetry() (context.Context, func()) {
	span := trace.SpanFromContext(h.ctx)
	ctx, span := span.TracerProvider().Tracer(h.service).Start(h.ctx, h.service)
	return ctx, func() {
		span.End()
	}
}

func (h handle) Exec(query string, args ...interface{}) (sql.Result, error) {
	defer func(start time.Time) {
		zerolog.Ctx(h.ctx).Debug().
			Str("storage_service", h.service).
			Str("query", query).
			Any("arguments", args).
			TimeDiff("execution_time", time.Now(), start).
			Msg("exec")
	}(time.Now())

	ctx, deferFn := h.telemetry()
	defer deferFn()

	return h.ext.ExecContext(ctx, query, args...)
}

func (h handle) Query(query string, args ...interface{}) (*sql.Rows, error) {
	defer func(start time.Time) {
		zerolog.Ctx(h.ctx).Debug().
			Str("storage_service", h.service).
			Str("query", query).
			Any("arguments", args).
			TimeDiff("execution_time", time.Now(), start).
			Msg("query")
	}(time.Now())

	ctx, deferFn := h.telemetry()
	defer deferFn()

	return h.ext.QueryContext(ctx, query, args...)
}

func (h handle) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	defer func(start time.Time) {
		zerolog.Ctx(h.ctx).Debug().
			Str("storage_service", h.service).
			Str("query", query).
			Any("arguments", args).
			TimeDiff("execution_time", time.Now(), start).
			Msg("queryx")
	}(time.Now())

	ctx, deferFn := h.telemetry()
	defer deferFn()

	return h.ext.QueryxContext(ctx, query, args...)
}

func (h handle) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	defer func(start time.Time) {
		zerolog.Ctx(h.ctx).Debug().
			Str("storage_service", h.service).
			Str("query", query).
			Any("arguments", args).
			TimeDiff("execution_time", time.Now(), start).
			Msg("query_rowx")
	}(time.Now())

	ctx, deferFn := h.telemetry()
	defer deferFn()

	return h.ext.QueryRowxContext(ctx, query, args...)
}

func (h handle) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	defer func(start time.Time) {
		zerolog.Ctx(h.ctx).Debug().
			Str("storage_service", h.service).
			Str("query", query).
			Any("arguments", arg).
			TimeDiff("execution_time", time.Now(), start).
			Msg("bind_named")
	}(time.Now())

	return h.ext.BindNamed(query, arg)
}

func (h handle) Rebind(query string) string {
	defer func(start time.Time) {
		zerolog.Ctx(h.ctx).Debug().
			Str("storage_service", h.service).
			Str("query", query).
			TimeDiff("execution_time", time.Now(), start).
			Msg("rebind")
	}(time.Now())

	return h.ext.Rebind(query)
}

func (h handle) DriverName() string {
	return h.ext.DriverName()
}

var _ sqlx.Execer = handle{}
var _ sqlx.Queryer = handle{}
var _ sqlx.Ext = handle{}

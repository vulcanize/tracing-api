package cache

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //postgres driver
	"github.com/vulcanize/ipld-eth-indexer/pkg/postgres"
	"github.com/vulcanize/tracing-api/pkg/eth/tracer"
)

// Service for saving trace data
type Service struct {
	db *sqlx.DB
}

// New create new service
func New() (*Service, error) {
	dbCfg := dbConfig()
	db, err := sqlx.Connect("postgres", postgres.DbConnectionString(dbCfg))
	if err != nil {
		return nil, postgres.ErrDBConnectionFailed(err)
	}
	if dbCfg.MaxOpen > 0 {
		db.SetMaxOpenConns(dbCfg.MaxOpen)
	}
	if dbCfg.MaxIdle > 0 {
		db.SetMaxIdleConns(dbCfg.MaxIdle)
	}
	if dbCfg.MaxLifetime > 0 {
		lifetime := time.Duration(dbCfg.MaxLifetime) * time.Second
		db.SetConnMaxLifetime(lifetime)
	}
	return &Service{db}, nil
}

// TxTraceGraph infomation about callstack
type TxTraceGraph struct {
	TxHash      common.Hash    `json:"txHash"`
	TxIndex     uint64         `json:"txIndex"`
	BlockHash   common.Hash    `json:"blockHash"`
	BlockNumber uint64         `json:"blockNumber"`
	Frames      []tracer.Frame `json:"frames"`
}

func (srv *Service) getTxid(tx *sqlx.Tx, hash string) (uint64, error) {
	stmt := `SELECT id FROM trace.transaction WHERE tx_hash = $1`
	var txID uint64
	return txID, tx.QueryRowx(stmt, hash).Scan(&txID)
}

func (srv *Service) existTx(tx *sqlx.Tx, hash string) (bool, error) {
	_, err := srv.getTxid(tx, hash)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// UpsertTX update information about transaction
func (srv *Service) UpsertTX(tx *sqlx.Tx, txHash string, txIndex uint64, bHash string, bNumber uint64) (uint64, error) {
	stmt := `
		INSERT INTO trace.transaction (tx_hash, index, block_hash, block_number)
			VALUES ($1, $2, $3, $4)
		ON CONFLICT (tx_hash) DO UPDATE SET (index, block_hash, block_number) = ($2, $3, $4)
		RETURNING id
	`
	var txID uint64
	return txID, tx.QueryRowx(stmt, txHash, txIndex, bHash, bNumber).Scan(&txID)
}

// SaveTxTraceGraph save callstack to database
func (srv *Service) SaveTxTraceGraph(data *TxTraceGraph) error {
	tx, err := srv.db.Beginx()
	if err != nil {
		return err
	}

	exist, err := srv.existTx(tx, data.TxHash.Hex())
	if exist || err != nil {
		tx.Rollback()
		return err
	}

	txID, err := srv.UpsertTX(
		tx,
		data.TxHash.Hex(),
		data.TxIndex,
		data.BlockHash.Hex(),
		data.BlockNumber,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	stmt := `
		INSERT INTO trace.call (opcode, src, dst, input, output, value, gas_used, transaction_id)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	for _, frame := range data.Frames {
		_, err := tx.Exec(
			stmt,
			byte(frame.Op),
			frame.From.Hex(),
			frame.To.Hex(),
			[]byte(frame.Input),
			[]byte(frame.Output),
			frame.Value.Uint64(),
			frame.Cost,
			txID,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

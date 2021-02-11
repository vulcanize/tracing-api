-- +goose Up
create table if not exists trace.transaction(
	id serial not null
		constraint transaction_pk
			primary key,
	tx_hash varchar(66) unique not null,
	index integer not null,
	block_number integer not null,
	block_hash varchar(66) not null
);
comment on column trace.transaction.tx_hash is 'Transction hash';
comment on column trace.transaction.index is 'Transaction index';

-- +goose Down
DROP TABLE trace.transaction;
-- +goose Up
create table if not exists trace.call(
  id serial not null
		constraint call_pk
			primary key,
	opcode bytea not null,
	src varchar(66) not null,
	dst varchar(66) not null,
	input bytea not null,
	output bytea not null,
	value numeric not null,
	gas_used numeric not null,
	transaction_id integer not null
		constraint call_transaction_id_fk
			references trace.transaction
);
comment on table trace.call is 'Internal calls';
comment on column trace.call.opcode is 'Solidity Opcode';
comment on column trace.call.src is 'sender of internal tx';
comment on column trace.call.input is 'Input of internal transaction. First 4 bytes are keccak256 hash of method signature';
comment on column trace.call.output is 'Result of internal transaction';

-- +goose Down
DROP TABLE trace.call;
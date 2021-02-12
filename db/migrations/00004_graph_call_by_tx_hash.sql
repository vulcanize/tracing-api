-- +goose Up
CREATE FUNCTION "graphCallByTxHash"(txHash varchar(66)) returns SETOF RECORD
    stable
    language sql
as
$$
SELECT
    gt.block_number,
    gt.block_hash,
    gt.tx_hash,
    gt.index,
    gc.src,
    gc.dst,
    gc.input,
    gc.output,
    gc.value,
    gc.gas_used
FROM trace.graph_transaction gt
         INNER JOIN trace.graph_call gc ON (gc.transaction_id = gt.id)
WHERE gt.tx_hash = $1
$$;
-- +goose Down
DROP FUNCTION "graphCallByTxHash"(varchar(66));
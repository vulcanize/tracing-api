import { SyncCall } from '../generated/BlockNumStorage/BlockNumStorage'
import { SetCall } from '../generated/UintStorage/UintStorage'
import { Sync, StorageSet, Frame } from '../generated/schema'
import { ethereum } from '@graphprotocol/graph-ts';

function toStr(value: ethereum.Value): string {
  switch (value.kind) {
    case ethereum.ValueKind.INT:
    case ethereum.ValueKind.UINT:
      return value.toBigInt().toString();
    case ethereum.ValueKind.STRING:
      return value.toString()
  }
  return "";
}

function stringify(params: ethereum.EventParam[]): string {
  return '[' + params.reduce<string>(function (str, prm, i, prms) {
    let last = prms.length - 1;
    let start = '{';
    let name = '"name":' + '"' + prm.name + '",';
    let value = '"value":' + '"' + toStr(prm.value) + '"';
    let end = '}'
    if (i < last) {
      end += ','
    }
    return str + start + name + value + end;
  }, '') + ']';
}

export function handleSync(call: SyncCall): void {
  let id = call.transaction.hash.toHex();

  let sync = new Sync(id);
  sync.key = call.inputs.key;
  sync.value = call.outputs.value0;
  sync.save();

  let frm = new Frame('sync');
  frm.index = call.transaction.index;
  frm.blockHash = call.block.hash;
  frm.blockNumber = call.block.number;
  frm.from = call.from;
  frm.to = call.to;
  frm.input = stringify(call.inputValues);
  frm.output = stringify(call.outputValues);
  frm.save();
}

export function handleSet(call: SetCall): void {
  let id = call.transaction.hash.toHex();
  let data = new StorageSet(id);
  data.key = call.inputs.key;
  data.value = call.inputs._value;
  data.save();

  let frm = new Frame('set');
  frm.index = call.transaction.index;
  frm.blockHash = call.block.hash;
  frm.blockNumber = call.block.number;
  frm.from = call.from;
  frm.to = call.to;
  frm.input = stringify(call.inputValues);
  frm.output = stringify(call.outputValues);
  frm.save();
}
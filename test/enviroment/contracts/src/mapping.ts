import { SyncCall } from '../generated/BlockNumStorage/BlockNumStorage'
import { SetCall } from '../generated/UintStorage/UintStorage'
import { Sync, StorageSet } from '../generated/schema'


export function handleSync(call: SyncCall): void {
  let id = call.transaction.hash.toHex();
  let sync = new Sync(id);
  sync.key = call.inputs.key;
  sync.value = call.outputs.value0;
  sync.save();
}

export function handleSet(call: SetCall): void {
  let id = call.transaction.hash.toHex();
  let data = new StorageSet(id);
  data.key = call.inputs.key;
  data.value = call.inputs._value;
  data.save();
}
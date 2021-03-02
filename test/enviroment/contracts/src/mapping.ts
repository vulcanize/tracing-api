import { SyncCall } from '../generated/BlockNumStorage/BlockNumStorage'
import { Sync } from '../generated/schema'

export function handleSync(call: SyncCall): void {
  let id = call.transaction.hash.toHex();
  let sync = new Sync(id);
  sync.key = call.inputs.key;
  sync.value = call.outputs.value0;
  sync.save();
}
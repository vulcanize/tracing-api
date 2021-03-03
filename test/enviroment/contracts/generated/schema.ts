// THIS IS AN AUTOGENERATED FILE. DO NOT EDIT THIS FILE DIRECTLY.

import {
  TypedMap,
  Entity,
  Value,
  ValueKind,
  store,
  Address,
  Bytes,
  BigInt,
  BigDecimal
} from "@graphprotocol/graph-ts";

export class Sync extends Entity {
  constructor(id: string) {
    super();
    this.set("id", Value.fromString(id));
  }

  save(): void {
    let id = this.get("id");
    assert(id !== null, "Cannot save Sync entity without an ID");
    assert(
      id.kind == ValueKind.STRING,
      "Cannot save Sync entity with non-string ID. " +
        'Considering using .toHex() to convert the "id" to a string.'
    );
    store.set("Sync", id.toString(), this);
  }

  static load(id: string): Sync | null {
    return store.get("Sync", id) as Sync | null;
  }

  get id(): string {
    let value = this.get("id");
    return value.toString();
  }

  set id(value: string) {
    this.set("id", Value.fromString(value));
  }

  get key(): string | null {
    let value = this.get("key");
    if (value === null || value.kind == ValueKind.NULL) {
      return null;
    } else {
      return value.toString();
    }
  }

  set key(value: string | null) {
    if (value === null) {
      this.unset("key");
    } else {
      this.set("key", Value.fromString(value as string));
    }
  }

  get value(): BigInt | null {
    let value = this.get("value");
    if (value === null || value.kind == ValueKind.NULL) {
      return null;
    } else {
      return value.toBigInt();
    }
  }

  set value(value: BigInt | null) {
    if (value === null) {
      this.unset("value");
    } else {
      this.set("value", Value.fromBigInt(value as BigInt));
    }
  }
}

export class StorageSet extends Entity {
  constructor(id: string) {
    super();
    this.set("id", Value.fromString(id));
  }

  save(): void {
    let id = this.get("id");
    assert(id !== null, "Cannot save StorageSet entity without an ID");
    assert(
      id.kind == ValueKind.STRING,
      "Cannot save StorageSet entity with non-string ID. " +
        'Considering using .toHex() to convert the "id" to a string.'
    );
    store.set("StorageSet", id.toString(), this);
  }

  static load(id: string): StorageSet | null {
    return store.get("StorageSet", id) as StorageSet | null;
  }

  get id(): string {
    let value = this.get("id");
    return value.toString();
  }

  set id(value: string) {
    this.set("id", Value.fromString(value));
  }

  get key(): string | null {
    let value = this.get("key");
    if (value === null || value.kind == ValueKind.NULL) {
      return null;
    } else {
      return value.toString();
    }
  }

  set key(value: string | null) {
    if (value === null) {
      this.unset("key");
    } else {
      this.set("key", Value.fromString(value as string));
    }
  }

  get value(): BigInt | null {
    let value = this.get("value");
    if (value === null || value.kind == ValueKind.NULL) {
      return null;
    } else {
      return value.toBigInt();
    }
  }

  set value(value: BigInt | null) {
    if (value === null) {
      this.unset("value");
    } else {
      this.set("value", Value.fromBigInt(value as BigInt));
    }
  }
}

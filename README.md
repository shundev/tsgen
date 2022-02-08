# TSGEN - TypeScript API client generator

no library, pure Go.

## Build

```
make
```

## Usage

```
./tsgen <file>
```

## Samples

### sample1

```
make sample1
```

↓

```ts
export type SimpleRequest = {
  userID: string;
};

export type MoreComplexRequest = {
  id: string;
  userID: string;
  age: number;
  now: string;
};

export type PointerRequest = {
  age: number | null;
};

export type ArrayRequest = {
  ids: (string)[];
};

export type TypeDefRequest = {
  requestStatus: string;
  requestKind: number | null;
};
```

### sample2

```
make test2
```

↓

```ts
export type FetchRequest = {
  ID: string;
};

export type FetchResponse = {
  ID: string;
  groupID: string;
  name: string;
  shortName: string | null;
  age: number;
  stringSample: string;
  intSample: number;
  strings: (string)[];
  timeSample: string;
  timeNullableSample: string | null;
  timesSample: (string | null)[];
  numbersSample: (number)[];
};
```


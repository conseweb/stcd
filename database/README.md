database
========

[![Build Status](http://img.shields.io/travis/btcsuite/btcd.svg)]
(https://travis-ci.org/btcsuite/btcd)  [![ISC License]
(http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)]
(http://godoc.org/github.com/conseweb/stcd/database)

Package database provides a database interface for the bitcoin block chain and
transactions.

Please note that this package is intended to enable btcd to support different
database backends and is not something that a client can directly access as only
one entity can have the database open at a time (for most database backends),
and that entity will be btcd.

When a client wants programmatic access to the data provided by btcd, they'll
likely want to use the [btcrpcclient](https://github.com/conseweb/btcrpcclient)
package which makes use of the [JSON-RPC API]
(https://github.com/conseweb/stcd/tree/master/docs/json_rpc_api.md).

## Installation and Updating

```bash
$ go get -u github.com/conseweb/stcd/database
```

## Examples

* [CreateDB Example]
  (http://godoc.org/github.com/conseweb/stcd/database#example-CreateDB)  
  Demonstrates creating a new database and inserting the genesis block into it.

* [NewestSha Example]
  (http://godoc.org/github.com/conseweb/stcd/database#example-Db--NewestSha)  
  Demonstrates  querying the database for the most recent best block height and
  hash.

## TODO
- Increase test coverage to 100%

## GPG Verification Key

All official release tags are signed by Conformal so users can ensure the code
has not been tampered with and is coming from the btcsuite developers.  To
verify the signature perform the following:

- Download the public key from the Conformal website at
  https://opensource.conformal.com/GIT-GPG-KEY-conformal.txt

- Import the public key into your GPG keyring:
  ```bash
  gpg --import GIT-GPG-KEY-conformal.txt
  ```

- Verify the release tag with the following command where `TAG_NAME` is a
  placeholder for the specific tag:
  ```bash
  git tag -v TAG_NAME
  ```

## License

Package database is licensed under the [copyfree](http://copyfree.org) ISC
License.

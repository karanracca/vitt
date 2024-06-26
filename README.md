# Vitt

Vitt is hackable plain text double-entry accounting system inspired by [ledger](https://ledger-cli.org/) and follows the [plain text accounting](https://plaintextaccounting.org/) philosophy.

> Plain text accounting is a way of doing bookkeeping and accounting with plain text files and scriptable, command-line-friendly software based on readable, version-controllable, future-proof plain text files.

Advantages of using plain text files for accounting:

- Fast, lightweight and free from vendor lock-in.
- Local and privacy friendly, not cloud-based. Your private financial data remains under your control.
- Scripting-friendly and automatable.

## Installation

### Build from source

Requirments:

- [Go](https://go.dev/)

To build and install an executable, first navigate to Vitt main directory and run command: `go install`

Note: Make sure the Go install path is added to your system's shell path. Refer [Go Docs](https://go.dev/doc/tutorial/compile-install) for more details.

## Usage

TODO...

Note: Vitt does not support comments in the ledger file. The reason for this is to support sorting functionality for transactions. Since, for sorting, Vitt needs to re-write the transactions after they are sorted, its difficult to keep track of comments in the file. Instead, Vitt supports `Notes` for each transaction.

## Rewrite
- Support sub-accounts.
- Support changes to multiple (>2) accounts in single transactions.
- Features
  - Balance Sheet ???
  - Trial Balance ???
  - 

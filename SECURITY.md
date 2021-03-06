# Security

> **IMPORTANT**: If you find a security issue, you can contact our team directly at
security@tendermint.com, or report it to our [bug bounty program](https://hackerone.com/tendermint) on HackerOne. *DO NOT* open a public issue on the repository.

## Bug Bounty

As part of our [Coordinated Vulnerability Disclosure Policy](https://tendermint.com/security), we operate a
[bug bounty program](https://hackerone.com/tendermint) with Hacker One.

See the policy linked above for more details on submissions and rewards and read
this [blog post](https://blog.cosmos.network/bug-bounty-program-for-tendermint-cosmos-833c67693586) for the program scope.

The following is a list of examples of the kinds of bugs we're most interested
in for the Cosmos SDK. See [here](https://github.com/tendermint/tendermint/blob/master/SECURITY.md) for vulnerabilities we are interested
in for Tendermint and other lower-level libraries (eg. [IAVL](https://github.com/tendermint/iavl)).

### Core packages

- [`/baseapp`](https://github.com/DFWallet/anatha/tree/master/baseapp)
- [`/crypto`](https://github.com/DFWallet/anatha/tree/master/crypto)
- [`/types`](https://github.com/DFWallet/anatha/tree/master/types)
- [`/store`](https://github.com/DFWallet/anatha/tree/master/store)

### Modules

- [`x/auth`](https://github.com/DFWallet/anatha/tree/master/x/auth)
- [`x/bank`](https://github.com/DFWallet/anatha/tree/master/x/bank)
- [`x/staking`](https://github.com/DFWallet/anatha/tree/master/x/staking)
- [`x/slashing`](https://github.com/DFWallet/anatha/tree/master/x/slashing)
- [`x/evidence`](https://github.com/DFWallet/anatha/tree/master/x/evidence)
- [`x/distribution`](https://github.com/DFWallet/anatha/tree/master/x/distribution)
- [`x/supply`](https://github.com/DFWallet/anatha/tree/master/x/supply)
- [`x/ibc`](https://github.com/DFWallet/anatha/tree/ibc-alpha/x/ibc) (currently in alpha mode)

We are interested in bugs in other modules, however the above are most likely to
have significant vulnerabilities, due to the complexity / nuance involved. We
also recommend you to read the [specification](https://github.com/DFWallet/anatha/blob/master/docs/building-modules/README.md) of each module before digging into
the code.

### How we process Tx parameters

- Integer operations on tx parameters, especially `sdk.Int` / `sdk.Dec`
- Gas calculation & parameter choices
- Tx signature verification (see [`x/auth/ante`](https://github.com/DFWallet/anatha/tree/master/x/auth/ante))
- Possible Node DoS vectors (perhaps due to gas weighting / non constant timing)

### Handling private keys

- HD key derivation, local and Ledger, and all key-management functionality
- Side-channel attack vectors with our implementations
  - e.g. key exfiltration based on time or memory-access patterns when decrypting privkey

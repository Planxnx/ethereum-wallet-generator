<h3 align="center">
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go-eth" heigth="100" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go-eth" heigth="100" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go-eth" heigth="100" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go-eth" heigth="100" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go-eth" heigth="100" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go-eth" heigth="100" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go-eth" heigth="100" />
</h3>

# eth-wallet-gen

> Ethereum Wallet Generator in Go (golang). Implements the [go-ethereum-hdwallet](https://github.com/miguelmota/go-ethereum-hdwallet).

[![Golang](https://badges.aleen42.com/src/golang.svg)](https://golang.org/)
[![license](https://img.shields.io/badge/license-WTFPL%20--%20Do%20What%20the%20Fuck%20You%20Want%20to%20Public%20License-green.svg)](https://github.com/Planxnx/eth-wallet-gen/blob/main/LICENSE)

#### Easy & Fast Way to generate multiple ETH wallets at once!

- Tiny sizes and Fastest Speed with Golang.
- Embeded Database Supported! (with SQLite3).
- Filtering by Letters Supported! (resovled only addresses that you want).
- Speed up! with customable concurrency numbers.
- Auto generated BIP-39 mnemonic using 128-256 bits of entropy (12, 24 Word Seed Phrase) (Default is 256 bits).
- Default Hierarchical Deterministic Path - m/44'/60'/0'/0 .
- You can benchmark generate speed by DryRun.
- We recommend every user of this application audit and verify any underlying code for its validity and suitability.
- DYOR!

## Installation

```
$ go get -u github.com/Planxnx/eth-wallet-gen
```

## Usage

```
Usage of eth-wallet-gen:
  -n         int    set number of wallets to generate (default 10)
  -db        string set sqlite output path eg. wallets.d
  -c         int    set number of concurrency (default 1)
  -bit       int    set number of entropy bits [128, 256] (default 256)
  -contains  string used to check the given letters present in the given string or not (used for filtered address)
  -strict    bool   strict contains mode (required contains to use)
  -dryrun    bool   generate wallet without result (used for benchamark speed)
```

## Example

**Simple with stdout:**

```txt
$ eth-wallet-gen

===============ETH Wallet Generator===============

10 / 10 | [█████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████] | 100.00% | ? p/s | resovled: 10

Address                                    Seed
------------------------------------------ ----------------------------------------------------------------------------------------------------------------------------------------------------------------
0xc982Dd2E1E9E980A40725367916000793c258aB4 beauty spring lend endless unique thing neutral dignity soup beyond park pact accident mosquito barely tuition memory rather salt spend disease execute list input
0x39F85fb3EFd4Ee8191cE2531AE61d8F2D7C8716f judge match embark suspect wonder sea skull immense ahead galaxy tree recycle lyrics ridge slide physical derive equip clever improve recipe quality tattoo link
0xb3125E8e8Ac9cd54D42DddA9f04D149717AeaaEC pupil ugly festival cruise bar shuffle ball mansion unhappy knee chunk spell fetch rude usage wait picture glue effort wrong angry awake common sample
0x12De8E6eE2A3c82188f12e0ebA984977E49f6E42 hamster treat alcohol thunder reopen demise sick burger beauty reflect bird simple few win female moral paddle version awful develop tell cake that sphere
0x65b25D613C1fAf0a2D48d23757Ed8242f3Ef19f4 flight safe pitch digital main civil pumpkin trick harbor announce drastic nerve super net credit brother swift soldier tonight bonus beyond jelly way video
0x831A0855726f322c61F943b5fb2D212d5451EBb4 olympic people exchange defense lizard maple doctor wool scene ask broken bitter moment sweet help slide off buyer guilt boost trial fame ride method
0x8031c4f8d2A64b1458D749e772A0B365E783A61E increase detect together doll include security insect flash arena deputy orchard poem pact dove atom review wash fashion lonely globe over visa remind toddler
0xD8e65744C6F0DA597b34BDEFAE67e8B1eD1ed9F1 word hour assist warrior hold number right when city off frequent tube enrich steel dentist provide million reject dune ship pudding candy annual almost
0x2CC9df2B0210415Bf15dde9883e89Ad73a548d2f involve people rural grief business case fun injury noodle ritual slender flash predict prosper weird expire remind tank knock anger pool network change style
0x6c6575237168e923F1a2d2Ef6fb7E1cc62404A8D bargain include street tent unique vague animal axis turn hockey scatter attitude naive couple adjust cement deny actor average odor estate happy barrel birth


Resolved Speed: 502.78 w/s
Total Duration: 24.832485ms
Total Wallet Resolved: 10 w

```

**With contains and strict options:**

```txt
$ eth-wallet-gen -n 50000 -contains 0x00,777 -strict
===============ETH Wallet Generator===============

50000 / 50000 | [██████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████] | 100.00% | 447 p/s | resovled: 1

Address                                    Seed
------------------------------------------ ----------------------------------------------------------------------------------------------------------------------------------------------------------------
0x00d6118f53777ecFbF430E40121345D91Fb6018b pen cliff toward mushroom stairs finish filter basic slogan exotic bomb senior drip brush coffee include lady tent finish stable evoke wolf lobster frame


Resolved Speed: 0.01 w/s
Total Duration: 1m52.063956141s
Total Wallet Resolved: 1 w
```

**Speed up with Concuurecny and stored to SQLite3:**

```txt
$ eth-wallet-gen -n 50000 -c 100 -db wallets.db -contain 0x777
===============ETH Wallet Generator===============

50000 / 50000 | [█████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████] | 100.00% | 1674 p/s | resovled: 16

Resolved Speed: 0.44 w/s
Total Duration: 31.177676099s
Total Wallet Resolved: 16 w

Copyright (C) 2021 Planxnx <planxthanee@gmail.com>

```

## Contributing

Pull requests are welcome!

For contributions please create a new branch and submit a pull request for review.

## License

This repo is released under the [WTFPL](http://www.wtfpl.net/) – Do What the Fuck You Want to Public License.

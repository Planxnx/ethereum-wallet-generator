<br>
<h3 align="center">
  
  <img src="https://user-images.githubusercontent.com/37617738/120125455-13de3a00-c1e3-11eb-9a51-707e2dcefdaa.png" alt="eth" height="70" />
   &nbsp&nbsp
  <img src="https://user-images.githubusercontent.com/37617738/120122724-aaefc580-c1d4-11eb-9343-234eb8fb3ab9.png" alt="eth-wallet-gen" height="50" />
</h3>
<br/>
<h3 align="center">
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go-eth" height="50" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go-eth" height="60" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go-eth" height="70" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go-eth" height="80" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go-eth" height="90" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go-eth" height="100" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go-eth" height="110" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go-eth" height="120" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go-eth" height="130" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go-eth" height="140" />
</h3>
<br>

> **Multiple Ethereum Wallet Generator in Go (golang).** Implements the [go-ethereum-hdwallet](https://github.com/miguelmota/go-ethereum-hdwallet).

[![Golang](https://badges.aleen42.com/src/golang.svg)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/Planxnx/eth-wallet-gen)](https://goreportcard.com/report/github.com/Planxnx/eth-wallet-gen)
[![Build Status](https://travis-ci.com/Planxnx/eth-wallet-gen.svg?branch=main)](https://travis-ci.com/Planxnx/eth-wallet-gen)
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/planxthanee/eth-wallet-gen)](https://hub.docker.com/r/planxthanee/eth-wallet-gen)
![Docker Image Size (tag)](https://img.shields.io/docker/image-size/planxthanee/eth-wallet-gen/latest)
![GitHub issues](https://img.shields.io/github/issues/Planxnx/eth-wallet-gen)
[![DeepSource](https://deepsource.io/gh/Planxnx/eth-wallet-gen.svg/?label=active+issues)](https://deepsource.io/gh/Planxnx/eth-wallet-gen/?ref=repository-badge)
[![license](https://img.shields.io/badge/license-WTFPL%20--%20Do%20What%20the%20Fuck%20You%20Want%20to%20Public%20License-green.svg)](https://github.com/Planxnx/eth-wallet-gen/blob/main/LICENSE)

#### Easy & Fast Way to generate multiple Ethereum Wallets at once!

- Tiny sizes and Fastest Speed with Golang 🚀 (required go 1.14 or higher)
- No Go? No Problem! [Docker images](https://hub.docker.com/r/planxthanee/eth-wallet-gen) are provided for you 🐳
- Embeded Database Supported! (with SQLite3).
- Filtering by RegEx or Letters Supported! [regex, prefix, suffix, contains] (resovled only addresses that you want) 🔥
- Speed up!! with customable concurrency numbers ⚡️
- ∞ Infinity wallets generated! (set number to -1 to active infinite loop) ∞
- Auto generated BIP-39 mnemonic using 128-256 bits of entropy (12, 24 Word Seed Phrase) (Default is 256 bits).
- Default Hierarchical Deterministic Path - m/44'/60'/0'/0 .
- You can benchmark generate speed by DryRun 📈
- [Golang Package](https://github.com/Planxnx/eth-wallet-gen/blob/main/generator) for import to your projects.
- We recommend every user of this application audit and verify any underlying code for its validity and suitability. 👮🏻‍♂️

## Installation

<img  align="right" src="https://user-images.githubusercontent.com/37617738/120122855-b1cb0800-c1d5-11eb-9502-8d64bb275337.png" height="140" alt="gopher" />

```console
$ go get -u github.com/Planxnx/eth-wallet-gen
or
$ docker pull planxthanee/eth-wallet-gen:latest
```

## Usage

```console
Usage of eth-wallet-gen:
  -n          int    set number of wallets to generate (default 10) (set number to -1 for Infinite loop ∞)
  -db         string set sqlite output file name eg. wallets.db (db file will create in /db)
  -c          int    set concurrency value (default 1)
  -bit        int    set number of entropy bits [128, 256] (default 256)
  -strict     bool   strict contains mode (required contains to use)
  -contains   string used to check if the given letters are present in the given string (support for multiple characters)
  -prefix     string used to check if the given letters are present in the prefix string (support for single character)
  -suffix     string used to check if the given letters are present in the suffix string (support for single character)
  -regex      string used to check the given letters present in the regex format (eg. ^0x99 or ^0x00)
  -dryrun     bool   generate wallet without a result (used for benchmark speed)
  -compatible bool   logging compatible mode (turn this on to fix logging glitch)
```

## Example

### **Simple with stdout:**

```console
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

### **Using filtered by contains and strict options:**

```console
$ eth-wallet-gen -n 50000 -contains 0x00,777 -strict
===============ETH Wallet Generator===============

50000 / 50000 | [██████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████] | 100.00% | 447 p/s | resovled: 1

Address                                    Seed
------------------------------------------ ----------------------------------------------------------------------------------------------------------------------------------------------------------------
0x00d6118f53777ecFbF430E40121345D91Fb6018b pen cliff toward mushroom stairs finish filter basic slogan exotic bomb senior drip brush coffee include lady tent finish stable evoke wolf lobster frame
0x0077e7063389F6F7fF8216394ed23D259367dbCb fuel wolf embark tip this accident vague face cave echo shift pear between very child draw version face noodle head bubble oblige supreme slot


Resolved Speed: 0.01 w/s
Total Duration: 1m52.063956141s
Total Wallet Resolved: 1 w
```

### **Speed up with by using the Concuurecny and Stored to SQLite3:**

```console
$ eth-wallet-gen -n 50000 -c 500 -db wallet.db -prefix 0x
===============ETH Wallet Generator===============

50000 / 50000 | [█████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████] | 100.00% | 2311 p/s | resovled: 16

Resolved Speed: 2289 w/s
Total Duration: 23.177676099s
Total Wallet Resolved: 50000 w
```

#### **Using Docker (recommend using concurrency for speed up):**

```console
$ docker run --rm -v $PWD:/db planxthanee/eth-wallet-gen -n 1000 -db wallet.db
===============ETH Wallet Generator===============

  100% |██████████████████████████████████████| (50000/50000, 1240 w/s) [1m26s:1s]

Resolved Speed: 1231.89 w/s
Total Duration: 42.78s
Total Wallet Resolved: 50000 w

```

## Contributing ![eth](https://user-images.githubusercontent.com/37617738/120125730-1d1bd680-c1e4-11eb-83ad-45664245cae9.png)

Pull requests are welcome!

For contributions please create a new branch and submit a pull request for review.

## License

This repo is released under the [WTFPL](http://www.wtfpl.net/) – Do What the Fuck You Want to Public License.

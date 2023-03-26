# Ethereum Wallet Generator
<br>
<h3 align="center">
  
  <img src="https://user-images.githubusercontent.com/37617738/120125455-13de3a00-c1e3-11eb-9a51-707e2dcefdaa.png" alt="Ethereum Wallet Generator" height="70" />
   &nbsp&nbsp
  <img src="https://user-images.githubusercontent.com/37617738/120122724-aaefc580-c1d4-11eb-9343-234eb8fb3ab9.png" alt="Ethereum Wallet Generator" height="50" />
</h3>
<br/>
<h3 align="center">
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go ethereum" height="50" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go ethereum" height="60" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go ethereum" height="70" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go ethereum" height="80" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go ethereum" height="90" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go ethereum" height="100" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go ethereum" height="110" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go ethereum" height="120" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go ethereum" height="130" />
  <img src="https://user-images.githubusercontent.com/37617738/120087436-1886ed80-c112-11eb-945f-8065957a1dd0.png" alt="go ethereum" height="140" />
</h3>
<br>

> **Multiple Ethereum and Crypto Wallets Generator written in Go💰** <br>Generate a thousand crypto wallets (public key and mnemonic seed) in a sec ⚡️<br>Find beautiful and awesome wallet addresses 🎨

[![Golang](https://badges.aleen42.com/src/golang.svg)](https://golang.org/)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/1d765b63df4b4266bdcf653d5a024458)](https://www.codacy.com/gh/Planxnx/eth-wallet-gen/dashboard?utm_source=github.com&utm_medium=referral&utm_content=Planxnx/eth-wallet-gen&utm_campaign=Badge_Grade)
![Docker Image Size (tag)](https://img.shields.io/docker/image-size/planxthanee/eth-wallet-gen/latest)
[![Code Analysis & Tests](https://github.com/Planxnx/eth-wallet-gen/actions/workflows/code-analysis.yml/badge.svg)](https://github.com/Planxnx/eth-wallet-gen/actions/workflows/code-analysis.yml)
![GitHub issues](https://img.shields.io/github/issues/Planxnx/eth-wallet-gen)
[![DeepSource](https://deepsource.io/gh/Planxnx/eth-wallet-gen.svg/?label=active+issues)](https://deepsource.io/gh/Planxnx/eth-wallet-gen/?ref=repository-badge)
[![license](https://img.shields.io/badge/license-WTFPL%20--%20Do%20What%20the%20Fuck%20You%20Want%20to%20Public%20License-green.svg)](https://github.com/Planxnx/eth-wallet-gen/blob/main/LICENSE)

## Easy & Fast  Way to generate a thousands beauty Ethereum Wallets ⚡️🎨
![ethereum and crypto wallets generated](https://user-images.githubusercontent.com/37617738/227807144-c1dc59ae-94fd-4fdf-9678-bf8c12e58cd4.png)

- Awesome and Beautiful wallet addresses filtering supported! [regex, prefix, suffix, contains] 🎨
- ∞ Infinite wallet generating! (set number to -1 to active infinite loop) ∞
- Generate word seed phrase with BIP-39 mnemonic (support 12, 24 Word Seed Phrase) (Default is 128 bits for 12 words).
- Embedded Database Supported! (with SQLite3). It's easiest to generate, manage, search a billion wallets without any pain.
- Tiny Sizes and Superior Speed with Golang 🚀 (required go 1.19 or higher)
- No Go? No Problem! [Docker images](https://hub.docker.com/r/planxthanee/eth-wallet-gen) are provided for you 🐳
- Speed things up!! with customizable concurrency numbers ⚡️
- You can benchmark generating speed by setting the `isDryrun` flag 📈
- Default (HD Wallet)Hierarchical Deterministic Path - m/44'/60'/0'/0 .
- [Golang Package](https://github.com/Planxnx/eth-wallet-gen/blob/main/generator) to import within your projects.
- We recommend every user of this application audit and verify every source code in this repository and every imported dependecies for its validity and clearness. 👮🏻‍♂️




## Installation

<img  align="right" src="https://user-images.githubusercontent.com/37617738/120122855-b1cb0800-c1d5-11eb-9502-8d64bb275337.png" height="140" alt="gopher" />

```console
$ go install github.com/Planxnx/eth-wallet-gen
or
$ docker pull planxthanee/eth-wallet-gen:latest
```

## Usage

```console
Usage of eth-wallet-gen:
  -n          int    set number of wallets to generate (default 10) (set number to -1 for Infinite loop ∞)
  -db         string set sqlite output file name eg. wallets.db (db file will create in `/db` folder)
  -c          int    set concurrency value (default 1)
  -bit        int    set number of entropy bits [128 for 12 words, 256 for 24 words] (default 128)
  -strict     bool   strict contains mode, resolve only the addresses that contain all the given letters (required contains to use)
  -contains   string used to check if the given letters are present in the given string (support for multiple characters)
  -prefix     string used to check if the given letters are present in the prefix string (support for single character)
  -suffix     string used to check if the given letters are present in the suffix string (support for single character)
  -regex      string used to check the given letters present in the regex format (eg. ^0x99 or ^0x00)
  -dryrun     bool   generate wallet without a result (used for benchmark speed)
  -compatible bool   logging compatible mode (turn this on to fix logging glitch)
```

## Examples

### **Simple stdout:**

```console
$ eth-wallet-gen

===============ETH Wallet Generator===============

10 / 10 | [█████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████] | 100.00% | 503 p/s | resovled: 10

Address                                    Seed
------------------------------------------ ----------------------------------------------------------------------------------------------------------------------------------------------------------------
0x239ffa10fcd89b2359a5bd8c27c866cfad8eb75a lecture edge end come west mountain van wing zebra trumpet size wool
0x3addecebd6c63be1730205d249681a179e3c768b need decide earth farm punch crush banana unfold income month bread unhappy
0xc4f55b38e6e586cf974eb005e07482fd40274a26 hundred hundred canvas casual staff candy sign travel sort chat travel space
0xe8df7efc452801dc7c75137136c76006bbc2e6d6 gospel father funny pair catalog today champion maple valid feed loop write
0xdf2809a480e29a883a69beb6dedff095984f09eb poet impulse can undo vital stadium tattoo labor trap now blanket assume
0xabc91fd93be63474c14699a1697533410115824c aisle almost miracle coach practice ostrich thing solution ask kiss idle object
0xc9af163bbac66c1c75f3c5f67f758eed1c6077ba funny shift guilt lucky fringe install sugar forget wagon famous inject evoke
0xcf959644c8ee3c20ac9fbecc85610de067cca890 cupboard analyst remove sausage frame engage visual crowd deny boy firm stick
0xa8904e447afb9e0d9b601669aeca53c9b66fe058 sentence skin april wool huge dad bitter loyal perfect again document boring
0x107a842b628b999827730e4543314c6681c72b93 turkey shove mountain yellow ugly shoot crouch donor topple girl polar pelican


Resolved Speed: 502.78 w/s
Total Duration: 24.832485ms
Total Wallet Resolved: 10 w

```

### **Filter with contains options:**

```console
$ eth-wallet-gen -n 100 -contains 0x000,777
===============ETH Wallet Generator===============

100 / 100 | [██████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████] | 100.00% | 5073 p/s | resovled: 2

Address                                    Seed
------------------------------------------ ----------------------------------------------------------------------------------------------------------------------------------------------------------------
0x0004b4067dfe53903d208c032c7bc113ee02c562 economy scene object friend ribbon hungry area thank arrest unfair push hungry
0xac9466777a844f96bf9ac7dafc9bdfeab3e5c8b4 dutch camp banana horse capable border arrest stadium math sport decade help


Resolved Speed: 9.08 w/s
Total Duration: 220.403958ms
Total Wallet Resolved: 2 w
```

### **24 word seed prhase and filter with contains and strict options:**

```console
$ eth-wallet-gen -n 50000 -contains 0x00,777 -strict -bit 256
===============ETH Wallet Generator===============

50000 / 50000 | [██████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████] | 100.00% | 803 p/s | resovled: 2

Address                                    Seed
------------------------------------------ ----------------------------------------------------------------------------------------------------------------------------------------------------------------
0x00d6118f53777ecFbF430E40121345D91Fb6018b pen cliff toward mushroom stairs finish filter basic slogan exotic bomb senior drip brush coffee include lady tent finish stable evoke wolf lobster frame
0x0077e7063389F6F7fF8216394ed23D259367dbCb fuel wolf embark tip this accident vague face cave echo shift pear between very child draw version face noodle head bubble oblige supreme slot


Resolved Speed: 0.02 w/s
Total Duration: 1m2.3567961s
Total Wallet Resolved: 2 w
```

### **Speed up wallet generation by using Concurrency and storing to SQLite3:**

```console
$ eth-wallet-gen -n 50000 -c 12 -db 0x77.db -prefix 0x77
===============ETH Wallet Generator===============

50000 / 50000 | [█████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████] | 100.00% | 5384 p/s | resovled: 50000

Resolved Speed: 5266.02 w/s
Total Duration: 9.4908578s
Total Wallet Resolved: 50000 w
```
![image](https://user-images.githubusercontent.com/37617738/227806706-02a8a7fa-7d2b-43ca-b89b-c21cc51835ff.png)


### **Use Docker (recommend using concurrency for speed up):**

```console
$ docker run --rm -v $PWD:/db planxthanee/eth-wallet-gen -n 50000 -db wallet.db -c 8
===============ETH Wallet Generator===============

  100% |██████████████████████████████████████| (50000/50000, 4651 w/s) [10s:95ms]

Resolved Speed: 4563.50 w/s
Total Duration: 10.9545307s
Total Wallet Resolved: 50000 w

```

## Contributing ![eth](https://user-images.githubusercontent.com/37617738/120125730-1d1bd680-c1e4-11eb-83ad-45664245cae9.png)

Pull requests are welcome!

For contributions please create a new branch and submit a pull request for review.

## License

This repo is released under the [WTFPL](http://www.wtfpl.net/) – Do What the Fuck You Want to Public License.

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
[![license](https://img.shields.io/badge/license-WTFPL%20--%20Do%20What%20the%20Fuck%20You%20Want%20to%20Public%20License-green.svg)](https://github.com/modood/btckeygen/blob/master/LICENSE)

#### Easy & Fast Way to generate multiple ETH wallets at once! 

- Tiny and Fastest wallet generator with Golang.
- Embeded Database Supported! (with SQLite3).
- Adjust speed with customable concurrency numbers.
- Filtering by Letters Supported! (resovled only addresses that you want).
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
  -contain   string used to check the given letters present in the given string or not (used for filtered address)
  -dryrun    bool   generate wallet without result (used for benchamark speed)
```

## Example

**Simple with stdout:**

```txt
$ eth-wallet-gen

===============ETH Wallet Generator===============

10 / 10 | [████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████] | 100.00% | 10 p/s | resovled: 10

Address                                    Seed
------------------------------------------ ----------------------------------------------------------------------------------------------------------------------------------------------------------------
0xbf0835798a3246f689e75725C5BA5fB66C8C0a1d trend shoe loan civil melt please forget spread lava sad kiwi sunset donate expire match joy crew bring fruit chief lion peanut ketchup initial
0x18d9f2DC63F7Ec0C93242bd67153DD09e32073d0 cinnamon total reject merit budget fee boring file charge hawk rice pulp isolate mask small cycle bounce hidden remove desk budget avoid auto wonder
0x725b35b4ebe480f055770773Afb9f1D1254D4473 unfair fruit thought accuse steel confirm iron sort weather orchard rice remove jazz work rebel you tobacco stable follow pig oil slogan potato nominee
0x17637A027cfF8412Fa12964F3FfC16CdB0Fa86E8 rhythm orient coyote level become over whale behave merge company private steel sort galaxy cargo admit rain possible luxury denial good devote raise sausage
0xd06697f0299fA909475FA3Ab202ec08167003f73 craft purse liberty fix monitor glow carry speed price slight bunker crystal find exotic tag drink vessel remember hill digital omit away idea already
0x3620D1fCfc4a88E3ef53b7e4708b225FC0A05e04 help quarter merry romance banner mammal display together velvet denial empower family word silly there custom palm retire call seminar uncle basket range armed
0x7ac8A0F23ab05AD665d7c5787E6a827A8458Aa67 program bleak vivid vote comic they world bind antenna city laundry duck group half cause rookie unlock diesel steak march noise correct sudden sphere
0xa2a1F5667aBD6ae8d0BAD89FfEe367f5d7d1AE57 chat return pledge old win wedding notice teach pattern name bean argue thrive true barely wine traffic bubble crunch always what puppy install off
0xfa5846738CD0db45bF2c1154F4b5BD2245038E5B warrior section between champion curious about tube toy sail symbol grab exhaust ordinary poet universe grit dwarf soap clarify typical chalk solid mask hand
0x5aba5511D8F814B4d9a3877beF1cC0c2C151Ec8d good elevator scheme course wine believe spare august turkey solar label ability arrive dune picture large point fall tail reflect photo develop limb olympic


Resolved Speed: 155.22 w/s
Total Duration: 64.448369ms
Total Wallet Resolved: 10 w

```

**With SQLite3 and options:**

```txt
$ eth-wallet-gen -n 10000 -c 50 -db wallets.db -contain 0x77

===============ETH Wallet Generator===============

10000 / 10000 | [████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████] | 100.00% | 500 p/s | resovled: 42

Resolved Speed: 2.08 w/s
Total Duration: 20.212770838s
Total Wallet Resolved: 42 w

```

## Contributing

Pull requests are welcome!

For contributions please create a new branch and submit a pull request for review.

## License

This repo is released under the [WTFPL](http://www.wtfpl.net/) – Do What the Fuck You Want to Public License.

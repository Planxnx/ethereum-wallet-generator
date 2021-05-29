# eth-wallet-gen

[![license](https://img.shields.io/badge/license-WTFPL%20--%20Do%20What%20the%20Fuck%20You%20Want%20to%20Public%20License-green.svg)](https://github.com/modood/btckeygen/blob/master/LICENSE)

Easy & Fast way to generate multiple eth wallets at once! (Implements the [go-ethereum-hdwallet](https://github.com/miguelmota/go-ethereum-hdwallet))

\*note: HD Path - m/44'/60'/0'/0

## Usage

```
Usage of eth-wallet-gen:
  -n         int    set number of wallets to generate (default 10)
  -c         int    set number of concurrency (default 1)
  -bit       int    set number of entropy bits [128, 256] (default 256)
  -contain   string used to check the given letters present in the given string or not
  -l         bool   show logging data in stdout
```

## Contributing

Pull requests are welcome!

For contributions please create a new branch and submit a pull request for review.

## License

This repo is released under the [WTFPL](http://www.wtfpl.net/) – Do What the Fuck You Want to Public License.

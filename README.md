# QTUM Address Converter

Install

```
go get -u github.com/hayeah/qtumaddr
```

Convert hex address to base58 address (test net):

```
qtumaddr --hex fe59cbc1704e89a698571413a81f0de9d8f00c69 --test
qgkGMtMJbK921UZpwYt3bMhE9d7nNMyUcC
```

Convert hex address to base58 address (main net):

```
qtumaddr --hex fe59cbc1704e89a698571413a81f0de9d8f00c69
QjnsK9sSa8QhJbvTRYEHXapT8M9JHhSLqY
```

Convert base58 address to hex address:

```
qtumaddr --base58 qgkGMtMJbK921UZpwYt3bMhE9d7nNMyUcC
fe59cbc1704e89a698571413a81f0de9d8f00c69
```

```
qtumaddr --base58 QjnsK9sSa8QhJbvTRYEHXapT8M9JHhSLqY
fe59cbc1704e89a698571413a81f0de9d8f00c69
```

## Help

```
usage: qtumaddr [<flags>]

Flags:
      --help           Show context-sensitive help (also try --help-long and --help-man).
  -h, --hex=HEX        address to convert in hex
  -b, --base58=BASE58  address to convert in hex
  -t, --test           generate base58 address for testnet
```

# whois

This is a fork of [whois-go](https://github.com/likexian/whois-go) with tld database from [ruby whois](https://github.com/weppos/whois)

## Overview

whois.go: A golang module for domain whois query.

whois: A golang cli command for domain whois query.

*Works for most domain extensions most of the time.*

## Installation

    go get github.com/glebtv/whois

## Importing

    import (
        "github.com/glebtv/whois"
    )

## Documentation

    func Whois(domain string) (whois.Result)

## Example

    result := whois.Whois("example.com")
    if err == nil {
        spew.Dump(result)
    }

## Whois info parser in Go

Parser is called on whois result automatically.

Also some fields from whois (nameservers, created date, expiration date) have additional guess processing.

Please refer to [whois-parser](https://github.com/likexian/whois-parser)

## LICENSE

Copyright 2018, GlebV

Copyright 2014-2018, Li Kexian

Apache License, Version 2.0

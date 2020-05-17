# Usage

aips (pronounced apes) returns the download URL for Microsoft [Azure IP Ranges](https://www.microsoft.com/en-us/download/details.aspx?id=56519).

Basic:

```
$ go get github.com/jasonwalsh/aips
$ aips
https://download.microsoft.com/download/7/1/D/71D86715-5596-4529-9B13-DA13A5DE5B63/ServiceTags_Public_20200511.json
```

Write to file:

```
$ curl --silent -o ranges.json $(aips)
```

Pipe to [jq](https://stedolan.github.io/jq/):

```
$ curl --silent $(aips) | jq
```

# Contributing

Create a [pull request](https://github.com/jasonwalsh/aips/pulls).

# License

[MIT License](LICENSE)

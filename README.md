# envject

 A command-line tool for injecting environment variables to any file. Useful for any configuration and CI.

Inspired by the functionality of [variable-injector](https://github.com/LucianoPAlmeida/variable-injector) but written in go.

## Installation

### Install from source

```shell
go run main.go --file sample.config.txt --debug true
```

### Grab the binary from the release assets

- https://github.com/michaelhenry/envject/releases

```shell
./envject --file sample.config.txt
```

### Using Homebrew

```shell
brew tap michaelhenry/envject
brew install envject
```

```shell
envject --file sample.config.txt
```


## LICENSE

MIT

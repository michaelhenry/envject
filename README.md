# envject

A command-line tool for injecting environment variables or secrets to any file that has to be compiled especially in iOS.

I have written a [short article](https://www.iamkel.net/posts/6-ios-secrets-handling) regarding this.

And this is inspired by the functionality of [variable-injector](https://github.com/LucianoPAlmeida/variable-injector) but written in go and not dependent with the swift toolchain version.

## Demo

![envject](https://user-images.githubusercontent.com/717992/222741865-e8c51ba1-3660-4c07-a02f-8a630b5a577d.gif)

> So inside the CI, this can be executed during pre-compilation.

## How to use

### Use from source

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

# swift

A demo project for using envject in a swift project.

## To run the demo project

Please see the [Secrets.swift](Sources/Demo/Secrets.swift) and unit test case [DemoTests](Tests/DemoTests/DemoTests.swift).

### Without string obfuscation

```bash
X_API_KEY="Some secret key" go run ../../main.go --file Sources/Demo/Secrets.swift
swift test
```

### With string obfuscation

```bash
X_API_KEY="Some secret key" go run ../../main.go --file Sources/Demo/Secrets.swift --obfuscate-for=swift
swift test
```
// swift-tools-version: 5.7

import PackageDescription

let package = Package(
    name: "Demo",
    products: [
        .library(
            name: "Demo",
            targets: ["Demo"]),
    ],
    targets: [
        .target(
            name: "Demo",
            dependencies: []),
        .testTarget(
            name: "DemoTests",
            dependencies: ["Demo"]),
    ]
)

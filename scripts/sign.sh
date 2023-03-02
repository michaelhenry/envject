#!/bin/bash

# Code sign
/usr/bin/codesign --force -s "$DEV_ID" bin/envject-macos-arm64 -v
/usr/bin/codesign --force -s "$DEV_ID" bin/envject-macos-amd64 -v

# Verify signature
/usr/bin/codesign --verify --deep --strict bin/envject-macos-arm64
/usr/bin/codesign --verify --deep --strict bin/envject-macos-amd64
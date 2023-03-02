#!/bin/bash

echo $MACOS_DEV_CERT | base64 --decode > $DEV_CERT
security create-keychain -p $KEYCHAIN_PW $KEYCHAIN_NAME
security list-keychains -d user -s login.keychain $KEYCHAIN_NAME
security default-keychain -s $KEYCHAIN_NAME
security unlock-keychain -p $KEYCHAIN_PW $KEYCHAIN_NAME
security import $DEV_CERT -k $KEYCHAIN_NAME -P $MACOS_DEV_CERT_PW  -T /usr/bin/codesign
security set-key-partition-list -S apple-tool:,apple:,codesign: -s -k $KEYCHAIN_PW $KEYCHAIN_NAME
security find-identity -v
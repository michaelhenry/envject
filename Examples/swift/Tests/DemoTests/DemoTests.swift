import XCTest
@testable import Demo

final class DemoTests: XCTestCase {
    func testDeobfuscation() throws {
        XCTAssertEqual(Secrets.apiKey, "Some secret key")
    }

    func testDeobfuscationWithSymbols() throws {
        XCTAssertEqual(Secrets.otherApiKey, "j3kj3,,33rz,m234434443flll787483W22345333@32swift testswift test@@@&&")
    }
}

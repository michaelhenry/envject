import XCTest
@testable import Demo

final class DemoTests: XCTestCase {
    func testDeobfuscation() throws {
        XCTAssertEqual(Secrets.apiKey, "Some secret key")
    }
}

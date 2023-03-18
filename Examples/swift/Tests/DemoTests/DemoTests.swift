import XCTest
@testable import Demo

final class DemoTests: XCTestCase {
    func testUnobfuscated() throws {
        XCTAssertEqual(Secrets.apiKey, "Some secret key")
    }
}

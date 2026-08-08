import XCTest
import SwiftTreeSitter
import TreeSitterLsl

final class TreeSitterLslTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_lsl())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading LSL grammar")
    }
}

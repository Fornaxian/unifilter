package unifilter

// Unicode control characters
var controlRunes = []rune{
	// Control block 0
	'\u0000', // NUL: Null character
	'\u0001', // SOH: Start of Heading
	'\u0002', // STX: Start of Text
	'\u0003', // ETX: End-of-text character
	'\u0004', // EOT: End-of-transmission character
	'\u0005', // ENQ: Enquiry character
	'\u0006', // ACK: Acknowledge character
	'\u0007', // BEL: Bell character
	'\u0008', // BS:  Backspace
	'\u0009', // HT:  Horizontal tab
	'\u000A', // LF:  Line feed
	'\u000B', // VT:  Vertical tab
	'\u000C', // FF:  Form feed
	'\u000D', // CR:  Carriage return
	'\u000E', // SO:  Shift Out
	'\u000F', // SI:  Shift In
	'\u0010', // DLE: Data Link Escape
	'\u0011', // DC1: Device Control 1
	'\u0012', // DC2: Device Control 2
	'\u0013', // DC3: Device Control 3
	'\u0014', // DC4: Device Control 4
	'\u0015', // NAK: Negative-acknowledge character
	'\u0016', // SYN: Synchronous Idle
	'\u0017', // ETB: End of Transmission Block
	'\u0018', // CAN: Cancel character
	'\u0019', // EM:  End of Medium
	'\u001A', // SUB: Substitute character
	'\u001B', // ESC: Escape character
	'\u001C', // FS:  File Separator
	'\u001D', // GS:  Group Separator
	'\u001E', // RS:  Record Separator
	'\u001F', // US:  Unit Separator

	'\u007F', // DEL : Delete character, deletes whatever character comes after it

	// Control block 1
	'\u0080', // PAD:  Padding Character
	'\u0081', // HOP:  High Octet Preset
	'\u0082', // BPH:  Break Permitted Here
	'\u0083', // NBH:  No Break Here
	'\u0084', // IND:  Index
	'\u0085', // NEL:  Next Line
	'\u0086', // SSA:  Start of Selected Area
	'\u0087', // ESA:  End of Selected Area
	'\u0088', // HTS:  Character Tabulation Set
	'\u0089', // HTJ:  Character Tabulation with Justification
	'\u008A', // VTS:  Line Tabulation Set
	'\u008B', // PLD:  Partial Line Forward
	'\u008C', // PLU:  Partial Line Backward
	'\u008D', // RI:   Reverse Line Feed
	'\u008E', // SS2:  Single-Shift Two
	'\u008F', // SS3:  Single-Shift Three
	'\u0090', // DCS:  Device Control String
	'\u0091', // PU1:  Private Use 1
	'\u0092', // PU2:  Private Use 2
	'\u0093', // STS:  Set Transmit State
	'\u0094', // CCH:  Cancel character
	'\u0095', // MW:   Message Waiting
	'\u0096', // SPA:  Start of Protected Area
	'\u0097', // EPA:  End of Protected Area
	'\u0098', // SOS:  Start of String
	'\u0099', // SGCI: Single Graphic Character Introducer
	'\u009A', // SCI:  Single Character Intro Introducer
	'\u009B', // CSI:  Control Sequence Introducer
	'\u009C', // ST:   String Terminator
	'\u009D', // OSC:  Operating System Command
	'\u009E', // PM:   Private Message
	'\u009F', // APC:  Application Program Command
}

// Invisible punctuation
var nonPrintingRunes = []rune{
	'\u200B', // ZWSP: Zero-width space
	'\u200C', // ZWNJ: Zero-width non-joiner
	'\u200D', // ZWJ:  Zero-width joiner
	'\u200E', // LRM:  Left-to-right mark
	'\u200F', // RLM:  Right-to-left mark
	'\u2029', // PS:   Paragraph separator
	'\u202A', // LRE:  Left-to-right embedding
	'\u202B', // RLE:  Right-to-left embedding
	'\u202C', // PDF:  Pop directional formatting
	'\u202D', // LRO:  Left-to-right override
	'\u202E', // RLO:  Right-to-left override
	'\uFEFF', // BOM:  Byte order mark
	'\uFFF9', // IAA:  Interlinear annotation anchor
	'\uFFFA', // IAS:  Interlinear annotation separator
	'\uFFFB', // IAT:  Interlinear annotation terminator
	'\uFFFE', // Noncharacter
	'\uFFFF', // Noncharacter
}

// Characters which are illegal in path names
var illegalPathRunesUnix = []rune{
	'\u0000', // NUL, used for determining the end of a path
	'\u002F', // Forward slash, path separator
}

var illegalPathRunesWindows = []rune{
	'\u0000', // NUL, used for determining the end of a path
	'\u0022', // " Quotation mark
	'\u002A', // * Asterisk
	'\u002F', // / Forward slash
	'\u003A', // : Colon
	'\u003C', // < Less than sign
	'\u003E', // > Greater than sign
	'\u003F', // ? Question mark
	'\u005C', // \ Backward slash
	'\u007C', // | Pipe
}

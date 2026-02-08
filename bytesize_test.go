package bytesize

import (
	"fmt"
	"testing"
)

// TestIsValidUnit tests the IsValidUnit function with various unit strings
func TestIsValidUnit(t *testing.T) {
	tests := []struct {
		unit     string
		expected bool
		name     string
	}{
		// Valid short decimal units
		{"b", true, "byte"},
		{"kb", true, "kilobyte"},
		{"mb", true, "megabyte"},
		{"gb", true, "gigabyte"},
		{"tb", true, "terabyte"},
		{"pb", true, "petabyte"},
		{"eb", true, "exabyte"},
		{"zb", true, "zettabyte"},
		{"yb", true, "yottabyte"},
		{"rb", true, "ronnabyte"},
		{"qb", true, "quettabyte"},

		// Valid short binary units
		{"kib", true, "kibibyte"},
		{"mib", true, "mebibyte"},
		{"gib", true, "gibibyte"},
		{"tib", true, "tebibyte"},
		{"pib", true, "pebibyte"},
		{"eib", true, "exbibyte"},
		{"zib", true, "zebibyte"},
		{"yib", true, "yobibyte"},
		{"rib", true, "ronnibyte"},
		{"qib", true, "quettibyte"},

		// Valid long decimal names
		{"byte", true, "byte (long)"},
		{"bytes", true, "bytes (long)"},
		{"kilobyte", true, "kilobyte (long)"},
		{"kilobytes", true, "kilobytes (long)"},
		{"megabyte", true, "megabyte (long)"},
		{"megabytes", true, "megabytes (long)"},
		{"gigabyte", true, "gigabyte (long)"},
		{"gigabytes", true, "gigabytes (long)"},
		{"terabyte", true, "terabyte (long)"},
		{"terabytes", true, "terabytes (long)"},
		{"petabyte", true, "petabyte (long)"},
		{"petabytes", true, "petabytes (long)"},
		{"exabyte", true, "exabyte (long)"},
		{"exabytes", true, "exabytes (long)"},
		{"zettabyte", true, "zettabyte (long)"},
		{"zettabytes", true, "zettabytes (long)"},
		{"yottabyte", true, "yottabyte (long)"},
		{"yottabytes", true, "yottabytes (long)"},
		{"ronnabyte", true, "ronnabyte (long)"},
		{"ronnabytes", true, "ronnabytes (long)"},
		{"quettabyte", true, "quettabyte (long)"},
		{"quettabytes", true, "quettabytes (long)"},

		// Valid long binary names
		{"kibibyte", true, "kibibyte (long)"},
		{"kibibytes", true, "kibibytes (long)"},
		{"mebibyte", true, "mebibyte (long)"},
		{"mebibytes", true, "mebibytes (long)"},
		{"gibibyte", true, "gibibyte (long)"},
		{"gibibytes", true, "gibibytes (long)"},
		{"tebibyte", true, "tebibyte (long)"},
		{"tebibytes", true, "tebibytes (long)"},
		{"pebibyte", true, "pebibyte (long)"},
		{"pebibytes", true, "pebibytes (long)"},
		{"exbibyte", true, "exbibyte (long)"},
		{"exbibytes", true, "exbibytes (long)"},
		{"zebibyte", true, "zebibyte (long)"},
		{"zebibytes", true, "zebibytes (long)"},
		{"yobibyte", true, "yobibyte (long)"},
		{"yobibytes", true, "yobibytes (long)"},
		{"ronnibyte", true, "ronnibyte (long)"},
		{"ronnibytes", true, "ronnibytes (long)"},
		{"quettibyte", true, "quettibyte (long)"},
		{"quettibytes", true, "quettibytes (long)"},

		// Case insensitivity
		{"KB", true, "uppercase"},
		{"Kb", true, "mixed case"},
		{"KiB", true, "mixed case binary"},
		{"MEGABYTE", true, "uppercase long"},
		{"GigaByte", true, "mixed case long"},

		// Whitespace handling
		{" kb", true, "leading space"},
		{"kb ", true, "trailing space"},
		{"  kb  ", true, "both spaces"},
		{"\tkb", true, "tab"},
		{"\tKiB\t", true, "tab and mixed case"},

		// Invalid units
		{"x", false, "invalid single character"},
		{"xb", false, "invalid unit"},
		{"kilobit", false, "kilobit (not supported)"},
		{"megabit", false, "megabit (not supported)"},
		{"k", false, "k without b"},
		{"ki", false, "ki without b"},
		{"", false, "empty string"},
		{"   ", false, "only spaces"},
		{"123", false, "just number"},
		{"kb2", false, "unit with number"},
		{"gigabytee", false, "typo"},
		{"kilobytes2", false, "long name with number"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidUnit(tt.unit)
			if result != tt.expected {
				t.Errorf("IsValidUnit(%q) = %v, expected %v", tt.unit, result, tt.expected)
			}
		})
	}
}

// TestParseBasicUnits tests parsing of basic byte units
func TestParseBasicUnits(t *testing.T) {
	tests := []struct {
		input    string
		expected Bytes
	}{
		// Single byte
		{"1b", B},
		{"1 b", B},
		{"  1  b  ", B},

		// Decimal units
		{"1 kb", KB},
		{"1 KB", KB},
		{"1 Kb", KB},
		{"2 MB", Bytes(Uint128(MB).Mul64(2))},
		{"3 GB", Bytes(Uint128(GB).Mul64(3))},
		{"4 TB", Bytes(Uint128(TB).Mul64(4))},
		{"5 PB", Bytes(Uint128(PB).Mul64(5))},
		{"6 EB", Bytes(Uint128(EB).Mul64(6))},
		{"7 ZB", Bytes(Uint128(ZB).Mul64(7))},
		{"8 YB", Bytes(Uint128(YB).Mul64(8))},
		{"9 RB", Bytes(Uint128(RB).Mul64(9))},
		{"10 QB", Bytes(Uint128(QB).Mul64(10))},

		// Binary units
		{"1 KiB", KiB},
		{"1 kib", KiB},
		{"2 MiB", Bytes(Uint128(MiB).Mul64(2))},
		{"3 GiB", Bytes(Uint128(GiB).Mul64(3))},
		{"4 TiB", Bytes(Uint128(TiB).Mul64(4))},
		{"5 PiB", Bytes(Uint128(PiB).Mul64(5))},
		{"6 EiB", Bytes(Uint128(EiB).Mul64(6))},
		{"7 ZiB", Bytes(Uint128(ZiB).Mul64(7))},
		{"8 YiB", Bytes(Uint128(YiB).Mul64(8))},
		{"9 RiB", Bytes(Uint128(RiB).Mul64(9))},
		{"10 QiB", Bytes(Uint128(QiB).Mul64(10))},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Fatalf("Parse(%q) error = %v, want nil", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("Parse(%q) = {%d, %d}, want {%d, %d}",
					tt.input, result.Lo, result.Hi, tt.expected.Lo, tt.expected.Hi)
			}
		})
	}
}

// TestParseLongNames tests parsing with long unit names
func TestParseLongNames(t *testing.T) {
	tests := []struct {
		input    string
		expected Bytes
	}{
		// Byte variations
		{"1 byte", B},
		{"1 bytes", B},
		{"2 bytes", Bytes(Uint128(B).Mul64(2))},

		// Long decimal names
		{"1 kilobyte", KB},
		{"1 kilobytes", KB},
		{"1 megabyte", MB},
		{"1 megabytes", MB},
		{"1 gigabyte", GB},
		{"1 gigabytes", GB},
		{"1 terabyte", TB},
		{"1 terabytes", TB},
		{"1 petabyte", PB},
		{"1 petabytes", PB},
		{"1 exabyte", EB},
		{"1 exabytes", EB},
		{"1 zettabyte", ZB},
		{"1 zettabytes", ZB},
		{"1 yottabyte", YB},
		{"1 yottabytes", YB},
		{"1 ronnabyte", RB},
		{"1 ronnabytes", RB},
		{"1 quettabyte", QB},
		{"1 quettabytes", QB},

		// Long binary names
		{"1 kibibyte", KiB},
		{"1 kibibytes", KiB},
		{"1 mebibyte", MiB},
		{"1 mebibytes", MiB},
		{"1 gibibyte", GiB},
		{"1 gibibytes", GiB},
		{"1 tebibyte", TiB},
		{"1 tebibytes", TiB},
		{"1 pebibyte", PiB},
		{"1 pebibytes", PiB},
		{"1 exbibyte", EiB},
		{"1 exbibytes", EiB},
		{"1 zebibyte", ZiB},
		{"1 zebibytes", ZiB},
		{"1 yobibyte", YiB},
		{"1 yobibytes", YiB},
		{"1 ronnibyte", RiB},
		{"1 ronnibytes", RiB},
		{"1 quettibyte", QiB},
		{"1 quettibytes", QiB},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Fatalf("Parse(%q) error = %v, want nil", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("Parse(%q) = {%d, %d}, want {%d, %d}",
					tt.input, result.Lo, result.Hi, tt.expected.Lo, tt.expected.Hi)
			}
		})
	}
}

// TestParseFloatingPoint tests parsing with floating point numbers
func TestParseFloatingPoint(t *testing.T) {
	tests := []struct {
		input   string
		checkFn func(result Bytes) bool
		name    string
	}{
		{
			input: "1.5 KB",
			checkFn: func(result Bytes) bool {
				// 1.5 KB = 1500 bytes
				expected := uint64(1500)
				return result.Lo == expected && result.Hi == 0
			},
			name: "1.5 KB should be 1500 bytes",
		},
		{
			input: "0.5 KB",
			checkFn: func(result Bytes) bool {
				// 0.5 KB = 500 bytes
				expected := uint64(500)
				return result.Lo == expected && result.Hi == 0
			},
			name: "0.5 KB should be 500 bytes",
		},
		{
			input: "2.5 MB",
			checkFn: func(result Bytes) bool {
				// 2.5 MB = 2.5 * 1e6 = 2500000 bytes
				expected := uint64(2500000)
				return result.Lo == expected && result.Hi == 0
			},
			name: "2.5 MB should be 2500000 bytes",
		},
		{
			input: "0.1 GB",
			checkFn: func(result Bytes) bool {
				// 0.1 GB = 1e8 bytes
				expected := uint64(100000000)
				return result.Lo == expected && result.Hi == 0
			},
			name: "0.1 GB should be 100000000 bytes",
		},
		{
			input: "3.14159 KB",
			checkFn: func(result Bytes) bool {
				// Should be approximately 3141 bytes (3.14159 * 1000)
				// Due to float precision, we allow a small range
				return result.Lo >= 3141 && result.Lo <= 3142
			},
			name: "3.14159 KB should be approximately 3141 bytes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Fatalf("Parse(%q) error = %v, want nil", tt.input, err)
			}
			if !tt.checkFn(result) {
				t.Errorf("Parse(%q) = {%d, %d}, validation failed: %s",
					tt.input, result.Lo, result.Hi, tt.name)
			}
		})
	}
}

// TestParseWhitespace tests parsing with various whitespace patterns
func TestParseWhitespace(t *testing.T) {
	tests := []struct {
		input    string
		expected Bytes
	}{
		{"1b", B},
		{"1 b", B},
		{"1  b", B},
		{"  1b", B},
		{"1b  ", B},
		{"  1  b  ", B},
		{"\t1\tMB", MB},
		{"\n1\nGB", GB},
		{" \t 5 \n KB \t ", Bytes(Uint128(KB).Mul64(5))},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("input=%q", tt.input), func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Fatalf("Parse(%q) error = %v, want nil", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("Parse(%q) = {%d, %d}, want {%d, %d}",
					tt.input, result.Lo, result.Hi, tt.expected.Lo, tt.expected.Hi)
			}
		})
	}
}

// TestParseCaseInsensitivity tests that parsing is case-insensitive
func TestParseCaseInsensitivity(t *testing.T) {
	tests := []struct {
		input    string
		expected Bytes
	}{
		{"1b", B},
		{"1B", B},
		{"1 kb", KB},
		{"1 KB", KB},
		{"1 Kb", KB},
		{"1 kB", KB},
		{"1 kib", KiB},
		{"1 KIB", KiB},
		{"1 KiB", KiB},
		{"1 megabyte", MB},
		{"1 MEGABYTE", MB},
		{"1 MegaByte", MB},
		{"1 KIBIBYTE", KiB},
		{"1 Kibibyte", KiB},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Fatalf("Parse(%q) error = %v, want nil", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("Parse(%q) = {%d, %d}, want {%d, %d}",
					tt.input, result.Lo, result.Hi, tt.expected.Lo, tt.expected.Hi)
			}
		})
	}
}

// TestParseZeroValues tests parsing of zero values
func TestParseZeroValues(t *testing.T) {
	tests := []struct {
		input    string
		expected Bytes
	}{
		{"0 b", Bytes{}},
		{"0 B", Bytes{}},
		{"0 KB", Bytes{}},
		{"0 MB", Bytes{}},
		{"0 GiB", Bytes{}},
		{"0.0 GB", Bytes{}},
		{"0 bytes", Bytes{}},
		{"0 megabytes", Bytes{}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Fatalf("Parse(%q) error = %v, want nil", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("Parse(%q) = {%d, %d}, want {%d, %d}",
					tt.input, result.Lo, result.Hi, tt.expected.Lo, tt.expected.Hi)
			}
		})
	}
}

// TestParseLargeValues tests parsing of large values
func TestParseLargeValues(t *testing.T) {
	tests := []struct {
		input   string
		checkFn func(result Bytes) bool
		name    string
	}{
		{
			input: "1000 YB",
			checkFn: func(result Bytes) bool {
				// 1000 YB is a huge number, just verify it's not zero
				return result.Lo > 0 || result.Hi > 0
			},
			name: "1000 YB should be a large non-zero value",
		},
		{
			input: "999 QB",
			checkFn: func(result Bytes) bool {
				// 999 QB is an extremely large number
				return result.Lo > 0 || result.Hi > 0
			},
			name: "999 QB should be a large non-zero value",
		},
		{
			input: "1 QiB",
			checkFn: func(result Bytes) bool {
				// QiB = 2^100
				return result.Lo > 0 || result.Hi > 0
			},
			name: "1 QiB should be a large non-zero value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Fatalf("Parse(%q) error = %v, want nil", tt.input, err)
			}
			if !tt.checkFn(result) {
				t.Errorf("Parse(%q) validation failed: %s", tt.input, tt.name)
			}
		})
	}
}

// TestParseErrors tests error cases
func TestParseErrors(t *testing.T) {
	tests := []struct {
		input       string
		expectedErr string
	}{
		// Empty/whitespace strings
		{"", "empty string"},
		{" ", "empty string"},
		{"\t", "empty string"},
		{"\n", "empty string"},

		// Invalid formats
		{"abc", "invalid number"},
		{"MB", "invalid number"},
		{"1.2.3 KB", "multiple decimal points"},
		{" . MB", "invalid number"},

		// Negative values
		{"-1 B", "negative value"},
		{"-5 MB", "negative value"},
		{"-0.1 GB", "negative value"},

		// Unknown units
		{"10 XB", "unknown unit"},
		{"5 unknown", "unknown unit"},
		{"100 zz", "unknown unit"},
		{"1 kilobit", "unknown unit"},
		{"1 megabit", "unknown unit"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("input=%q", tt.input), func(t *testing.T) {
			result, err := Parse(tt.input)
			if err == nil {
				t.Fatalf("Parse(%q) should have errored, got {%d, %d}", tt.input, result.Lo, result.Hi)
			}
		})
	}
}

// TestParseBoundaryValues tests boundary conditions
func TestParseBoundaryValues(t *testing.T) {
	tests := []struct {
		input   string
		checkFn func(result Bytes) bool
		name    string
	}{
		{
			input: "1 B",
			checkFn: func(result Bytes) bool {
				return result.Lo == 1 && result.Hi == 0
			},
			name: "1 byte",
		},
		{
			input: "18446744073709551615 B",
			checkFn: func(result Bytes) bool {
				return result.Lo == 18446744073709551615 && result.Hi == 0
			},
			name: "MaxUint64 bytes",
		},
		{
			input: "0.00001 KB",
			checkFn: func(result Bytes) bool {
				// 0.00001 KB = 0.01 bytes, should round to 0
				return result.Lo == 0 && result.Hi == 0
			},
			name: "Very small fractional value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Fatalf("Parse(%q) error = %v, want nil", tt.input, err)
			}
			if !tt.checkFn(result) {
				t.Errorf("Parse(%q) = %v, validation failed: %s", tt.input, result, tt.name)
			}
		})
	}
}

// TestParseConsistency tests that parsing and conversion are consistent
func TestParseConsistency(t *testing.T) {
	tests := []struct {
		input    string
		multiple uint64
		unit     Bytes
	}{
		{"5 KB", 5, KB},
		{"10 MB", 10, MB},
		{"3 GiB", 3, GiB},
		{"7 TB", 7, TB},
		{"2 KiB", 2, KiB},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			parsed, err := Parse(tt.input)
			if err != nil {
				t.Fatalf("Parse(%q) error = %v", tt.input, err)
			}

			// Calculate expected value
			expected := Bytes(Uint128(tt.unit).Mul64(tt.multiple))

			if parsed != expected {
				t.Errorf("Parse(%q) = {%d, %d}, want {%d, %d}",
					tt.input, parsed.Lo, parsed.Hi, expected.Lo, expected.Hi)
			}
		})
	}
}

// FuzzParse is a fuzzing test for the Parse function
func FuzzParse(f *testing.F) {
	// Add seed corpus
	seedInputs := []string{
		"0 b",
		"1 B",
		"10 KB",
		"100 MB",
		"1000 GB",
		"1 KiB",
		"1 kilobyte",
		"1.5 MB",
		"999 QB",
		"0.001 KB",
		"1e2 MB",
		"",
		"invalid",
		"-5 MB",
		"1.2.3 KB",
		"unknown unit",
		"1 2 3 MB",
		"   10   MB   ",
		"\t50\tGB\n",
	}

	for _, seed := range seedInputs {
		f.Add(seed)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		// Call Parse - it should not panic
		result, err := Parse(input)

		// If it succeeded, verify the result is valid
		if err == nil {
			// Result should be a valid Bytes value (no corruption)
			// Just ensure the function didn't panic and returned something
			_ = result

			// Verify negative results are not possible for valid parse
			// (Bytes is unsigned)
		}
		// If it errored, that's also fine - just ensure no panic occurred
	})
}

// BenchmarkParseBasic benchmarks parsing simple byte values
func BenchmarkParseBasic(b *testing.B) {
	for b.Loop() {
		Parse("100 B")
	}
}

// BenchmarkParseDecimal benchmarks parsing decimal unit (KB)
func BenchmarkParseDecimal(b *testing.B) {
	for b.Loop() {
		Parse("512 MB")
	}
}

// BenchmarkParseBinary benchmarks parsing binary unit (KiB)
func BenchmarkParseBinary(b *testing.B) {
	for b.Loop() {
		Parse("256 GiB")
	}
}

// BenchmarkParseLongName benchmarks parsing with long unit name
func BenchmarkParseLongName(b *testing.B) {
	for b.Loop() {
		Parse("10 gigabyte")
	}
}

// BenchmarkParseFloatingPoint benchmarks parsing floating-point values
func BenchmarkParseFloatingPoint(b *testing.B) {
	for b.Loop() {
		Parse("3.14159 KB")
	}
}

// BenchmarkParseLargeNumber benchmarks parsing large numeric values
func BenchmarkParseLargeNumber(b *testing.B) {
	for b.Loop() {
		Parse("999999999 TB")
	}
}

// BenchmarkParseWithWhitespace benchmarks parsing with extra whitespace
func BenchmarkParseWithWhitespace(b *testing.B) {
	for b.Loop() {
		Parse("   500   MB   ")
	}
}

// BenchmarkParseNoWhitespace benchmarks parsing without any whitespace
func BenchmarkParseNoWhitespace(b *testing.B) {
	for b.Loop() {
		Parse("512GB")
	}
}

// BenchmarkParseLargeUnit benchmarks parsing with very large units (QB)
func BenchmarkParseLargeUnit(b *testing.B) {
	for b.Loop() {
		Parse("5 QB")
	}
}

// BenchmarkParseError benchmarks the overhead of error handling (invalid input)
func BenchmarkParseError(b *testing.B) {
	for b.Loop() {
		Parse("invalid")
	}
}

// BenchmarkParseParallel benchmarks Parse function with parallel execution
func BenchmarkParseParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Parse("256 MB")
		}
	})
}

// ============ Format Function Tests ============

// TestFormatBasicBytes tests formatting basic byte values
func TestFormatBasicBytes(t *testing.T) {
	tests := []struct {
		input   Bytes
		checkFn func(string) bool
		name    string
	}{
		{
			input: B,
			checkFn: func(s string) bool {
				return s == "1.00 B"
			},
			name: "1 B",
		},
		{
			input: Bytes(Uint128(B).Mul64(512)),
			checkFn: func(s string) bool {
				return s == "512.00 B"
			},
			name: "512 B",
		},
		{
			input: Bytes(Uint128(B).Mul64(1024)),
			checkFn: func(s string) bool {
				return s == "1.02 KB"
			},
			name: "1024 bytes (should be 1.02 KB)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.Format()
			if err != nil {
				t.Fatalf("Format() error = %v", err)
			}
			if !tt.checkFn(result) {
				t.Errorf("Format() = %q, validation failed for test: %s", result, tt.name)
			}
		})
	}
}

// TestFormatDecimalUnits tests formatting with decimal (SI) units
func TestFormatDecimalUnits(t *testing.T) {
	tests := []struct {
		input    Bytes
		expected string
		name     string
	}{
		{KB, "1.00 KB", "1 KB"},
		{MB, "1.00 MB", "1 MB"},
		{GB, "1.00 GB", "1 GB"},
		{TB, "1.00 TB", "1 TB"},
		{PB, "1.00 PB", "1 PB"},
		{EB, "1.00 EB", "1 EB"},
		{ZB, "1.00 ZB", "1 ZB"},
		{YB, "1.00 YB", "1 YB"},
		{RB, "1.00 RB", "1 RB"},
		{QB, "1.00 QB", "1 QB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.Format(WithDecimalUnits(true))
			if err != nil {
				t.Fatalf("Format() error = %v", err)
			}
			if result != tt.expected {
				t.Errorf("Format() = %q, want %q", result, tt.expected)
			}
		})
	}
}

// TestFormatBinaryUnits tests formatting with binary (IEC) units
func TestFormatBinaryUnits(t *testing.T) {
	tests := []struct {
		input    Bytes
		expected string
		name     string
	}{
		{KiB, "1.00 KiB", "1 KiB"},
		{MiB, "1.00 MiB", "1 MiB"},
		{GiB, "1.00 GiB", "1 GiB"},
		{TiB, "1.00 TiB", "1 TiB"},
		{PiB, "1.00 PiB", "1 PiB"},
		{EiB, "1.00 EiB", "1 EiB"},
		{ZiB, "1.00 ZiB", "1 ZiB"},
		{YiB, "1.00 YiB", "1 YiB"},
		{RiB, "1.00 RiB", "1 RiB"},
		{QiB, "1.00 QiB", "1 QiB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.Format(WithDecimalUnits(false))
			if err != nil {
				t.Fatalf("Format() error = %v", err)
			}
			if result != tt.expected {
				t.Errorf("Format() = %q, want %q", result, tt.expected)
			}
		})
	}
}

// TestFormatLongNames tests formatting with long unit names
func TestFormatLongNames(t *testing.T) {
	tests := []struct {
		input   Bytes
		checkFn func(string) bool
		name    string
	}{
		{
			input: KB,
			checkFn: func(s string) bool {
				return s == "1.00 Kilobyte"
			},
			name: "1 KB with long name",
		},
		{
			input: MB,
			checkFn: func(s string) bool {
				return s == "1.00 Megabyte"
			},
			name: "1 MB with long name",
		},
		{
			input: GiB,
			checkFn: func(s string) bool {
				return s == "1.07 Gigabytes"
			},
			name: "1 GiB with long name (should be 1.07 Gigabytes)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.Format(WithLongUnits(true))
			if err != nil {
				t.Fatalf("Format() error = %v", err)
			}
			if !tt.checkFn(result) {
				t.Errorf("Format() = %q ; input = %v, validation failed for test: %s", result, Uint128(tt.input), tt.name)
			}
		})
	}
}

// TestFormatShortNames tests formatting with short unit names (default)
func TestFormatShortNames(t *testing.T) {
	tests := []struct {
		input    Bytes
		expected string
		name     string
	}{
		{KB, "1.00 KB", "1 KB short name"},
		{MB, "1.00 MB", "1 MB short name"},
		{GiB, "1.07 GB", "1 GiB short name (should be 1.07 GB)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.Format(WithLongUnits(false))
			if err != nil {
				t.Fatalf("Format() error = %v", err)
			}
			if result != tt.expected {
				t.Errorf("Format() = %q, input = %v, want %q", result, Uint128(tt.input), tt.expected)
			}
		})
	}
}

// TestFormatCustomFormatString tests formatting with custom format strings
func TestFormatCustomFormatString(t *testing.T) {
	tests := []struct {
		input     Bytes
		formatStr string
		expected  string
		name      string
	}{
		{
			input:     Bytes(Uint128(KB).Mul64(10)),
			formatStr: "%.0f %s",
			expected:  "10 KB",
			name:      "custom format %.0f %s",
		},
		{
			input:     Bytes(Uint128(MB).Mul64(5)),
			formatStr: "%.3f %s",
			expected:  "5.000 MB",
			name:      "custom format %.3f %s",
		},
		{
			input:     Bytes(Uint128(GB).Mul64(2)),
			formatStr: "%[2]s: %.1[1]f",
			expected:  "GB: 2.0",
			name:      "custom format with unit first",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.Format(WithFormatString(tt.formatStr))
			if err != nil {
				t.Fatalf("Format() error = %v", err)
			}
			if result != tt.expected {
				t.Errorf("Format() = %q, want %q", result, tt.expected)
			}
		})
	}
}

// TestFormatForcedUnit tests formatting with a forced unit
func TestFormatForcedUnit(t *testing.T) {
	tests := []struct {
		input      Bytes
		forcedUnit Bytes
		expected   string
		name       string
	}{
		{
			input:      Bytes(Uint128(B).Mul64(500)),
			forcedUnit: KB,
			expected:   "0.50 KB",
			name:       "500 bytes forced to KB",
		},
		{
			input:      Bytes(Uint128(MB).Mul64(1024)),
			forcedUnit: GB,
			expected:   "1.02 GB",
			name:       "1024 MB forced to GB",
		},
		{
			input:      Bytes(Uint128(B).Mul64(1000000)),
			forcedUnit: MB,
			expected:   "1.00 MB",
			name:       "1000000 bytes forced to MB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.Format(WithForcedUnit(tt.forcedUnit))
			if err != nil {
				t.Fatalf("Format() error = %v", err)
			}
			if result != tt.expected {
				t.Errorf("Format() = %q, want %q", result, tt.expected)
			}
		})
	}
}

// TestFormatCombinedOptions tests using multiple format options together
func TestFormatCombinedOptions(t *testing.T) {
	tests := []struct {
		input    Bytes
		opts     []FormatOption
		expected string
		name     string
	}{
		{
			input:    KB,
			opts:     []FormatOption{WithLongUnits(true), WithDecimalUnits(true)},
			expected: "1.00 Kilobyte",
			name:     "KB with long decimal names",
		},
		{
			input:    GiB,
			opts:     []FormatOption{WithLongUnits(true), WithDecimalUnits(false)},
			expected: "1.00 Gibibyte",
			name:     "GiB with long binary names",
		},
		{
			input:    Bytes(Uint128(MB).Mul64(512)),
			opts:     []FormatOption{WithFormatString("%.1f %s"), WithForcedUnit(MB)},
			expected: "512.0 MB",
			name:     "512 MB with custom format and forced unit",
		},
		{
			input:    Bytes(Uint128(KB).Mul64(2560)),
			opts:     []FormatOption{WithFormatString("%.0f %s"), WithForcedUnit(MB), WithDecimalUnits(true)},
			expected: "3 MB",
			name:     "2560 KB = 2.56 MB rounded",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.Format(tt.opts...)
			if err != nil {
				t.Fatalf("Format() error = %v", err)
			}
			if result != tt.expected {
				t.Errorf("Format() = %q, want %q", result, tt.expected)
			}
		})
	}
}

// TestFormatErrors tests error handling in format options
func TestFormatErrors(t *testing.T) {
	tests := []struct {
		opts []FormatOption
		name string
	}{
		{
			opts: []FormatOption{WithFormatString("")},
			name: "empty format string",
		},
		{
			opts: []FormatOption{WithForcedUnit(Bytes{12345, 67890})},
			name: "invalid forced unit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := B.Format(tt.opts...)
			if err == nil {
				t.Errorf("Format() should have errored, got %q", result)
			}
		})
	}
}

// TestFormatAutoUnitSelection tests automatic unit selection for various sizes
func TestFormatAutoUnitSelection(t *testing.T) {
	tests := []struct {
		input   Bytes
		checkFn func(string) bool
		name    string
	}{
		{
			input: Bytes(Uint128(B).Mul64(512)),
			checkFn: func(s string) bool {
				return s == "512.00 B"
			},
			name: "512 bytes should use B",
		},
		{
			input: Bytes(Uint128(KB).Mul64(1)),
			checkFn: func(s string) bool {
				return s == "1.00 KB"
			},
			name: "1 KB",
		},
		{
			input: Bytes(Uint128(MB).Mul64(5)),
			checkFn: func(s string) bool {
				return s == "5.00 MB"
			},
			name: "5 MB",
		},
		{
			input: Bytes(Uint128(GB).Mul64(100)),
			checkFn: func(s string) bool {
				return s == "100.00 GB"
			},
			name: "100 GB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.Format()
			if err != nil {
				t.Fatalf("Format() error = %v", err)
			}
			if !tt.checkFn(result) {
				t.Errorf("Format() = %q, validation failed for test: %s", result, tt.name)
			}
		})
	}
}

// TestFormatPluralization tests correct pluralization of unit names
func TestFormatPluralization(t *testing.T) {
	tests := []struct {
		input   Bytes
		opts    []FormatOption
		checkFn func(string) bool
		name    string
	}{
		{
			input: B,
			opts:  []FormatOption{WithLongUnits(true)},
			checkFn: func(s string) bool {
				return s == "1.00 Byte"
			},
			name: "1 byte (singular)",
		},
		{
			input: Bytes(Uint128(B).Mul64(2)),
			opts:  []FormatOption{WithLongUnits(true)},
			checkFn: func(s string) bool {
				return s == "2.00 Bytes"
			},
			name: "2 bytes (plural)",
		},
		{
			input: KB,
			opts:  []FormatOption{WithLongUnits(true)},
			checkFn: func(s string) bool {
				return s == "1.00 Kilobyte"
			},
			name: "1 kilobyte (singular)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.Format(tt.opts...)
			if err != nil {
				t.Fatalf("Format() error = %v", err)
			}
			if !tt.checkFn(result) {
				t.Errorf("Format() = %q, validation failed for test: %s", result, tt.name)
			}
		})
	}
}

// TestFormatZeroValue tests formatting of zero value
func TestFormatZeroValue(t *testing.T) {
	result, err := Bytes{}.Format()
	if err != nil {
		t.Fatalf("Format() error = %v", err)
	}
	if result == "" {
		t.Errorf("Format() should not return empty string for zero value")
	}
}

// TestFormatLargeValues tests formatting of very large values
func TestFormatLargeValues(t *testing.T) {
	tests := []struct {
		input   Bytes
		checkFn func(string) bool
		name    string
	}{
		{
			input: QB,
			checkFn: func(s string) bool {
				return s == "1.00 QB"
			},
			name: "1 QB",
		},
		{
			input: QiB,
			checkFn: func(s string) bool {
				return s == "1.27 QB"
			},
			name: "1 QiB (should be 1.27 QB)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.Format()
			if err != nil {
				t.Fatalf("Format() error = %v", err)
			}
			if !tt.checkFn(result) {
				t.Errorf("Format() = %q, validation failed for test: %s", result, tt.name)
			}
		})
	}
}

// FuzzFormat is a fuzzing test for the Format function with various options
func FuzzFormat(f *testing.F) {
	seedInputs := []string{
		"0 b",
		"1 B",
		"100 KB",
		"1 MB",
		"1 GiB",
		"500 TB",
		"999 QB",
		"1 kilobyte",
		"0.5 MB",
	}

	for _, seed := range seedInputs {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, input string) {
		// First parse the input
		parsed, err := Parse(input)
		if err != nil {
			return // Skip invalid inputs
		}

		// Test various format option combinations
		optionCombos := [][]FormatOption{
			{},
			{WithDecimalUnits(true)},
			{WithDecimalUnits(false)},
			{WithLongUnits(true)},
			{WithLongUnits(false)},
			{WithFormatString("%.1f %s")},
			{WithFormatString("%.3f %s")},
			{WithDecimalUnits(true), WithLongUnits(true)},
			{WithDecimalUnits(false), WithLongUnits(true)},
		}

		for _, opts := range optionCombos {
			result, err := parsed.Format(opts...)
			if err != nil {
				t.Errorf("Format(%v) error = %v", opts, err)
				continue
			}

			// Result should not be empty
			if result == "" {
				t.Errorf("Format(%v) returned empty string for input %q", opts, input)
			}
		}
	})
}

// ============ Format Function Benchmarks ============

// BenchmarkFormatDefault benchmarks formatting with default options
func BenchmarkFormatDefault(b *testing.B) {
	value := Bytes(Uint128(MB).Mul64(512))

	for b.Loop() {
		value.Format()
	}
}

// BenchmarkFormatShortDecimal benchmarks formatting with short decimal units
func BenchmarkFormatShortDecimal(b *testing.B) {
	value := Bytes(Uint128(MB).Mul64(512))

	for b.Loop() {
		value.Format(WithDecimalUnits(true), WithLongUnits(false))
	}
}

// BenchmarkFormatShortBinary benchmarks formatting with short binary units
func BenchmarkFormatShortBinary(b *testing.B) {
	value := Bytes(Uint128(MiB).Mul64(512))

	for b.Loop() {
		value.Format(WithDecimalUnits(false), WithLongUnits(false))
	}
}

// BenchmarkFormatLongDecimal benchmarks formatting with long decimal unit names
func BenchmarkFormatLongDecimal(b *testing.B) {
	value := Bytes(Uint128(MB).Mul64(512))

	for b.Loop() {
		value.Format(WithDecimalUnits(true), WithLongUnits(true))
	}
}

// BenchmarkFormatLongBinary benchmarks formatting with long binary unit names
func BenchmarkFormatLongBinary(b *testing.B) {
	value := Bytes(Uint128(MiB).Mul64(512))

	for b.Loop() {
		value.Format(WithDecimalUnits(false), WithLongUnits(true))
	}
}

// BenchmarkFormatCustomFormat benchmarks formatting with custom format string
func BenchmarkFormatCustomFormat(b *testing.B) {
	value := Bytes(Uint128(MB).Mul64(512))

	for b.Loop() {
		value.Format(WithFormatString("%.0f %s"))
	}
}

// BenchmarkFormatForcedUnit benchmarks formatting with a forced unit
func BenchmarkFormatForcedUnit(b *testing.B) {
	value := Bytes(Uint128(B).Mul64(500000000))

	for b.Loop() {
		value.Format(WithForcedUnit(MB))
	}
}

// BenchmarkFormatSmallValue benchmarks formatting small byte values
func BenchmarkFormatSmallValue(b *testing.B) {
	value := B

	for b.Loop() {
		value.Format()
	}
}

// BenchmarkFormatLargeValue benchmarks formatting very large values
func BenchmarkFormatLargeValue(b *testing.B) {
	value := QB

	for b.Loop() {
		value.Format()
	}
}

// BenchmarkFormatMultipleOptions benchmarks formatting with multiple options
func BenchmarkFormatMultipleOptions(b *testing.B) {
	value := Bytes(Uint128(GiB).Mul64(42))

	for b.Loop() {
		value.Format(WithDecimalUnits(false), WithLongUnits(true), WithFormatString("%.2f %s"))
	}
}

// BenchmarkFormatParallel benchmarks Format function with parallel execution
func BenchmarkFormatParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		value := Bytes(Uint128(MB).Mul64(256))
		for pb.Next() {
			value.Format()
		}
	})
}

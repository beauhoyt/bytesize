// Package bytesize provides a type for representing byte sizes with support
// for both decimal (SI) and binary (IEC) units, as well as parsing from
// strings and formatting to human-readable strings.
package bytesize

import (
	"fmt"
	"math"
	"math/big"
	"slices"
	"strings"
	"unicode"
)

// Bytes represents a byte size as a 128-bit unsigned integer, allowing for
// very large sizes up to 2^128 - 1 bytes.
type Bytes Uint128

// Decimal byte size units (powers of 10).
var (
	None = Bytes{0, 0}
	One  = Bytes{1, 0}

	B  = One
	KB = Bytes(Uint128(B).Mul64(1e3))  // 1e3
	MB = Bytes(Uint128(KB).Mul64(1e3)) // 1e6
	GB = Bytes(Uint128(MB).Mul64(1e3)) // 1e9
	TB = Bytes(Uint128(GB).Mul64(1e3)) // 1e12
	PB = Bytes(Uint128(TB).Mul64(1e3)) // 1e15
	EB = Bytes(Uint128(PB).Mul64(1e3)) // 1e18
	ZB = Bytes(Uint128(EB).Mul64(1e3)) // 1e21
	YB = Bytes(Uint128(ZB).Mul64(1e3)) // 1e24
	RB = Bytes(Uint128(YB).Mul64(1e3)) // 1e27
	QB = Bytes(Uint128(RB).Mul64(1e3)) // 1e30
)

// LongDecimal maps decimal byte size units to their long names.
var LongDecimal = map[Bytes]string{
	KB: "Kilobyte",
	MB: "Megabyte",
	GB: "Gigabyte",
	TB: "Terabyte",
	PB: "Petabyte",
	EB: "Exabyte",
	ZB: "Zettabyte",
	YB: "Yottabyte",
	RB: "Ronnabyte",
	QB: "Quettabyte",
}

// ShortDecimal maps decimal byte size units to their short names.
var ShortDecimal = map[Bytes]string{
	KB: "KB",
	MB: "MB",
	GB: "GB",
	TB: "TB",
	PB: "PB",
	EB: "EB",
	ZB: "ZB",
	YB: "YB",
	RB: "RB",
	QB: "QB",
}

// Binary byte size units (powers of 2).
var (
	KiB = Bytes{1024, 0}
	MiB = Bytes{uint64(math.Pow(1024, 2)), 0}
	GiB = Bytes{uint64(math.Pow(1024, 3)), 0}
	TiB = Bytes{uint64(math.Pow(1024, 4)), 0}
	PiB = Bytes{uint64(math.Pow(1024, 5)), 0}
	EiB = Bytes{uint64(math.Pow(1024, 6)), 0}
	// ZB (2^70) and YB (2^80) cannot be represented as a single
	// uint64, so we use the high bits.
	// 2^70 = 2^(64+6) = 2^64 * 2^6 = (1 << 6) in the high bits.
	ZiB = Bytes{0, 1 << 6}
	// 2^80 = 2^(64+16) = 2^64 * 2^16 = (1 << 16) in the high bits.
	YiB = Bytes{0, 1 << 16}
	// 2^90 = 2^(64+26) = 2^64 * 2^26 = (1 << 26) in the high bits.
	RiB = Bytes{0, 1 << 26}
	// 2^100 = 2^(64+36) = 2^64 * 2^36 = (1 << 36) in the high bits.
	QiB = Bytes{0, 1 << 36}
)

// LongBinary maps binary byte size units to their long names.
var LongBinary = map[Bytes]string{
	KiB: "Kibibyte",
	MiB: "Mebibyte",
	GiB: "Gibibyte",
	TiB: "Tebibyte",
	PiB: "Pebibyte",
	EiB: "Exbibyte",
	ZiB: "Zebibyte",
	YiB: "Yobibyte",
	RiB: "Ronnibyte",
	QiB: "Quettibyte",
}

// ShortBinary maps binary byte size units to their short names.
var ShortBinary = map[Bytes]string{
	KiB: "KiB",
	MiB: "MiB",
	GiB: "GiB",
	TiB: "TiB",
	PiB: "PiB",
	EiB: "EiB",
	ZiB: "ZiB",
	YiB: "YiB",
	RiB: "RiB",
	QiB: "QiB",
}

// ValidUnits lists all supported unit strings for parsing.
var ValidUnits = []string{
	"b",
	"kb", "mb", "gb", "tb", "pb", "eb", "zb", "yb", "rb", "qb",
	"kib", "mib", "gib", "tib", "pib", "eib", "zib", "yib", "rib", "qib",
	"byte", "bytes",
	"kilobyte", "kilobytes", "megabyte", "megabytes", "gigabyte", "gigabytes", "terabyte", "terabytes", "petabyte", "petabytes",
	"exabyte", "exabytes", "zettabyte", "zettabytes", "yottabyte", "yottabytes", "ronnabyte", "ronnabytes", "quettabyte", "quettabytes",
	"kibibyte", "kibibytes", "mebibyte", "mebibytes", "gibibyte", "gibibytes", "tebibyte", "tebibytes", "pebibyte", "pebibytes",
	"exbibyte", "exbibytes", "zebibyte", "zebibytes", "yobibyte", "yobibytes", "ronnibyte", "ronnibytes", "quettibyte", "quettibytes",
}

// IsValidUnit checks if the provided unit string is a valid unit for
// parsing byte sizes.
func IsValidUnit(unit string) bool {
	unit = strings.ToLower(strings.TrimSpace(unit))
	return slices.Contains(ValidUnits, unit)
}

// Parse parses a string representation of a byte size (e.g., "10 MB",
// "5.5 GiB", "100 kilobytes", "2.34 Tebibytes") returns the corresponding
// Bytes value.
func Parse(s string) (Bytes, error) {
	// Trim whitespace
	s = strings.TrimSpace(s)
	if s == "" {
		return Bytes{}, fmt.Errorf("empty string")
	}

	numRunes, unitRunes, err := getNumAndUnitRunes(s)
	if err != nil {
		return Bytes{}, fmt.Errorf("error parsing number and unit: %v", err)
	}

	multiplier, err := getMultiplierForUnitString(string(unitRunes))
	if err != nil {
		return Bytes{}, err
	}

	// Parse the numeric part using big.Rat for arbitrary precision
	numStr := string(numRunes)
	if numStr == "" {
		return Bytes{}, fmt.Errorf("invalid number: empty numeric part")
	}

	numRat := new(big.Rat)
	_, ok := numRat.SetString(numStr)
	if !ok {
		return Bytes{}, fmt.Errorf("invalid number: %s", numStr)
	}

	if numRat.Sign() < 0 {
		return Bytes{}, fmt.Errorf("negative value: %s", numStr)
	}

	// Convert multiplier to big.Int
	multiplierInt := big.NewInt(0).SetUint64(Uint128(multiplier).Lo)
	if Uint128(multiplier).Hi > 0 {
		// Reconstruct full 128-bit number: (Hi << 64) | Lo
		multiplierInt.SetUint64(Uint128(multiplier).Hi)
		multiplierInt.Lsh(multiplierInt, 64)
		multiplierInt.Or(multiplierInt, big.NewInt(0).SetUint64(Uint128(multiplier).Lo))
	}

	// Multiply the number by the multiplier: result = numRat * multiplier
	resultRat := new(big.Rat).Mul(numRat, new(big.Rat).SetInt(multiplierInt))

	// Get the integer and fractional parts by dividing numerator by denominator
	resultInt := new(big.Int).Div(resultRat.Num(), resultRat.Denom())

	// Check if result overflows 128 bits
	if resultInt.BitLen() > 128 {
		return Bytes{}, fmt.Errorf("value overflows Uint128: result is %d bits", resultInt.BitLen())
	}

	if resultInt.Sign() < 0 {
		// This should never happen since we check for negative input, but
		// just in case, handle it gracefully
		return Bytes{}, fmt.Errorf("fatal: negative result from positive inputs")
	}

	// Convert big.Int to Uint128 (Lo and Hi)
	// Extract Lo (lower 64 bits)
	loInt := new(big.Int).And(resultInt, big.NewInt(-1).SetUint64(^uint64(0)))
	lo := loInt.Uint64()

	// Extract Hi (upper 64 bits)
	hiInt := new(big.Int).Rsh(resultInt, 64)
	hi := hiInt.Uint64()

	result := Uint128{lo, hi}
	return Bytes(result), nil
}

// getNumAndUnitRunes separates the numeric part and the unit part of the
// input string.
func getNumAndUnitRunes(s string) ([]rune, []rune, error) {
	foundDecimalPoint := false
	var numRunes, unitRunes []rune

	for _, r := range s {
		// 1. Skip spaces between number and unit
		if unicode.IsSpace(r) {
			continue
		}
		// 2. If we hit a number or decimal point, it's part of the number
		if r == '-' || (r >= '0' && r <= '9') || r == '.' {
			if r == '.' {
				if foundDecimalPoint {
					return nil, nil, fmt.Errorf("invalid number: multiple decimal points in %s", s)
				}
				foundDecimalPoint = true
			}
			numRunes = append(numRunes, r)
		} else {
			// 3. The rest is the unit
			unitRunes = append(unitRunes, r)
		}
	}

	return numRunes, unitRunes, nil
}

// getMultiplierForUnitString returns the multiplier Bytes value corresponding
// to the given unit string.
func getMultiplierForUnitString(unitStr string) (Bytes, error) {
	unitStr = strings.ToLower(strings.TrimSpace(unitStr))
	// Check decimal units (short names first, then long names)
	switch unitStr {
	// Short unit names
	// Decimal units
	case "b":
		return B, nil
	case "kb":
		return KB, nil
	case "mb":
		return MB, nil
	case "gb":
		return GB, nil
	case "tb":
		return TB, nil
	case "pb":
		return PB, nil
	case "eb":
		return EB, nil
	case "zb":
		return ZB, nil
	case "yb":
		return YB, nil
	case "rb":
		return RB, nil
	case "qb":
		return QB, nil

	// Binary units
	case "kib":
		return KiB, nil
	case "mib":
		return MiB, nil
	case "gib":
		return GiB, nil
	case "tib":
		return TiB, nil
	case "pib":
		return PiB, nil
	case "eib":
		return EiB, nil
	case "zib":
		return ZiB, nil
	case "yib":
		return YiB, nil
	case "rib":
		return RiB, nil
	case "qib":
		return QiB, nil

	// Long decimal names
	case "byte", "bytes":
		return B, nil
	case "kilobyte", "kilobytes":
		return KB, nil
	case "megabyte", "megabytes":
		return MB, nil
	case "gigabyte", "gigabytes":
		return GB, nil
	case "terabyte", "terabytes":
		return TB, nil
	case "petabyte", "petabytes":
		return PB, nil
	case "exabyte", "exabytes":
		return EB, nil
	case "zettabyte", "zettabytes":
		return ZB, nil
	case "yottabyte", "yottabytes":
		return YB, nil
	case "ronnabyte", "ronnabytes":
		return RB, nil
	case "quettabyte", "quettabytes":
		return QB, nil

	// Long binary names
	case "kibibyte", "kibibytes":
		return KiB, nil
	case "mebibyte", "mebibytes":
		return MiB, nil
	case "gibibyte", "gibibytes":
		return GiB, nil
	case "tebibyte", "tebibytes":
		return TiB, nil
	case "pebibyte", "pebibytes":
		return PiB, nil
	case "exbibyte", "exbibytes":
		return EiB, nil
	case "zebibyte", "zebibytes":
		return ZiB, nil
	case "yobibyte", "yobibytes":
		return YiB, nil
	case "ronnibyte", "ronnibytes":
		return RiB, nil
	case "quettibyte", "quettibytes":
		return QiB, nil
	default:
		return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
	}
}

// Set implements the flag.Value interface for Bytes.
func (b *Bytes) Set(s string) error {
	parsed, err := Parse(s)
	if err != nil {
		return err
	}
	*b = parsed
	return nil
}

// Get implements the flag.Getter interface for Bytes.
func (b *Bytes) Get() any {
	return Bytes(*b)
}

// Type implements the flag.Value interface for Bytes.
func (b *Bytes) Type() string {
	return "bytesize.Bytes"
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Bytes.
func (b *Bytes) UnmarshalText(text []byte) error {
	return b.Set(string(text))
}

type formatOptions struct {
	// Format string for formatting, defaults to "%.2f %s"
	formatStr string

	// Forced unit for formatting, nil if automatic
	forcedUnitType *Bytes

	// Use long unit names if true, short unit names if false
	longUnits bool

	// Use decimal (SI) units if true, binary (IEC) units if false
	decimalUnits bool
}

const (
	// DefaultFormatStr is the default format string for formatting byte
	// sizes, which includes two decimal places and the unit.
	DefaultFormatStr = "%.2f %s"
	// DefaultLongUnits indicates whether to use long unit names, such
	// as "Megabyte" instead of "MB", though the default is to use short unit
	// names.
	DefaultLongUnits = false
	// DefaultDecimalUnits indicates whether to use decimal (SI) units by default
	DefaultDecimalUnits = true
)

func newFormatOptions() *formatOptions {
	return &formatOptions{
		formatStr:    DefaultFormatStr,
		longUnits:    DefaultLongUnits,
		decimalUnits: DefaultDecimalUnits,
	}
}

// FormatOption defines a functional option for configuring the formatting
// of byte sizes.
type FormatOption func(*formatOptions) error

// WithFormatString allows you to specify a custom format string for
// formatting byte sizes. The format string should include two verbs:
// one for the value (e.g., %.2f) and one for the unit (e.g., %s).
func WithFormatString(formatStr string) FormatOption {
	return func(opts *formatOptions) error {
		if formatStr == "" {
			return fmt.Errorf("format string cannot be empty")
		}
		opts.formatStr = formatStr
		return nil
	}
}

// WithForcedUnit allows you to specify a specific unit to use when formatting
// byte sizes. If not set, the formatting will automatically choose the most
// appropriate unit based on the value.
func WithForcedUnit(unit Bytes) FormatOption {
	return func(opts *formatOptions) error {
		switch unit {
		case B, KB, MB, GB, TB, PB, EB, ZB, YB, RB, QB:
			opts.decimalUnits = true
			opts.forcedUnitType = &unit
			return nil
		case KiB, MiB, GiB, TiB, PiB, EiB, ZiB, YiB, RiB, QiB:
			opts.decimalUnits = false
			opts.forcedUnitType = &unit
			return nil
		default:
			return fmt.Errorf("invalid forced unit: %v", unit)
		}
	}
}

// WithLongUnits allows you to specify whether to use long unit names (e.g.,
// "Megabyte") or short unit names (e.g., "MB") when formatting byte sizes.
func WithLongUnits(longUnits bool) FormatOption {
	return func(opts *formatOptions) error {
		opts.longUnits = longUnits
		return nil
	}
}

// WithDecimalUnits allows you to specify whether to use decimal (SI) units
// or binary (IEC) units when formatting byte sizes. If true, it will use
// decimal units (KB, MB, etc.); if false, it will use binary units (KiB,
// MiB, etc.).
func WithDecimalUnits(decimalUnits bool) FormatOption {
	return func(opts *formatOptions) error {
		opts.decimalUnits = decimalUnits
		return nil
	}
}

func (b Bytes) String() string {
	str, err := b.Format()
	if err != nil {
		// This should never happen since we're using default options,
		// but just in case, return a fallback string
		return fmt.Sprintf("%d B", Uint128(b).Lo)
	}
	return str
}

// Format formats the Bytes value as a human-readable string using the
// specified options. It returns the formatted string or an error if any
// of the options are invalid.
func (b Bytes) Format(opts ...FormatOption) (string, error) {
	return b.format(opts...)
}

func (b Bytes) format(opts ...FormatOption) (string, error) {
	formatOptions := newFormatOptions()
	for _, opt := range opts {
		if err := opt(formatOptions); err != nil {
			return "", err
		}
	}

	// Select the appropriate unit maps
	var unitMap map[Bytes]string
	var unitSlice []Bytes

	if formatOptions.decimalUnits {
		if formatOptions.longUnits {
			unitMap = LongDecimal
		} else {
			unitMap = ShortDecimal
		}
		unitSlice = []Bytes{QB, RB, YB, ZB, EB, PB, TB, GB, MB, KB, B}
	} else {
		if formatOptions.longUnits {
			unitMap = LongBinary
		} else {
			unitMap = ShortBinary
		}
		unitSlice = []Bytes{QiB, RiB, YiB, ZiB, EiB, PiB, TiB, GiB, MiB, KiB, B}
	}

	// Determine which unit to use
	var bestUnit Bytes

	if formatOptions.forcedUnitType != nil {
		bestUnit = *formatOptions.forcedUnitType
	} else {
		// Find the best unit by finding the largest unit <= b
		for _, unit := range unitSlice {
			if Uint128(b).Cmp(Uint128(unit)) >= 0 {
				bestUnit = unit
				break
			}
		}
		// If no unit was found (b is less than all units), use bytes
		if bestUnit.Lo == 0 && bestUnit.Hi == 0 {
			bestUnit = B
		}
	}

	// Calculate the value in the chosen unit using big.Float for precision
	bBig := big.NewInt(0).SetUint64(Uint128(b).Lo)
	if Uint128(b).Hi > 0 {
		bBig.SetUint64(Uint128(b).Hi)
		bBig.Lsh(bBig, 64)
		bBig.Add(bBig, big.NewInt(0).SetUint64(Uint128(b).Lo))
	}

	unitBig := big.NewInt(0).SetUint64(Uint128(bestUnit).Lo)
	if Uint128(bestUnit).Hi > 0 {
		unitBig.SetUint64(Uint128(bestUnit).Hi)
		unitBig.Lsh(unitBig, 64)
		unitBig.Add(unitBig, big.NewInt(0).SetUint64(Uint128(bestUnit).Lo))
	}

	// Use big.Float to calculate the value with proper precision
	bFloat := big.NewFloat(0).SetInt(bBig)
	unitFloat := big.NewFloat(0).SetInt(unitBig)
	value := big.NewFloat(0).Quo(bFloat, unitFloat)

	// Get the unit name
	// fmt.Printf("UnitMap: %v\n", unitMap)
	unitName, found := unitMap[bestUnit]
	if !found {
		if formatOptions.longUnits {
			unitName = "Byte"
		} else {
			unitName = "B"
		}
	}
	if formatOptions.longUnits && value.Cmp(big.NewFloat(1)) != 0 {
		unitName += "s"
	}

	return fmt.Sprintf(formatOptions.formatStr, value, unitName), nil
}

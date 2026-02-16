package bytesize

import (
	"fmt"
	"strings"
)

// UnitStringToBytes maps valid unit strings to their corresponding Bytes
// multiplier values for parsing.
var UnitStringToBytes = map[string]Bytes{
	"b":  B,
	"kb": KB, "mb": MB, "gb": GB, "tb": TB, "pb": PB, "eb": EB, "zb": ZB, "yb": YB, "rb": RB, "qb": QB,
	"kib": KiB, "mib": MiB, "gib": GiB, "tib": TiB, "pib": PiB, "eib": EiB, "zib": ZiB, "yib": YiB, "rib": RiB, "qib": QiB,
	"byte": B, "bytes": B,
	"kilobyte": KB, "kilobytes": KB, "megabyte": MB, "megabytes": MB, "gigabyte": GB, "gigabytes": GB,
	"terabyte": TB, "terabytes": TB, "petabyte": PB, "petabytes": PB, "exabyte": EB, "exabytes": EB, "zettabyte": ZB, "zettabytes": ZB, "yottabyte": YB, "yottabytes": YB,
	"ronnabyte": RB, "ronnabytes": RB, "quettabyte": QB, "quettabytes": QB,
	"kibibyte": KiB, "kibibytes": KiB, "mebibyte": MiB, "mebibytes": MiB, "gibibyte": GiB, "gibibytes": GiB,
	"tebibyte": TiB, "tebibytes": TiB, "pebibyte": PiB, "pebibytes": PiB, "exbibyte": EiB, "exbibytes": EiB,
	"zebibyte": ZiB, "zebibytes": ZiB, "yobibyte": YiB, "yobibytes": YiB, "ronnibyte": RiB, "ronnibytes": RiB, "quettibyte": QiB, "quettibytes": QiB,
}

// getMultiplierForUnit returns the multiplier Bytes value corresponding to the
// given unit string. It looks up the unit string in the UnitStringToBytes map
// and returns the corresponding multiplier, or an error if the unit is unknown.
func getMultiplierByUnitStringMapVersion(unitStr string) (Bytes, error) {
	unitStr = strings.ToLower(strings.TrimSpace(unitStr))
	multiplier, found := UnitStringToBytes[unitStr]
	if !found {
		return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
	}
	return multiplier, nil
}

// getMultiplierByUnitStringNestedSwitchesVersion is an alternative implementation of
// getMultiplierByUnitString that uses a switch statement to determine the multiplier
// based on the first few characters of the unit string. This approach may be more
// efficient than a map lookup for a small number of units, but it is less flexible
// and harder to maintain if new units are added in the future.
func getMultiplierByUnitStringNestedSwitchesVersion(unitStr string) (Bytes, error) {
	unitStr = strings.ToLower(strings.TrimSpace(unitStr))
	switch unitStr[0] {
	case 'b':
		return B, nil
	case 'k':
		switch unitStr[1] {
		case 'b':
			return KB, nil
		case 'i':
			switch unitStr[2] {
			case 'b':
				return KiB, nil
			case 'l':
				return KB, nil
			default:
				return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
			}
		default:
			return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
		}
	case 'm':
		switch unitStr[1] {
		case 'b':
			return MB, nil
		case 'i':
			return MiB, nil
		case 'e':
			switch unitStr[2] {
			case 'b':
				return MiB, nil
			case 'g':
				return MB, nil
			default:
				return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
			}
		default:
			return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
		}
	case 'g':
		switch unitStr[1] {
		case 'b':
			return GB, nil
		case 'i':
			switch unitStr[2] {
			case 'b':
				return GiB, nil
			case 'g':
				return GB, nil
			default:
				return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
			}
		default:
			return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
		}
	case 't':
		switch unitStr[1] {
		case 'b':
			return TB, nil
		case 'i':
			return TiB, nil
		case 'e':
			switch unitStr[2] {
			case 'b':
				return TiB, nil
			case 'r':
				return TB, nil
			default:
				return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
			}
		default:
			return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
		}
	case 'p':
		switch unitStr[1] {
		case 'b':
			return PB, nil
		case 'i':
			return PiB, nil
		case 'e':
			switch unitStr[2] {
			case 'b':
				return PiB, nil
			case 't':
				return PB, nil
			default:
				return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
			}
		default:
			return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
		}
	case 'e':
		switch unitStr[1] {
		case 'b':
			return EB, nil
		case 'i':
			return EiB, nil
		case 'x':
			switch unitStr[2] {
			case 'b':
				return EiB, nil
			case 'a':
				return EB, nil
			default:
				return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
			}
		default:
			return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
		}
	case 'z':
		switch unitStr[1] {
		case 'b':
			return ZB, nil
		case 'i':
			return ZiB, nil
		case 'e':
			switch unitStr[2] {
			case 'b':
				return ZiB, nil
			case 't':
				return ZB, nil
			default:
				return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
			}
		default:
			return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
		}
	case 'y':
		switch unitStr[1] {
		case 'b':
			return YB, nil
		case 'i':
			return YiB, nil
		case 'o':
			switch unitStr[2] {
			case 'b':
				return YiB, nil
			case 't':
				return YB, nil
			default:
				return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
			}
		default:
			return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
		}
	case 'r':
		switch unitStr[1] {
		case 'b':
			return RB, nil
		case 'i':
			return RiB, nil
		case 'o':
			switch unitStr[2] {
			case 'n':
				switch unitStr[3] {
				case 'n':
					switch unitStr[4] {
					case 'i':
						return RiB, nil
					case 'a':
						return RB, nil
					default:
						return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
					}
				default:
					return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
				}
			default:
				return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
			}
		default:
			return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
		}
	case 'q':
		switch unitStr[1] {
		case 'b':
			return QB, nil
		case 'i':
			return QiB, nil
		case 'u':
			switch unitStr[2] {
			case 'e':
				switch unitStr[3] {
				case 't':
					switch unitStr[4] {
					case 't':
						switch unitStr[5] {
						case 'i':
							return QiB, nil
						case 'a':
							return QB, nil
						default:
							return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
						}
					default:
						return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
					}
				default:
					return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
				}
			default:
				return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
			}
		default:
			return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
		}
	default:
		return Bytes{}, fmt.Errorf("unknown unit: %s", unitStr)
	}
}

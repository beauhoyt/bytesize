package bytesize

import "testing"

func TestGetMultiplierByUnitString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Bytes
		wantErr bool
	}{
		{"Valid short bytes", "B", B, false},
		{"Valid KB", "KB", KB, false},
		{"Valid KiB", "KiB", KiB, false},
		{"Valid kilobyte", "Kilobyte", KB, false},
		{"Valid kibibyte", "Kibibyte", KiB, false},
		{"Invalid typoed kilobyte", "Kiobyte", Bytes{}, true},
		{"Invalid typoed kilobyte", "Klobyte", Bytes{}, true},
		{"Valid MB", "MB", MB, false},
		{"Valid MiB", "MiB", MiB, false},
		{"Valid megabyte", "Megabyte", MB, false},
		{"Valid mebibyte", "Mebibyte", MiB, false},
		{"Invalid typoed megabyte", "Meabtye", Bytes{}, true},
		{"Invalid typoed megabyte", "Mgabtye", Bytes{}, true},
		{"Valid GB", "GB", GB, false},
		{"Valid GiB", "GiB", GiB, false},
		{"Valid gigabyte", "Gigabyte", GB, false},
		{"Valid gibibyte", "Gibibyte", GiB, false},
		{"Invalid typoed gigabyte", "Giabtye", Bytes{}, true},
		{"Invalid typoed gigabyte", "Ggabtye", Bytes{}, true},
		{"Valid TB", "TB", TB, false},
		{"Valid TiB", "TiB", TiB, false},
		{"Valid terabyte", "Terabyte", TB, false},
		{"Valid tebibyte", "Tebibyte", TiB, false},
		{"Invalid typoed terabyte", "Teabtye", Bytes{}, true},
		{"Invalid typoed terabyte", "Trabyte", Bytes{}, true},
		{"Valid PB", "PB", PB, false},
		{"Valid PiB", "PiB", PiB, false},
		{"Valid petabyte", "Petabyte", PB, false},
		{"Valid pebibyte", "Pebibyte", PiB, false},
		{"Invalid typoed petabyte", "Peabtye", Bytes{}, true},
		{"Invalid typoed petabyte", "Ptabtye", Bytes{}, true},
		{"Valid EB", "EB", EB, false},
		{"Valid EiB", "EiB", EiB, false},
		{"Valid exabyte", "Exabyte", EB, false},
		{"Valid exbibyte", "Exbibyte", EiB, false},
		{"Invalid typoed exabyte", "Exibyte", Bytes{}, true},
		{"Invalid typoed exabyte", "Eabyte", Bytes{}, true},
		{"Valid ZB", "ZB", ZB, false},
		{"Valid ZiB", "ZiB", ZiB, false},
		{"Valid zettabyte", "Zettabyte", ZB, false},
		{"Valid zebibyte", "Zebibyte", ZiB, false},
		{"Invalid typoed zettabyte", "Zeabyte", Bytes{}, true},
		{"Invalid typoed zettabyte", "Zttabyte", Bytes{}, true},
		{"Valid YB", "YB", YB, false},
		{"Valid YiB", "YiB", YiB, false},
		{"Valid yottabyte", "Yottabyte", YB, false},
		{"Valid yobibyte", "Yobibyte", YiB, false},
		{"Invalid typoed yottabyte", "Yoabtye", Bytes{}, true},
		{"Invalid typoed yottabyte", "Yttabyte", Bytes{}, true},
		{"Valid RB", "RB", RB, false},
		{"Valid RiB", "RiB", RiB, false},
		{"Valid ronnabyte", "Ronnabyte", RB, false},
		{"Valid ronnibyte", "Ronnibyte", RiB, false},
		{"Invalid typoed ronnabyte", "Ronnbyte", Bytes{}, true},
		{"Invalid typoed ronnabyte", "Ronabyte", Bytes{}, true},
		{"Invalid typoed ronnabyte", "Romabyte", Bytes{}, true},
		{"Invalid typoed ronnabyte", "Rnnabtye", Bytes{}, true},
		{"Valid QB", "QB", QB, false},
		{"Valid QiB", "QiB", QiB, false},
		{"Valid quettabyte", "Quettabyte", QB, false},
		{"Valid quettibyte", "Quettibyte", QiB, false},
		{"Invalid typoed quettabyte", "Quettbyte", Bytes{}, true},
		{"Invalid typoed quettabyte", "Quetabyte", Bytes{}, true},
		{"Invalid typoed quettabyte", "Queabytee", Bytes{}, true},
		{"Invalid typoed quettabyte", "Quttabyte", Bytes{}, true},
		{"Invalid typoed quettabyte", "Qettabtye", Bytes{}, true},
		{"Invalid unit", "InvalidUnit", Bytes{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getMultiplierByUnitString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMultiplierByUnitString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getMultiplierByUnitString() = %v, want %v", got, tt.want)
			}
		})

		t.Run("Nested switches version - "+tt.name, func(t *testing.T) {
			got, err := getMultiplierByUnitStringNestedSwitchesVersion(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMultiplierByUnitStringNestedSwitchesVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getMultiplierByUnitStringNestedSwitchesVersion() = %v, want %v", got, tt.want)
			}
		})

		t.Run("Map version - "+tt.name, func(t *testing.T) {
			got, err := getMultiplierByUnitStringMapVersion(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMultiplierByUnitStringMapVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getMultiplierByUnitStringMapVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGetMultiplierByUnitString_LongDecimal(b *testing.B) {
	for b.Loop() {
		getMultiplierByUnitString("Quettabyte")
	}
}

func BenchmarkGetMultiplierByUnitString_LongBinary(b *testing.B) {
	for b.Loop() {
		getMultiplierByUnitString("Quettibyte")
	}
}

func BenchmarkGetMultiplierByUnitString_ShortDecimal(b *testing.B) {
	for b.Loop() {
		getMultiplierByUnitString("QB")
	}
}

func BenchmarkGetMultiplierByUnitString_ShortBinary(b *testing.B) {
	for b.Loop() {
		getMultiplierByUnitString("QiB")
	}
}

func BenchmarkGetMultiplierByUnitStringNestedSwitchesVersion_LongDecimal(b *testing.B) {
	for b.Loop() {
		getMultiplierByUnitStringNestedSwitchesVersion("Quettabyte")
	}
}

func BenchmarkGetMultiplierByUnitStringNestedSwitchesVersion_LongBinary(b *testing.B) {
	for b.Loop() {
		getMultiplierByUnitStringNestedSwitchesVersion("Quettibyte")
	}
}

func BenchmarkGetMultiplierByUnitStringNestedSwitchesVersion_ShortDecimal(b *testing.B) {
	for b.Loop() {
		getMultiplierByUnitStringNestedSwitchesVersion("QB")
	}
}

func BenchmarkGetMultiplierByUnitStringNestedSwitchesVersion_ShortBinary(b *testing.B) {
	for b.Loop() {
		getMultiplierByUnitStringNestedSwitchesVersion("QiB")
	}
}

func BenchmarkGetMultiplierByUnitStringMapVersion_LongDecimal(b *testing.B) {
	for b.Loop() {
		getMultiplierByUnitStringMapVersion("Quettabyte")
	}
}

func BenchmarkGetMultiplierByUnitStringMapVersion_LongBinary(b *testing.B) {
	for b.Loop() {
		getMultiplierByUnitStringMapVersion("Quettibyte")
	}
}

func BenchmarkGetMultiplierByUnitStringMapVersion_ShortDecimal(b *testing.B) {
	for b.Loop() {
		getMultiplierByUnitStringMapVersion("QB")
	}
}

func BenchmarkGetMultiplierByUnitStringMapVersion_ShortBinary(b *testing.B) {
	for b.Loop() {
		getMultiplierByUnitStringMapVersion("QiB")
	}
}

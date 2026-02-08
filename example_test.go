package bytesize_test

import (
	"fmt"

	"github.com/beauhoyt/bytesize"
)

func ExampleBytes() {
	bs := bytesize.Bytes{123456789, 0}
	fmt.Printf("%s\n", bs)
	// Output: 123.46 MB
}

func ExampleParse() {
	bs, err := bytesize.Parse("123.45 MB")
	if err != nil {
		fmt.Printf("Error parsing byte size: %v", err)
		return
	}
	fmt.Printf("%s\n", bs)
	// Output: 123.45 MB
}

func ExampleIsValidUnit() {
	fmt.Printf("Is 'MB' a valid unit? %v\n", bytesize.IsValidUnit("MB"))
	fmt.Printf("Is 'invalid' a valid unit? %v\n", bytesize.IsValidUnit("invalid"))
	// Output:
	// Is 'MB' a valid unit? true
	// Is 'invalid' a valid unit? false
}

func ExampleBytes_Format() {
	bs := bytesize.Bytes{1234567890, 0}
	formatStr, err := bs.Format(bytesize.WithForcedUnit(bytesize.MiB))
	if err != nil {
		fmt.Printf("Error formatting byte size: %v\n", err)
		return
	}
	fmt.Printf("%s\n", formatStr)
	// Output: 1177.38 MiB
}

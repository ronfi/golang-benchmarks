// Package parse benchmarks parsing.
package parse

import (
	"fmt"
	"math/big"
	"strconv"
	"testing"
)

var smallStr = "35"
var bigStr = "999999999999999"

func BenchmarkAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val, _ := strconv.Atoi(smallStr)
		_ = val
	}
}

func BenchmarkAtoiParseInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val, _ := strconv.ParseInt(smallStr, 0, 64)
		_ = val
	}
}

func BenchmarkAtoiBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val, _ := strconv.Atoi(bigStr)
		_ = val
	}
}

func BenchmarkAtoiParseIntBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val, _ := strconv.ParseInt(bigStr, 0, 64)
		_ = val
	}
}

func BenchmarkParseBool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := strconv.ParseBool("true")
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkParseInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := strconv.ParseInt("1337", 10, 64)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkParseFloat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := strconv.ParseFloat("3.141592653589793238462643383", 64)
		if err != nil {
			panic(err)
		}
	}
}

// ToEth wei to eth
func ToEth(amount *big.Int) string {
	compactAmount := big.NewInt(0)
	reminder := big.NewInt(0)
	divisor := big.NewInt(1e18)
	compactAmount.QuoRem(amount, divisor, reminder)
	return fmt.Sprintf("%v.%018s", compactAmount.String(), reminder.String())
}

func BenchmarkParseBig2Float1(b *testing.B) {
	gasPrice := big.NewInt(0).Mul(big.NewInt(5), big.NewInt(1e9))
	txFee := new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(500000000))
	for n := 0; n < b.N; n++ {
		fval := float64(new(big.Int).Div(txFee, big.NewInt(1e9)).Int64()) / 1e9
		if fval == 0.0 {
			panic("exit")
		}
	}
}

func BenchmarkParseBig2Float2(b *testing.B) {
	gasPrice := big.NewInt(0).Mul(big.NewInt(5), big.NewInt(1e9))
	txFee := new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(500000000))
	for n := 0; n < b.N; n++ {
		_, err := strconv.ParseFloat(ToEth(txFee), 64)
		if err != nil {
			panic(err)
		}
	}
}

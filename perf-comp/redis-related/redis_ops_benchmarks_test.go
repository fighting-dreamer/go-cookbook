package redis_related

import (
	"testing"
)

var rc RedisCli

func init() {
	rc = NewRedisCli()
}

func BenchmarkGetOnMultipleKeysSerial(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getOnMultipleKeysSerial(&rc)
	}
}

func BenchmarkGetOnMultipleKeysConncurent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getOnMultipleKeysConncurent(&rc)
	}
}

func BenchmarkMGetOnKeys(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mgetOnKeys(&rc)
	}
}
//EVAL "return {redis.call('get', 'qwerty1'), redis.call('get', 'qwerty2'), redis.call('get', 'qwerty3'), redis.call('get', 'qwerty4'), redis.call('get', 'qwerty5'), redis.call('get', 'qwerty0', 1)}" 0

func BenchmarkEval(b *testing.B) {
	for n := 0; n < b.N; n++ {
		evalOnKeys(&rc)
	}
}
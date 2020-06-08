package hello

import (
	quoteV3 "rsc.io/quote/v3"
)

func Hello() string {
	return quoteV3.HelloV3()
}

func Proverb() string {
	return quoteV3.Concurrency()
}

func Glass() string {
	return quoteV3.GlassV3()
}

func Opt() string {
	return quoteV3.OptV3()
}

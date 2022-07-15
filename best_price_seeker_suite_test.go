package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBestPriceSeeker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BestPriceSeeker Suite")
}

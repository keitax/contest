package library_test

import (
	"github.com/keitax/contest/library"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("functions", func() {
	Describe("NextPerm()", func() {
		Context("called until returns true", func() {
			var (
				l   []int
				got [][]int
			)

			BeforeEach(func() {
				l = []int{1, 2, 2, 3}
				got = [][]int{append([]int(nil), l...)}
				for library.NextPerm(l) {
					got = append(got, append([]int(nil), l...))
				}
			})

			It("generates next permutations", func() {
				Expect(got).To(Equal([][]int{
					{1, 2, 2, 3},
					{1, 2, 3, 2},
					{1, 3, 2, 2},
					{2, 1, 2, 3},
					{2, 1, 3, 2},
					{2, 2, 1, 3},
					{2, 2, 3, 1},
					{2, 3, 1, 2},
					{2, 3, 2, 1},
					{3, 1, 2, 2},
					{3, 2, 1, 2},
					{3, 2, 2, 1},
				}))
			})
		})
	})
})

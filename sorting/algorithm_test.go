// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package sorting

import (
	"math/rand"
	"sort"
	"testing"
)

type sortFn func(data sort.Interface)

func testSortFn(t *testing.T, fn sortFn, fnName string) {
	for _, test := range []sort.Interface{
		sort.IntSlice([]int{660, 14, 796, 336, 223, 594, 419, 574, 372, 103, 991, 718, 436, 351, 844, 277, 668, 250, 330, 86}),
		sort.Float64Slice([]float64{213.237, 458.642, 978.311, 547.351, 57.8992, 245.518, 445.638, 251.79, 960.202, 100.069, 483.136, 407.858, 496.913, 562.943, 557.959, 219.648, 164.599, 843.304, 671.732, 222.676}),
		sort.StringSlice([]string{"assaults", "brackish", "monarchism", "ascribe", "mechanize", "andiron", "overpricing", "jading", "hauliers", "snug", "zodiac", "credit", "tremendous", "palavered", "hibiscuses", "amplest", "interrogated", "geologic", "unorthodoxy", "propagated"}),
	} {
		// Make a copy to avoid modification of the original slice.
		var data sort.Interface
		switch v := test.(type) {
		case sort.IntSlice:
			data = append(sort.IntSlice(nil), v...)
		case sort.Float64Slice:
			data = append(sort.Float64Slice(nil), v...)
		case sort.StringSlice:
			data = append(sort.StringSlice(nil), v...)
		default:
			t.Errorf("missing case: cannot copy %T", v)
			continue
		}
		fn(data)
		if !sort.IsSorted(data) {
			t.Errorf("%s(%v)", fnName, test)
			t.Errorf(" got %v", data)
			sort.Sort(data)
			t.Errorf("want %v", data)
		}
	}
}

func TestBubbleSort(t *testing.T) { testSortFn(t, BubbleSort, "BubbleSort") }

func benchRandom(b *testing.B, fn sortFn, size int) {
	b.StopTimer()
	ints := sort.IntSlice(rand.New(rand.NewSource(int64(size))).Perm(size))
	for i := 0; i < b.N; i++ {
		data := append(sort.IntSlice(nil), ints...)
		b.StartTimer()
		fn(data)
		b.StopTimer()
	}
}

func benchNearlySorted(b *testing.B, fn sortFn, size int) {
	b.StopTimer()
	ints := sort.IntSlice(rand.New(rand.NewSource(int64(size))).Perm(size))
	sort.Sort(ints)
	for i := 2; i < len(ints); i += 4 {
		ints[i], ints[i-2] = ints[i-2], ints[i]
	}
	for i := 0; i < b.N; i++ {
		data := append(sort.IntSlice(nil), ints...)
		b.StartTimer()
		fn(data)
		b.StopTimer()
	}
}

func benchSorted(b *testing.B, fn sortFn, size int) {
	b.StopTimer()
	ints := sort.IntSlice(rand.New(rand.NewSource(int64(size))).Perm(size))
	sort.Sort(ints)
	for i := 0; i < b.N; i++ {
		data := append(sort.IntSlice(nil), ints...)
		b.StartTimer()
		fn(data)
		b.StopTimer()
	}
}

func benchReversed(b *testing.B, fn sortFn, size int) {
	b.StopTimer()
	ints := sort.IntSlice(rand.New(rand.NewSource(int64(size))).Perm(size))
	sort.Sort(sort.Reverse(ints))
	for i := 2; i < len(ints); i += 4 {
		ints[i], ints[i-2] = ints[i-2], ints[i]
	}
	for i := 0; i < b.N; i++ {
		data := append(sort.IntSlice(nil), ints...)
		b.StartTimer()
		fn(data)
		b.StopTimer()
	}
}

func BenchmarkBubbleSortRandom1e1(b *testing.B)       { benchRandom(b, BubbleSort, 1e1) }
func BenchmarkBubbleSortRandom1e2(b *testing.B)       { benchRandom(b, BubbleSort, 1e2) }
func BenchmarkBubbleSortRandom1e3(b *testing.B)       { benchRandom(b, BubbleSort, 1e3) }
func BenchmarkBubbleSortNearlySorted1e1(b *testing.B) { benchNearlySorted(b, BubbleSort, 1e1) }
func BenchmarkBubbleSortNearlySorted1e2(b *testing.B) { benchNearlySorted(b, BubbleSort, 1e2) }
func BenchmarkBubbleSortNearlySorted1e3(b *testing.B) { benchNearlySorted(b, BubbleSort, 1e3) }
func BenchmarkBubbleSortSorted1e1(b *testing.B)       { benchSorted(b, BubbleSort, 1e1) }
func BenchmarkBubbleSortSorted1e2(b *testing.B)       { benchSorted(b, BubbleSort, 1e2) }
func BenchmarkBubbleSortSorted1e3(b *testing.B)       { benchSorted(b, BubbleSort, 1e3) }
func BenchmarkBubbleSortReversed1e1(b *testing.B)     { benchReversed(b, BubbleSort, 1e1) }
func BenchmarkBubbleSortReversed1e2(b *testing.B)     { benchReversed(b, BubbleSort, 1e2) }
func BenchmarkBubbleSortReversed1e3(b *testing.B)     { benchReversed(b, BubbleSort, 1e3) }

package stringfactory

import (
	"fmt"
	"io"
	"math"
)

var Char_0_9 = []uint8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var Char_a_z = []uint8{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
	'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var Char_A_Z = []uint8{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

type Stringfactory struct {
	totalNums     uint64
	generatedNums uint64
	prefix        string
	items         []uint8
	itemsNums     uint32
	strLen        int
	counters      []uint32
	onlyPrefix    bool
}

func New(prefix string, strLen int, items []uint8) *Stringfactory {
	strFac := new(Stringfactory)
	if strLen == 0 || len(items) == 0 {
		if prefix != "" {
			strFac.totalNums = 1
			strFac.prefix = prefix
			strFac.onlyPrefix = true
		}
		return strFac
	}

	strFac.onlyPrefix = false
	strFac.prefix = prefix
	strFac.itemsNums = uint32(len(items))
	strFac.items = items
	strFac.strLen = strLen
	strFac.totalNums = uint64(math.Pow(float64(strFac.itemsNums), float64(strLen)))
	for i := 0; i < strFac.strLen; i++ {
		strFac.counters = append(strFac.counters, 0)
	}
	return strFac
}

func (s *Stringfactory) makeOneString() string {
	if s.generatedNums == s.totalNums {
		return ""
	}

	if s.onlyPrefix {
		s.generatedNums++
		return s.prefix
	}

	ret := s.prefix
	for i := 0; i < s.strLen; i++ {
		ret += string(s.items[s.counters[i]])
	}
	s.generatedNums++

	// calc every counters
	for i := s.strLen - 1; i >= 0; i-- {
		if s.counters[i] == s.itemsNums-1 {
			s.counters[i] = 0
			continue
		} else {
			s.counters[i]++
			break
		}
	}

	return ret
}

func (s *Stringfactory) Reset() {
	*s = Stringfactory{}
}

func (s *Stringfactory) GetFirstString() string {
	return s.makeOneString()
}

func (s *Stringfactory) GetNextString() string {
	return s.makeOneString()
}

func (s Stringfactory) TotalNumber() uint64 {
	return s.totalNums
}

func (s Stringfactory) GeneratedNumber() uint64 {
	return s.generatedNums
}

func (s Stringfactory) ShowInfo(w io.Writer) {
	fmt.Fprintf(w, "StringFactory information : \n")
	fmt.Fprintf(w, "prefix      : %v\n", s.prefix)
	fmt.Fprintf(w, "totalNums   : %v\n", s.totalNums)
	fmt.Fprintf(w, "getNums		: %v\n", s.generatedNums)
	fmt.Fprintf(w, "itemsNums	: %v\n", s.itemsNums)
	fmt.Fprintf(w, "onlyPrefix	: %v\n", s.onlyPrefix)
	if len(s.counters) != 0 {
		var str string
		for _, cnt := range s.counters {
			str += string(cnt) + ""
		}
		fmt.Fprintf(w, "counters   : %v\n", str)
	}
}

package stringfactory

import (
	"testing"
)

func TestBase(t *testing.T) {
	s := New("", 1, Char_0_9)
	expected_str := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	if s.TotalNumber() != 10 {
		t.Errorf(`totalNums is error, %v`, s.TotalNumber())
	}

	str := s.GetFirstString()
	i := -1
	for str != "" {
		i++
		if str != expected_str[i] || int(s.GeneratedNumber()) != i+1 {
			t.Errorf(`str %v, generated number : %v, expected(%v, %v)`, str, s.GeneratedNumber(), expected_str[i], i+1)
		}
		str = s.GetNextString()
	}
}

func TestBaseWithPrefix(t *testing.T) {
	s := New("aa", 1, Char_0_9)
	expected_str := []string{"aa0", "aa1", "aa2", "aa3", "aa4", "aa5", "aa6", "aa7", "aa8", "aa9"}

	if s.TotalNumber() != 10 {
		t.Errorf(`totalNums is error, %v`, s.TotalNumber())
	}

	str := s.GetFirstString()
	i := -1
	for str != "" {
		i++
		if str != expected_str[i] || int(s.GeneratedNumber()) != i+1 {
			t.Errorf(`str %v, generated number : %v, expected(%v, %v)`, str, s.GeneratedNumber(), expected_str[i], i+1)
		}
		str = s.GetNextString()
	}
}

func TestTwoItems(t *testing.T) {
	s := New("", 2, Char_0_9)
	expected_str := []string{
		"00", "01", "02", "03", "04", "05", "06", "07", "08", "09",
		"10", "11", "12", "13", "14", "15", "16", "17", "18", "19",
		"20", "21", "22", "23", "24", "25", "26", "27", "28", "29",
		"30", "31", "32", "33", "34", "35", "36", "37", "38", "39",
		"40", "41", "42", "43", "44", "45", "46", "47", "48", "49",
		"50", "51", "52", "53", "54", "55", "56", "57", "58", "59",
		"60", "61", "62", "63", "64", "65", "66", "67", "68", "69",
		"70", "71", "72", "73", "74", "75", "76", "77", "78", "79",
		"80", "81", "82", "83", "84", "85", "86", "87", "88", "89",
		"90", "91", "92", "93", "94", "95", "96", "97", "98", "99"}

	if s.TotalNumber() != 100 {
		t.Errorf(`totalNums is error, %v`, s.TotalNumber())
	}

	str := s.GetFirstString()
	i := -1
	for str != "" {
		i++
		if str != expected_str[i] || int(s.GeneratedNumber()) != i+1 {
			t.Errorf(`str %v, generated number : %v, expected(%v, %v)`, str, s.GeneratedNumber(), expected_str[i], i+1)
		}
		str = s.GetNextString()
	}
}

func TestNoPrefixAndNoItems(t *testing.T) {
	s := New("", 0, []uint8{})

	if s.TotalNumber() != 0 {
		t.Errorf(`s should be zero, %v`, s.TotalNumber())
	}

	s1 := New("a", 0, []uint8{})
	if s1.TotalNumber() != 1 {
		t.Errorf(`s1 should be 1, %v`, s1.TotalNumber())
	}
	str := s1.GetFirstString()
	if str != "a" {
		t.Errorf(`s1 str should be "a", %v`, str)
	}
	str = s1.GetNextString()
	if str != "" {
		t.Errorf(`s1 str should be "", %v`, str)
	}

	s2 := New("abcd", 0, []uint8{'1', '2'})
	if s2.TotalNumber() != 1 {
		t.Errorf(`s2 should be 1, %v`, s2.TotalNumber())
	}

	//s2.ShowInfo(os.Stdout)
}

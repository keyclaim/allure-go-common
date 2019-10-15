package beans

import (
	"encoding/xml"
	"time"
)

const NsModel = `urn:model.allure.qatools.yandex.ru`

type Suite struct {
	XMLName   xml.Name    `xml:"ns2:test-suite"`
	NsAttr    string      `xml:"xmlns:ns2,attr"`
	Start     int64       `xml:"start,attr"`
	End       int64       `xml:"end,attr"`
	Name      string      `xml:"name"`
	Title     string      `xml:"title"`
	TestCases []*TestCase `xml:"test-cases>test-case"`
}

func NewSuite(name string, start time.Time) *Suite {
	s := new(Suite)

	s.NsAttr = NsModel
	s.Name = name
	s.Title = name

	if !start.IsZero() {
		s.Start = start.UTC().UnixNano() / int64(time.Millisecond)
	} else {
		s.Start = time.Now().UTC().UnixNano() / int64(time.Millisecond)
	}

	return s
}

// set end time for suite
func (s *Suite) EndSuite(endTime time.Time) {
	if !endTime.IsZero() {
		//strict UTC
		s.End = endTime.UTC().UnixNano() / int64(time.Millisecond)
	} else {
		s.End = time.Now().UTC().UnixNano() / int64(time.Millisecond)
	}
}

//suite has test-cases?
func (s Suite) HasTests() bool {
	return len(s.TestCases) > 0
}

//add test in suite
func (s *Suite) AddTest(test *TestCase) {
	if len(s.TestCases) > 0 {
		test.Prev = s.TestCases[len(s.TestCases)-1]
	}

	s.TestCases = append(s.TestCases, test)
}

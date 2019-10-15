package beans

import (
	"strings"
	"time"
)

type TestCase struct {
	Status      string        `xml:"status,attr"`
	Start       int64         `xml:"start,attr"`
	Stop        int64         `xml:"stop,attr"`
	Name        string        `xml:"name"`
	Steps       []*Step       `xml:"steps>step"`
	Labels      []*Label      `xml:"labels>label"`
	Attachments []*Attachment `xml:"attachments>attachment"`
	Desc        string        `xml:"description"`
	Prev        *TestCase
	Failure     struct {
		Msg   string `xml:"message"`
		Trace string `xml:"stack-trace"`
	} `xml:"failure,omitempty"`
}

//start new test case
func NewTestCase(name string, start time.Time) *TestCase {
	test := new(TestCase)
	test.Name = name

	if !start.IsZero() {
		test.Start = start.UnixNano() / int64(time.Millisecond)
	} else {
		test.Start = time.Now().UnixNano() / int64(time.Millisecond)
	}

	return test
}

type Label struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

func (t *TestCase) SetDescription(desc string) {
	t.Desc = desc
}

func (t *TestCase) AddLabel(label *Label) {
	t.Labels = append(t.Labels, label)
}

func (t *TestCase) AddStep(step *Step) {
	t.Steps = append(t.Steps, step)
}

func (t *TestCase) AddAttachment(attach *Attachment) {
	t.Attachments = append(t.Attachments, attach)
}

func (t *TestCase) End(status string, err error, end time.Time) {
	if !end.IsZero() {
		t.Stop = end.UnixNano() / int64(time.Millisecond)
	} else {
		t.Stop = time.Now().UnixNano() / int64(time.Millisecond)
	}
	t.Status = status
	if err != nil {
		msg := strings.Split("\n", err.Error())
		t.Failure.Msg = msg[0]
		t.Failure.Trace = strings.Join(msg[1:], "\n")
	}
}

package forth

import (
	"forth/internal/stdcapture"
	"testing"

	"github.com/stretchr/testify/suite"
)

type forthTestSuite struct {
	suite.Suite
	capturer *stdcapture.StdoutCapture
}

func TestRunner(t *testing.T) {
	suite.Run(t, new(forthTestSuite))
}

func (t *forthTestSuite) SetupTest() {
	t.capturer = stdcapture.New()
}

func (t *forthTestSuite) TearDownTest() {
	t.capturer = nil
}

func (t *forthTestSuite) TestScriptResult() {
	// Start capturing str output
	t.capturer.StartCapture()
	defer t.capturer.StopCapture()

	// testing two words (functions), add, sub and print
	err := Run(`:fna 37; 15 30 + . 50 * :fn 80 .; fn fn . fna .`)

	output := t.capturer.StopCapture()
	expectedResult := 
`45
80
80
80
37
`

	t.Equal(expectedResult, output)

	t.Nil(err)
}

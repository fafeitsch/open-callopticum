package report

import (
	"bytes"
	"github.com/fafeitsch/open-callopticum/cdrcsv"
	"testing"
)

func TestApplyCountings(t *testing.T) {
	def, err := ParseDefinitionFromFile("../mockdata/reportDefinition.json")
	file, err := cdrcsv.ReadWithoutHeaderFromFile("../mockdata/smallcdr.csv")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	expectedMap := make(map[string]int)
	expectedMap["production_calls"] = 17
	expectedMap["employees"] = 4
	expectedMap["evening_hours"] = 3
	actualMap := applyCountings(def.Countings, file.Records)
	if len(expectedMap) != len(actualMap) {
		t.Errorf("Expected fields were %d, but actual fields were %d.", len(expectedMap), len(actualMap))
		return
	}
	for key, value := range expectedMap {
		if actualMap[key] != value {
			t.Errorf("Expected value for field \"%s\" was %d, but actual was %d.", key, value, actualMap[key])
		}
	}
}

func TestGenerateReport(t *testing.T) {
	writer := new(bytes.Buffer)
	settings := Settings{
		Writer:        writer,
		ReportDefFile: "../mockdata/reportDefinition.json",
		TemplateFile:  "../mockdata/reportTemplate.tmpl",
		CdrFile:       "../mockdata/smallcdr.csv",
	}
	err := GenerateReport(settings)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	actual := writer.String()
	expected := `This is a sample call detail report:

Production Calls: 17
Calls in the evening hours: 3
Calls from Magdalene Greenman and Farlie Brager: 4`
	if actual != expected {
		t.Errorf("Expected text:\n%s\n\nActual text:\n%s", expected, actual)
	}
}

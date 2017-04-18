package assist

import "testing"
import "github.com/DATA-DOG/godog/gherkin"

func TestParseInstance(t *testing.T) {
	table := &gherkin.DataTable{
		Rows: []*gherkin.TableRow{
			{Cells: []*gherkin.TableCell{
				{Value: "First Name"},
				{Value: "John"},
			}},
			{Cells: []*gherkin.TableCell{
				{Value: "Last Name"},
				{Value: "Doe"},
			}},
			{Cells: []*gherkin.TableCell{
				{Value: "Age"},
				{Value: "35"},
			}},
		},
	}

	result := ParseInstance(table)
	if result["First Name"] != "John" {
		t.Error("expected First Name to be John, but got", result["First Name"])
	}
	if result["Last Name"] != "Doe" {
		t.Error("expected Last Name to be Doe, but got", result["Last Name"])
	}
	if result["Age"] != "35" {
		t.Error("expected Age to be 35, but got", result["Age"])
	}
}

func TestParseList(t *testing.T) {
	table := &gherkin.DataTable{
		Rows: []*gherkin.TableRow{
			{Cells: []*gherkin.TableCell{
				{Value: "First Name"},
				{Value: "Last Name"},
				{Value: "Age"},
			}},
			{Cells: []*gherkin.TableCell{
				{Value: "John"},
				{Value: "Doe"},
				{Value: "35"},
			}},
			{Cells: []*gherkin.TableCell{
				{Value: "Gwen"},
				{Value: "Stacy"},
				{Value: "22"},
			}},
		},
	}

	result := ParseList(table)
	if len(result) != 2 {
		t.Error("expected 2 lines in the table")
	}

	john := result[0]
	if john["First Name"] != "John" || john["Last Name"] != "Doe" || john["Age"] != "35" {
		t.Error("expected John Doe, age 35, but got", john)
	}

	gwen := result[1]
	if gwen["First Name"] != "Gwen" || gwen["Last Name"] != "Stacy" || gwen["Age"] != "22" {
		t.Error("expected Gwen Stacy, age 22, but got", gwen)
	}
}

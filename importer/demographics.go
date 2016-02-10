package importer

import (
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/projectcypress/cdatools/models"
)

func ExtractDemographics(patient *models.Record, patientElement xml.Node) {
	var firstNameXPath = xpath.Compile("cda:name/cda:given")
	patient.First = FirstElementContent(firstNameXPath, patientElement)
	var lastNameXPath = xpath.Compile("cda:name/cda:family")
	patient.Last = FirstElementContent(lastNameXPath, patientElement)
	var genderXPath = xpath.Compile("cda:administrativeGenderCode/@code")
	patient.Gender = FirstElementContent(genderXPath, patientElement)
	var birthTimeXPath = xpath.Compile("cda:birthTime/@value")
	patient.Birthdate = GetTimestamp(birthTimeXPath, patientElement)
	var raceXPath = xpath.Compile("cda:raceCode/@code")
	patient.Race.Code = FirstElementContent(raceXPath, patientElement)
	var raceCodeSetXPath = xpath.Compile("cda:raceCode/@codeSystemName")
	patient.Race.CodeSet = FirstElementContent(raceCodeSetXPath, patientElement)
	var ethnicityXPath = xpath.Compile("cda:ethnicGroupCode/@code")
	patient.Ethnicity.Code = FirstElementContent(ethnicityXPath, patientElement)
	var ethnicityCodeSetXPath = xpath.Compile("cda:ethnicGroupCode/@codeSystemName")
	patient.Ethnicity.CodeSet = FirstElementContent(ethnicityCodeSetXPath, patientElement)
}

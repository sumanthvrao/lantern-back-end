package validation

import (
	"github.com/onc-healthit/lantern-back-end/endpointmanager/pkg/capabilityparser"
	"github.com/onc-healthit/lantern-back-end/endpointmanager/pkg/endpointmanager"
)

// from https://www.hl7.org/fhir/codesystem-FHIR-version.html
// looking at official and release versions only
var dstu2 = []string{"1.0.1", "1.0.2"}
var stu3 = []string{"3.0.0", "3.0.1"}
var r4 = []string{"4.0.0", "4.0.1"}

// Validator @TODO this will be updated once I know what I'm doing
type Validator interface {
	RunValidation(capabilityparser.CapabilityStatement, int, []string, string) endpointmanager.Validation
	CapStatExists(capabilityparser.CapabilityStatement) endpointmanager.Rule
	MimeTypeValid([]string, string) endpointmanager.Rule
	HTTPResponseValid(int) endpointmanager.Rule
}

// GetValidationForVersion checks the given fhir version and then runs the validation checks
// specific to that version
// To note: All but the newR4Val() function returns the base validation currently
func GetValidationForVersion(fhirVersion string) Validator {
	if fhirVersion == "" {
		return newUnknownVal()
	}

	if contains(dstu2, fhirVersion) {
		return newDSTU2Val()
	} else if contains(stu3, fhirVersion) {
		return newSTU3Val()
	} else if contains(r4, fhirVersion) {
		return newR4Val()
	}

	return newUnknownVal()
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

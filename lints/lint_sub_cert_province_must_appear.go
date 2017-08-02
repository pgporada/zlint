package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertProvinceMustAppear struct {
	// Internal data here
}

func (l *subCertProvinceMustAppear) Initialize() error {
	return nil
}

func (l *subCertProvinceMustAppear) CheckApplies(c *x509.Certificate) bool {
	//Check if GivenName or Surname fields are filled out
	return util.IsSubscriberCert(c)
}

func (l *subCertProvinceMustAppear) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.Subject.GivenName != "" || len(c.Subject.Organization) > 0 || c.Subject.Surname != "" {
		if len(c.Subject.Locality) == 0 {
			if len(c.Subject.Province) == 0 {
				return ResultStruct{Result: Error}, nil
			} else {
				return ResultStruct{Result: Pass}, nil
			}
		}
	}
	return ResultStruct{Result: NA}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_province_must_appear",
		Description:   "Subscriber Certificate: subject:stateOrProvinceName MUST appeear if the subject:organizationName, subject:givenName, or subject:surname fields are present and subject:localityName is absent.",
		Providence:    "CAB: 7.1.4.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCertProvinceMustAppear{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.ESubCertProvinceMustAppear = result },
	})
}
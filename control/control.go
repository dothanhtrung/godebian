/*
 * Copyright (C) 2017 Do Thanh Trung <dothanhtrung.16@gmail.com>
 */

package control

import (
	"gitlab.com/kimtinh/godebian/deb822"
)

type Control struct {
	Source           string
	Maintainer       map[string]string
	Section          string
	Priority         string
	StandardsVersion string
	BuildDepends     []map[string]string
	XSTestsuite      string
	Packages         []deb822.Package
}

/*
 * Parse debian/control
 */
func Parse(pathfile string) Control {
	var control Control
	sections, _ := deb822.FindAll(pathfile, "", "")

	control.Source = sections[0]["Source"]
	control.Section = sections[0]["Section"]
	control.Priority = sections[0]["Priority"]
	control.StandardsVersion = sections[0]["Standards-Version"]
	control.XSTestsuite = sections[0]["XS-Testsuite"]

	for i := 1; i < len(sections); i++ {
		control.Packages = append(control.Packages, deb822.Deb822ToPackage(sections[i]))
	}

	return control
}

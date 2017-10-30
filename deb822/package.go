/*
 * Copyright (C) 2017 Do Thanh Trung <dothanhtrung.16@gmail.com>
 */

package deb822

import (
	"strings"
)

type Package struct {
	Package          string
	Binary           string
	Version          string
	Maintainer       map[string]string
	Uploaders        []map[string]string
	BuildDepends     []map[string]string
	Architecture     []string
	StandardsVersion string
	Format           string
	Files            []map[string]string
	VcsBrowser       string
	VcsSvn           string
	VcsGit           string
	ChecksumsSha256  []map[string]string
	Homepage         string
	PackageList      string
	Directory        string
	Priority         string
	Section          string
}

func Deb822ToPackage(section map[string]string) Package {
	var pkg Package
	pkg.Package = section["Package"]
	pkg.Binary = section["Binary"]
	pkg.Version = section["Version"]
	pkg.StandardsVersion = section["Standards-Version"]
	pkg.Format = section["Format"]
	pkg.VcsBrowser = section["Vcs-Browser"]
	pkg.VcsSvn = section["Vcs-Svn"]
	pkg.VcsGit = section["Vcs-Git"]
	pkg.Homepage = section["Homepage"]
	pkg.PackageList = section["Package-List"]
	pkg.Directory = section["Directory"]
	pkg.Priority = section["Priority"]
	pkg.Section = section["Section"]

	maintainer := strings.Split(section["Maintainer"], "<")
	maintainer_name := strings.TrimSpace(maintainer[0])
	maintainer_mail := strings.TrimSpace(strings.Split(maintainer[1], ">")[0])
	pkg.Maintainer = map[string]string{"Name": maintainer_name, "Email": maintainer_mail}

	uploaders := strings.Split(section["Uploaders"], ",")
	for _, uploader := range uploaders {
		namemail := strings.Split(uploader, "<")

		if len(namemail) > 1 {
			uploader_name := strings.TrimSpace(namemail[0])
			uploader_mail := strings.TrimSpace(strings.Split(namemail[1], ">")[0])
			pkg.Uploaders = append(pkg.Uploaders, map[string]string{"Name": uploader_name,
				"Email": uploader_mail})
		}
	}

	builddeps := strings.Split(section["Build-Depends"], ",")
	for _, dep := range builddeps {
		namecondition := strings.Split(dep, "(")
		if len(namecondition) > 1 {
			packagename := strings.TrimSpace(namecondition[0])
			condition := strings.TrimSpace(strings.Split(namecondition[1], ")")[0])
			pkg.BuildDepends = append(pkg.BuildDepends, map[string]string{"Package": packagename,
				"Condition": condition})
		}
	}

	pkg.Architecture = strings.Fields(section["Architecture"])

	files := strings.Split(section["Files"], "\n")
	for _, file := range files {
		msn := strings.Fields(file)
		pkg.Files = append(pkg.Files, map[string]string{"md5sum": msn[0],
			"size": msn[1],
			"name": msn[2]})
	}

	sha256s := strings.Split(section["Checksums-Sha256"], "\n")
	for _, sha256 := range sha256s {
		ssn := strings.Fields(sha256)
		pkg.ChecksumsSha256 = append(pkg.ChecksumsSha256, map[string]string{"sha256": ssn[0],
			"size": ssn[1],
			"name": ssn[2]})
	}

	return pkg
}

func FindPackage(pathfile, skey, svalue string, limit uint) ([]Package, error) {
	sections, err := find(pathfile, skey, svalue, limit)
	var pkgs []Package
	for _, section := range sections {
		pkgs = append(pkgs, Deb822ToPackage(section))
	}
	return pkgs, err
}

func FindAllPackage(pathfile, skey, svalue string) ([]Package, error) {
	sections, err := find(pathfile, skey, svalue, 0)
	var pkgs []Package
	for _, section := range sections {
		pkgs = append(pkgs, Deb822ToPackage(section))
	}
	return pkgs, err
}

func FindOnePackage(pathfile, skey, svalue string) (Package, error) {
	section, err := FindOne(pathfile, skey, svalue)
	pkg := Deb822ToPackage(section)
	return pkg, err
}

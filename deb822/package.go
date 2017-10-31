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
	Description      string
	Directory        string
	Essential        string
	Format           string
	Homepage         string
	MultiArch        string
	PackageList      string
	Priority         string
	Section          string
	StandardsVersion string
	VcsBrowser       string
	VcsGit           string
	VcsSvn           string
	Version          string

	Architecture []string
	Conflicts    []string
	Replaces     []string

	Maintainer map[string]string

	BuildDepends    []map[string]string
	ChecksumsSha256 []map[string]string
	Depends         []map[string]string
	Files           []map[string]string
	PreDepends      []map[string]string
	Uploaders       []map[string]string
}

func Deb822ToPackage(section map[string]string) Package {
	var pkg Package
	pkg.Package = section["Package"]
	pkg.Binary = section["Binary"]
	pkg.Description = section["Description"]
	pkg.Directory = section["Directory"]
	pkg.Essential = section["Essential"]
	pkg.Format = section["Format"]
	pkg.Homepage = section["Homepage"]
	pkg.MultiArch = section["Multi-Arch"]
	pkg.PackageList = section["Package-List"]
	pkg.Priority = section["Priority"]
	pkg.Section = section["Section"]
	pkg.StandardsVersion = section["Standards-Version"]
	pkg.VcsBrowser = section["Vcs-Browser"]
	pkg.VcsSvn = section["Vcs-Svn"]
	pkg.VcsGit = section["Vcs-Git"]
	pkg.Version = section["Version"]

	pkg.Architecture = strings.Fields(section["Architecture"])

	conflicts := strings.Split(section["Conflicts"], ",")
	for _, conflict := range conflicts {
		pkg.Conflicts = append(pkg.Conflicts, conflict)
	}

	replaces := strings.Split(section["Replaces"], ",")
	for _, replace := range replaces {
		pkg.Replaces = append(pkg.Replaces, replace)
	}

	maintainer := strings.Split(section["Maintainer"], "<")
	if len(maintainer) > 1 {
		maintainer_name := strings.TrimSpace(maintainer[0])
		maintainer_mail := strings.TrimSpace(strings.Split(maintainer[1], ">")[0])
		pkg.Maintainer = map[string]string{"Name": maintainer_name, "Email": maintainer_mail}
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

	sha256s := strings.Split(section["Checksums-Sha256"], "\n")
	for _, sha256 := range sha256s {
		ssn := strings.Fields(sha256)
		if len(ssn) > 2 {
			pkg.ChecksumsSha256 = append(pkg.ChecksumsSha256, map[string]string{"sha256": ssn[0],
				"size": ssn[1],
				"name": ssn[2]})
		}
	}

	deps := strings.Split(section["Depends"], ",")
	for _, dep := range deps {
		namecondition := strings.Split(dep, "(")
		if len(namecondition) > 1 {
			packagename := strings.TrimSpace(namecondition[0])
			condition := strings.TrimSpace(strings.Split(namecondition[1], ")")[0])
			pkg.Depends = append(pkg.Depends, map[string]string{"Package": packagename,
				"Condition": condition})
		}
	}

	files := strings.Split(section["Files"], "\n")
	for _, file := range files {
		msn := strings.Fields(file)
		if len(msn) > 2 {
			pkg.Files = append(pkg.Files, map[string]string{"md5sum": msn[0],
				"size": msn[1],
				"name": msn[2]})
		}
	}

	predeps := strings.Split(section["Pre-Depends"], ",")
	for _, dep := range predeps {
		namecondition := strings.Split(dep, "(")
		if len(namecondition) > 1 {
			packagename := strings.TrimSpace(namecondition[0])
			condition := strings.TrimSpace(strings.Split(namecondition[1], ")")[0])
			pkg.PreDepends = append(pkg.PreDepends, map[string]string{"Package": packagename,
				"Condition": condition})
		}
	}

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

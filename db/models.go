// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql/driver"
	"fmt"
)

type Pkgtype string

const (
	PkgtypeGo   Pkgtype = "go"
	PkgtypeRust Pkgtype = "rust"
	PkgtypeJava Pkgtype = "java"
	PkgtypeNode Pkgtype = "node"
)

func (e *Pkgtype) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Pkgtype(s)
	case string:
		*e = Pkgtype(s)
	default:
		return fmt.Errorf("unsupported scan type for Pkgtype: %T", src)
	}
	return nil
}

type NullPkgtype struct {
	Pkgtype Pkgtype
	Valid   bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPkgtype) Scan(value interface{}) error {
	if value == nil {
		ns.Pkgtype, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Pkgtype.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPkgtype) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.Pkgtype, nil
}

type Vcstype string

const (
	VcstypeBzr    Vcstype = "bzr"
	VcstypeFossil Vcstype = "fossil"
	VcstypeGit    Vcstype = "git"
	VcstypeHg     Vcstype = "hg"
	VcstypeSvn    Vcstype = "svn"
)

func (e *Vcstype) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Vcstype(s)
	case string:
		*e = Vcstype(s)
	default:
		return fmt.Errorf("unsupported scan type for Vcstype: %T", src)
	}
	return nil
}

type NullVcstype struct {
	Vcstype Vcstype
	Valid   bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVcstype) Scan(value interface{}) error {
	if value == nil {
		ns.Vcstype, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Vcstype.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVcstype) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.Vcstype, nil
}

type Package struct {
	ID          int64   `db:"id" json:"id"`
	PkgName     string  `db:"pkg_name" json:"pkg_name"`
	PkgType     Pkgtype `db:"pkg_type" json:"pkg_type"`
	Vcs         Vcstype `db:"vcs" json:"vcs"`
	Url         string  `db:"url" json:"url"`
	Description string  `db:"description" json:"description"`
}

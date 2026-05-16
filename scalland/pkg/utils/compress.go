package utils

import (
	"strings"

	"github.com/mholt/archiver/v3"
)

// CompressTBZ2 - a function to compress one or more file(s) into a compressed TBZ2 (.tar.bz2) archive
func (u *Utils) CompressTBZ2(sources []string, destination string) error {
	a := archiver.NewTarBz2()
	a.CompressionLevel = 1 // This is the best compression-level for BZ2 format
	err := a.Archive(sources, destination)
	if err != nil {
		u.Logger.Errorf("Error archiving file(s) %#v into %s. Details: %s", sources, destination, err.Error())
	}
	return err
}

// CompressTXZ - a function to compress one or more file(s) into a compressed TXZ (.tar.xz) archive
func (u *Utils) CompressTXZ(sources []string, destination string) error {
	a := archiver.NewTarXz()
	err := a.Archive(sources, destination)
	if err != nil {
		u.Logger.Errorf("Error archiving file(s) %s into %s. Details: %s", strings.Join(sources, ","), destination, err.Error())
	}
	return err
}

// ExtractTBZ2 unarchives TBZ2
func (u *Utils) ExtractTBZ2(source string, destination string) error {
	a := archiver.NewTarBz2()
	err := a.Unarchive(source, destination)
	if err != nil {
		u.Logger.Errorf("Error Decompressing file: %s into %s. Reason: %s", source, destination, err.Error())
	}
	return err
}

// ExtractTXZ unarchives TXZ
func (u *Utils) ExtractTXZ(source string, destination string) error {
	a := archiver.NewTarXz()
	err := a.Unarchive(source, destination)
	if err != nil {
		u.Logger.Errorf("Error Decompressing file: %s into %s. Reason: %s", source, destination, err.Error())
	}
	return err
}

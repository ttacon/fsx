package fsx

import (
	"errors"
	"os"
)

type File struct {
	Name string
	Mode os.FileMode
}

func (f File) Exists() error {
	finfo, err := os.Stat(f.Name)
	if err != nil {
		return err
	}

	if finfo.IsDir() {
		return errors.New("found a directory, not a file")
	}
	return nil
}

func (f File) Create() error {
	return f.create(false)
}

func (f File) ForceCreate() error {
	return f.create(true)
}

func (f File) create(force bool) error {
	if err := f.Exists(); err == nil && !force {
		return errors.New("file already exists")
	}

	ff, err := os.Create(f.Name)
	if err != nil {
		return err
	}

	if f.Mode > 0 && f.Mode != 0666 {
		if err = ff.Chmod(f.Mode); err != nil {
			ff.Close()
			return err
		}
	}

	return ff.Close()
}

type Dir struct {
	Name string
	Mode os.FileMode
}

func (d Dir) Exists() error {
	finfo, err := os.Stat(d.Name)
	if err != nil {
		return err
	}

	if !finfo.IsDir() {
		return errors.New("found a file, not a directory")
	}
	return nil
}

func (d Dir) Create() error {
	if d.Mode == 0 {
		d.Mode = os.ModeDir | 0766
	}
	return os.MkdirAll(d.Name, d.Mode)
}

type Symlink struct {
	Src    string // The name we want to make.
	Target string // The name we want to point to.
}

func (s Symlink) Exists() error {
	if found, err := os.Readlink(s.Src); err != nil {
		return err
	} else if found != s.Target {
		return errors.New("found symlink but it points somewhere else")
	}
	return nil
}

func (s Symlink) Create() error {
	return os.Symlink(s.Target, s.Src)
}

type FSXable interface {
	Exists() error
	Create() error
}

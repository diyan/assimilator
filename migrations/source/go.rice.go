package source

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	rice "github.com/GeertJohan/go.rice"
	"github.com/mattes/migrate/source"
)

func init() {
	source.Register("go.rice", &GoRice{})
}

type GoRice struct {
	box        *rice.Box
	migrations *source.Migrations
}

func (r *GoRice) Open(url string) (source.Driver, error) {
	return nil, errors.New("not yet implemented")
}

func WithInstance(instance interface{}) (source.Driver, error) {
	box, ok := instance.(*rice.Box)
	if !ok {
		return nil, errors.Errorf("unexpected type %T, *rice.Box is expected", instance)
	}
	driver := GoRice{
		box:        box,
		migrations: source.NewMigrations(),
	}
	err := box.Walk("", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".sql" {
			m, err := source.Parse(path)
			if err != nil {
				return errors.Errorf("error parse file %v", path)
			}
			if !driver.migrations.Append(m) {
				return errors.Errorf("error append migration %v", path)
			}
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "error scan file tree")
	}
	return &driver, nil
}

func (r *GoRice) Close() error {
	return nil
}

func (r *GoRice) First() (version uint, err error) {
	v, ok := r.migrations.First()
	if !ok {
		return 0, &os.PathError{
			Op:   "first",
			Path: r.box.Name(),
			Err:  os.ErrNotExist,
		}
	}
	return v, nil
}

func (r *GoRice) Prev(version uint) (prevVersion uint, err error) {
	v, ok := r.migrations.Prev(version)
	if !ok {
		return 0, &os.PathError{
			Op:   fmt.Sprintf("prev for version %v", version),
			Path: r.box.Name(),
			Err:  os.ErrNotExist,
		}
	}
	return v, nil
}

func (r *GoRice) Next(version uint) (nextVersion uint, err error) {
	v, ok := r.migrations.Next(version)
	if !ok {
		return 0, &os.PathError{
			Op:   fmt.Sprintf("next for version %v", version),
			Path: r.box.Name(),
			Err:  os.ErrNotExist,
		}
	}
	return v, nil
}

func (r *GoRice) ReadUp(version uint) (reader io.ReadCloser, identifier string, err error) {
	if m, ok := r.migrations.Up(version); ok {
		reader, err := r.box.Open(m.Raw)
		if err != nil {
			return nil, "", err
		}
		return reader, m.Identifier, nil
	}
	return nil, "", &os.PathError{
		Op:   fmt.Sprintf("read version %v", version),
		Path: r.box.Name(),
		Err:  os.ErrNotExist,
	}
}

func (r *GoRice) ReadDown(version uint) (reader io.ReadCloser, identifier string, err error) {
	if m, ok := r.migrations.Down(version); ok {
		reader, err := r.box.Open(m.Raw)
		if err != nil {
			return nil, "", err
		}
		return reader, m.Identifier, nil
	}
	return nil, "", &os.PathError{
		Op:   fmt.Sprintf("read version %v", version),
		Path: r.box.Name(),
		Err:  os.ErrNotExist,
	}
}

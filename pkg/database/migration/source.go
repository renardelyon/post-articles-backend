package migration

import (
	"net/url"

	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/file"
)

func FileDriver(fileName string) SourceFunc {
	return func() (string, source.Driver, error) {
		uri := url.URL{
			Scheme: fileScheme,
			Path:   fileName,
		}
		fileSource := &file.File{}
		drv, err := fileSource.Open(uri.String())
		return fileScheme, drv, err
	}
}

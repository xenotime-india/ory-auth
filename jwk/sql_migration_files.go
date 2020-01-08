// Code generated for package jwk by go-bindata DO NOT EDIT. (@generated)
// sources:
// migrations/sql/cockroach/4.sql
// migrations/sql/mysql/4.sql
// migrations/sql/postgres/4.sql
// migrations/sql/shared/1.sql
// migrations/sql/shared/2.sql
// migrations/sql/shared/3.sql
// migrations/sql/tests/.gitkeep
// migrations/sql/tests/1_test.sql
// migrations/sql/tests/2_test.sql
// migrations/sql/tests/3_test.sql
// migrations/sql/tests/4_test.sql
package jwk

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _migrationsSqlCockroach4Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\x51\x4b\xc3\x30\x14\x85\x9f\x7b\x7f\xc5\x79\x6c\xb1\x03\x11\xf6\xb4\xa7\x68\xef\x20\xd8\xb5\x35\x4d\x70\x7b\x1a\x61\x09\x2e\x16\xbb\x91\x86\xcd\xfd\x7b\x51\x64\x0c\x04\x5f\xef\x3d\xdf\x81\xf3\xcd\x66\xb8\xfb\x08\x6f\xd1\x26\x0f\x73\xa4\x27\xc5\x42\x33\xb4\x78\xac\x19\x72\x89\xa6\xd5\xe0\xb5\xec\x75\x8f\xfd\xc5\x45\xbb\x7d\x3f\x0f\xc8\x29\x9b\x82\xc3\xc9\xc6\xdd\xde\xc6\xfc\x61\x3e\x2f\x7e\x82\x8d\xa9\xeb\x92\xb2\xe1\x9f\xdf\xc9\xc7\x29\x1c\x46\x84\x31\x5d\xcf\xa8\x78\x29\x4c\xad\x71\xff\x0d\xfb\x8b\xb3\xc9\x22\xf9\xcf\x74\x0b\xee\xa2\xb7\xc9\xbb\xad\x4d\xd0\x72\xc5\xbd\x16\xab\xee\x6f\x43\xd3\xbe\xe6\x45\x49\xd9\x71\x40\xcf\x4a\x8a\x1a\x9d\x92\x2b\xa1\x36\x78\xe6\x4d\x49\x99\x69\xe4\x8b\x61\xc8\xa6\xe2\x35\xf2\x29\xb8\x12\x43\x70\x05\x15\x0b\xa2\x5b\x13\xd5\xe1\x3c\x52\xa5\xda\xee\xd7\xc4\x75\xfb\x82\xbe\x02\x00\x00\xff\xff\xe3\x74\x1b\x99\x31\x01\x00\x00")

func migrationsSqlCockroach4SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlCockroach4Sql,
		"migrations/sql/cockroach/4.sql",
	)
}

func migrationsSqlCockroach4Sql() (*asset, error) {
	bytes, err := migrationsSqlCockroach4SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/cockroach/4.sql", size: 305, mode: os.FileMode(420), modTime: time.Unix(1575291413, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlMysql4Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xb1\x6e\x83\x30\x14\x45\x77\xbe\xe2\x8e\xad\x5a\xbe\x80\xc9\xc5\x4f\x95\x55\x78\xa6\xae\x2d\x95\xc9\x42\x72\xd5\x38\x56\x12\x42\x12\x91\xfc\x7d\xa4\x2c\x81\x01\xf6\xab\x7b\xce\xc9\x73\xbc\xed\xe2\xff\xd0\x9d\xff\xe0\xfa\x4c\x54\x96\x0c\xac\xf8\xa8\x08\x9b\x5b\x18\x3a\xbf\x1d\x13\xa4\xd1\x0d\x1a\xa3\x6a\x61\x5a\x7c\x51\x5b\x64\xa5\x21\x61\x09\x8e\xd5\xb7\x23\x28\x96\xf4\xfb\xdc\xfb\x18\xae\x3e\x06\x7f\x39\x42\xf3\xe4\xe6\xe5\x14\xc3\x3b\x52\x0c\xaf\xc5\x02\x48\x48\x89\x3e\x41\xb1\x85\xe3\x1f\xf5\xc9\x24\x21\x9c\xd5\x5e\x71\x69\xa8\x26\xb6\x73\x8d\x6c\xaa\x2f\x0f\xe3\x7e\x2d\xa0\xd4\x95\xab\x19\x7d\x5a\xa2\x3f\x56\x8b\x2d\x6b\xce\x13\xa9\x59\xe5\x3d\x00\x00\xff\xff\x5d\xb2\x7a\x0d\x5e\x01\x00\x00")

func migrationsSqlMysql4SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlMysql4Sql,
		"migrations/sql/mysql/4.sql",
	)
}

func migrationsSqlMysql4Sql() (*asset, error) {
	bytes, err := migrationsSqlMysql4SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/mysql/4.sql", size: 350, mode: os.FileMode(420), modTime: time.Unix(1575291413, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlPostgres4Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x8f\xb1\x6a\xc3\x30\x10\x86\x77\x3d\xc5\x8d\x36\xad\x9f\x40\x93\x6a\xdd\x20\x2a\xcb\xae\x2c\x41\x3d\x09\x83\x4c\xab\x8a\xb6\x8a\x93\xe0\xf8\xed\x03\x21\x10\x2d\x0e\x19\x32\x1e\x77\xff\xfd\xdf\x57\x55\xf0\xf2\x1b\xbe\xe6\xf1\x30\x81\x4d\x84\x49\x83\x1a\x0c\x7b\x93\x08\xdf\xab\x9f\x47\xf7\xb3\x44\xe0\xba\xed\xa0\x6e\x55\x6f\x34\x13\xca\xdc\x36\x2e\xc5\x69\xa5\x1b\x29\xc6\x39\xa4\x08\x3d\x6a\xc1\xe4\xbd\xa3\x4e\x8b\x86\xe9\x01\xde\x71\x80\x22\xc5\x92\x92\x5a\x23\x33\x08\x56\x89\x0f\x8b\x20\x14\xc7\xcf\xac\x34\xf8\x93\x0b\xde\x1d\x77\xd0\xaa\xec\x55\xb1\x0f\xfe\x15\x62\xf0\x25\x25\x24\xf7\xe2\xff\xcb\xdf\x53\xcd\xae\x29\x69\x1b\x05\x29\x52\x72\x99\x37\x29\x1f\x56\xcf\xf8\xcf\x01\x00\x00\xff\xff\x54\x0e\x5a\x87\x97\x01\x00\x00")

func migrationsSqlPostgres4SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlPostgres4Sql,
		"migrations/sql/postgres/4.sql",
	)
}

func migrationsSqlPostgres4Sql() (*asset, error) {
	bytes, err := migrationsSqlPostgres4SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/postgres/4.sql", size: 407, mode: os.FileMode(420), modTime: time.Unix(1575291413, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlShared1Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\x41\x6b\x84\x30\x14\x84\xcf\xbe\x5f\x31\x47\xa5\xbb\x50\x0a\x7b\xda\x53\x5a\x23\x48\x53\x95\x18\xa1\x9e\x4a\x68\x42\x4d\xa5\x5a\x62\xd0\xfa\xef\x8b\xa5\x88\x97\x7d\xc7\x37\xf3\xc1\x7c\xe7\x33\xee\xbe\xdc\x87\xd7\xc1\xa2\xf9\xa6\x27\xc9\x99\xe2\x50\xec\x51\x70\xe4\x19\x8a\x52\x81\xbf\xe6\xb5\xaa\xd1\xad\xc6\xeb\xb7\xcf\xa5\x47\x4c\xd1\xe4\x0c\xb6\x9b\xb5\x7f\xef\xb4\x8f\x1f\x2e\x97\xe4\xaf\x5c\x34\x42\x9c\x28\xea\x9d\x41\x74\x2b\x9c\xad\x9f\xdc\x38\xc0\x0d\x61\x7f\x23\xe5\x19\x6b\x84\xc2\xfd\x46\xdb\xd5\xe8\xa0\x11\xec\x4f\x38\x82\x95\xcc\x5f\x98\x6c\xf1\xcc\x5b\xc4\x93\x33\x27\xf4\xce\x24\x94\x5c\x89\x8e\x1e\xe9\xb8\x0c\x94\xca\xb2\xfa\xf7\xd8\x97\x5f\xe9\x37\x00\x00\xff\xff\x2d\x5b\x1e\xa7\xef\x00\x00\x00")

func migrationsSqlShared1SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlShared1Sql,
		"migrations/sql/shared/1.sql",
	)
}

func migrationsSqlShared1Sql() (*asset, error) {
	bytes, err := migrationsSqlShared1SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/shared/1.sql", size: 239, mode: os.FileMode(420), modTime: time.Unix(1575291413, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlShared2Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xc8\xa8\x4c\x29\x4a\x8c\xcf\x2a\xcf\x56\x70\x74\x71\x51\x48\x2e\x4a\x4d\x2c\x49\x4d\x89\x4f\x2c\x51\x08\xf1\xf4\x75\x0d\x0e\x71\xf4\x0d\x50\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\x51\x70\x71\x75\x73\x0c\xf5\x09\x51\xf0\xf3\x0f\xd7\xd0\xb4\xe6\xe2\x42\x36\xd6\x25\xbf\x3c\x0f\x87\xc1\x2e\x41\xfe\x01\x0a\xce\xfe\x3e\xa1\xbe\x7e\x48\x16\x58\x73\x01\x02\x00\x00\xff\xff\x5c\xc3\x60\x41\x96\x00\x00\x00")

func migrationsSqlShared2SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlShared2Sql,
		"migrations/sql/shared/2.sql",
	)
}

func migrationsSqlShared2Sql() (*asset, error) {
	bytes, err := migrationsSqlShared2SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/shared/2.sql", size: 150, mode: os.FileMode(420), modTime: time.Unix(1575291413, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlShared3Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\x71\xf5\x71\x0d\x71\x55\x70\x0b\xf2\xf7\x55\xc8\xa8\x4c\x29\x4a\x8c\xcf\x2a\xcf\x56\x08\xf7\x70\x0d\x72\x55\x28\xce\x4c\xb1\x55\x07\x0b\xea\xe5\x17\xa4\xe6\x65\xa6\xe8\x65\xa6\xe8\x96\xe4\x67\xa7\xe6\xa9\x5b\x73\x71\x21\x1b\xe4\x92\x5f\x9e\xc7\x05\x08\x00\x00\xff\xff\xaa\xd8\x2c\x82\x5a\x00\x00\x00")

func migrationsSqlShared3SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlShared3Sql,
		"migrations/sql/shared/3.sql",
	)
}

func migrationsSqlShared3Sql() (*asset, error) {
	bytes, err := migrationsSqlShared3SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/shared/3.sql", size: 90, mode: os.FileMode(420), modTime: time.Unix(1575291413, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTestsGitkeep = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func migrationsSqlTestsGitkeepBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTestsGitkeep,
		"migrations/sql/tests/.gitkeep",
	)
}

func migrationsSqlTestsGitkeep() (*asset, error) {
	bytes, err := migrationsSqlTestsGitkeepBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/.gitkeep", size: 0, mode: os.FileMode(420), modTime: time.Unix(1575291413, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTests1_testSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xf2\xf4\x0b\x76\x0d\x0a\x51\xf0\xf4\x0b\xf1\x57\xc8\xa8\x4c\x29\x4a\x8c\xcf\x2a\xcf\x56\xd0\x28\xce\x4c\xd1\x51\xc8\x06\x11\x65\xa9\x45\xc5\x99\xf9\x79\x3a\x0a\xd9\xa9\x95\x29\x89\x25\x89\x9a\x0a\x61\x8e\x3e\xa1\xae\xc1\x0a\x1a\xea\x86\xba\xc5\x99\x29\xea\x3a\x0a\xea\x86\xba\xd9\x60\x86\x81\x8e\x82\x7a\x71\x7e\x6e\xaa\x6e\x76\x6a\xa5\xba\xa6\x35\x17\x17\xb2\x65\x2e\xf9\xe5\x79\x5c\x80\x00\x00\x00\xff\xff\x4c\x05\x9c\x3d\x7e\x00\x00\x00")

func migrationsSqlTests1_testSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTests1_testSql,
		"migrations/sql/tests/1_test.sql",
	)
}

func migrationsSqlTests1_testSql() (*asset, error) {
	bytes, err := migrationsSqlTests1_testSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/1_test.sql", size: 126, mode: os.FileMode(420), modTime: time.Unix(1575291413, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTests2_testSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcd\xb1\xaa\xc2\x30\x14\x87\xf1\x3d\x4f\xf1\xdf\xd2\x72\x4f\xe0\xd2\xd5\x49\xb0\x43\x41\x5a\xb0\xad\x8e\xe5\x60\x0e\x1a\x43\x5b\x49\x82\x25\x6f\x2f\x3a\xb9\x7c\x7c\xdb\xcf\x18\xfc\xcd\xee\x16\x38\x09\xc6\xa7\x6a\xda\xbe\x3e\x0d\x68\xda\xa1\xc3\x3d\xdb\xc0\xd3\x63\xf3\x28\xa2\xb3\x04\xff\xc9\x4b\x42\x74\xeb\x42\xf0\x92\x2d\x27\x26\x5c\x83\x70\x12\x3b\x71\x2a\x71\xde\x1f\xc7\xba\x47\xa1\x2b\x13\x9d\xd5\x04\x5d\x19\xff\x9d\x7f\x82\x8e\xeb\x2c\xc6\x4b\xd6\x84\xb6\xbb\x14\x65\xb9\x53\xea\xd7\x3f\xac\xdb\xa2\xde\x01\x00\x00\xff\xff\x58\xd8\xae\x35\x91\x00\x00\x00")

func migrationsSqlTests2_testSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTests2_testSql,
		"migrations/sql/tests/2_test.sql",
	)
}

func migrationsSqlTests2_testSql() (*asset, error) {
	bytes, err := migrationsSqlTests2_testSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/2_test.sql", size: 145, mode: os.FileMode(420), modTime: time.Unix(1575291413, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTests3_testSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcd\xb1\xaa\xc2\x30\x14\x87\xf1\x3d\x4f\xf1\xdf\xd2\x72\x4f\xe0\x42\x47\x27\xc1\x0e\x05\x69\xc1\xb6\x3a\x96\x83\x39\x68\x0c\x6d\x25\x09\x96\xbc\xbd\xe8\xe4\xf2\xf1\x6d\x3f\x63\xf0\x37\xbb\x5b\xe0\x24\x18\x9f\xaa\x69\xfb\xfa\x34\xa0\x69\x87\x0e\xf7\x6c\x03\x4f\x8f\xcd\xa3\x88\xce\x12\xfc\x27\x2f\x09\xd1\xad\x0b\xc1\x4b\xb6\x9c\x98\x70\x0d\xc2\x49\xec\xc4\xa9\xc4\x79\x7f\x1c\xeb\x1e\x85\xae\x4c\x74\x56\x13\x74\x65\xfc\x77\xfe\x09\x3a\xae\xb3\x18\x2f\x59\x13\xda\xee\x52\x94\xe5\x4e\xa9\x5f\xff\xb0\x6e\x8b\x7a\x07\x00\x00\xff\xff\xd8\xf8\x31\xf0\x91\x00\x00\x00")

func migrationsSqlTests3_testSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTests3_testSql,
		"migrations/sql/tests/3_test.sql",
	)
}

func migrationsSqlTests3_testSql() (*asset, error) {
	bytes, err := migrationsSqlTests3_testSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/3_test.sql", size: 145, mode: os.FileMode(420), modTime: time.Unix(1575291413, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTests4_testSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcd\xb1\xaa\xc2\x30\x14\x87\xf1\x3d\x4f\xf1\xdf\xd2\x72\x4f\xe0\x0e\xdd\x9c\x04\x3b\x14\xa4\x05\xdb\xea\x58\x0e\xe6\xa0\x31\xb4\x95\x24\x58\xf2\xf6\xa2\x93\xcb\xc7\xb7\xfd\x8c\xc1\xdf\xec\x6e\x81\x93\x60\x7c\xaa\xa6\xed\xeb\xd3\x80\xa6\x1d\x3a\xdc\xb3\x0d\x3c\x3d\x36\x8f\x22\x3a\x4b\xf0\x9f\xbc\x24\x44\xb7\x2e\x04\x2f\xd9\x72\x62\xc2\x35\x08\x27\xb1\x13\xa7\x12\xe7\xfd\x71\xac\x7b\x14\xba\x32\xd1\x59\x4d\xd0\x95\xf1\xdf\xf9\x27\xe8\xb8\xce\x62\xbc\x64\x4d\x68\xbb\x4b\x51\x96\x3b\xa5\x7e\xfd\xc3\xba\x2d\xea\x1d\x00\x00\xff\xff\x9b\x11\x7f\xc5\x91\x00\x00\x00")

func migrationsSqlTests4_testSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTests4_testSql,
		"migrations/sql/tests/4_test.sql",
	)
}

func migrationsSqlTests4_testSql() (*asset, error) {
	bytes, err := migrationsSqlTests4_testSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/4_test.sql", size: 145, mode: os.FileMode(420), modTime: time.Unix(1575291413, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"migrations/sql/cockroach/4.sql":  migrationsSqlCockroach4Sql,
	"migrations/sql/mysql/4.sql":      migrationsSqlMysql4Sql,
	"migrations/sql/postgres/4.sql":   migrationsSqlPostgres4Sql,
	"migrations/sql/shared/1.sql":     migrationsSqlShared1Sql,
	"migrations/sql/shared/2.sql":     migrationsSqlShared2Sql,
	"migrations/sql/shared/3.sql":     migrationsSqlShared3Sql,
	"migrations/sql/tests/.gitkeep":   migrationsSqlTestsGitkeep,
	"migrations/sql/tests/1_test.sql": migrationsSqlTests1_testSql,
	"migrations/sql/tests/2_test.sql": migrationsSqlTests2_testSql,
	"migrations/sql/tests/3_test.sql": migrationsSqlTests3_testSql,
	"migrations/sql/tests/4_test.sql": migrationsSqlTests4_testSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"migrations": &bintree{nil, map[string]*bintree{
		"sql": &bintree{nil, map[string]*bintree{
			"cockroach": &bintree{nil, map[string]*bintree{
				"4.sql": &bintree{migrationsSqlCockroach4Sql, map[string]*bintree{}},
			}},
			"mysql": &bintree{nil, map[string]*bintree{
				"4.sql": &bintree{migrationsSqlMysql4Sql, map[string]*bintree{}},
			}},
			"postgres": &bintree{nil, map[string]*bintree{
				"4.sql": &bintree{migrationsSqlPostgres4Sql, map[string]*bintree{}},
			}},
			"shared": &bintree{nil, map[string]*bintree{
				"1.sql": &bintree{migrationsSqlShared1Sql, map[string]*bintree{}},
				"2.sql": &bintree{migrationsSqlShared2Sql, map[string]*bintree{}},
				"3.sql": &bintree{migrationsSqlShared3Sql, map[string]*bintree{}},
			}},
			"tests": &bintree{nil, map[string]*bintree{
				".gitkeep":   &bintree{migrationsSqlTestsGitkeep, map[string]*bintree{}},
				"1_test.sql": &bintree{migrationsSqlTests1_testSql, map[string]*bintree{}},
				"2_test.sql": &bintree{migrationsSqlTests2_testSql, map[string]*bintree{}},
				"3_test.sql": &bintree{migrationsSqlTests3_testSql, map[string]*bintree{}},
				"4_test.sql": &bintree{migrationsSqlTests4_testSql, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

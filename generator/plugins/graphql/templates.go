// Code generated by go-bindata.
// sources:
// templates/schemas_body.gohtml
// templates/schemas_head.gohtml
// templates/types_body.gohtml
// templates/types_head.gohtml
// DO NOT EDIT!

package graphql

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesSchemas_bodyGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xcb\x6e\xdb\x3a\x10\x5d\xd3\x5f\x31\x30\x8c\x0b\xfb\x42\xd1\x07\x18\xc8\xca\x68\xd3\x2e\x9a\x47\x93\x5d\x51\x14\x8c\x3c\x56\xd4\xd2\xa4\x4b\x51\x2e\x82\x01\xff\xbd\xe0\x53\x94\x9d\x18\xd5\x8a\x9a\xe7\x39\x67\x38\x34\xaf\x07\x04\xa2\x45\xfd\xd8\xbc\xe0\x9e\xdf\xf2\x3d\x5a\x1b\xce\x1b\xd1\xa1\x34\x3d\xf4\x46\x0f\x8d\x01\x9a\x31\x22\xd0\x5c\xb6\x08\x8b\x1e\xf5\xb1\x6b\x10\xd6\xd7\xb0\xa8\x1f\xc3\x4f\x0f\x57\xd6\xce\x18\x23\x4a\xee\x3a\x94\x0b\x85\x80\xa8\x55\x4f\xae\x5d\x76\x07\xc7\x8d\xb7\xba\x54\x22\x40\xb9\xf5\x65\xec\x6c\x37\xc8\x06\x6e\xd0\xbc\x0d\x6e\xd9\x88\xfe\x32\xee\x0a\xba\x17\xf8\x9f\xa8\x93\x06\x75\x83\x07\xa3\x74\x7f\xff\xab\xb5\xb6\xfe\x3c\x5a\x3e\x71\xb9\x15\xa8\x2b\x30\x1a\x88\x8c\xe6\x0d\xea\x10\xf4\xe4\xcf\x2b\x58\x12\xb5\xbf\x45\xb0\x85\xfa\x15\xa0\xd6\x4a\xaf\xfe\x5d\x91\x23\x77\xd5\x4f\x54\xf9\xd8\xa1\xd8\xf6\x70\x0d\x44\xd1\xb1\x51\x32\x68\xad\xf4\x58\x6f\x61\xad\xe3\x5a\xbf\xa3\xaa\x63\x49\x74\x05\xdd\x0e\x16\x11\xf3\x07\xc9\x9f\x05\x7a\x19\x1d\x2f\xe7\x8d\xaa\xae\x22\x96\x1f\xbe\xeb\x9b\x78\x26\x43\x28\xe8\xa9\xe7\x9f\xd8\x98\xc0\xee\xce\x9f\x4f\xc8\x85\x80\x58\xcb\xd7\xcf\xb2\xdd\xe2\x9f\x90\x52\x6a\x19\x2c\x1b\x25\x77\x5d\x4b\x33\xc6\x98\xcb\x5c\xc3\xfc\xb4\xd4\xbc\x72\x4e\x22\xc7\x50\x66\x20\xf5\xc3\xa0\x0c\x6e\x37\x6a\xbf\x77\x57\x6b\x3e\x8f\x60\x18\x8b\xa6\x75\x81\x69\x12\x6b\x6d\x2a\x98\x59\x32\x16\xb8\xaf\x4b\xd0\xc1\xe4\x91\xc5\xee\x02\x65\x6e\x1f\x87\x97\x9a\x16\x42\xed\xc4\xd6\xab\xf4\x4e\x60\xac\xe5\xc2\xd2\x2d\xf1\x4e\xb8\xf0\x39\x4d\x5c\x7c\x14\xc4\x53\x2b\xf2\x27\xf3\xfb\x76\x12\xfc\xbd\x1a\xfb\xa2\xe8\xb1\x40\xc2\xce\xea\xfe\x77\xca\x9f\x52\x68\x31\x9d\x22\xa3\xca\x6e\xb7\xc1\x19\xd7\x5d\x39\xbf\x31\xe6\x2b\xf6\x4a\x1c\x71\x0d\x6e\xb3\x97\x87\x52\xec\xe8\xba\xe7\x9a\xef\xfb\x15\x2c\xfd\xce\xee\x78\x83\x64\xcb\x6d\x4b\x9f\x46\x33\x68\x19\x1f\x26\xb2\x2e\x48\x76\x22\xfb\xc7\x9e\xb6\x64\x3f\x8e\xfb\xec\xff\x4c\x9c\xb9\x54\x41\xcf\x4b\xaa\x44\x4d\x72\x68\x6a\x96\xc4\x18\xdf\x0d\xa3\x3b\xd9\x46\x77\xc4\x34\x45\xe0\x8d\x6e\x3f\x4b\x73\xa4\x39\xdd\xa4\xf8\xfc\x9d\xbd\x4a\xe3\x26\x3d\x0c\xa8\x5f\xfd\x30\x6a\x7f\x0c\xe3\x08\x93\x88\xb7\xaf\xfe\x32\x18\x6e\x3a\x25\x83\x2f\xa1\x48\xd6\x90\x3c\x8d\xc9\xf9\x19\x9e\x5d\xcd\xec\xdf\x00\x00\x00\xff\xff\x5a\xb6\x42\x9c\x3f\x06\x00\x00")

func templatesSchemas_bodyGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesSchemas_bodyGohtml,
		"templates/schemas_body.gohtml",
	)
}

func templatesSchemas_bodyGohtml() (*asset, error) {
	bytes, err := templatesSchemas_bodyGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/schemas_body.gohtml", size: 1599, mode: os.FileMode(420), modTime: time.Unix(1531402025, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSchemas_headGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xca\xb1\x4a\xc5\x30\x14\xc6\xf1\xd9\x3c\xc5\xc7\xa5\x83\x0e\x26\x20\x4e\x82\x83\x70\x1d\xee\x62\x1d\xf2\x02\xa7\xed\x31\x0d\xb6\x49\x4d\x52\x44\x0e\xe7\xdd\x05\x2d\xdc\xed\xcf\xef\xfb\x9c\x83\x9f\x63\xc5\x47\x5c\x18\xdf\x54\x11\x38\x71\xa1\xc6\x13\x86\x1f\x84\xd8\xe6\x7d\xb0\x63\x5e\x5d\xa5\xb6\x97\xf4\xc8\xc5\x6d\x25\xb7\xfc\x10\xbe\x16\x8b\x73\x8f\xb7\xde\xe3\xf5\x7c\xf1\xb8\x78\xb3\xd1\xf8\x49\x81\x21\xd2\xd9\xa3\x55\x8d\x89\xeb\x96\x4b\xc3\xad\x11\x29\x94\x02\xa3\x3b\xe4\xe9\x19\x9d\xfd\xef\x8a\x7b\x55\x73\x23\x72\x6c\xf6\x65\x89\x54\x55\x71\xba\xd2\x3b\xb5\x59\xf5\x64\x44\x38\x4d\x7f\xff\x3b\xf3\x1b\x00\x00\xff\xff\x76\x6d\x2f\x2d\xc0\x00\x00\x00")

func templatesSchemas_headGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesSchemas_headGohtml,
		"templates/schemas_head.gohtml",
	)
}

func templatesSchemas_headGohtml() (*asset, error) {
	bytes, err := templatesSchemas_headGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/schemas_head.gohtml", size: 192, mode: os.FileMode(420), modTime: time.Unix(1531402025, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTypes_bodyGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\xdd\x73\xe3\xb6\x11\x7f\x26\xff\x8a\x0d\x9b\xdc\x90\x19\x46\xee\xb3\x32\x7a\xf0\xb8\x71\x7a\x93\xde\x47\xcf\x9e\xe4\xc1\xf1\x38\x38\x0a\x92\x50\x51\x20\x0d\x52\xb2\x3d\x1c\xfe\xef\x1d\x7c\x91\x00\x09\x52\x94\x63\xb7\xd7\xce\xf1\xe1\x62\x81\x8b\xc5\x62\xf1\xdb\xc5\x7e\x30\x67\x67\xf0\x13\xdd\xef\x0a\xbf\xaa\x18\xa2\x6b\x0c\xdf\x62\xba\xdf\xc1\x7c\x01\xb3\x4b\x92\xe2\x99\x78\x09\x3f\xd4\xb5\xef\x1d\x10\x83\xaa\x12\xef\x67\xbf\x22\x46\xd0\xe7\x14\xbf\x47\x3b\x5c\xd7\xb0\x80\xaa\x5a\xdf\xa7\x1f\xb7\xeb\xba\x9e\xbd\xc7\x0f\x7c\x56\x68\x0c\xf1\xdf\x17\x19\x5d\x91\x75\xe5\x7b\x1e\x9f\x34\x07\xf5\x04\x9a\xe5\xcf\x0c\xe5\x9b\x7f\xfe\x43\x72\x0c\x62\xdf\xf3\xaa\x0a\xc8\x4a\x0a\x34\xbb\xc8\x76\x3b\x4c\x4b\x29\x89\xe7\xfd\x0d\x17\x09\x23\x79\x49\x32\x3a\x6f\x84\x52\x34\x75\xad\x26\x63\xba\x54\xf4\xbf\xa2\x74\x8f\x8b\x39\x74\x44\x12\xc3\x52\xae\x77\x28\xaf\x7c\x30\x9e\x46\x1f\x07\x4e\xc4\x15\xa2\x77\xce\x59\x69\x39\x3c\x2e\xbe\xa0\x98\x29\xc1\xe7\xf0\x66\x78\x95\x4a\xcc\x91\xe2\x08\xb9\xe5\x54\xf1\x5b\x8a\xed\xe9\x6d\x53\xbd\x74\xb3\xf5\x3f\x82\xe0\x8f\x66\xdd\xbe\x0a\x2c\x62\x93\x59\xab\x06\xcf\x93\xc3\x55\xd5\x8e\xf1\x91\x3a\xf2\xdb\x21\xdf\x3f\x3b\x83\xb7\x34\xdf\x97\x90\x7d\xfe\x17\x4e\x4a\xbf\xaa\x40\xe9\x42\x0e\xb4\xe8\x10\x64\x1f\xc4\xa0\x0d\x12\x49\x78\x0c\x26\xc6\x74\x13\x2d\xc6\xb0\x56\x9b\x3e\x15\x09\x9d\xa0\x5d\xa1\x8b\x1a\x4d\x78\x49\x70\xba\xb4\x4f\xbc\xc7\x56\xd0\xf0\x83\xaf\xe5\xbc\x3a\xf2\xbd\xd5\x9e\x26\x40\x28\x29\xc3\xa8\x12\x28\x52\x3b\x5f\x71\x5a\x81\x02\xb5\xb0\x5c\x40\x2b\x76\x68\xcb\x8a\x2c\x8c\x6e\xb8\xcc\x82\x89\x06\xca\x2d\x2c\x2c\xa8\x18\xe2\x89\x49\xd5\x47\x46\x0e\xa8\xc4\xed\x96\xad\xe9\x31\x5c\x3f\xe5\x02\x44\x09\x4a\x53\x25\xe0\x8c\x8f\xc1\xb7\xb5\x10\xca\x38\xe5\xda\x37\x70\xd0\x3d\xdf\x02\x18\x2e\xb2\xf4\x80\x59\x61\x1c\xb5\x1e\x73\x1e\xf6\x27\x3d\x41\x30\x14\x3a\xab\xaa\x66\xca\xec\x72\x4f\x13\x8e\x4b\x29\x6a\xa8\xec\x78\x76\xcd\x50\x82\xd9\x4f\x94\xab\x67\x09\x75\x0d\x25\xc7\x4a\x29\x46\xa5\x0e\x24\x45\x0c\x42\xf4\xba\x4e\xca\x47\xbe\xbf\xf2\x51\xbe\xbd\xc8\x68\x89\x1f\xcb\x18\x08\x10\x5a\x62\xb6\x42\x09\xae\xea\x08\xc2\x3b\x7e\xcc\x99\xdc\x7b\x23\xc4\x87\x7d\x99\xef\xcb\x9f\xc5\x70\x5d\xc7\xc0\x30\x63\x80\x19\xcb\x18\x3f\x59\xb7\x4c\xf2\x34\x8b\x1c\x51\xbe\xed\x92\xcd\x2e\x18\x46\x25\xbe\xd8\x90\x74\x79\x95\x23\x7a\xc9\xb2\x9d\x92\x22\x4c\xca\xc7\x58\x1c\xcb\xc0\xb6\x83\xc8\xf7\xbc\x25\x5e\x61\x06\x9c\xe1\xec\x92\x50\x52\x6c\xc2\x76\x94\xab\x4d\xa2\xcc\x23\x2b\xc8\xb9\x78\xf3\x05\x30\x9c\x64\x07\xcc\xc2\xe8\x47\x39\xf4\xcd\x02\x28\x49\x41\x7a\x0e\xc1\xe8\x0a\x97\xd7\x68\x1d\x06\x62\x2f\x41\x0c\x41\xc9\xf6\x38\x88\xec\xf1\xbb\x1d\x2e\x0a\xb4\xc6\x41\x2c\xd8\x74\xdf\x16\x25\x4a\xb6\x41\x0c\x45\xc9\x08\x5d\x87\x55\xb5\xc4\x9f\xf7\x6b\xa9\xe6\x2b\xfe\x2e\x8c\x22\x2e\xa9\x57\x2b\xf1\xd8\xcb\xc8\xc2\xd9\xcc\x7e\xe2\x83\x61\xcb\xbf\xe6\x4a\x31\xd0\x4a\x56\x40\x60\xd1\xae\xc5\x70\xb9\x67\x94\xff\x8c\xf9\x3f\x1c\xce\x1e\x62\xeb\x82\xab\x8b\xcc\xc2\x1d\xca\x6f\xe4\x3e\x6e\x4d\x58\xf8\xde\x1d\x2c\x80\xd3\x49\xaf\xc4\x70\xb1\x4f\x4b\x58\x00\xc5\x0f\xa1\xc6\xcb\x65\xc6\xde\xe3\x87\x41\xd4\x08\xb9\xa0\xe7\x02\xda\x13\x37\x9d\x00\x59\x89\xd5\x0c\x4b\x57\xce\x49\x18\x8e\x20\x6d\x4c\xdf\xd4\xa4\x84\x22\x29\xce\x19\x43\x4f\xda\x90\xa5\x04\x8d\xdf\x26\x02\x90\x53\xd9\xcf\xc2\x9b\x8e\x2a\x3c\xae\x45\xbe\xff\x59\x33\x5b\xee\xd4\x98\x07\x0b\xd8\xa1\x2d\x0e\x5b\x5b\x32\x45\xe1\x26\x94\x62\x1a\x12\x2a\x0f\xce\x5b\x65\x0c\x48\x0c\x07\x94\x0a\xd8\x0a\x1d\x11\xaa\xc0\xa1\xed\x4b\x72\xd0\xee\xe2\x37\x52\x6e\xc4\xe1\x43\x73\x8b\x1d\x62\x50\xc0\xb7\x1d\x99\xb8\x12\xf5\x3c\x08\x0e\x28\x0d\xb8\x5f\x53\xb3\xc8\x0a\x7a\x78\xf4\x6c\xa0\x54\x95\xc0\x5e\x21\x31\xfd\x1b\x43\x79\x88\x19\x8b\x21\x58\x21\xc2\x4d\xbd\xcc\xb4\xdb\x03\x62\x38\x43\x10\xcb\x07\x91\x62\xa9\x17\x3c\xaa\xbc\x1b\xc2\xfd\xf9\xa1\xbd\x74\xd3\x02\x1b\x97\xf5\xc4\xf9\x27\xe8\xa0\x7b\xb1\xfb\x8e\x75\x8f\x1c\x42\x23\xde\xb4\x43\x08\x93\x8c\x26\xa8\x84\x40\xc0\xf0\xf7\x20\x80\x31\x1c\x42\xf0\x7b\x70\x1b\x44\xad\xc0\xee\x33\x7b\xf1\x23\x53\xab\x4d\x41\xfb\xc1\x77\x1e\xd6\x94\xa9\xaf\xa7\x27\xfb\x5c\xad\x5f\xb5\x6f\xfd\x6e\x62\x93\x8c\xe2\x6c\x65\x3b\xa6\x0f\x14\x7f\x58\x59\xde\xa9\xa1\x26\x74\x89\x1f\x63\x2b\xa2\xe1\xf3\x7b\x01\x0d\x3f\xaf\x7b\x45\x0e\x7f\xd5\xe3\x64\x05\xc7\xfc\xcf\x5d\x0c\xd9\xf6\x14\x77\xf5\x23\xa7\x7f\xf3\xe6\x38\xe3\x16\x3c\xd0\x79\x26\x20\xbd\x3b\x85\x3f\xa7\x01\x7f\xfc\x18\xef\xd4\x21\xba\xd6\xe9\x62\xdf\x45\xc3\x9f\x67\x5a\x43\xc6\x4f\xdb\xb2\x86\xa3\xba\x0c\x22\xa7\x10\x7d\xf1\x4d\xf3\x70\xaa\xf0\x75\xb5\xd7\x86\x05\xa3\x47\x7e\x5e\x14\x64\x4d\x09\x5d\x73\x3d\xe5\x78\xf8\xc4\x5b\xe3\x96\xa8\x3f\x6e\xdc\x3d\xd6\xc1\x21\x18\x10\x75\x5c\x53\x53\x96\x3e\x38\xb9\xb6\xfe\xa0\xae\x2a\xbd\x86\x5c\xec\xab\x3d\x7e\xb5\x47\x5b\x85\x5f\xed\x71\xa2\xa6\x5e\xca\x1e\x75\x5a\xad\x32\x6e\x79\x35\xcb\x3f\x7c\x1d\xdf\xc8\xb5\x64\xfa\x62\x25\xe3\x67\x67\x20\xd7\xd5\xc9\xb8\xb3\xda\xf2\xad\xcc\xc0\x25\xe5\xf3\xeb\x2d\xfd\x52\x8b\x5d\x65\xd1\xa5\xb9\x91\xfa\x8a\xe7\x39\x2a\x2b\x72\xa8\x52\xa5\xa4\x7e\x0d\xa5\x9f\x41\x9d\x56\x44\x39\x5f\x2e\x05\xa5\x94\x33\x74\xd4\x42\xde\x74\xa5\x91\xf1\xe5\x50\xed\x44\xbc\x1c\xab\x9f\x48\x0a\x65\x33\x73\x99\xaa\xe7\xe6\x96\xd5\xab\x8f\x88\xa1\x5d\x11\x41\x68\xa4\x5b\xb1\xaa\x32\x18\x7e\xc5\x13\xff\x14\x0f\xa4\x4c\x36\x50\xb0\x84\xeb\x20\x9f\x5d\x65\x7b\x96\xe0\x59\x58\x3e\xe5\x38\xd2\x01\x71\x82\x0a\x0c\xdf\xb7\x39\x98\x3e\x07\x95\x84\xcd\x75\xae\x43\x56\x82\xd1\xa2\x93\x02\xf5\x93\x65\x3b\x91\xf1\x44\xd6\xfc\x7d\xc1\x12\x3d\x20\x22\x3d\xad\x1f\x8c\x97\x17\xa8\x28\x8d\xcc\xa5\xe1\xd8\x28\x91\x13\x5c\x67\xa2\xae\xd3\xf3\x32\x10\x14\xc2\x18\x23\x6b\xf1\xf6\xd2\xea\x33\x1d\x64\xd1\xe1\x60\x65\x3b\x4a\x4d\xc7\xb5\xf4\xd2\x9b\x63\xc9\x0b\x6c\x4f\x31\xe9\xf2\x30\xf7\xd7\xa4\x22\x43\xf7\xcf\x7b\xfc\x10\x06\x85\x00\x10\x64\x2b\xd8\xd3\x2d\xcd\x1e\x28\x70\x28\xa9\x4c\x48\x62\x58\x24\xff\x96\xf6\x46\xcc\xf1\x1d\xca\xbf\x5a\xe4\x7f\xce\x22\x5d\x97\x92\x7e\x7a\x66\xaa\x4a\x58\x70\x73\xeb\xae\x77\x69\xba\x55\xc6\x60\x8b\x9f\x44\x61\x46\xf6\x2d\xe4\x69\x8f\x58\x9a\x25\x6b\x01\x0b\xe0\xf7\x2a\x5d\x86\x0c\x17\x31\xb8\x57\x6b\x67\x78\xc1\x16\x3f\x05\x73\x00\xb1\xaa\x31\x2c\xd6\x0f\xe6\x52\x8e\xf6\x45\x1d\xf5\x5c\x52\x7b\x47\x9a\x0a\x9a\x68\xe1\xaf\xa3\x17\x65\xa2\x5f\xa4\x66\x5e\xcf\x37\x74\x23\x93\x77\x28\x2f\xac\x52\x8b\x3b\x3a\x91\xc1\xc9\x3b\x94\xff\x7f\xb4\x83\xae\x37\x7b\xba\x0d\x65\x75\x7e\xd2\x84\x4e\x02\xd1\xf8\xfe\x09\x9d\x27\xdf\xc0\xc9\x58\x33\x48\x6d\xdc\x76\x90\x6a\xcf\xbf\xe0\x27\x49\xd9\x34\x7f\xe4\xbe\x3d\x2d\x50\x03\xb8\xe7\xae\x20\xac\xc2\xb9\x86\x7e\x5a\x67\x56\x47\x4d\x33\xed\x18\x98\x26\xf6\x9d\x6c\x60\xd9\xad\x27\xef\x0b\xe8\x3d\x71\x3f\xe0\xe8\x3f\xfd\x82\x9f\xb4\xb3\xba\x75\xbc\x16\x3a\x1d\xea\x4e\x8d\xb6\xa7\x5e\xb2\x3f\x35\xd0\xa0\xea\x77\xa8\x26\xb7\xa8\xfe\x0b\x3d\x2a\xe5\x11\x9d\x5d\xaa\x17\x6a\x53\xc9\x25\x44\xa3\xca\x0a\xd5\xba\xad\x2a\xc7\x65\x2f\xc8\xa8\xea\x55\xf5\xfa\x32\xaa\x2b\x35\x57\xbd\x97\x17\xc0\x12\xe7\xaa\x7a\x33\xc4\xd9\x9c\xb9\x83\x05\x10\x5f\x5c\x9d\xe2\x2d\xa7\x1a\xe9\xa2\x79\xde\x36\x96\x55\x85\x03\x4a\x6f\x84\xaf\xba\x8d\xe5\xdf\xd2\xaf\xdc\x0a\xa6\x31\x70\xbe\xdb\xf8\xd0\x54\x8f\x2d\xf1\x87\xbb\x0f\xdb\x6d\xaf\xe8\xe3\x9a\x08\xc1\xb6\xed\x7e\x38\x7b\x09\xe3\xc5\x9a\xd5\x60\xb5\xe6\x2f\xdf\x2d\xb9\x0d\x03\x4e\xb1\xf8\xce\x82\xef\x30\x06\x12\x99\x7d\x15\x2b\xc0\xdf\x6e\x4f\x11\xd6\x8a\xec\x3b\x8a\xb1\x2a\x33\x7d\xd5\x1c\xfa\xf5\x30\xf7\xd4\xa6\x06\xf2\xea\xca\x91\x47\x3e\xaa\x9e\xc3\xe1\x34\x81\x2d\x05\x49\x7b\xb8\xd9\x6e\x6f\x17\x87\x83\xb2\x1e\x47\x0d\xa5\xff\x45\x83\xb8\x5e\xb2\xe3\xa5\x94\xe6\x4e\xf9\x02\x8a\x29\x3a\x17\x19\x2e\xa9\xb4\x1f\xe8\xf4\xab\x2a\x13\x53\x33\x09\x67\x77\x3a\xa6\x24\x14\x24\x7e\x2f\x01\x1b\x8a\x2f\x04\xe9\x9f\xce\xc4\x84\x73\xee\x26\x5e\x23\x5e\xc8\x9d\x5f\xb9\x92\xab\xda\x37\xde\x14\x2c\x69\x9c\x96\x22\x10\x1b\x50\xc1\xef\x44\x25\x6a\xd8\x8f\xaa\x51\x11\x0d\x2a\xd2\x11\x46\xfd\xef\xa9\x52\xf9\xfc\xbe\x32\x6b\xf3\x1b\xb2\xb3\x33\xb8\xc2\xec\x40\x12\x6c\x5a\x61\x21\x87\x5a\x33\xd4\x34\x46\x2c\xf7\x33\x2e\xab\x4a\x53\xaa\xf2\x80\x22\x7b\x87\xcb\x4d\xb6\x2c\xc2\xc4\x48\x0b\x35\xe1\x05\x4a\xd3\xb7\x7a\x9b\x3c\xa0\x22\x1b\x9e\xb3\x8b\x9d\x27\x38\x2f\x1b\x77\xf7\xb6\x1d\xf9\x3b\xa2\xcb\x14\x33\x18\x8c\x11\xe3\x81\x20\x51\xc7\x88\x51\xdf\x62\xc1\x88\xde\xb4\x6c\x4a\xee\xd6\xc5\x75\xf3\x83\xd6\xd8\x8d\x3e\xed\x4e\xcc\x12\xe5\x99\x01\x46\xf2\xeb\x43\x49\xe7\xfe\xfc\xd0\x28\xbc\x98\xee\xc8\x9e\xa2\x33\x50\x1b\xb3\x8a\x46\xb9\x2c\xe9\x2c\xed\x5a\x4c\xb3\x4f\x45\x7a\xce\xd6\x7b\x7e\x47\x14\x66\x25\xec\x9c\xad\x1d\x8e\x4d\x9a\x95\x9e\xc0\x05\xec\xd6\x3f\xcc\x3a\x15\x62\x6b\xa1\x86\xb1\x75\xa4\x2e\x10\x5b\xbb\x15\xa1\xe7\x38\xb3\x1c\x3e\xcb\xc8\x6b\x34\xc3\x5e\xcd\xd1\x33\x37\x6e\xbf\x3a\xc1\x84\xef\xc0\x32\x62\x23\xf0\x1f\x69\x44\xf1\x87\x67\x24\xc2\xc2\x55\x84\x3f\x4a\xcc\x83\xb1\xa4\x7c\x34\xea\xa7\xc3\xf9\x44\x13\x27\x1f\x4f\x2a\xf2\x36\x11\x0a\x7a\x66\x3a\xeb\x42\x0b\xf4\x75\xdf\x7c\x4f\x33\x98\x70\x18\xef\xcc\xac\xc3\x1b\x0e\xec\x5f\x32\xba\xb7\xab\x2e\x2a\xce\x57\x8a\xeb\x60\x80\x47\xfc\x9b\x7e\x91\xcf\x36\x85\x4f\xf8\x7e\x8f\x8b\x26\x57\x1d\x6c\x41\x59\x88\x67\xf8\xbe\x17\xee\x0d\xf0\x0b\x72\x8e\x68\x59\xbd\xab\xaa\x1f\xc4\xf7\xc2\x59\x39\x44\xae\x02\x4b\x2e\x85\xba\x34\x42\x4a\xd2\x48\x37\xa9\x8e\x09\x06\x27\x34\x4c\x3b\x1b\x6a\x2f\x14\xcc\xd8\xa4\x49\xd3\xc4\xe9\x94\xd7\xd5\xb6\x2f\x52\x82\x69\x29\x9d\x24\xbf\x0f\xb8\xa2\x92\x00\x02\x86\xef\xcd\x8f\xc7\x5c\xdf\x67\x09\xb7\x7c\xcf\x35\xef\xfa\x3a\xd0\x56\xab\x4e\xb3\xa6\xee\xff\xcf\x48\xd9\xf1\x3f\x0d\x08\x95\x2f\x78\xe3\xba\xe1\x94\x85\xb6\xd8\x54\xd7\xe7\xdc\x61\xb2\x41\xeb\xee\xa4\x44\x63\x17\x84\xe7\x49\x27\x36\x87\x3c\x76\x94\x2a\x5b\xf4\x92\x8d\xf6\x7a\x1c\xa5\x21\x11\xa5\x08\x61\xd6\x42\x70\xe7\xbd\xdc\xf8\x15\x8a\x1f\x4b\x70\x51\x18\x2c\xdf\xd2\x43\xb6\xc5\x2c\x82\x50\xe5\xcf\xdd\xb8\xc8\x8e\x8d\x5e\xcc\x40\x5d\xe7\xf9\xaa\xb6\x79\x04\xb1\xd2\xbe\xa6\x22\xd6\x6a\x3f\xb8\xe0\x15\x19\x4e\xce\xe9\x70\x9f\xf9\x3d\x04\xe2\xfa\x18\xac\x6e\x93\x8d\x88\xde\x14\x4c\x04\x8e\xa6\x62\x85\xdb\xac\x75\xf4\x83\xe0\xe1\x2b\x48\x84\x3f\x03\x3b\x4c\x7f\x41\xc3\xf0\xfd\xcc\xf8\x3e\xd6\xed\x17\xbc\x56\x89\xdf\x64\x5b\xf3\xba\x1a\x2f\xd9\x57\xd5\x6a\x57\xaa\xf2\x56\xce\x08\x2d\x57\x61\xf0\xc9\xd0\x20\x18\xbb\x52\x1e\xb0\x80\xcf\x68\xc9\xa5\xe2\xeb\x8b\x02\x7f\xf8\xdd\x75\x34\x83\xab\x4d\xb6\x4f\x97\xf0\x59\x04\x39\x63\xd2\x8a\xdb\xf0\xde\xb8\x05\x4d\xc8\x19\xf6\xf2\x11\x3d\xa5\x19\x5a\x8a\x4b\xf3\x62\x83\x93\xed\x44\x9b\x91\xce\xba\x18\xba\xd7\x06\x7d\xa1\xfb\x8b\x0e\xd7\xf3\x9c\x9b\x09\x3a\x9f\xf6\xbc\xdc\xe5\x24\xbf\x8f\xb4\xf6\xe8\x52\x5d\xc0\x70\x11\x88\xf6\xd2\x54\x71\x6d\xf7\x65\xf2\x3c\x4f\x12\x5c\x14\x23\x1f\x55\x0d\x3d\xdc\xbe\x2c\x4e\xd0\x3b\x1f\xe7\x3a\x4a\xf8\x53\x44\x1f\xfd\x0e\xe7\x14\xf9\x18\x2e\x4e\x5a\xd7\x76\x70\x26\xbc\xad\xde\xda\x14\x00\x18\x55\xae\x29\xab\x9f\x78\xf9\x77\xaf\xfe\x21\xd7\xac\x53\x10\xf5\xdf\xce\xd7\x48\x7e\xb7\x18\xd7\xa2\x7c\xe4\x7f\x08\xfa\x77\x00\x00\x00\xff\xff\x77\xb4\xb7\x26\x04\x38\x00\x00")

func templatesTypes_bodyGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTypes_bodyGohtml,
		"templates/types_body.gohtml",
	)
}

func templatesTypes_bodyGohtml() (*asset, error) {
	bytes, err := templatesTypes_bodyGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/types_body.gohtml", size: 14340, mode: os.FileMode(420), modTime: time.Unix(1531736988, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTypes_headGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xca\xb1\x4a\xc5\x30\x14\xc6\xf1\xd9\x3c\xc5\xc7\xa5\x83\x0e\x26\x20\x4e\x82\x83\x70\x1d\xee\x62\x1d\xf2\x02\xa7\xed\x31\x0d\xb6\x49\x4d\x52\x44\x0e\xe7\xdd\x05\x2d\xdc\xed\xcf\xef\xfb\x9c\x83\x9f\x63\xc5\x47\x5c\x18\xdf\x54\x11\x38\x71\xa1\xc6\x13\x86\x1f\x84\xd8\xe6\x7d\xb0\x63\x5e\x5d\xa5\xb6\x97\xf4\xc8\xc5\x6d\x25\xb7\xfc\x10\xbe\x16\x8b\x73\x8f\xb7\xde\xe3\xf5\x7c\xf1\xb8\x78\xb3\xd1\xf8\x49\x81\x21\xd2\xd9\xa3\x55\x8d\x89\xeb\x96\x4b\xc3\xad\x11\x29\x94\x02\xa3\x3b\xe4\xe9\x19\x9d\xfd\xef\x8a\x7b\x55\x73\x23\x72\x6c\xf6\x65\x89\x54\x55\x71\xba\xd2\x3b\xb5\x59\xf5\x64\x44\x38\x4d\x7f\xff\x3b\xf3\x1b\x00\x00\xff\xff\x76\x6d\x2f\x2d\xc0\x00\x00\x00")

func templatesTypes_headGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTypes_headGohtml,
		"templates/types_head.gohtml",
	)
}

func templatesTypes_headGohtml() (*asset, error) {
	bytes, err := templatesTypes_headGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/types_head.gohtml", size: 192, mode: os.FileMode(420), modTime: time.Unix(1531402025, 0)}
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
	"templates/schemas_body.gohtml": templatesSchemas_bodyGohtml,
	"templates/schemas_head.gohtml": templatesSchemas_headGohtml,
	"templates/types_body.gohtml": templatesTypes_bodyGohtml,
	"templates/types_head.gohtml": templatesTypes_headGohtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"schemas_body.gohtml": &bintree{templatesSchemas_bodyGohtml, map[string]*bintree{}},
		"schemas_head.gohtml": &bintree{templatesSchemas_headGohtml, map[string]*bintree{}},
		"types_body.gohtml": &bintree{templatesTypes_bodyGohtml, map[string]*bintree{}},
		"types_head.gohtml": &bintree{templatesTypes_headGohtml, map[string]*bintree{}},
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


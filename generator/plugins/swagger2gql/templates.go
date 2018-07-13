// Code generated by go-bindata.
// sources:
// generator/plugins/swagger2gql/templates/array_value_resolver.gohtml
// DO NOT EDIT!

package swagger2gql

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

var _templatesArray_value_resolverGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x41\x6b\xe3\x30\x10\x85\xcf\xd2\xaf\x98\x98\xc0\xca\xac\xd7\x3f\x60\x21\xa7\xd2\x6b\x29\xa5\x24\x87\x10\xca\xe0\x8e\x5d\x61\x45\x0a\x23\xb9\x6d\x10\xfa\xef\x45\x72\xdc\xe6\x52\x4a\x7d\xf3\x48\xef\x9b\xf7\x9e\xfa\xc9\x76\xa0\x90\x07\xd0\x36\x10\xf7\xd8\x51\x4c\x35\xa8\x18\x3b\x34\x06\xd6\x2d\x93\x9f\x4c\x78\x3c\x9f\x28\xa5\x06\x88\xd9\x71\x0d\x51\x0a\x32\x74\x24\x1b\x7c\x03\x6e\x84\xff\x1b\x40\x1e\x5a\xb5\x3f\x5c\x53\xa4\xd0\x3d\xac\xdc\x98\xaf\x0b\xa6\x30\xb1\x05\xab\x4d\x03\x9f\xf0\x82\xf3\xf7\xe3\x90\x52\x7b\x47\x6f\xaa\x42\x1e\xa6\x8c\x05\xed\xc1\xba\x00\xc8\x8c\xe7\xaa\x96\x22\x49\xc1\xe4\xf3\xa2\x23\x8e\xf4\x9d\x3d\x43\x56\x2d\xc6\xea\x5a\x8a\xde\x31\xe8\x06\x2e\xa3\xac\x66\xb4\x03\x2d\x03\x9f\x8d\x3d\xc1\x66\xf9\x97\x42\xc4\x08\xba\xcf\xce\x0c\x1d\x1f\xc8\x3b\xf3\x4a\xbc\xd3\xe1\xe5\x96\x19\xfe\xa5\x24\x85\x10\x64\xb6\x68\x4a\x15\x19\xf8\x95\xe5\x4a\x01\xd5\x85\x58\x65\x8b\xce\x85\x9b\xf0\x3e\x8b\x75\x5f\x84\xab\x4d\x6e\x22\x4a\xb8\xfa\x7e\xae\x68\xc7\x78\x52\xc4\xdc\x40\xd5\xa1\xfd\x13\x80\xe7\x7d\x73\x4d\x4b\x8a\x5c\x97\x10\x65\x1b\x93\xdf\xeb\x43\x09\xb8\x45\x33\xc7\x23\xe3\x69\x89\x72\x39\xff\xfb\xdb\x18\x19\x63\x9f\x67\x4a\x79\x99\x62\x9c\xc9\x37\xd9\xbd\x4c\x2a\xc6\x75\x8b\x3c\xa4\x54\x7f\x04\x00\x00\xff\xff\x54\x31\x15\x4c\x61\x02\x00\x00")

func templatesArray_value_resolverGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesArray_value_resolverGohtml,
		"templates/array_value_resolver.gohtml",
	)
}

func templatesArray_value_resolverGohtml() (*asset, error) {
	bytes, err := templatesArray_value_resolverGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/array_value_resolver.gohtml", size: 609, mode: os.FileMode(420), modTime: time.Unix(1530883563, 0)}
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
	"templates/array_value_resolver.gohtml": templatesArray_value_resolverGohtml,
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
		"array_value_resolver.gohtml": &bintree{templatesArray_value_resolverGohtml, map[string]*bintree{}},
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

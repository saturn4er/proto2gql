// Code generated by go-bindata.
// sources:
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

var _templatesTypes_bodyGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5f\x73\xe3\xb6\x11\x7f\xa6\x3e\xc5\x86\x4d\x6e\xc8\x0e\x43\xf7\x59\x19\x3d\x78\xdc\xf8\x7a\x93\xfa\xee\x7a\xf6\x24\x0f\x8e\xc7\xc1\x51\x90\x84\x8a\x02\x69\x90\x92\xed\xe1\xf0\xbb\x77\xf0\x8f\x04\x48\x90\xa2\x5c\x5f\x9a\x4c\x4f\x0f\xc9\x19\x5c\x2c\x16\x8b\xdf\x2e\xf6\x0f\xce\xce\xe0\x47\xba\xdf\x15\xb3\xaa\x62\x88\xae\x31\x7c\x8b\xe9\x7e\x07\xf3\x05\xc4\x97\x24\xc5\xb1\xf8\x08\xdf\xd7\xf5\xcc\x3b\x20\x06\x55\x25\xbe\xc7\x3f\x23\x46\xd0\xe7\x14\xbf\x47\x3b\x5c\xd7\xb0\x80\xaa\x5a\x3f\xa4\x1f\xb7\xeb\xba\x8e\xdf\xe3\x47\x3e\x2b\x30\x86\xf8\xdf\x17\x19\x5d\x91\x75\x35\xf3\x3c\x3e\x69\x0e\xea\xe7\x6b\x96\x6f\x19\xca\x37\xff\xfa\xa7\xe4\xe8\x47\x33\xcf\xab\x2a\x20\x2b\x29\x50\x7c\x91\xed\x76\x98\x96\x52\x12\xcf\xfb\x3b\x2e\x12\x46\xf2\x92\x64\x74\xde\x08\xa5\x68\xea\x5a\x4d\xc6\x74\xa9\xe8\x7f\x46\xe9\x1e\x17\x73\xe8\x88\x24\x86\xa5\x5c\x57\x28\xaf\x66\x60\xfc\x1a\x7d\x1c\x38\x11\x57\x88\xde\x39\x67\xa5\xe5\xf0\xb8\xf8\x82\x22\x56\x82\xcf\xe1\xcd\xf0\x2a\x95\x98\x23\xc5\x11\x72\xcb\xa9\xe2\x6f\x29\xb6\xa7\xb7\x4d\xf5\xd2\xcd\xd6\x7f\xf3\xfd\xdf\x9a\x75\xfb\x2a\xb0\x88\x4d\x66\xad\x1a\x3c\x4f\x0e\x57\x55\x3b\xc6\x47\xea\x70\xd6\x0e\xcd\x66\x67\x67\xf0\x8e\xe6\xfb\x12\xb2\xcf\xff\xc6\x49\x39\xab\x2a\x50\xba\x90\x03\x2d\x3a\x04\xd9\x07\x31\x68\x83\x44\x12\x1e\x83\x89\x31\xdd\x44\x8b\x31\xac\xd5\xa6\x4f\x45\x42\xc7\x6f\x57\xe8\xa2\x46\x13\x5e\x12\x9c\x2e\xed\x13\xef\xb1\x15\x34\x57\x28\xbf\xd9\xec\xe9\x36\x58\xed\x69\x12\x84\x93\x26\x80\x0d\x15\x86\xcb\x3d\xa3\x93\x66\xda\x13\x2d\x9c\xad\x38\x89\xc0\x99\xda\x9a\xdc\x82\x50\x6b\x77\x92\x36\x1c\x31\xc7\x8d\x3c\x43\x04\xc1\x48\x69\xf2\xe6\x39\x17\xc8\x4b\x50\x9a\xaa\x35\x63\x3e\x06\xdf\xd6\x1c\x33\x7d\xe9\x34\x2c\xcc\xd1\xf6\xaf\x3a\x94\x73\x04\x82\x1a\xa4\x75\x11\x54\x00\xc3\x45\x96\x1e\x30\x2b\x0c\x30\xe9\x31\x27\x9c\x3e\xe9\x09\x82\x21\x3f\x1b\x0e\x2b\x3d\x25\xbe\xdc\xd3\x84\x23\x5f\x6e\x3d\x50\x9e\x22\xbe\x61\x28\xc1\xec\x47\xca\x31\xb7\x84\xba\x86\x92\xa3\xb1\x14\xa3\x52\x2d\x92\x22\x92\x1b\xab\xeb\xa4\x7c\xe2\xca\x28\x9f\xe4\xd7\x8b\x8c\x96\xf8\xa9\x8c\x80\x00\xa1\x25\x66\x2b\x94\xe0\xaa\x0e\x21\xb8\xe7\xa7\x9b\x49\x45\x35\x42\x7c\xd8\x97\xf9\xbe\x7c\x2b\x86\xeb\x3a\x02\x86\x19\x03\xcc\x58\xc6\xc2\x6a\xe6\xb9\x65\x92\x46\x57\xe4\x88\xf2\x6d\x97\x2c\xbe\x60\x18\x95\xf8\x62\x43\xd2\xe5\x75\x8e\xe8\x25\xcb\x76\x4a\x8a\x20\x29\x9f\x22\x71\xcc\x03\xdb\xf6\xc3\x99\xe7\x2d\xf1\x0a\x33\xe0\x0c\xe3\x4b\x42\x49\xb1\x09\xda\x51\x09\x69\xe1\x74\xc8\x0a\x72\x2e\xde\x7c\x01\x0c\x27\xd9\x01\xb3\x20\xfc\x41\x0e\x7d\xb3\x00\x4a\x52\x90\xbe\x49\x30\xba\xc6\xe5\x0d\x5a\x07\xbe\xd8\x8b\x1f\x81\x5f\xb2\x3d\xf6\x43\x7b\xfc\x7e\x87\x8b\x02\xad\xb1\x1f\x09\x36\xdd\xaf\x45\x89\x92\xad\x1f\x41\x51\x32\x42\xd7\x41\x55\x2d\xf1\xe7\xfd\x5a\xaa\xf9\x9a\x7f\x0b\xc2\x90\x4b\xea\xd5\x4a\x3c\xf6\x3a\xb2\x70\x36\xf1\x8f\x7c\x30\x68\xf9\xd7\x5c\x29\x86\xd7\x23\x2b\x20\xb0\x68\xd7\x52\x16\x4c\x49\x1a\xf1\xff\xcc\xf8\x24\xc4\xd6\x05\x57\x17\x89\x83\x1d\xca\x6f\xe5\x3e\xee\x4c\x58\xcc\xbc\x7b\x58\x00\xa7\x93\x7e\x8f\xe1\x62\x9f\x96\xb0\x00\x8a\x1f\x03\x8d\x97\xcb\x8c\xbd\xc7\x8f\x83\xa8\x11\x72\x41\xcf\x05\xb4\x27\xde\x3a\x01\xa1\x25\xbe\xda\x6d\x6b\xfa\xca\xfd\x09\xc3\x11\xa4\x0a\x19\x77\x96\x26\x25\x14\x49\x71\xce\x18\x7a\xd6\x56\x2f\x25\x68\x6e\x06\x22\x00\x39\x95\x7d\x1c\xdc\x76\x54\xe1\x71\x2d\xf2\xfd\xc7\xcd\x6c\xb9\x53\x63\x1e\x2c\x60\x87\xb6\x38\x68\x6d\xc9\x14\x85\x9b\x50\x8a\x69\x40\xa8\x3c\x38\x6f\x95\x31\x20\x11\x1c\x50\x2a\x60\x2b\x74\x44\xa8\x02\x87\xb6\x2f\xc9\x41\xbb\x8b\x5f\x48\xb9\x11\x87\x0f\xcd\x3d\x79\x88\x40\x01\xdf\xf6\x7a\xe2\xd2\xd5\xf3\xc0\x3f\xa0\xd4\xe7\x4e\x50\xcd\x22\x2b\xe8\xe1\xd1\xb3\x81\x52\x55\x02\x7b\x85\xc4\xf4\x2f\x0c\xe5\x01\x66\x2c\x02\x7f\x85\x08\x37\xf5\x32\xd3\x6e\x0f\x88\xe1\x0c\x41\x2c\xef\x87\x8a\xa5\x5e\xf0\xa8\xf2\x6e\xc9\x1d\x2c\xe0\xd0\x5e\xeb\x69\x81\x8d\x70\x60\xe2\xfc\x13\x74\xd0\x0d\x1d\x66\x8e\x75\x8f\x1c\x42\x23\xde\xb4\x43\x08\x92\x8c\x26\xa8\x04\x5f\xc0\xf0\x57\xdf\x87\x31\x1c\x82\xff\xab\x7f\xe7\x87\xad\xc0\xee\x33\x7b\xf5\x23\x53\xab\x4d\x41\xfb\x61\xe6\x3c\xac\x29\x53\xbf\x9c\x9e\xec\x73\xb5\xfe\xaa\x67\xd6\xdf\x4d\x6c\x92\x51\x9c\xad\x6c\xc7\xf4\x81\xe2\x0f\x2b\xcb\x3b\x35\xd4\x84\x2e\xf1\x53\x64\x45\x34\x7c\xbe\xed\xcb\xbc\xaa\xe2\xe7\xf5\xa0\xc8\xe1\x6f\x7a\x9c\xac\xe0\x98\xff\xb9\x8f\x20\xdb\x9e\xe2\xae\x7e\xe0\xf4\x6f\xde\x1c\x67\xdc\x82\xa7\x1f\x09\x1d\x45\xba\x2b\x4a\x3b\x0d\xf8\xe3\xc7\x78\xaf\x0e\xd1\xb5\x4e\x17\xfb\x2e\x1a\x23\x56\x3d\xd5\x1a\x32\x7e\xda\x96\x35\x1c\xd5\xa5\x1f\x3a\x85\xe8\x8b\x6f\x9a\x87\x53\x85\x5f\x56\x7b\xee\x10\xb7\x77\xe4\xe7\x45\x41\xd6\x94\xd0\x35\xd7\x53\x8e\x87\x4f\xbc\x35\x6e\x89\xfa\xe3\xc6\xdd\x63\xed\x1f\xfc\x01\x51\xc7\x35\x35\x65\xe9\x83\x93\x6b\xeb\x0f\xea\xaa\xd2\x6b\xc8\xc5\xbe\xda\xe3\x57\x7b\xb4\x55\xf8\xd5\x1e\x27\x6a\xea\xb5\xec\x51\xa7\xd5\xaa\x72\x23\xaf\x66\xf9\x8f\x99\x8e\x6f\xe4\x5a\x32\x7d\xa9\xcd\x64\xfc\xec\x0c\xe4\xba\x3a\x19\x77\xd6\x73\xbe\x95\x19\xb8\xa4\x7c\x79\x45\xa7\x5f\xcc\xb1\xeb\x38\xba\xf8\x37\x52\xc1\xf1\x3c\x47\xed\x46\x0e\x55\xaa\x58\xe5\x89\x7a\x00\xa1\xa4\x94\xd9\xad\x2b\x83\xea\x17\x51\x44\xcc\xe1\xde\x47\x7c\xbe\x5c\x1a\x55\x92\xa0\x57\x5b\x89\xac\xda\x8a\x20\x95\xf1\x65\xbb\x1d\x9b\x5e\x7c\x1c\x2b\xb6\x48\x0a\x65\x33\x73\x99\xaa\xe7\xe6\x96\xd5\xa7\x8f\x88\xa1\x5d\x11\x42\x60\xa4\x5b\x91\xaa\x32\x18\x7e\xc5\x13\xff\x29\x1e\x49\x99\x6c\xa0\x60\x09\xd7\x41\x1e\x5f\x67\x7b\x96\xe0\x38\x28\x9f\x73\x1c\xea\x80\x38\x41\x05\x86\xbf\xb6\x39\x98\x3e\x07\x95\x84\xcd\x75\xae\x43\x56\x82\xd1\xa2\x93\x02\xf5\x93\x65\x3b\x91\x91\x81\x9d\x56\x07\xc6\xcb\x0b\x54\x94\x46\xa2\xd2\x30\x68\x74\xc6\x09\x6e\xb2\xba\x0e\x0a\x96\xb4\x41\xf1\xdb\x4c\x42\xe7\x2d\x2e\x4b\xcc\xea\x3a\xb4\x56\x6b\x6f\xa9\x2e\xdb\x51\x26\x1d\x1e\x56\x82\xa3\x34\x73\x5c\x31\x7f\xa2\x0d\x5a\xfb\x6b\xb2\x8f\xa1\x2b\xe7\x3d\x7e\x0c\xfc\x42\x60\x06\xb2\x15\xec\xe9\x96\x66\x8f\x14\x38\x7a\x54\xf2\x23\x61\x2b\xf2\x7d\x4b\x7b\x23\x16\x78\x85\xf2\xaf\x46\xf8\xbb\x1b\xa1\xaa\x47\xc1\xed\x9d\xbb\x78\xa5\xe9\x56\x19\x83\x2d\x7e\x16\x55\x16\xd9\xe6\x90\xe7\x38\x0a\x32\x4b\x8e\x02\x16\xc0\xaf\x49\xba\x0c\x18\x2e\x22\x70\xaf\xd7\xce\xf0\xfc\x2d\x7e\xf6\xe7\x00\x62\x5d\x63\x58\x48\xe0\xcf\xa5\x24\xed\x87\x3a\xec\x6d\xae\xbd\xf2\xcc\xcd\x4f\xb4\xde\xff\x3f\xcd\x7c\x39\xbb\xef\x06\x1a\x57\x28\x2f\xac\xca\x89\x3b\xd8\x90\xb1\xc6\x15\xca\xbf\xf6\x8f\xe0\xd4\xfe\x91\x81\x93\x93\xdb\x3d\x6a\xcf\x3f\xe1\x67\x49\xd9\x69\xfc\x78\x5a\xa0\x06\x70\x2f\x5d\x41\xe4\x03\xce\x35\xf4\xef\x58\x0b\xc9\x09\xa6\x89\x6d\x24\x1b\x58\x76\x27\xc9\xfb\x03\xb4\x92\xb8\x1f\x70\xb4\x93\x7e\xc2\xcf\xda\x59\xdd\x39\x3e\x0b\x9d\x0e\x35\x9b\x46\xbb\x4d\xaf\xd9\x6e\x1a\xe8\x37\xf5\x1b\x4e\x93\x3b\x4e\xff\x83\x96\x93\xf2\x88\xce\xa6\xd3\x2b\x75\x9d\xe4\x12\xa2\xef\x64\x85\x61\xdd\xce\x93\xe3\x22\x17\x64\x54\xb5\x9e\x7a\x6d\x16\xd5\x64\x9a\xab\x56\xca\x2b\x60\x89\x73\x55\xad\x16\xe2\xec\xb5\xdc\xc3\x02\xc8\x4c\x5c\x9d\xe2\x2b\xa7\x1a\x69\x8a\x79\xde\x36\x92\x45\x82\x03\x4a\x6f\x85\xaf\xba\x8b\xe4\xbf\xa5\x5f\xb9\x13\x4c\x23\xe0\x7c\xb7\xd1\xa1\x29\x06\x5b\xe2\x0f\x37\x13\xb6\xdb\x5e\x0d\xc7\x35\x11\xfc\x6d\xdb\xcc\x70\xb6\x06\xc6\x6b\x2f\xab\xc1\xe2\xcb\x5f\xbe\x5b\x72\x1b\x06\x9c\x62\xf1\x30\x83\xef\x30\x02\x12\x9a\x6d\x12\x2b\x78\xdf\x6e\x4f\x11\xd6\x8a\xda\x3b\x8a\xb1\x0a\x2d\x7d\xd5\x1c\xfa\xe5\x2d\xf7\xd4\xa6\xa4\xf1\xc5\x95\x23\x8f\x7c\x54\x3d\x87\xc3\x69\x02\x5b\x0a\x92\xf6\x70\xbb\xdd\xde\x2d\x0e\x07\x65\x3d\x8e\x92\x88\x1d\xaa\x34\xd7\x4b\x76\xbc\x32\xd2\xdc\x29\x7f\x80\xda\x88\xce\x33\x86\x2b\x24\xed\x8b\x9e\x7e\x91\x64\x62\xda\x25\xe1\xec\x4e\xb5\x94\x84\x82\x64\xd6\x4b\xae\x86\xe2\x0b\x41\xfa\x5f\x67\x59\xc2\x39\x77\x93\xaa\x11\x2f\xe4\xce\x9d\x5c\x89\x53\x3d\xf3\xac\xf4\xba\x71\x5a\x8a\x40\x6c\x40\x05\xbf\x13\x95\xa8\x61\x3f\xaa\x46\x45\x34\xa8\x48\x47\x18\xf5\xe7\x53\xa5\xf2\xf9\x7d\x65\xd6\xe6\xa3\xb3\xb3\x33\xb8\xc6\xec\x40\x12\x6c\x5a\x61\x21\x87\x5a\x33\xd4\x34\x46\x2c\xf7\x16\x97\x55\xa5\x29\x55\xea\xaf\xc8\xae\x70\xb9\xc9\x96\x45\x90\x18\x69\xa1\x26\xbc\x40\x69\xfa\x4e\x6f\x93\x07\x54\x64\xc3\xf3\x71\xb1\xf3\x04\xe7\x65\xe3\xee\xde\xb5\x23\xff\x40\x74\x99\x62\x06\x83\x31\x62\x34\x10\x24\xea\x18\x31\xec\x5b\x2c\x18\xd1\x9b\x96\x4d\xc9\xdd\xba\xb8\x6e\x7e\xd0\x1a\xbb\xd1\x76\xdd\x89\x59\xa2\xf4\x32\xc0\x48\x3e\x57\x94\x74\xee\x57\x63\x46\x51\xc5\x74\x47\xf6\x14\x9d\x81\xda\x98\x55\x34\xca\x65\x49\x67\x69\xd7\x59\x9a\x7d\x2a\xd2\x73\xb6\xde\xf3\x3b\xa2\x30\xab\x5c\xe7\x6c\xed\x70\x6c\xd2\xac\xf4\x04\x2e\x60\xb7\x6a\x6e\xd6\xa0\x10\x5b\x0b\x35\x8c\xad\x23\x75\x81\xd8\xda\xad\x08\x3d\xc7\x99\xe5\xf0\x59\x46\x5e\xa3\x19\xf6\xea\x89\x9e\xb9\x71\xfb\xd3\x09\x26\x7c\x0f\x96\x11\x1b\x81\xff\x48\x5f\x89\xff\x78\x46\x22\x2c\x5c\x45\xf8\xa3\xc4\x3c\x18\x4b\xca\x27\xa3\x36\x3a\x9c\x4f\x34\x71\xf2\xf1\xa4\x22\x6f\x13\x21\xbf\x67\xa6\x71\x17\x5a\xa0\xaf\xfb\xe6\x79\xcc\x60\xc2\x61\x7c\x33\xb3\x0e\x6f\x38\xb0\x7f\xcd\xe8\xde\xae\xba\xa8\x38\x5f\x29\xae\x83\x01\x1e\xf1\x6f\xfa\x05\x3c\xdb\x14\x3e\xe1\x87\x3d\x2e\x9a\x5c\x75\xb0\xa3\x64\x21\x9e\xe1\x87\x5e\xb8\x37\xc0\xcf\xcf\x39\xa2\x0b\x11\x42\x55\xd5\xf7\xe2\x81\x71\x56\x0e\x91\xab\xc0\x92\x4b\xa1\x2e\x8d\x80\x92\x34\xd4\x3d\xa7\x63\x82\xc1\x09\xfd\xcf\xce\x86\xda\x0b\x05\x33\x36\x69\xd2\x34\x71\x14\xe7\x24\xee\xe8\xe9\x22\x25\x98\x96\xd2\x4d\xf2\x1b\x81\xab\x8a\xe1\x07\xf3\x25\x98\xeb\xb1\x95\x70\xca\x0f\x5c\xef\xae\xa7\x7e\xb6\x52\x75\x92\x35\x75\xf7\x2f\x97\xb1\xe3\x7b\x1a\x00\x2a\x3f\xf0\xc6\x75\xbb\x29\xeb\x6c\x71\xa9\xae\xce\xb9\xc3\x5c\xfd\xd6\xd5\x49\x69\xc6\x2e\x07\xcf\x93\x0e\x6c\x0e\x79\xe4\x28\x53\xb6\xc8\x25\x1b\xed\xf1\x38\x42\x03\x22\xca\x10\xc2\xa4\x85\xe0\xce\x3b\xb9\xf1\x29\x14\x3f\x95\xe0\xa2\x30\x58\xbe\xa3\x87\x6c\x8b\x59\x08\x81\xca\x9d\xbb\x31\x91\x1d\x17\xbd\x9a\x71\xaa\x4b\xfb\xf7\xb3\xcb\x23\x78\x95\xb6\x35\x15\xaf\x56\x5b\xc1\x05\xaf\xd0\x70\x70\x4e\x67\xfb\xc2\xa7\x0d\x88\xeb\x63\xb0\xb2\x4d\x36\x22\x72\x53\x30\x11\x38\x9a\x8a\x15\x6e\xb1\xd6\xd1\x0f\x82\x87\xaf\x20\x11\xfe\x02\xec\x30\xfd\x18\x86\xe1\x87\xd8\x78\xea\xea\xf6\x0a\x5e\xab\xc4\x6f\xb2\xad\x79\x55\x8d\x97\xeb\xab\x6a\xb5\x2b\x55\x69\x2b\x67\x84\x96\xab\xc0\xff\x64\x68\x10\x8c\x5d\x29\xef\x57\xc0\x67\xb4\xe4\x52\xf1\xf5\x45\x71\x3f\xf8\xee\x26\x8c\xe1\x7a\x93\xed\xd3\x25\x7c\x16\x01\xce\x98\xb4\xe2\x26\x7c\x30\x6e\x40\x13\x72\x86\xbd\x7c\x44\xcf\x69\x86\x96\xe2\xc2\xbc\xd8\xe0\x64\x3b\xd1\x66\xa4\xa3\x2e\x1a\xcf\x30\xcd\x13\xba\x9f\x66\xb8\x7e\x2f\xb9\x93\xa0\xf3\x46\xe7\xf5\xae\x25\xf9\xd0\xd1\xda\x9f\x4b\x71\x3e\xc3\x85\x2f\x1a\x4b\x53\xc5\xb5\x9d\x97\xc9\xf3\x3c\x49\x70\x51\x8c\xbc\x8e\x1a\xfa\x71\xeb\xb2\x38\x41\x2f\xe2\x70\xae\xa3\x84\x3f\x45\xf4\xd1\x07\x35\xa7\xc8\xc7\x70\x71\xd2\xba\xb6\x7b\x33\xc1\x6d\x75\xd5\xa6\x00\xc0\xa8\x6f\x4d\x59\xfd\xa4\x8b\xbf\x7b\xed\x0f\xb9\x65\x9d\x7a\xa8\xff\x77\x1e\x15\xcd\xba\x45\xb8\x16\xe3\x76\xc1\xda\x2a\x9b\xfd\x27\x00\x00\xff\xff\x8a\x6f\x45\x18\x2d\x38\x00\x00")

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

	info := bindataFileInfo{name: "templates/types_body.gohtml", size: 14381, mode: os.FileMode(420), modTime: time.Unix(1531229601, 0)}
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

	info := bindataFileInfo{name: "templates/types_head.gohtml", size: 192, mode: os.FileMode(420), modTime: time.Unix(1530882049, 0)}
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

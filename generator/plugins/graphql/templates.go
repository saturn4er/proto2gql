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

var _templatesTypes_bodyGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5f\x73\xdb\x36\x12\x7f\xa6\x3e\xc5\x96\x97\x66\xc8\x1b\x96\xbe\x67\x75\xf4\xe0\xf1\xc5\xb9\x4c\xcf\x49\x2e\xf6\xb4\x0f\xae\xc7\x45\x28\x48\xc2\x89\x02\x69\x90\x92\xed\xe1\xf0\xbb\xdf\xe0\x1f\x09\x90\x20\x45\xe5\x9c\x5e\x3b\x17\x3d\xb4\x31\xb8\x58\x2c\x16\xbf\x5d\xec\x1f\x9c\x9d\xc1\x1b\xba\xdf\x15\xb3\xaa\x62\x88\xae\x31\xbc\xc2\x74\xbf\x83\xf9\x02\xe2\x4b\x92\xe2\x58\x7c\x84\x1f\xea\x7a\xe6\x1d\x10\x83\xaa\x12\xdf\xe3\x9f\x11\x23\xe8\x73\x8a\xdf\xa3\x1d\xae\x6b\x58\x40\x55\xad\x1f\xd2\x8f\xdb\x75\x5d\xc7\xef\xf1\x23\x9f\x15\x18\x43\xfc\xef\x8b\x8c\xae\xc8\xba\x9a\x79\x1e\x9f\x34\x07\xf5\xf3\x35\xcb\xb7\x0c\xe5\x9b\x7f\xfd\x53\x72\xf4\xa3\x99\xe7\x55\x15\x90\x95\x14\x28\xbe\xc8\x76\x3b\x4c\x4b\x29\x89\xe7\xfd\x1d\x17\x09\x23\x79\x49\x32\x3a\x6f\x84\x52\x34\x75\xad\x26\x63\xba\x54\xf4\x3f\xa3\x74\x8f\x8b\x39\x74\x44\x12\xc3\x52\xae\x2b\x94\x57\x33\x30\x7e\x8d\x3e\x0e\x9c\x88\x2b\x44\xef\x9c\xb3\xd2\x72\x78\x5c\x7c\x41\x11\x2b\xc1\xe7\xf0\x7a\x78\x95\x4a\xcc\x91\xe2\x08\xb9\xe5\x54\xf1\xb7\x14\xdb\xd3\xdb\xa6\x7a\xe9\x66\xeb\xbf\xf9\xfe\x6f\xcd\xba\x7d\x15\x58\xc4\x26\xb3\x56\x0d\x9e\x27\x87\xab\xaa\x1d\xe3\x23\x75\x38\x6b\x87\x66\xb3\xb3\x33\x78\x47\xf3\x7d\x09\xd9\xe7\x7f\xe3\xa4\x9c\x55\x15\x28\x5d\xc8\x81\x16\x1d\x82\xec\x83\x18\xb4\x41\x22\x09\x8f\xc1\xc4\x98\x6e\xa2\xc5\x18\xd6\x6a\xd3\xa7\x22\xa1\xe3\xb7\x2b\x74\x51\xa3\x09\x2f\x09\x4e\x97\xf6\x89\xf7\xd8\x0a\x9a\x2b\x94\xdf\x6c\xf6\x74\x1b\xac\xf6\x34\x09\xc2\x49\x13\xc0\x86\x0a\xc3\xe5\x9e\xd1\x49\x33\xed\x89\x16\xce\x56\x9c\x44\xe0\x4c\x6d\x4d\x6e\x41\xa8\xb5\x3b\x49\x1b\x8e\x98\xe3\x46\x9e\x21\x82\x60\xa4\x34\x79\xf3\x9c\x0b\xe4\x25\x28\x4d\xd5\x9a\x31\x1f\x83\x57\x35\xc7\x4c\x5f\x3a\x0d\x0b\x73\xb4\xfd\xab\x0e\xe5\x1c\x81\xa0\x06\x69\x5d\x04\x15\xc0\x70\x91\xa5\x07\xcc\x0a\x03\x4c\x7a\xcc\x09\xa7\x4f\x7a\x82\x60\xc8\xcf\x86\xc3\x4a\x4f\x89\x2f\xf7\x34\xe1\xc8\x97\x5b\x0f\x94\xa7\x88\x6f\x18\x4a\x30\x7b\x43\x39\xe6\x96\x50\xd7\x50\x72\x34\x96\x62\x54\xaa\x45\x52\x44\x72\x63\x75\x9d\x94\x4f\x5c\x19\xe5\x93\xfc\x7a\x91\xd1\x12\x3f\x95\x11\x10\x20\xb4\xc4\x6c\x85\x12\x5c\xd5\x21\x04\xf7\xfc\x74\x33\xa9\xa8\x46\x88\x0f\xfb\x32\xdf\x97\x6f\xc5\x70\x5d\x47\xc0\x30\x63\x80\x19\xcb\x58\x58\xcd\x3c\xb7\x4c\xd2\xe8\x8a\x1c\x51\xbe\xed\x92\xc5\x17\x0c\xa3\x12\x5f\x6c\x48\xba\xbc\xce\x11\xbd\x64\xd9\x4e\x49\x11\x24\xe5\x53\x24\x8e\x79\x60\xdb\x7e\x38\xf3\xbc\x25\x5e\x61\x06\x9c\x61\x7c\x49\x28\x29\x36\x41\x3b\x2a\x21\x2d\x9c\x0e\x59\x41\xce\xc5\x9b\x2f\x80\xe1\x24\x3b\x60\x16\x84\x3f\xca\xa1\xef\x16\x40\x49\x0a\xd2\x37\x09\x46\xd7\xb8\xbc\x41\xeb\xc0\x17\x7b\xf1\x23\xf0\x4b\xb6\xc7\x7e\x68\x8f\xdf\xef\x70\x51\xa0\x35\xf6\x23\xc1\xa6\xfb\xb5\x28\x51\xb2\xf5\x23\x28\x4a\x46\xe8\x3a\xa8\xaa\x25\xfe\xbc\x5f\x4b\x35\x5f\xf3\x6f\x41\x18\x72\x49\xbd\x5a\x89\xc7\x5e\x46\x16\xce\x26\x7e\xc3\x07\x83\x96\x7f\xcd\x95\x62\x78\x3d\xb2\x02\x02\x8b\x76\x2d\x65\xc1\x94\xa4\x11\xff\xcf\x8c\x4f\x42\x6c\x5d\x70\x75\x91\x38\xd8\xa1\xfc\x56\xee\xe3\xce\x84\xc5\xcc\xbb\x87\x05\x70\x3a\xe9\xf7\x18\x2e\xf6\x69\x09\x0b\xa0\xf8\x31\xd0\x78\xb9\xcc\xd8\x7b\xfc\x38\x88\x1a\x21\x17\xf4\x5c\x40\x7b\xe2\xad\x13\x10\x5a\xe2\xab\xdd\xb6\xa6\xaf\xdc\x9f\x30\x1c\x41\xaa\x90\x71\x67\x69\x52\x42\x91\x14\xe7\x8c\xa1\x67\x6d\xf5\x52\x82\xe6\x66\x20\x02\x90\x53\xd9\xc7\xc1\x6d\x47\x15\x1e\xd7\x22\xdf\x7f\xdc\xcc\x96\x3b\x35\xe6\xc1\x02\x76\x68\x8b\x83\xd6\x96\x4c\x51\xb8\x09\xa5\x98\x06\x84\xca\x83\xf3\x56\x19\x03\x12\xc1\x01\xa5\x02\xb6\x42\x47\x84\x2a\x70\x68\xfb\x92\x1c\xb4\xbb\xf8\x85\x94\x1b\x71\xf8\xd0\xdc\x93\x87\x08\x14\xf0\x6d\xaf\x27\x2e\x5d\x3d\x0f\xfc\x03\x4a\x7d\xee\x04\xd5\x2c\xb2\x82\x1e\x1e\x3d\x1b\x28\x55\x25\xb0\x57\x48\x4c\xff\xc2\x50\x1e\x60\xc6\x22\xf0\x57\x88\x70\x53\x2f\x33\xed\xf6\x80\x18\xce\x10\xc4\xf2\x7e\xa8\x58\xea\x05\x8f\x2a\xef\x96\xdc\xc1\x02\x0e\xed\xb5\x9e\x16\xd8\x08\x07\x26\xce\x3f\x41\x07\xdd\xd0\x61\xe6\x58\xf7\xc8\x21\x34\xe2\x4d\x3b\x84\x20\xc9\x68\x82\x4a\xf0\x05\x0c\x7f\xf5\x7d\x18\xc3\x21\xf8\xbf\xfa\x77\x7e\xd8\x0a\xec\x3e\xb3\x17\x3f\x32\xb5\xda\x14\xb4\x1f\x66\xce\xc3\x9a\x32\xf5\xeb\xe9\xc9\x3e\x57\xeb\xaf\x7a\x66\xfd\xdd\xc4\x26\x19\xc5\xd9\xca\x76\x4c\x1f\x28\xfe\xb0\xb2\xbc\x53\x43\x4d\xe8\x12\x3f\x45\x56\x44\xc3\xe7\xdb\xbe\xcc\xab\x2a\x7e\x5e\x0f\x8a\x1c\xfe\xa6\xc7\xc9\x0a\x8e\xf9\x9f\xfb\x08\xb2\xed\x29\xee\xea\x47\x4e\xff\xfa\xf5\x71\xc6\x2d\x78\xfa\x91\xd0\x51\xa4\xbb\xa2\xb4\xd3\x80\x3f\x7e\x8c\xf7\xea\x10\x5d\xeb\x74\xb1\xef\xa2\x31\x62\xd5\x53\xad\x21\xe3\xa7\x6d\x59\xc3\x51\x5d\xfa\xa1\x53\x88\xbe\xf8\xa6\x79\x38\x55\xf8\x75\xb5\xe7\x0e\x71\x7b\x47\x7e\x5e\x14\x64\x4d\x09\x5d\x73\x3d\xe5\x78\xf8\xc4\x5b\xe3\x96\xa8\x3f\x6e\xdc\x3d\xd6\xfe\xc1\x1f\x10\x75\x5c\x53\x53\x96\x3e\x38\xb9\xb6\xfe\xa0\xae\x2a\xbd\x86\x5c\xec\x9b\x3d\x7e\xb3\x47\x5b\x85\xdf\xec\x71\xa2\xa6\x5e\xca\x1e\x75\x5a\xad\x2a\x37\xf2\x6a\x96\xff\x98\xe9\xf8\x46\xae\x25\xd3\x97\xda\x4c\xc6\xcf\xce\x40\xae\xab\x93\x71\x67\x3d\xe7\x95\xcc\xc0\x25\xe5\x97\x57\x74\xfa\xc5\x1c\xbb\x8e\xa3\x8b\x7f\x23\x15\x1c\xcf\x73\xd4\x6e\xe4\x50\xa5\x8a\x55\x9e\xa8\x07\x10\x4a\x4a\x99\xdd\xba\x32\xa8\x7e\x11\x45\xc4\x1c\xee\x7d\xc4\xe7\xcb\xa5\x51\x25\x09\x7a\xb5\x95\xc8\xaa\xad\x08\x52\x19\x5f\xb6\xdb\xb1\xe9\xc5\xc7\xb1\x62\x8b\xa4\x50\x36\x33\x97\xa9\x7a\x6e\x6e\x59\x7d\xfa\x88\x18\xda\x15\x21\x04\x46\xba\x15\xa9\x2a\x83\xe1\x57\x3c\xf1\x9f\xe2\x91\x94\xc9\x06\x0a\x96\x70\x1d\xe4\xf1\x75\xb6\x67\x09\x8e\x83\xf2\x39\xc7\xa1\x0e\x88\x13\x54\x60\xf8\x6b\x9b\x83\xe9\x73\x50\x49\xd8\x5c\xe7\x3a\x64\x25\x18\x2d\x3a\x29\x50\x3f\x59\xb6\x13\x19\x19\xd8\x69\x75\x60\xbc\xbc\x40\x45\x69\x24\x2a\x0d\x83\x46\x67\x9c\xe0\x26\xab\xeb\xa0\x60\x49\x1b\x14\xbf\xcd\x24\x74\xde\xe2\xb2\xc4\xac\xae\x43\x6b\xb5\xf6\x96\xea\xb2\x1d\x65\xd2\xe1\x61\x25\x38\x4a\x33\xc7\x15\xf3\x27\xda\xa0\xb5\xbf\x26\xfb\x18\xba\x72\xde\xe3\xc7\xc0\x2f\x04\x66\x20\x5b\xc1\x9e\x6e\x69\xf6\x48\x81\xa3\x47\x25\x3f\x12\xb6\x22\xdf\xb7\xb4\x37\x62\x81\x57\x28\xff\x66\x84\xbf\xbb\x11\xaa\x7a\x14\xdc\xde\xb9\x8b\x57\x9a\x6e\x95\x31\xd8\xe2\x67\x51\x65\x91\x6d\x0e\x79\x8e\xa3\x20\xb3\xe4\x28\x60\x01\xfc\x9a\xa4\xcb\x80\xe1\x22\x02\xf7\x7a\xed\x0c\xcf\xdf\xe2\x67\x7f\x0e\x20\xd6\x35\x86\x85\x04\xfe\x5c\x4a\xd2\x7e\xa8\xc3\xde\xe6\xda\x2b\xcf\xdc\xfc\x44\xeb\xfd\xff\xd3\xcc\xd7\xb3\xfb\x6e\xa0\x71\x85\xf2\xc2\xaa\x9c\xb8\x83\x0d\x19\x6b\x5c\xa1\xfc\x5b\xff\x08\x4e\xed\x1f\x19\x38\x39\xb9\xdd\xa3\xf6\xfc\x13\x7e\x96\x94\x9d\xc6\x8f\xa7\x05\x6a\x00\xf7\xa5\x2b\x88\x7c\xc0\xb9\x86\xfe\x1d\x6b\x21\x39\xc1\x34\xb1\x8d\x64\x03\xcb\xee\x24\x79\x7f\x80\x56\x12\xf7\x03\x8e\x76\xd2\x4f\xf8\x59\x3b\xab\x3b\xc7\x67\xa1\xd3\xa1\x66\xd3\x68\xb7\xe9\x25\xdb\x4d\x03\xfd\xa6\x7e\xc3\x69\x72\xc7\xe9\x7f\xd0\x72\x52\x1e\xd1\xd9\x74\x7a\xa1\xae\x93\x5c\x42\xf4\x9d\xac\x30\xac\xdb\x79\x72\x5c\xe4\x82\x8c\xaa\xd6\x53\xaf\xcd\xa2\x9a\x4c\x73\xd5\x4a\x79\x01\x2c\x71\xae\xaa\xd5\x42\x9c\xbd\x96\x7b\x58\x00\x99\x89\xab\x53\x7c\xe5\x54\x23\x4d\x31\xcf\xdb\x46\xb2\x48\x70\x40\xe9\xad\xf0\x55\x77\x91\xfc\xb7\xf4\x2b\x77\x82\x69\x04\x9c\xef\x36\x3a\x34\xc5\x60\x4b\xfc\xe1\x66\xc2\x76\xdb\xab\xe1\xb8\x26\x82\xbf\x6d\x9b\x19\xce\xd6\xc0\x78\xed\x65\x35\x58\x7c\xf9\xcb\xf7\x4b\x6e\xc3\x80\x53\x2c\x1e\x66\xf0\x1d\x46\x40\x42\xb3\x4d\x62\x05\xef\xdb\xed\x29\xc2\x5a\x51\x7b\x47\x31\x56\xa1\xa5\xaf\x9a\x43\xbf\xbc\xe5\x9e\xda\x94\x34\xbe\xba\x72\xe4\x91\x8f\xaa\xe7\x70\x38\x4d\x60\x4b\x41\xd2\x1e\x6e\xb7\xdb\xbb\xc5\xe1\xa0\xac\xc7\x51\x12\xb1\x43\x95\xe6\x7a\xc9\x8e\x57\x46\x9a\x3b\xe5\x0f\x50\x1b\xd1\x79\xc6\x70\x85\xa4\x7d\xd1\xd3\x2f\x92\x4c\x4c\xbb\x24\x9c\xdd\xa9\x96\x92\x50\x90\xcc\x7a\xc9\xd5\x50\x7c\x21\x48\xff\xeb\x2c\x4b\x38\xe7\x6e\x52\x35\xe2\x85\xdc\xb9\x93\x2b\x71\xaa\x67\x9e\x95\x5e\x37\x4e\x4b\x11\x88\x0d\xa8\xe0\x77\xa2\x12\x35\xec\x47\xd5\xa8\x88\x06\x15\xe9\x08\xa3\xfe\x7c\xaa\x54\x3e\xbf\xaf\xcc\xda\x7c\x74\x76\x76\x06\xd7\x98\x1d\x48\x82\x4d\x2b\x2c\xe4\x50\x6b\x86\x9a\xc6\x88\xe5\xde\xe2\xb2\xaa\x34\xa5\x4a\xfd\x15\xd9\x15\x2e\x37\xd9\xb2\x08\x12\x23\x2d\xd4\x84\x17\x28\x4d\xdf\xe9\x6d\xf2\x80\x8a\x6c\x78\x3e\x2e\x76\x9e\xe0\xbc\x6c\xdc\xdd\xbb\x76\xe4\x1f\x88\x2e\x53\xcc\x60\x30\x46\x8c\x06\x82\x44\x1d\x23\x86\x7d\x8b\x05\x23\x7a\xd3\xb2\x29\xb9\x5b\x17\xd7\xcd\x0f\x5a\x63\x37\xda\xae\x3b\x31\x4b\x94\x5e\x06\x18\xc9\xe7\x8a\x92\xce\xfd\x6a\xcc\x28\xaa\x98\xee\xc8\x9e\xa2\x33\x50\x1b\xb3\x8a\x46\xb9\x2c\xe9\x2c\xed\x3a\x4b\xb3\x4f\x45\x7a\xce\xd6\x7b\x7e\x47\x14\x66\x95\xeb\x9c\xad\x1d\x8e\x4d\x9a\x95\x9e\xc0\x05\xec\x56\xcd\xcd\x1a\x14\x62\x6b\xa1\x86\xb1\x75\xa4\x2e\x10\x5b\xbb\x15\xa1\xe7\x38\xb3\x1c\x3e\xcb\xc8\x6b\x34\xc3\x5e\x3d\xd1\x33\x37\x6e\x7f\x3a\xc1\x84\xef\xc1\x32\x62\x23\xf0\x1f\xe9\x2b\xf1\x1f\xcf\x48\x84\x85\xab\x08\x7f\x94\x98\x07\x63\x49\xf9\x64\xd4\x46\x87\xf3\x89\x26\x4e\x3e\x9e\x54\xe4\x6d\x22\xe4\xf7\xcc\x34\xee\x42\x0b\xf4\x75\xdf\x3c\x8f\x19\x4c\x38\x8c\x6f\x66\xd6\xe1\x0d\x07\xf6\x2f\x19\xdd\xdb\x55\x17\x15\xe7\x2b\xc5\x75\x30\xc0\x23\xfe\x4d\xbf\x80\x67\x9b\xc2\x27\xfc\xb0\xc7\x45\x93\xab\x0e\x76\x94\x2c\xc4\x33\xfc\xd0\x0b\xf7\x06\xf8\xf9\x39\x47\x74\x21\x42\xa8\xaa\xfa\x41\x3c\x30\xce\xca\x21\x72\x15\x58\x72\x29\xd4\xa5\x11\x50\x92\x86\xba\xe7\x74\x4c\x30\x38\xa1\xff\xd9\xd9\x50\x7b\xa1\x60\xc6\x26\x4d\x9a\x26\x8e\xe2\x9c\x18\x88\xe3\x17\x80\xf4\x8f\x75\x6d\xc2\x94\xe1\x87\xd0\x32\xe9\xb1\x0e\x9f\x4b\xfc\x29\x8b\xb8\x5e\x08\xda\x67\xa1\x72\xb3\x70\xcc\xbb\x34\x10\x53\x96\xfe\xda\x75\x7f\xa9\x35\x5b\xe4\xa9\xcb\x71\xee\x30\x48\xbf\x75\x66\x52\xea\x31\xf7\xef\x79\xd2\x45\xcd\x21\x8f\x1c\x85\xc8\x16\x9b\x64\xa3\x7d\x1a\xc7\x60\x40\x44\xa1\x41\x18\xad\x10\xdc\x79\xeb\x1a\x9a\x7a\x2a\xc1\x45\x61\xb0\x7c\x47\x0f\xd9\x16\xb3\x10\x02\x95\x1d\x77\xa3\x1e\x3b\xf2\x79\x31\xf3\x53\xd7\xf2\xef\x67\x79\x2e\x58\x36\xee\x48\x5b\xcf\x54\x68\x59\x8d\x03\x17\xbc\x42\xc3\x85\x39\xdd\xe9\x17\x3e\x5e\x40\x5c\x1f\x83\xb5\x6b\xb2\x11\x56\xa3\x60\x22\x70\x34\x15\x2b\x0c\x3f\xd8\x47\x3f\x08\x9e\xd6\x2e\xbf\x00\x3b\x4c\x3f\x77\x61\xf8\x21\x36\x1e\xb3\x3a\xb5\x3c\x33\x2e\xa5\xef\xb2\xad\x79\x19\x8d\x17\xe4\xab\x6a\xb5\x2b\x55\xf1\x2a\x67\x84\x96\xab\xc0\xff\x64\x68\x10\x8c\x5d\x29\xff\x56\xc0\x67\xb4\xe4\x52\xf1\xf5\x45\xf9\x3e\xf8\xfe\x26\x8c\xe1\x7a\x93\xed\xd3\x25\x7c\x16\x21\xcc\x98\xb4\xbe\x74\x7f\xad\xd0\x26\xe4\x0c\x7b\xf9\x88\x9e\xd3\x0c\x2d\xc5\x95\x78\xb1\xc1\xc9\x76\xa2\xcd\x48\x57\x5c\x34\x9e\x61\xd0\x55\xf2\xc3\x8e\xa5\x7b\x31\x8e\xd7\xfd\x82\xa5\xfb\xfb\x92\xab\x07\x3a\x4f\x71\x5e\xee\xf6\x91\xef\x19\x2d\x17\xe1\xd2\x9e\xcf\x70\xe1\x8b\xfe\xd1\x54\x71\x6d\x0f\x66\xf2\x3c\x4f\x12\x5c\x14\x23\x8f\xa0\x86\x7e\x4a\xeb\x2d\x27\xe8\x05\x16\xce\x75\x94\xf0\xa7\x88\x3e\xf9\x56\x3d\x26\x1f\xc3\xc5\x49\xeb\xda\x3e\xce\x44\xb8\xd5\x3c\x9b\x02\x00\xa3\x8c\x35\x65\xf5\x63\xe1\xc1\x10\xe6\x5b\x0b\x1c\x72\xd0\x3a\xcd\x50\xff\xef\x3c\x20\x9a\x75\x0b\x6e\x2d\xd0\xed\xe2\xb4\x55\x22\xfb\x4f\x00\x00\x00\xff\xff\xa5\x05\x14\x19\x19\x38\x00\x00")

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

	info := bindataFileInfo{name: "templates/types_body.gohtml", size: 14361, mode: os.FileMode(420), modTime: time.Unix(1530789088, 0)}
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

	info := bindataFileInfo{name: "templates/types_head.gohtml", size: 192, mode: os.FileMode(420), modTime: time.Unix(1530624944, 0)}
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


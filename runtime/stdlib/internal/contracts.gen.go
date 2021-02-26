// Code generated by go-bindata. DO NOT EDIT.
// sources:
// contracts/crypto.cdc (5.061kB)

package internal

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _contractsCryptoCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\x5d\x4f\xe4\x36\x17\xbe\x9f\x5f\x71\xde\xbd\x02\xbd\xa3\x81\xa2\xaa\xaa\x46\xca\xa2\x6d\xbb\x6d\x11\xad\x54\x41\xd9\x5e\x20\xb4\x6b\x26\x27\x13\x8b\xac\x33\xb2\x9d\x81\x74\x34\xff\xbd\xb2\x9d\xf8\x3b\x30\x54\x74\x2e\x20\x89\x9f\xf3\xf5\xf8\x9c\x13\x9f\xcc\x36\xdd\x3d\x08\xc9\xbb\x95\x04\xca\x24\xf2\x8a\xac\x10\xae\xe9\x9a\x11\xd9\x71\xfc\x84\x9c\x56\x14\x39\xc0\x6e\x36\x03\x00\x50\xf0\xaa\x63\xb0\x55\x0b\xfd\x91\x7e\xa6\x7e\x62\x94\x58\xc2\xed\xcd\x05\x93\xdf\xdf\xcd\xed\x9a\x24\xeb\x25\x5c\x4b\x4e\xd9\x7a\x1e\x08\x60\xf9\x13\x91\x24\x23\xb1\xe9\xee\x1b\xba\xba\xc4\x3e\xb3\x66\x2d\x7d\x68\xd6\x2d\xa7\xb2\xfe\xba\x74\xfe\xda\x67\x67\x4e\xa0\x26\xa2\xf6\xb0\xbf\xfa\xb7\x67\x1a\x75\xbc\x84\x1f\xda\xb6\x99\xed\x67\x79\x3a\x94\x48\x86\x03\xa5\xd8\x31\x50\xe6\x43\x21\x2f\x59\x1e\x24\x46\xe3\xab\x96\x49\x4e\x56\x12\x7e\xe4\xfd\x46\xb6\x59\x9b\x9f\x23\x63\xd3\x46\x9c\x7e\xd8\x59\x97\x38\xca\x8e\x33\x10\xd8\x54\x8b\x5a\x87\xa6\xff\x1d\x19\xa5\xea\x6f\xa0\xd1\x5e\x1e\x6b\x0d\x7b\xe7\xd0\x40\xd4\x25\xf6\xbf\x51\x21\x3f\x32\xc9\x7b\xcf\x8c\x42\x34\x28\xe1\x01\xfb\x0b\x56\xe2\xd3\x12\x2e\x98\x4c\x56\xbd\xad\xfe\x63\xbc\x3c\x4b\x50\x2f\xef\xa1\x8f\x7e\x44\xba\xae\xe5\x12\x6e\x7e\xa6\x4f\xdf\x7d\x9b\x2c\x53\x71\x85\xdb\xf6\x01\xcb\x61\xdb\x2d\x80\x32\x2a\xdd\x86\xaa\x5f\xe0\xfb\x3c\x58\xca\x3a\x1e\x42\x9e\xf5\x3a\x84\x86\x2e\x87\x6b\xb1\xbf\xe3\xf3\x63\x8f\x6c\x5d\x19\x6a\x43\x47\x8f\xa1\xb0\xce\xa7\x20\xeb\x3b\x14\x2e\x8e\x14\x16\xf8\x0f\x45\x18\x4f\x0a\x37\x31\x40\x31\x04\x93\x02\x6c\x20\x50\xb8\xa0\x2c\x6c\xff\x7c\x7a\x8d\x95\xa0\x57\x39\xdd\xea\xad\x44\x26\x39\x45\xb1\x84\x5b\x3f\x09\xef\xa2\x2d\xcd\x12\x35\x88\x42\x01\xb7\x77\x9e\x0f\xf6\xf2\xe4\xe4\x04\x3e\x94\xa5\x00\x02\x0c\x1f\x15\x99\xf0\x48\x65\x0d\xb2\x46\x58\xd3\x2d\xb2\x38\xcc\xb1\x44\x49\x59\x86\x49\xf4\xf9\x3f\xcd\x15\x97\x0e\xcb\xb8\x12\x03\x29\xbf\x14\xa1\x08\x38\x58\x34\xc8\xd6\xb2\x4e\xe0\xa8\xf5\x14\x81\xda\x30\x34\x08\x6a\x64\xbc\x9a\x27\x18\x8f\x00\x7b\x99\xa2\x22\x0e\x82\xdb\x14\x3d\xd2\x60\xfe\xa7\xeb\x5e\xd9\x54\xa4\x11\x18\x00\x8e\x27\x13\x62\x41\x36\x1b\x64\xe5\x91\x0e\x3e\x84\x0d\x8d\x53\xaf\x4c\xa5\xcc\x95\xc6\x08\x9d\x26\x2a\x67\x88\xf4\x32\x86\x6a\x76\x80\x56\x40\x25\xe0\x13\x15\x52\x2c\x22\x69\x53\x1f\x0f\xd8\x0b\x20\x1c\x81\x34\x8f\xa4\x17\x83\x65\x2c\xe7\x70\xdf\x69\x85\x3d\xd4\x64\x8b\xf0\xc5\x06\xf9\x05\x2a\x8a\x4d\x09\x02\x25\xc8\x16\x24\xef\x30\xc9\xcc\x35\xca\xa3\xa0\xa3\x45\x29\x73\x1e\xd5\x09\xad\x5c\xc6\xbc\xcf\xa6\x4c\x24\xe0\x91\xc4\x68\x13\x2c\xed\x67\x39\x2a\x7d\x95\xb7\xa3\xad\xc9\x72\xfc\x9d\xf0\x87\xe7\x98\x05\x6e\xb8\x30\x2c\x95\x2d\x0a\x60\xad\x84\x12\x1b\x94\x08\x34\x2d\x55\x83\x8f\x38\x79\x3b\x12\x22\x02\xfc\x3b\x55\x5e\xab\x8e\x73\x64\x43\xb5\x16\x2f\x71\x01\x51\x9e\x3a\xc8\x2b\x2a\xd4\x37\xb9\x38\xa8\x5c\x03\x89\xc3\x6b\x37\x10\x3b\xb0\x90\x03\x99\x03\xaa\x3a\xc8\x71\x08\x8a\x7a\xaa\x28\x79\x87\x6a\x43\x5d\xde\xd8\xd3\xa4\x29\xb7\x2d\x69\x68\x09\x55\xcb\x23\x08\x96\xfa\x5c\x94\x24\x10\x15\x9f\x94\x44\x48\xb9\xd5\x79\x8d\xd2\xbd\x99\xec\x19\xf5\x6e\x9e\xa0\xa3\x83\xb0\xdf\xd2\xd5\x1b\x3f\x6e\xe5\x5b\xc2\x8d\xa7\x7f\x69\x92\xc4\xf8\x2a\x80\x02\x4e\x17\xa7\x69\xdb\x17\x88\xec\x52\x6f\x36\x5d\xa9\x97\xe5\xee\x82\x49\xa3\x79\x0f\x05\xec\xa2\xd2\x54\xd1\xdb\x10\x80\xb2\x20\x9e\xd8\x15\xc3\x30\x7c\x64\x42\x81\xc7\xd2\x34\xd5\x48\x85\xf1\x32\x95\xa0\x95\x53\xba\x78\x7d\x75\x81\x6b\x20\x69\x67\x87\xb4\xd9\xc4\x4e\x52\xe1\x79\x59\x13\xd3\x25\x48\xc3\x91\x94\x3d\xdc\xa3\xda\x73\x44\x96\x77\x3b\xa0\xf2\x36\x8d\xe2\x0e\xce\xcf\x8d\x57\x6f\xe7\xf8\x15\xae\x5a\x5e\x46\xec\x3e\x12\x31\xe1\xe6\x01\x3e\x16\xa6\x78\x72\xc6\x7e\x41\xd3\x57\xc9\x4a\x76\xa4\x51\x06\x53\xd8\x70\x96\x88\x9b\x56\xc6\xd2\x41\xd9\x62\x36\x60\x68\xde\x59\xda\x1f\xb0\xf7\x4e\x8f\x6f\x9f\x11\xe8\x67\xfc\x73\x69\xfb\x3f\x33\x96\x2d\x44\x3c\x21\x2f\xe2\x89\x38\xd8\x11\x37\x1d\x3b\x8e\xec\x55\xda\xe4\x60\x9c\x9a\x07\x6b\x65\xfb\x95\x50\x76\x8d\x1b\xc2\x89\xa4\x2d\xfb\x93\xac\x6f\x04\xf2\xbc\xa0\xdf\x50\xdc\x75\x1e\xeb\x35\x7a\x45\xb1\xbd\x7d\xae\xd3\xc3\xc4\x10\x1e\x2a\x50\x08\xb5\x98\x97\x0f\xdf\x16\x4a\x72\x7a\xb6\x80\xcc\xa8\x33\xfe\x5e\xbb\xe7\x7e\xd7\x84\x22\xbc\xfd\xbf\x8e\x20\x33\xbb\xe4\x8f\x2e\x81\xec\xfb\x02\xbe\x59\x9c\x1e\x3a\xca\xd8\x57\xc1\xab\xa7\xe5\xf4\x33\xcb\x2c\xc0\xe8\x71\x27\x1c\x5a\x33\x32\xff\x7e\x70\x74\x55\x52\x38\xbd\x13\x61\x8f\x33\xda\x44\xf2\x8e\x9f\x84\x22\x74\x52\x58\x4b\xd8\x25\x9f\xa3\xf6\xa1\x8c\xf9\x88\xb1\x84\x9d\xf9\x50\x33\x38\xa0\xb9\x38\x4c\xdd\x3c\x51\x71\xec\xbf\xe8\xc2\xd0\xed\x37\xb1\x22\x75\x36\x14\x31\x4a\x87\xc1\x19\xb9\x7f\x26\x81\x0b\x46\x25\x25\x0d\xfd\x1b\x61\xd5\x32\x21\x09\x93\x22\xb2\x38\x41\x1c\x14\xf0\xae\x13\xc8\xdf\x0d\x6c\xef\x67\xff\x04\x00\x00\xff\xff\xed\x71\x54\xfc\xc5\x13\x00\x00"

func contractsCryptoCdcBytes() ([]byte, error) {
	return bindataRead(
		_contractsCryptoCdc,
		"contracts/crypto.cdc",
	)
}

func contractsCryptoCdc() (*asset, error) {
	bytes, err := contractsCryptoCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "contracts/crypto.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x24, 0xf4, 0xc0, 0xa6, 0x69, 0xee, 0x2a, 0x43, 0x5e, 0x36, 0xdb, 0xd7, 0xf2, 0x5e, 0x98, 0x7a, 0x89, 0x93, 0xf5, 0x1, 0xdf, 0xc4, 0x31, 0x82, 0xf6, 0xaa, 0x78, 0xe2, 0x1, 0xed, 0xb7, 0xfe}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"contracts/crypto.cdc": contractsCryptoCdc,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"contracts": {nil, map[string]*bintree{
		"crypto.cdc": {contractsCryptoCdc, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}

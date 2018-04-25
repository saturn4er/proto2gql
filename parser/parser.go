package parser

import (
	"github.com/emicklei/proto"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

type Parser struct {
	ImportAliases map[string]string
	Paths         []string
	parsedFiles   []*File
}

func (f *Parser) parsedFile(filePath string) (*File, bool) {
	for _, f := range f.parsedFiles {
		if f.FilePath == filePath {
			return f, true
		}
	}
	return nil, false
}
func (f *Parser) importFilePath(filename string) (filePath string, err error) {
	if v, ok := f.ImportAliases[filename]; ok {
		filename = v
	}
	for _, path := range f.Paths {
		p := filepath.Join(path, filename)
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}
	return "", errors.Errorf("can't find import %s in any of %s", filename, f.Paths)
}

func (p *Parser) parseFileImports(file *File) error {
	for _, v := range file.protoFile.Elements {
		imprt, ok := v.(*proto.Import)
		if !ok {
			continue
		}
		imprtPath, err := p.importFilePath(imprt.Filename)
		if err != nil {
			return errors.Wrapf(err, "failed to resolve import(%s) file path", imprt.Filename)
		}
		absImprtPath, err := filepath.Abs(imprtPath)
		if err != nil {
			return errors.Wrap(err, "failed to resolve import absolute file path")
		}
		if fl, ok := p.parsedFile(absImprtPath); ok {
			file.Imports = append(file.Imports, fl)
			continue
		}
		importFile, err := p.Parse(absImprtPath)
		if err != nil {
			return errors.Wrap(err, "can't parse import")
		}
		file.Imports = append(file.Imports, importFile)
	}
	return nil
}

func (p *Parser) Parse(path string) (*File, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve file absolute path")
	}
	file, err := os.Open(absPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open file")
	}
	f, err := proto.NewParser(file).Parse()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse file")
	}
	result := &File{
		FilePath:  absPath,
		protoFile: f,
		PkgName:   resolveFilePkgName(f),
	}
	err = p.parseFileImports(result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse file imports")
	}
	result.parseMessages()
	result.parseEnums()
	err = result.parseServices()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse file services")
	}
	err = result.ParseMessagesFields()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse messages fields")
	}
	p.parsedFiles = append(p.parsedFiles, result)
	return result, nil
}

func New(importsAliases map[string]string, paths []string) *Parser {
	return &Parser{
		ImportAliases: importsAliases,
		Paths:         paths,
	}
}

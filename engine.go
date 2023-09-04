package gobatis

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobatis/gobatis/driver/clickhouse"
	"github.com/gobatis/gobatis/driver/mysql"
	"github.com/gobatis/gobatis/driver/postgresql"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func NewClickhouse(dsn string) *Engine {
	return NewEngine(NewDB(clickhouse.Clickhouse, dsn))
}

func NewPostgresql(dsn string) *Engine {
	return NewEngine(NewDB(postgresql.PGX, dsn))
}

func NewMySQL(dsn string) *Engine {
	return NewEngine(NewDB(mysql.MySQL, dsn))
}

func NewEngine(db *DB) *Engine {
	engine := &Engine{master: db, fragmentManager: newMethodManager()}
	return engine
}

type Engine struct {
	master          *DB
	slaves          []*DB
	logger          Logger
	fragmentManager *fragmentManager
}

func (p *Engine) Master() *DB {
	return p.master
}

func (p *Engine) SetTag(tag string) {
	reflect_tag = tag
}

func (p *Engine) UseJsonTag() {
	reflect_tag = "json"
}

func (p *Engine) SetLogLevel(level Level) {
	p.logger.SetLevel(level)
}

func (p *Engine) SetLogger(logger Logger) {
	p.logger = logger
	if p.master != nil {
		p.master.logger = logger
	}
	for _, v := range p.slaves {
		v.logger = logger
	}
}

func GetMappers(dir string) (files []string) {

	// 使用Walk函数遍历文件夹
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {

		} else {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
	}

	return files
}

func (p *Engine) Init(mapperDir string) (err error) {

	if p.logger == nil {
		p.logger = NewStdLogger()
		p.logger.SetLevel(InfoLevel)
	}

	err = p.parseMappers(mapperDir)

	if err != nil {
		return
	}

	err = p.master.initDB()
	p.master.logger = p.logger
	if err != nil {
		err = fmt.Errorf("init master db error: %s", err)
		return
	}
	return
}

func (p *Engine) Close() {
	if p.fragmentManager != nil {
		for _, v := range p.fragmentManager.all() {
			if v._stmt != nil {
				err := v._stmt.Close()
				if err != nil {
					p.logger.Errorf("[gobatis] close stmt error: %s", err)
				}
			}
		}
	}
	for _, v := range p.slaves {
		err := v.Close()
		if err != nil {
			p.logger.Errorf("[gobatis] close slave db error: %s", err)
		}
	}
	if p.master != nil {
		err := p.master.Close()
		if err != nil {
			p.logger.Errorf("[gobatis] close master db error: %s", err)
		}
	}
}

func (p *Engine) SQL(name string, args ...interface{}) {

}

func (p *Engine) BindMapper(ptr ...interface{}) (err error) {
	for _, v := range ptr {
		err = p.bindMapper(v)
		if err != nil {
			return
		}
	}
	return
}

func (p *Engine) bindMapper(mapper interface{}) (err error) {

	rv := reflect.ValueOf(mapper)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("exptect *struct, got: %s", rv.Type())
	}
	rv = rv.Elem()
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		if rv.Field(i).Kind() != reflect.Func {
			continue
		}
		must := false
		stmt := false
		id := rt.Field(i).Name
		if strings.HasPrefix(id, must_prefix) {
			id = strings.TrimPrefix(id, must_prefix)
			must = true
		}
		if strings.HasSuffix(id, stmt_suffix) {
			id = strings.TrimSuffix(id, stmt_suffix)
			stmt = true
		}
		if strings.HasSuffix(id, tx_suffix) {
			id = strings.TrimSuffix(id, tx_suffix)
		}
		m, ok := p.fragmentManager.get(id)
		if !ok {
			if must {
				return fmt.Errorf("%s.(Must)%s statement not defined", rt.Name(), id)
			}
			return fmt.Errorf("%s.%s statement not defined", rt.Name(), id)
		}
		m = m.fork()
		m.must = must
		m.stmt = stmt
		m.id = rt.Field(i).Name
		ft := rv.Field(i).Type()
		m.checkParameter(ft, rt.Name(), rv.Type().Field(i).Name)
		m.checkResult(ft, rt.Name(), rv.Type().Field(i).Name)
		m.proxy(rv.Field(i))
	}
	return
}

func (p *Engine) parseMappers(mapperDir string) (err error) {

	mappers := GetMappers(mapperDir)
	for _, path := range mappers {

		// 打开文件
		file, err := os.Open(path)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		var bs []byte
		bs, err = io.ReadAll(file)

		file.Close()

		if err != nil {
			err = fmt.Errorf("read %s content error: %s", path, err)
			return err
		}

		p.logger.Infof("[gobatis] register fragment: %s", path)
		err = parseMapper(p, path, string(bs))
		if err != nil {
			return err
		}
	}
	return
}

func (p *Engine) addFragment(file string, ctx antlr.ParserRuleContext, id string, node *xmlNode) {

	m, err := parseFragment(p.master, file, id, node)
	if err != nil {
		return
	}
	err = p.fragmentManager.add(m)
	if err != nil {
		throw(file, ctx, parseMapperErr).with(err)
	}
	return
}

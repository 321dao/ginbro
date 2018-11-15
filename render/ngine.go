package render

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

type ParseNgin struct {
	db             *sql.DB
	database       string
	ProjectPath    string
	ProjectPackage string
	MysqlAddr      string
	MysqlUser      string
	MysqlPassword  string
	MysqlDatabase  string
	MysqlCharset   string
	AppSecret      string
	AppAddr        string
	Tables         []*schema
}

func NewParseNgin(mysqlUser, mysqlPassword, mysqlAddr, mysqlDatabase, mysqlCharset, outDir, appAddr string) *ParseNgin {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", mysqlUser, mysqlPassword, mysqlAddr, mysqlDatabase, mysqlCharset)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}
	projectPath := path.Join(goPath, "src", outDir)
	return &ParseNgin{
		db,
		mysqlDatabase,
		projectPath,
		outDir,
		mysqlAddr,
		mysqlUser,
		mysqlPassword,
		mysqlDatabase,
		mysqlCharset,
		randomString(32),
		appAddr,
		[]*schema{},
	}
}
func (p *ParseNgin) Close() {
	p.db.Close()
}

func (p *ParseNgin) makeModuleDir(module string) {
	modulePath := path.Join(p.ProjectPath, module)
	if err := os.MkdirAll(modulePath, 0777); err != nil {
		log.Fatal(err)
	}
}

func (p *ParseNgin) copyFiles() error {

	tasks := map[string]string{
		"main.tpl":         "main.go",
		"config.toml.tpl":  "config.toml",
		"config.viper.tpl": "config/viper.go",
		"handlers.gin.tpl": "handlers/gin.go",
		"models.db.tpl":    "models/db.go",
	}
	for kk, vv := range tasks {
		if err := parseTmpl(kk, p.ProjectPackage, vv, p); err != nil {
			log.Println(err)
		}
	}
	return nil
}
func (p *ParseNgin) copyFolder() {
	tasks := []string{"swagger", "static"}
	for _, vv := range tasks {
		src := path.Join(goPath, "src", "github.com/dejavuzhou/ginbo/boilerplate", vv)
		dst := path.Join(p.ProjectPath, vv)
		if err := CopyDir(src, dst); err != nil {
			log.Println(err)
		}
	}
}

func (p *ParseNgin) GoFmt() {
	log.Println("running go fmt for the new project")

	runCmd("go", "fmt", p.ProjectPackage+"/...")
	log.Println("TODO::test your project")
	//mainPath := path.Join(p.ProjectPath,"main.go")
	//runCmd("go","run",mainPath)
	log.Printf("cd %s", p.ProjectPath)
	log.Println("go run main.go")

}

func runCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("%s", out)
	}
}

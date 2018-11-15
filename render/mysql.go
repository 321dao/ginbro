package render

import (
	"log"
)

func (p *ParseNgin) GenerateCodeFile() {
	p.copyFolder()
	//make project module dir
	for _, m := range []string{"config", "handlers", "models"} {
		p.makeModuleDir(m)
	}
	//copy file
	p.copyFiles()
	//find table where has id as primary key
	sQ := `SELECT DISTINCT information_schema.tables.table_name
		   FROM information_schema.tables JOIN information_schema.columns
           ON information_schema.columns.table_name = information_schema.tables.table_name
	       WHERE information_schema.columns.column_name = 'id' AND information_schema.tables.table_schema= ?`
	rows, err := p.db.Query(sQ, p.database)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if rows.Scan(&tableName) != nil {
			log.Println(err)
		} else {
			//create model and handler for every tableName
			s := p.generateCodeByTable(tableName, generateFiles)
			p.Tables = append(p.Tables, s)
		}
	}
	if err := parseTmpl("swagger.yaml", p.ProjectPackage, "swagger/doc.yml", p); err != nil {
		log.Println(err)
	}
	if err := parseTmpl("app_readme.md", p.ProjectPackage, "readme.md", p); err != nil {
		log.Println(err)
	}
}

func (p *ParseNgin) generateCodeByTable(tableName string, callback func(s *schema)) *schema {
	//log.Println(tableName)
	rawSql := "SELECT `COLUMN_NAME`,`DATA_TYPE`,`COLUMN_COMMENT`,`COLUMN_KEY`,`COLUMN_TYPE` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = ? AND TABLE_NAME = ?"
	rows, err := p.db.Query(rawSql, p.database, tableName)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	schema := newTable(tableName, p.ProjectPath, p.ProjectPackage)
	for rows.Next() {
		var cName, dType, cComment, cKey, cType string
		if rows.Scan(&cName, &dType, &cComment, &cKey, &cType) == nil {
			c := newCol(cName, dType, cComment, cKey, cType)
			schema.Columns = append(schema.Columns, c)
		} else {
			//create model and handler for every tableName
			log.Printf("scan table %s's schema failed", tableName)
		}
	}
	if callback != nil {
		callback(schema)
	}
	return schema
}
func generateFiles(s *schema) {
	s.generateModel()
	s.generateHandler()
}

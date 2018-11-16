package cmd

import (
	"github.com/dejavuzhou/ginbro/render"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

var mysqlUser, mysqlPassword, mysqlAddr, mysqlDatabase, mysqlCharset, outDir, appAddr string

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate a RESTful APIs app with gin and gorm for gophers",
	Long: `generate a RESTful APIs app with gin and gorm for gophers. For example:
	ginbo gen -u eric -p password -a "127.0.0.1:3306" -d "mydb"
`,
	Run: func(cmd *cobra.Command, args []string) {
		//SELECT table_name FROM information_schema.tables where table_schema='venom';
		ng := render.NewParseNgin(mysqlUser, mysqlPassword, mysqlAddr, mysqlDatabase, mysqlCharset, outDir, appAddr)
		ng.GenerateCodeFile()
		ng.GoFmt()
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	genCmd.PersistentFlags().StringVarP(&appAddr, "appAddr", "l", "127.0.0.1:5555", "app listen Address eg:mojotv.cn, use domain will support gin-TLS")

	genCmd.PersistentFlags().StringVarP(&mysqlUser, "user", "u", "root", "database user name")
	genCmd.PersistentFlags().StringVarP(&mysqlPassword, "password", "p", "Venom", "database password")
	genCmd.PersistentFlags().StringVarP(&mysqlAddr, "address", "a", "127.0.0.1:3306", "mysql host:port")
	genCmd.PersistentFlags().StringVarP(&mysqlDatabase, "database", "d", "dbname", "database name")
	genCmd.PersistentFlags().StringVarP(&mysqlCharset, "charset", "c", "utf8", "database charset")
	genCmd.PersistentFlags().StringVarP(&outDir, "out", "o", "github.com/dejavuzhou/gin-project", "golang project package name of your output project. eg: github."+
		"com/awesome/my_project, the project will be created at $GOPATH/src/github.com/awesome/my_project")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//create
}

package generate

import (
	"fmt"

	templates "github.com/golangast/gentil/generate/templates"
	gentil "github.com/golangast/gentil/utility/ff"
	temp "github.com/golangast/gentil/utility/temp"
	term "github.com/golangast/gentil/utility/term"
)

//p-path f-file s-script
//tie the viper config vars to params
func GenServer(p string, f string) {

	//header map for {{define "header"}} {{end}}
	m := make(map[string]string)
	header := fmt.Sprintf(`{{define "header"}}%s`, "")
	end := fmt.Sprintf(`{{end}}%s`, "")
	m["header"] = header
	m["end"] = end

	//footer map for {{define "footer"}} {{end}}
	mf := make(map[string]string)
	footer := fmt.Sprintf(`{{define "footer"}}%s`, "")
	endf := fmt.Sprintf(`{{end}}%s`, "")
	mf["footer"] = footer
	mf["end"] = endf

	//footer/header map for {{template "footer" .}} {{end}}
	mb := make(map[string]string)
	headerb := fmt.Sprintf(`{{template "header" .}}%s`, "")
	footerb := fmt.Sprintf(`{{template "footer" .}}%s`, "")
	mb["footer"] = footerb
	mb["header"] = headerb

	/* create folders*/
	err := gentil.Makefolder(p)
	if err != nil {
		fmt.Print(err)
	}
	err = gentil.Makefolder(p + "/templates")
	if err != nil {
		fmt.Print(err)
	}
	err = gentil.Makefolder(p + "/db")
	if err != nil {
		fmt.Print(err)
	}
	/* create files*/
	bfile, err := gentil.Makefile(p + "/templates/home.html")
	if err != nil {
		fmt.Print(err)
	}
	hfile, err := gentil.Makefile(p + "/templates/header.html")
	if err != nil {
		fmt.Print(err)
	}
	ffile, err := gentil.Makefile(p + "/templates/footer.html")
	if err != nil {
		fmt.Print(err)
	}
	sfile, err := gentil.Filefolder(p, f)
	if err != nil {
		fmt.Print(err)
	}
	dbfile, err := gentil.Makefile(p + "/db/database.db")
	if err != nil {
		fmt.Print(err)
	}

	/* write to files*/
	temp.Writetemplate(templates.Servertemp, sfile, nil)
	if err != nil {
		fmt.Print(err)
	}
	temp.Writetemplate(templates.Headertemp, hfile, m)
	if err != nil {
		fmt.Print(err)
	}
	temp.Writetemplate(templates.Footertemp, ffile, mf)
	if err != nil {
		fmt.Print(err)
	}
	temp.Writetemplate(templates.Bodytemp, bfile, mb)
	if err != nil {
		fmt.Print(err)
	}

	term.Pulldowneverything(p) //pulls dependencies and loads it
	if err != nil {
		fmt.Print(err)
	}

	bfile.Close()
	hfile.Close()
	ffile.Close()
	sfile.Close()
	dbfile.Close()
}

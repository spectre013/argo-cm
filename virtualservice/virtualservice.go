package virtualservice

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

type Values struct {
	Values map[string]string
}

type virtualService struct {
	vars []string
}

type VirtualService interface {
	Process(vars string, mvp bool) error
}


func NewVirtualService() VirtualService {
	return &virtualService{[]string{"DOMAIN", "CUSTOMER","WORKSPACE"}}
}


func ( v *virtualService) Process(vars string, mvp bool) error {
	v.vars = append(v.vars, strings.Split(vars,",")...)
	t, err  := v.getTemplates()
	if err != nil {
		return err
	}
	values := v.getEnvVariables(mvp)
	err = v.updateYaml(values,t)
	if err != nil {
		return err
	}

	return nil
}


func openfile(file string) (*os.File,error) {
	f, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return f, nil
}


func ( v *virtualService) updateYaml(values Values, t *template.Template) error {

	f, err := openfile("./virtualservice.yaml")
	defer f.Close()

	err = t.Execute(f, values)
	if err != nil {
		log.Print("execute: ", err)
		return err
	}

	return nil
}

func ( v *virtualService) getEnvVariables(mvp bool) Values {
	val := map[string]string{}
	values := Values{}
	for _,v := range v.vars {
		fmt.Println(v, os.Getenv(v))
		val[strings.ToLower(v)] = os.Getenv(v)
	}

	if mvp {
		val["hostname"] = fmt.Sprintf("-%s-%s",os.Getenv("WORKSPACE"),os.Getenv("CUSTOMER"))
	}

	values.Values = val

	return values
}


func ( v *virtualService) getTemplates() (*template.Template, error) {
	t, err := template.ParseFiles("virtualservice.yaml")
	if err != nil {
		log.Print(err)
		return t, err
	}
	return t, nil
}
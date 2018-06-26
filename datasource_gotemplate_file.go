package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"text/template"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/bradfitz/iter"
)

// ---------------------------------------

func hash(s string) string {
	sha := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sha[:])
}

type templateRenderError error

func renderFile(d *schema.ResourceData) (string, error) {

	var err error


	var data string // data from tf
	data = d.Get("data").(string)

	// unmarshal json from data into m
	var m = make(map[string]interface{}) // unmarshal data into m
	if err = json.Unmarshal([]byte(data), &m); err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"N": iter.N,
	}

	templateText := d.Get("template").(string)
	tt,err  := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		panic(err)
	}

	var contents bytes.Buffer // io.writer for template.Execute
	if tt != nil {
		err = tt.Execute(&contents, m)
		if err != nil { 
			return "", templateRenderError(fmt.Errorf("failed to render %v", err))
		}
	} else {
		return "", templateRenderError(fmt.Errorf("error: %v", err))
	}

	return contents.String(), nil
}

func dataSourceFileRead(d *schema.ResourceData, meta interface{}) error {
	rendered, err := renderFile(d)
	if err != nil {
		return err
	}
	d.Set("rendered", rendered)
	d.SetId(hash(rendered))
	return nil
}

func dataSourceFile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFileRead,

		Schema: map[string]*schema.Schema{
			"template": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "path to go template file",
			},
			"data": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "",
				Description:  "variables to substitute",
				ValidateFunc: nil,
			},
			"rendered": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "rendered template",
			},
		},
	}
}

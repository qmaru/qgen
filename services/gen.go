package services

import (
	"os"
	"path/filepath"
	"text/template"
)

type templates struct {
	templates []templateData
}

type templateData struct {
	Entries []templateCommon
}

type templateEntry struct {
	Name string
}

type templateCommon struct {
	FirstPath   string
	SecnondPath string
	ThirdPath   string
	File        string
	Template    *template.Template
}

type Project struct {
	Name     string
	Type     string
}

func (p *Project) mainTemplate() templateData {
	return templateData{
		Entries: []templateCommon{
			{
				FirstPath: "",
				File:      "main.go",
				Template:  MainEntry(),
			},
		},
	}
}

func (p *Project) apiTemplate() templateData {
	return templateData{
		Entries: []templateCommon{
			{
				FirstPath: "apis",
				File:      "router.go",
				Template:  ApisRouter(),
			},
			{
				FirstPath:   "apis",
				SecnondPath: "common",
				File:        "handler.go",
				Template:    ApisCommonHandler(),
			},
		},
	}
}

func (p *Project) cmdTemplate() templateData {
	return templateData{
		Entries: []templateCommon{
			{
				FirstPath: "cmds",
				File:      "root.go",
				Template:  CmdsRoot(),
			},
			{
				FirstPath:   "cmds",
				SecnondPath: "api",
				File:        "api.go",
				Template:    CmdsApiApi(),
			},
			{
				FirstPath:   "cmds",
				SecnondPath: "db",
				File:        "db.go",
				Template:    CmdsDbDb(),
			},
		},
	}
}

func (p *Project) configTemplate() templateData {
	return templateData{
		Entries: []templateCommon{
			{
				FirstPath: "configs",
				File:      "config.go",
				Template:  ConfigsConfig(),
			},
			{
				FirstPath: "configs",
				File:      "env.go",
				Template:  ConfigsEnv(),
			},
			{
				FirstPath: "configs",
				File:      "database.json",
				Template:  ConfigsDatabaseJson(),
			},
			{
				FirstPath: "configs",
				File:      "config.json",
				Template:  ConfigsConfigJSON(),
			},
		},
	}
}

func (p *Project) dbTemplate() templateData {
	return templateData{
		Entries: []templateCommon{
			{
				FirstPath: "dbs",
				File:      "dbs.go",
				Template:  DbDbsDbs(),
			},
			{
				FirstPath:   "dbs",
				SecnondPath: "models",
				File:        "model.go",
				Template:    DbDbsModelsModel(),
			},
		},
	}
}

func (p *Project) serviceTemplate() templateData {
	return templateData{
		Entries: []templateCommon{
			{
				FirstPath:   "services",
				SecnondPath: "common",
				ThirdPath:   "logs",
				File:        "common.go",
				Template:    ServiceServicesCommonLogsCommon(),
			},
			{
				FirstPath:   "services",
				SecnondPath: "common",
				ThirdPath:   "logs",
				File:        "logs.go",
				Template:    ServiceServicesCommonLogsLogs(),
			},
		},
	}
}

func (p *Project) utilsTemplate() templateData {
	return templateData{
		Entries: []templateCommon{
			{
				FirstPath: "utils",
				File:      "minireq.go",
				Template:  UtilsUtilsMinireq(),
			},
			{
				FirstPath: "utils",
				File:      "tools.go",
				Template:  UtilsUtilsTools(),
			},
			{
				FirstPath: "utils",
				File:      "version.go",
				Template:  UtilsUtilsVersion(),
			},
		},
	}
}

func (p *Project) Gen() error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}

	rootFolder := filepath.Join(root, p.Name)
	projectRoot, err := createFolder(rootFolder)
	if err != nil {
		return err
	}

	var tempData templates
	switch p.Type {
	case "web":
		tempData = templates{
			templates: []templateData{
				p.mainTemplate(),
				p.apiTemplate(),
				p.cmdTemplate(),
				p.configTemplate(),
				p.dbTemplate(),
				p.serviceTemplate(),
				p.utilsTemplate(),
			},
		}
	case "cli":
		tempData = templates{
			templates: []templateData{
				p.mainTemplate(),
				p.cmdTemplate(),
				p.configTemplate(),
				p.dbTemplate(),
				p.serviceTemplate(),
				p.utilsTemplate(),
			},
		}
	}

	for _, temps := range tempData.templates {
		for _, c := range temps.Entries {
			entrypath := filepath.Join(projectRoot, c.FirstPath, c.SecnondPath, c.ThirdPath)
			err = os.MkdirAll(entrypath, os.ModePerm)
			if err != nil {
				return err
			}

			entryfile := filepath.Join(entrypath, c.File)
			f, err := os.Create(entryfile)
			if err != nil {
				return err
			}
			defer f.Close()

			err = c.Template.Execute(f, p)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

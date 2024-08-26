.vscode
app
├── framework
| ├── echo-utils.go
| ├── framework-utils.go
| └── language.go
└── make/
| ├── app.mk
| ├── ci.mk
| ├── deploy.mk
| ├── dev.mk
| ├── help.mk
| └── terraform.mk
├── middlewares
| ├── i18n.go
| └── \*.go // other middlewares
├── routes
| ├── api-routes.go
| └── page-routes.go
├── scripts
| └── guiso.sh
├── templ
| ├── components
| | ├── \*.templ // components
| ├── pages
| | ├── error.templ
| | ├── index.templ
| | ├── \*.templ // other pages
| ├── ts
| | ├── main.ts
| | └── other-js-files.ts
| ├── types
| | └── \*.go // shared types
| ├── utils
| | ├── http-utils.go
| | ├── navigation-utils.go // here
| | └── templ-utils.go
| └── main.go
├── terraform/
│ ├── main.tf
│ ├── variables.tf
│ ├── outputs.tf
│ └── terraform.tfvars
public
├── css
├── images
├── out
├── \* // other static files folders
temp // temp folder
.commitlintrc.yaml
.dockerignore
.gitignore
.golangci.yml
.pre-commit-config.yaml
Dockerfile
esbuild.go
global.css
go.mod
go.sum
Makefile
README.md
tailwind.config.js

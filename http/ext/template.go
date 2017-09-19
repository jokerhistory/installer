package ext

//https://studygolang.com/articles/464
import (
	"os"
	"text/template"
)
//curl -i -X POST -H 'Content-type':'application/yaml' --data-binary @output.yml  45.55.29.141:8080/flow/v1/containerops/python_analysis_coala/flow/latest/yaml
//http://blog.csdn.net/sryan/article/details/52353937
func Ctemplate(tpString string value) {
	header := `name: cncf/python-for-cncf-ci/build-test-release-deploy
	number: 33
	title: Components For Python
	version: 4
	tag: latest
	timeout: 0
	stages:
		-
			type: start
			name: start
			title: Start
		-
			type: normal
			name: prometheus-test-build-release-python-ci-components
			title: Building, testing python project, compile then upload to Dockyard artifact repository.
			sequencing: sequence
			actions:`

	action := `-
	name: build-python-analysis-coala
	title: Build python project with \"make build\"
	jobs:
	-
		 type: component
		 kubectl: python/test.yaml
		 endpoint: hub.opshub.sh/containerops/analysis_java_gradle_checkstyle
		 resources:
			 cpu: 2
			 memory: 4G
		 timeout: 0
		 environments:
				- CO_DATA: \"version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json\"
-`  
			foot := `  -
			type: end
			name: end
			title: End`
	
		WriteFile(header,action,foot)
  
	string value := "123"; 
	tmpl, err := template.New("test").Parse("hello,111 {{.}}") //建立一个模板，内容是"hello, {{.}}"
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, value) //将string与模板合成，变量name的内容会替换掉{{.}}
	//合成结果放到os.Stdout里
	if err != nil {
		panic(err)
	}
}

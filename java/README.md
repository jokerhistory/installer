docker build -t analysis_java_gradle_checkstyle analysis/checkstyle 
docker build -t analysis_java_gradle_cpd analysis/cpd
docker build -t analysis_java_gradle_dependencies analysis/dependencies
docker build -t analysis_java_gradle_findbugs analysis/findbugs
docker build -t analysis_java_gradle_jdepend analysis/jdepend
docker build -t analysis_java_gradle_pmd analysis/pmd
docker build -t compile_java_gradle_war compile/war
docker build -t compile_java_gradle_jar compile/jar
docker build -t compile_java_gradle_ear compile/ear
docker build -t document_java_gradle_javadoc document/javadoc
docker build -t test_java_gradle_jacoco test/jacoco
docker build -t test_java_gradle_junit test/junit
docker build -t test_java_gradle_testng test/testng

docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" analysis_java_gradle_checkstyle
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=xml" analysis_java_gradle_cpd
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git" analysis_java_gradle_dependencies
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" analysis_java_gradle_findbugs
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" analysis_java_gradle_jdepend
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" analysis_java_gradle_pmd
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/rominirani/GradleWebAppSample.git target=https://hub.opshub.sh/binary/v1/lidian/test/binary/2.2.4/web.war" compile_java_gradle_war
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git target=https://hub.opshub.sh/binary/v1/containerops/component/binary/2.2.4/demo.jar" compile_java_gradle_jar
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git target=https://hub.opshub.sh/binary/v1/containerops/component/binary/2.2.4/demo.ear" compile_java_gradle_ear
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git target=https://hub.opshub.sh/binary/v1/lidian/test/binary/1.1.0/javadoc.tar" document_java_gradle_javadoc
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" test_java_gradle_jacoco
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" test_java_gradle_junit
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" test_java_gradle_testng

docker tag bd9a35ce4ffc hub.opshub.sh/containerops/analysis_java_gradle_checkstyle:latest
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/analysis_java_gradle_cpd:latest
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/analysis_java_gradle_dependencies:latest
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/analysis_java_gradle_findbugs:latest
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/analysis_java_gradle_jdepend:latest
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/analysis_java_gradle_pmd:latest
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/compile_java_gradle_war:latest
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/compile_java_gradle_jar:latest
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/compile_java_gradle_ear:latest
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/document_java_gradle_javadoc:latest
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/test_java_gradle_jacoco:latest
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/test_java_gradle_junit:latest
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/test_java_gradle_testng:latest

docker push hub.opshub.sh/containerops/analysis_java_gradle_checkstyle:latest
docker push hub.opshub.sh/containerops/analysis_java_gradle_cpd:latest
docker push hub.opshub.sh/containerops/analysis_java_gradle_dependencies:latest
docker push hub.opshub.sh/containerops/analysis_java_gradle_findbugs:latest
docker push hub.opshub.sh/containerops/analysis_java_gradle_jdepend:latest
docker push hub.opshub.sh/containerops/analysis_java_gradle_pmd:latest
docker push hub.opshub.sh/containerops/compile_java_gradle_war:latest
docker push hub.opshub.sh/containerops/compile_java_gradle_jar:latest
docker push hub.opshub.sh/containerops/compile_java_gradle_ear:latest
docker push hub.opshub.sh/containerops/document_java_gradle_javadoc:latest
docker push hub.opshub.sh/containerops/test_java_gradle_jacoco:latest
docker push hub.opshub.sh/containerops/test_java_gradle_junit:latest
docker push hub.opshub.sh/containerops/test_java_gradle_testng:latest

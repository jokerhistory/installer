docker build -t analysis_java_gradle_checkstyle analysis/checkstyle <br>
docker build -t analysis_java_gradle_cpd analysis/cpd<br>
docker build -t analysis_java_gradle_dependencies analysis/dependencies<br>
docker build -t analysis_java_gradle_findbugs analysis/findbugs<br>
docker build -t analysis_java_gradle_jdepend analysis/jdepend<br>
docker build -t analysis_java_gradle_pmd analysis/pmd<br>
docker build -t compile_java_gradle_war compile/war<br>
docker build -t compile_java_gradle_jar compile/jar<br>
docker build -t compile_java_gradle_ear compile/ear<br>
docker build -t document_java_gradle_javadoc document/javadoc<br>
docker build -t test_java_gradle_jacoco test/jacoco<br>
docker build -t test_java_gradle_junit test/junit<br>
docker build -t test_java_gradle_testng test/testng<br>

docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" analysis_java_gradle_checkstyle<br>
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=xml" analysis_java_gradle_cpd<br>
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git" analysis_java_gradle_dependencies<br>
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" analysis_java_gradle_findbugs<br>
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" analysis_java_gradle_jdepend<br>
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" analysis_java_gradle_pmd<br>
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/rominirani/GradleWebAppSample.git target=https://hub.opshub.sh/binary/v1/lidian/test/binary/2.2.4/web.war" compile_java_gradle_war<br>
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git target=https://hub.opshub.sh/binary/v1/containerops/component/binary/2.2.4/demo.jar" compile_java_gradle_jar<br>
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git target=https://hub.opshub.sh/binary/v1/containerops/component/binary/2.2.4/demo.ear" compile_java_gradle_ear<br>
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git target=https://hub.opshub.sh/binary/v1/lidian/test/binary/1.1.0/javadoc.tar" document_java_gradle_javadoc<br>
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" test_java_gradle_jacoco<br>
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" test_java_gradle_junit<br>
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" test_java_gradle_testng<br>

docker tag bd9a35ce4ffc hub.opshub.sh/containerops/analysis_java_gradle_checkstyle:latest<br>
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/analysis_java_gradle_cpd:latest<br>
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/analysis_java_gradle_dependencies:latest<br>
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/analysis_java_gradle_findbugs:latest<br>
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/analysis_java_gradle_jdepend:latest<br>
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/analysis_java_gradle_pmd:latest<br>
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/compile_java_gradle_war:latest<br>
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/compile_java_gradle_jar:latest<br>
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/compile_java_gradle_ear:latest<br>
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/document_java_gradle_javadoc:latest<br>
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/test_java_gradle_jacoco:latest<br>
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/test_java_gradle_junit:latest<br>
docker tag bd9a35ce4ffc hub.opshub.sh/containerops/test_java_gradle_testng:latest<br>

docker push hub.opshub.sh/containerops/analysis_java_gradle_checkstyle:latest<br>
docker push hub.opshub.sh/containerops/analysis_java_gradle_cpd:latest<br>
docker push hub.opshub.sh/containerops/analysis_java_gradle_dependencies:latest<br>
docker push hub.opshub.sh/containerops/analysis_java_gradle_findbugs:latest<br>
docker push hub.opshub.sh/containerops/analysis_java_gradle_jdepend:latest<br>
docker push hub.opshub.sh/containerops/analysis_java_gradle_pmd:latest<br>
docker push hub.opshub.sh/containerops/compile_java_gradle_war:latest<br>
docker push hub.opshub.sh/containerops/compile_java_gradle_jar:latest<br>
docker push hub.opshub.sh/containerops/compile_java_gradle_ear:latest<br>
docker push hub.opshub.sh/containerops/document_java_gradle_javadoc:latest<br>
docker push hub.opshub.sh/containerops/test_java_gradle_jacoco:latest<br>
docker push hub.opshub.sh/containerops/test_java_gradle_junit:latest<br>
docker push hub.opshub.sh/containerops/test_java_gradle_testng:latest<br>

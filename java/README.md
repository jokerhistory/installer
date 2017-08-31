docker build -t containerops/analysis/java_gradle_checkstyle analysis/checkstyle 
docker build -t containerops/analysis/java_gradle_cpd analysis/cpd
docker build -t containerops/analysis/java_gradle_dependencies analysis/dependencies
docker build -t containerops/analysis/java_gradle_findbugs analysis/findbugs
docker build -t containerops/analysis/java_gradle_jdepend analysis/jdepend
docker build -t containerops/analysis/java_gradle_pmd analysis/pmd
docker build -t huawei/compile/java_gradle_war compile/war
docker build -t huawei/compile/java_gradle_jar compile/jar
docker build -t huawei/compile/java_gradle_ear compile/ear
docker build -t huawei/document/java_gradle_javadoc document/javadoc
docker build -t containerops/test/java_gradle_jacoco test/jacoco
docker build -t containerops/test/java_gradle_junit test/junit
docker build -t containerops/test/java_gradle_testng test/testng

docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" containerops/analysis/java_gradle_checkstyle
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=xml" containerops/analysis/java_gradle_cpd
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git" containerops/analysis/java_gradle_dependencies
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" containerops/analysis/java_gradle_findbugs
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" containerops/analysis/java_gradle_jdepend
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" containerops/analysis/java_gradle_pmd
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/rominirani/GradleWebAppSample.git target=https://hub.opshub.sh/binary/v1/lidian/test/binary/2.2.4/web.war" huawei/compile/java_gradle_war
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git target=https://hub.opshub.sh/binary/v1/containerops/component/binary/2.2.4/demo.jar" huawei/compile/java_gradle_jar
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git target=https://hub.opshub.sh/binary/v1/containerops/component/binary/2.2.4/demo.ear" huawei/compile/java_gradle_ear
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git target=https://hub.opshub.sh/binary/v1/lidian/test/binary/1.1.0/javadoc.tar" huawei/document/java_gradle_javadoc
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" containerops/test/java_gradle_jacoco
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" containerops/test/java_gradle_junit
docker run --rm --env CO_DATA="version=gradle3 git-url=https://github.com/vanniuner/gradle-demo.git out-put-type=json" containerops/test/java_gradle_testng

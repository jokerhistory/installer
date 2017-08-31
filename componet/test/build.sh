#!/usr/bin/env bash
#  cd component/python
find  ./ -name 'test.md' | while read F;do
  DN=`analysis/$F`;
  echo $DN
  #NAME=`basename $DN`;
  #docker build -t containerops/coala1 -f ./analysis/coala/Dockerfile ./analysis/coala
  
  docker build -t containerops/$DN -f ./$DN/Dockerfile ./$DN;
    echo docker build -t containerops/$DN -f ./$DN/Dockerfile ./$DN

  #docker build -t containerops/$NAME -f $DN/Dockerfile $DN;

  #HASH=`docker images | grep containerops/$NAME | grep latest | grep -v hub.opshub.sh | awk '{print $3}'`;
  #docker tag $HASH hub.opshub.sh/containerops/$NAME:latest;
  #docker push hub.opshub.sh/containerops/$NAME:latest;
done

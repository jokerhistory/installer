#!/usr/bin/env bash
#  cd component/python
find  ./ -name 'README.md' | while read F;do
  DN=`dirname $F`;
  NAME=`basename $DN`;
  docker build -t containerops/$NAME -f $DN/Dockerfile $DN;
  HASH=`docker images | grep containerops/$NAME | grep latest | grep -v hub.opshub.sh | awk '{print $3}'`;
  docker tag $HASH hub.opshub.sh/containerops/$NAME:latest;
  docker push hub.opshub.sh/containerops/$NAME:latest;
done

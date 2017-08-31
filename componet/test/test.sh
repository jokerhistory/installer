#!/usr/bin/env bash
#  cd component/python
echo '#!/usr/bin/env bash' > all_test.sh
find ./ -name 'test.md' | while read F;do
  grep 'docker run' $F >> all_test.sh
done

chmod +x all_test.sh

./all_test.sh


    
while read line 
 do 
  DN=$line;
  fstr=`echo $DN | cut -d \/ -f 1`
	sstr=`echo $DN | cut -d \/ -f 2`
	echo $fstr $sstr
  docker build -t containerops/$DN -f ./$line/Dockerfile ./$DN;
	echo [docker build -t containerops/$line -f ./$line/Dockerfile ./$line/]
	
	sed "s/config:.*/config:abc/g" -i template
	
done < test2.md >>test3.md


docker build . -t tracker || exit
docker stop tracker | true
docker rm tracker | true
docker run -v `pwd`:/go/src/github.com/ericfengchao/treasure-hunting \
  -ti \
  --name tracker \
  -p 50050:50050 \
  tracker
#!/usr/bin/env bash

ORG_QNTY=3

while getopts n: opt; do
    case $opt in
        n)  ORG_QNTY=${OPTARG}
            ;;
    esac
done

if [ $ORG_QNTY != 3 -a $ORG_QNTY != 1 ]
then
  echo 'WARNING: The number of organizations allowed is either 3 or 1.'
  echo 'Defaulting to 3 organizations.'
  ORG_QNTY=3
fi

CCCG_PATH="../chaincode/collections.json"

./network.sh down -n $ORG_QNTY
rm -rf organizations/peerOrganizations
rm -rf organizations/ordererOrganizations
rm -rf organizations/rest-certs


download_binaries(){
  echo "Preparing to download fabric binaries..."
  curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh

  echo "Downloading fabric binaries..."
  ./install-fabric.sh --fabric-version 2.5.3 binary

  rm install-fabric.sh
}

FILE=bin
if [ ! -d "$FILE" ]; then
  echo "Directory $FILE not found"
  download_binaries
else
  cd bin;
  numFiles="$(ls -1 | wc -l)"
  if [ "$numFiles" -ne 10 ];
  then
    cd ..
    echo "Missing some fabric binaries"
    download_binaries
  else
    cd ..
  fi
fi

docker network create cc-tools-demo-net
./network.sh up createChannel -n $ORG_QNTY
./network.sh deployCC -ccn cc-tools-demo -ccp ../chaincode -ccl go -n $ORG_QNTY -cccg $CCCG_PATH
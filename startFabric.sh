#!/bin/bash

set -e

starttime=$(date +%s)
CC_SRC_LANGUAGE="go"
CC_SRC_PATH="../../chaincode/"
CC_Name="chaincode"
CC_init="InitCertification"
CC_asset="../../collections_config.json"

curl -sSL https://bit.ly/2ysbOFE | bash -s -- 2.3.1 1.4.9
export PATH=${PWD}/fabric-samples/bin:$PATH

# launch network; create channel and join peer to channel
pushd ./fabric-samples/test-network
./network.sh down
echo COMPOSE_PROJECT_NAME=docker >> .env
./network.sh up createChannel -ca -s couchdb
./network.sh deployCC -ccn ${CC_Name} -ccp ${CC_SRC_PATH} -cci ${CC_init} -ccl ${CC_SRC_LANGUAGE} -cccg ${CC_asset}

popd

cat <<EOF

Succesfully Start Fabric and Depoly Chaincode.

Total setup execution time : $(($(date +%s) - starttime)) secs ...

EOF

cd backend

yarn install

yarn start
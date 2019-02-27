#!/usr/bin/env bash
echo Building justondavies/go_keyring:build

sudo docker build                                                          \
  --network host                                                           \
  --file dockerfiles/all.docker                                            \
  --tag justondavies/go_keyring:build                                      \
  ./

sudo docker create                                                         \
  --name build_extract                                                     \
  justondavies/go_keyring:build

rm -rf ./build/browser*

sudo docker cp                                                             \
  build_extract:/go/src/github.com/justondavies/go_keyring/build           \
  ./

sudo docker rm -f build_extract

sudo chown -R $USER:$USER ./build
chmod -R 700 ./build

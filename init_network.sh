#!/bin/bash
if [ "$#" -ne 1 ]; then
    echo "ERROR: Peers number must be set for private ipfs network"
    echo "usage: init_network.sh \${peerNumber}"
    echo "For example: Run this command"
    echo "                 ./init_network.sh 3"
    echo "             A private ipfs network with 3 peers will be setup locally"
    exit 1
else
    size_of_net=$1
fi
exists=
hostip=$(hostname -I | awk '{split($0,a); print a[2]}')
peerIDs=$(ipfs config show | grep "PeerID" | awk '{split($0,a); print a[2]}' | sed -e 's/^"//' -e 's/"$//')
DeamonReady="Daemon is ready"
for ((node_count=1; node_count<=$1; node_count++))
do
    if [ ! "$(docker ps -a | grep "ipfs_node$node_count")" ]; then
        ipfs_staging=~/Documents/InsightDC/testnet_volumes/"docker_stage${node_count}"
        ipfs_data=~/Documents/InsightDC/testnet_volumes/"docker_stage${node_count}"
        FILE=~/.ipfs/swarm.key

        rm -rf $ipfs_staging
        mkdir -p $ipfs_staging

        rm -rf $ipfs_data
        mkdir -p $ipfs_data

        if ! [ -f "$FILE"  ]; then
          echo -e "/key/swarm/psk/1.0.0/\n/base16/\n`tr -dc 'a-f0-9' < /dev/urandom | head -c64`" > ~/.ipfs/swarm.key
        fi

        sleep 2
        swarmkey=$(cat ~/.ipfs/swarm.key)
        cp ~/.ipfs/swarm.key $ipfs_data
        docker run -d -it --name "ipfs_node${node_count}" \
                  -v $ipfs_staging:/export \
                  -v $ipfs_data:/data/ipfs \
                  -p $((4001 + ${node_count})):4001 \
                  -p $((5001 + ${node_count})):5001 \
                  -p 127.0.0.1:$((8080 + ${node_count})):8080 \
                  -e LIBP2P_FORCE_PNET=1 \
                  -e IPFS_SWARM_KEY="$swarmkey" \
                  ipfs/go-ipfs:latest \
                  daemon --writable --enable-pubsub-experiment --migrate=true
    else
        echo "conatiner ipfs_node${node_count} already exists."
        echo "starting container"
        docker start "ipfs_node${node_count}"
        exists=true
    fi

done
# if [ "$exists" == true ]; then
#   exit 1
# fi

node_count=1
while :
do
  if docker exec ipfs_node${node_count} ipfs config show | grep "PeerID" ; then
    if [ "$node_count" -eq "$1" ]; then
      break
    fi
    node_count=$(($node_count + 1))
    sleep 1
  else
    echo "no peerid yet?..."
  fi

done

for ((i=1; i<=$1; i++))
do
    echo "removing bootstraps"
    sleep 15
    docker exec "ipfs_node${i}" ipfs bootstrap rm --all
    echo -e "adding bootstrap...\n"
    docker exec "ipfs_node${i}" ipfs bootstrap add /ip4/$hostip/tcp/4001/ipfs/"$peerIDs"

done

1. 设置docker network
    docker network create --subnet 172.12.0.0/16 --gateway 172.12.0.1 zookeeper_kafka
2. 设置zoo.cfg
3. docker-compose.yml
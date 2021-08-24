# elastic-demo

## usage

1. 下载 `MovieLens` [最小测试数据集](https://grouplens.org/datasets/movielens)
2. 将数据集放入 `data` 目录下，因为 `/data/ml-latest-small` 目录与 `logstash` data 做了映射，这样的话，启动镜像就会执行导入数据操作，可修改 `/data/logstash/pipeline/logstash.conf` 来导入不同的数据
3. 进入 `docker-elastic` 目录, `docker-compose up -d` 启动容器
4. 进入kibana [http://127.0.0.1:5601](http://127.0.0.1:5601/app/home#/) 可视化数据

## documents

1. [elasticsearch 7.x guide](https://www.elastic.co/guide/en/elasticsearch/reference/7.x/index.html)
2. [logstash guide](https://www.elastic.co/guide/en/logstash/current/index.html)
3. [olivere/elastic](https://github.com/olivere/elastic)
4. [geektime-ELK](https://github.com/onebirdrocks/geektime-ELK)
5. [docker-elk](https://github.com/deviantony/docker-elk)
6. [go操作ElasticSearch](http://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/go%E6%93%8D%E4%BD%9Celasticsearch/)
7. [使用Elasticsearch(附Golang代码)](https://strconv.com/posts/use-elastic/)

---

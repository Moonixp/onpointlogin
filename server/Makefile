

build:
	docker build --rm . -t onpointserver && docker rmi `docker images --filter label=builder=true -q`

run:
	docker rm test-server-op
	docker run -e DB_IP='172.17.0.1'  -e DB_PASS='root' --name test-server-op  onpointserver	

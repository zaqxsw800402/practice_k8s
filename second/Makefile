docker_run: docker_delete
	docker run -d -e HOST="http://localhost" -e PORT=6666 --name second second

docker_delete:
	docker rm -f second

docker_build: docker_delete
	cd .. ; docker build -t second -f second/Dockerfile .

docker_push:
	docker tag second ichih/second
	docker push ichih/second

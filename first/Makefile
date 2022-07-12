docker_run: docker_delete
	docker run -d -p 6666:6666 --name first first

docker_delete:
	docker rm -f first

docker_build: docker_delete
	cd .. ; docker build -t first -f first/Dockerfile .

docker_push: docker_build
	docker tag first ichih/first
	docker push ichih/first

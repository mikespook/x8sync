build: build-dst build-src

run: run-dst run-src

build-dst:
	docker build -t rsync -f Dockerfile_dst .

run-dst:
	docker run -it --rm --entrypoint /bin/ash rsync

build-src:
	docker build -t sshd -f Dockerfile_src .

run-src:
	docker run -it --rm --entrypoint /bin/ash sshd

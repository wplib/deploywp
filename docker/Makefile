all:
	-@docker rm -f deploywp-0.9
	-@docker image rm -f gearboxworks/deploywp:0.9
	@cp ../dist/deploywp_linux_amd64/deploywp build/rootfs/usr/local/bin/
	@docker build -f Dockerfile -t gearboxworks/deploywp:0.9 .

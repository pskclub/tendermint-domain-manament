start:
	docker-compose up -d
stop:
	docker-compose down
logs:
	 docker logs tendermint_core -f
logs-abci:
	 docker logs tendermint_abci -f

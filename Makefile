start:
	docker-compose up -d --build
stop:
	docker-compose down
logs:
	 docker logs tendermint_core -f
logs-abci:
	 docker logs tendermint_abci -f
exec:
	docker exec -it tendermint_core bash
exec-app:
	docker exec -it tendermint_abci bash

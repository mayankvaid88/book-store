db:
	docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=root -d postgres
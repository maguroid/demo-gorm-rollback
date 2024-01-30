.PHONEY: db

db/down:
	docker-compose down

db/up: db/down
	docker-compose up --build -d

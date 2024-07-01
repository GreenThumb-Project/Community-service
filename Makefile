 CURRENT_DIR=$(shell pwd)
DBURL := postgres://postgres:1918@localhost:5432/user_service?sslmode=disable

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}


mig-up:
	migrate -path databases/migrations -database '${DBURL}' -verbose up

mig-down:
	migrate -path databases/migrations -database '${DBURL}' -verbose down

mig-force:
	migrate -path databases/migrations -database '${DBURL}' -verbose force 1


mig-create-user:
	migrate create -ext sql -dir databases/migrations -seq create_users_table

mig-create-schedule:
	migrate create -ext sql -dir databases/migrations -seq create_user_expertise_table
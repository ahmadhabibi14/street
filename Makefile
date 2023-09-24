
setup:
	go get -u -v github.com/kokizzu/gotro@latest
	go install github.com/cosmtrek/air@latest
	go install github.com/fatih/gomodifytags@latest
	go install github.com/kokizzu/replacer@latest
	go install github.com/akbarfa49/farify@latest
	go install golang.org/x/tools/cmd/goimports@latest
	curl -fsSL https://bun.sh/install | bash

local-tarantool:
	docker exec -it `docker ps | grep tarantool | cut -d ' ' -f 1` tarantoolctl connect userT:passT@localhost:3301
	# box.space -- list all tables
	# box.execute [[ SELECT * FROM "users" LIMIT 1 ]]
	# \set language sql
	# \set delimiter ;

local-clickhouse:
	docker exec -it `docker ps | grep clickhouse | cut -d ' ' -f 1` clickhouse-client --host localhost --port 9000 --user userC --password passC
	# SHOW TABLES -- list all tables
	# SELECT * FROM "actionLogs" LIMIT 1;

modtidy:
	sudo chmod -R a+rwx tmpdb && go mod tidy

fixtags:
	# fix struct tags
	go generate street/domain

orm:
	# generate ORM
	./gen-orm.sh

views:
	# generate views and routes
	./gen-views.sh

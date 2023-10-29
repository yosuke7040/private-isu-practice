all: bin/benchmarker

bin/benchmarker: *.go checker/*.go
	go build -o bin/benchmarker

start: 
	echo "" > ../webapp/logs/mysql/mysql-slow.log
	echo "" > ../webapp/logs/mysql/pt-query-digest.log
	echo "" > ../webapp/logs/nginx/access.log
	echo "" > ../webapp/logs/nginx/error.log
	docker run --rm --network host --add-host host.docker.internal:host-gateway -i private-isu-benchmarker /opt/go/bin/benchmarker -t http://host.docker.internal -u /opt/go/userdata -d
	grep -E '^{.*}$$' ../webapp/logs/nginx/access.log > ../webapp/logs/nginx/json_access.log
	pt-query-digest ../webapp/logs/mysql/mysql-slow.log > ../webapp/logs/mysql/pt-digest/$$(date +%Y%m%d%H%M%S).log

MIGRATIONS_COUNT=$(ls sql | wc -l)
TAG="0.1.$(echo $MIGRATIONS_COUNT)"

docker build -t ringob/service-user-hw2-postgres-migrator:$TAG .
docker push ringob/service-user-hw2-postgres-migrator:$TAG
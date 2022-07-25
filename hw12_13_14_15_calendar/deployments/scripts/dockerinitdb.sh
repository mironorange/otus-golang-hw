
CONTAINER_NAME='calendar-postgresql'
CONTAINER_ID=$(docker ps -qf "name=${CONTAINER_NAME}")

[ -z "$CONTAINER_ID" ] && echo "Container is not running." && exit 1

docker cp deployments/scripts/initdb.sh "$CONTAINER_ID:/tmp/initdb.sh"
docker exec "$CONTAINER_ID" bash /tmp/initdb.sh calendar

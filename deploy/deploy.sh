DOCKER_COMPOSE="docker-compose-"
EXT=".yml"


# shellcheck disable=SC1060
# shellcheck disable=SC1073
# shellcheck disable=SC1048
# shellcheck disable=SC1009
if [ "$#" == 0 ]; then
  echo "Please input app name which you want deploy"
else
  # shellcheck disable=SC2066
  # shellcheck disable=SC2034
  for app in "$*" ; do
    echo "start deploy service ""$app"""
    docker stack deploy -c "$DOCKER_COMPOSE""$app""$EXT" "mars"
  done
fi


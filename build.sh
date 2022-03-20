# shellcheck disable=SC1113
# /bin/bash

APP_AUTH_PATH="app/auth"
APP_SYSTEM_PATH="app/system"
APP_CHIEF_PATH="app/chief"

APP_AUTH_NAME="auth"
APP_SYSTEM_NAME="system"
APP_CHIEF_NAME="chief"


# shellcheck disable=SC1073
# shellcheck disable=SC1055
# shellcheck disable=SC1009
function auth() {
    # shellcheck disable=SC2164
    cd $APP_AUTH_PATH
    make build
    make docker
    make push
    cd ../../
}

function system() {
    # shellcheck disable=SC2164
    cd $APP_SYSTEM_PATH
    make build
    make docker
    make push
    cd ../../
}

function chief() {
    # shellcheck disable=SC2164
    cd $APP_CHIEF_PATH
    make build
    make docker
    make push
    cd ../../
}

# shellcheck disable=SC2205
if [ "$#" == 0 ]
then
  auth
  system
  chief
else
  # shellcheck disable=SC1060
  # shellcheck disable=SC1073
  # shellcheck disable=SC2066
  for app in "$*" ; do
    # shellcheck disable=SC1072
    # shellcheck disable=SC1020
    # shellcheck disable=SC1009
    if [ "$app" ==  $APP_AUTH_NAME ]; then
        auth
    fi
    if [ "$app" ==  $APP_SYSTEM_NAME ]; then
        system
    fi
    if [ "$app" ==  $APP_CHIEF_NAME ]; then
        chief
    fi
  done
fi



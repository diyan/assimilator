#!/usr/bin/env bash
# Use strict mode. http://redsymbol.net/articles/unofficial-bash-strict-mode/
set -euo pipefail
IFS=$'\n\t'

# Script must work if executed from other directory
DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
cd "$DIR"

VOLUMES_DIR="$DIR/../.volumes"
SENTRY_URL_PREFIX=http://localhost:9000

# WARN Use carefully! Comment out code below to debug from scratch.
#sudo rm -rf $VOLUMES_DIR/acme_sentry_db
#sudo rm -rf $VOLUMES_DIR/acme_sentry_redis

# NOTE docker-compose can not detect and wait for container readiness
#   That's why this script was written instead of compose.yml file

echo Forcibly delete existed containers before applying new configuration
echo The operation is not destructive because containers are ephemeral
docker rm -f acme_sentry_web || true
docker rm -f acme_sentry_worker || true
docker rm -f acme_sentry_redis || true
docker rm -f acme_sentry_db || true

# Pull container in advance, to do 'docker run' from expect with small timeouts
echo INFO pulling containers, this may take some time...
echo postgres:9.6-alpine redis:3.2-alpine slafs/sentry:8.0 | xargs -n 1 docker pull

expect << END
    # https://hub.docker.com/_/postgres/
    spawn docker run \
        -e POSTGRES_USER=sentry \
        -e POSTGRES_PASSWORD=RucLUS8A \
        -p 5432:5432 \
        -v "$VOLUMES_DIR/acme_sentry_db:/var/lib/postgresql/data" \
        --restart=unless-stopped \
        --name acme_sentry_db \
        postgres:9.6-alpine
    expect "database system is ready"
END
echo INFO acme_sentry_db is ready

expect << END
    spawn docker run \
        -v "$VOLUMES_DIR/acme_sentry_redis:/data" \
        --restart=unless-stopped \
        --name acme_sentry_redis \
        redis:3.2-alpine
    expect "server is now ready"
END
echo INFO acme_sentry_redis is ready

expect << END
    spawn docker run \
        -e SENTRY_URL_PREFIX=$SENTRY_URL_PREFIX \
        -e SENTRY_WEB_HOST=0.0.0.0 \
        -e SENTRY_WEB_PORT=9000 \
        -e SECRET_KEY=y3VAda4e \
        -e SENTRY_ADMIN_USERNAME=admin \
        -e SENTRY_ADMIN_PASSWORD=bHxTgk9K \
        -e SENTRY_ADMIN_EMAIL=alexey.diyan@gmail.com \
        -e DATABASE_URL=postgres://sentry:RucLUS8A@acme_sentry_db/sentry \
        -e SENTRY_REDIS_HOST=acme_sentry_redis \
        -e SENTRY_USE_REDIS_BUFFERS=True \
        -e CELERY_ALWAYS_EAGER=False \
        -e SENTRY_INITIAL_TEAM=ACME-Team \
        -e SENTRY_INITIAL_PROJECT=ACME \
        -e SENTRY_INITIAL_KEY=763a78a695424ed687cf8b7dc26d3161:763a78a695424ed687cf8b7dc26d3161 \
        -e SENTRY_INITIAL_PLATFORM=python \
        -p 9000:9000 \
        --link acme_sentry_db \
        --link acme_sentry_redis \
        --restart=unless-stopped \
        --name acme_sentry_web \
        diyan/sentry:8.12.0
    # sentryweb runs db migrations on start which takes time
    set timeout 60
    expect "Running service: 'http'"
END
echo INFO sentryweb is ready

expect << END
    spawn docker run \
        -e SENTRY_URL_PREFIX=$SENTRY_URL_PREFIX \
        -e SECRET_KEY=y3VAda4e \
        -e SENTRY_ADMIN_USERNAME=admin \
        -e SENTRY_ADMIN_PASSWORD=bHxTgk9K \
        -e SENTRY_ADMIN_EMAIL=alexey.diyan@gmail.com \
        -e DATABASE_URL=postgres://sentry:RucLUS8A@acme_sentry_db/sentry \
        -e SENTRY_REDIS_HOST=acme_sentry_redis \
        -e SENTRY_USE_REDIS_BUFFERS=True \
        -e CELERY_ALWAYS_EAGER=False \
        -e C_FORCE_ROOT=True \
        --link acme_sentry_db \
        --link acme_sentry_redis \
        --restart=unless-stopped \
        --name acme_sentry_worker \
        diyan/sentry:8.12.0 \
        run worker --concurrency=4
    expect "ready."
END
# TODO `sentry celery -B is dropped, consider use `celery run cron`
# TODO `sentry start` is deprecated, use `sentry run web`
echo INFO acme_sentry_worker is ready
echo DONE all sentry containers are ready, try to login on http://localhost:9000 with admin / bHxTgk9K

# TODO Consider use SENTRY_EMAIL_BACKEND=django.core.mail.backends.smtp.EmailBackend
# TODO Consider use SENTRY_USE_REMOTE_USER=True
# TODO SENTRY_WEB_HOST=0.0.0.0, SENTRY_WEB_PORT=9000, SENTRY_DOCKER_DO_DB_CHECK=True


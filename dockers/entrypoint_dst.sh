#!/bin/ash
SRC=/home/src/
DST=/home/dst/

_rsync() {
	rsync -az --update --ignore-existing $SRC $DST
}

_rsync_ssh() {
	rsync -az --update --ignore-existing -e "ssh -p $SSH_PORT" $SSH_USER@$SSH_HOST:$SRC $DST
}

if [ -z "$SSH_HOST" ]; then
	# local files syncing
	_rsync
	exit 0
fi

# remote ssh syncing
_rsync_ssh

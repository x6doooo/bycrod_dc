#!/usr/bin/env bash
#ps aux|grep -v grep|grep bin\/bycrod_center_prod|awk '{print $2}'|xargs kill -9
#./bin/bycrod_center_prod --conf=./conf/conf.prod.toml
kill -9 `pgrep -f bin/bycrod_dc`
nohup ./bin/bycrod_dc_prod --conf=./conf/conf.prod.toml $1 >/dev/null 2>&1 &


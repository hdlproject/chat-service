# Ejabberd
This repository is a proof of concept of Ejabberd. The main use case is a chat service.

## How to Run
```shell script
$ docker-compose up
$ docker exec -it ejabberd \
    bin/ejabberdapi register \
    --endpoint=http://127.0.0.1:5280/ \
    --jid=admin@localhost \
    --password=admin

```

TODO:
- Check why the `update_sql_schema` config as well as `ejabberdctl update_sql` command doesn't work
- 

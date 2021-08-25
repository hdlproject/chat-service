# Chat App
## How to Run
```shell script
$ docker-compose up
$ docker exec -it ejabberd \
    bin/ejabberdapi register \
    --endpoint=http://127.0.0.1:5280/ \
    --jid=admin@localhost \
    --password=admin

```
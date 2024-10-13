# Ejabberd
This repository is a proof of concept of Ejabberd. The main use case is a chat service.

## How to Run
```shell script
$ docker compose up
```
## How to Register Admin User
```shell script
$ docker exec -it ejabberd \
    bin/ejabberdctl register admin localhost admin123
```

## TODO:
- Check below error
```text
ejabberd  | 14:01:08.732 [critical] Internal error of module mod_last has occurred during start:
ejabberd  | ** Options: #{cache_size => 1000,cache_missed => true,db_type => sql,
ejabberd  |               cache_life_time => 3600000000,use_cache => true}
ejabberd  | ** exception error: {bad_option,{cache_missed,1000}}
ejabberd  |    in function  ets_cache:'-check_opts/1-fun-0-'/1 (/ejabberd/deps/cache_tab/src/ets_cache.erl, line 566)
ejabberd  |    in call from lists:map_1/2 (lists.erl, line 1564)
ejabberd  |    in call from lists:map/2 (lists.erl, line 1559)
ejabberd  |    in call from ets_cache:new/2 (/ejabberd/deps/cache_tab/src/ets_cache.erl, line 95)
ejabberd  |    in call from mod_last:start/2 (src/mod_last.erl, line 64)
ejabberd  |    in call from gen_mod:start_module/4 (src/gen_mod.erl, line 168)
ejabberd  |    in call from lists:foreach_1/2 (lists.erl, line 1686)
ejabberd  |    in call from gen_mod:start_link/0 (src/gen_mod.erl, line 95)
```

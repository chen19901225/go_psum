# go_psum


## examples

### show with detail

```
name   pid   mem     open_files  net_connections  cmdline
----   ---   ---     ----------  ---------------  -------
nginx  3196  1.02m   27          7                nginx: master process /www/server/nginx/sbin/nginx -c /www/server/nginx/conf/nginx.conf
nginx  3200  19.98m  24          7                nginx: worker process
nginx  3201  13.82m  25          8                nginx: worker process
nginx  3202  22.57m  24          7                nginx: worker process
nginx  3204  20.54m  24          7                nginx: worker process
nginx  3205  1.42m   18          1                nginx: cache manager process
```

### search by one name

```
[vagrant@localhost go_psum]$ sudo ./go_psum --name=redis   --show=0
name   count  mem     open_files  net_connections
----   -----  ---     ----------  ---------------
redis  1      22.66m  40          34
```


### search multi

```
[vagrant@localhost go_psum]$ sudo ./go_psum --name=redis,nginx   --show=0
name   count  mem     open_files  net_connections
----   -----  ---     ----------  ---------------
nginx  6      79.35m  142         37
redis  1      22.68m  40          34
```

### with exclude

```
[vagrant@localhost go_psum]$ sudo ./go_psum --name=redis,nginx --exclude=redis  --show=0
name   count  mem     open_files  net_connections
----   -----  ---     ----------  ---------------
nginx  6      79.35m  142         37
```

### use with watch


```
Every 2.0s: sudo ./go_psum --name=redis,nginx --exclude=redis --show=0                                                                                                                          Sun Jan  3 14:04:25 2021

name   count  mem     open_files  net_connections
----   -----  ---     ----------  ---------------
nginx  6      79.35m  142         37


```
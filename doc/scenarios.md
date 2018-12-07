# Scenarios to support

## Master is started for the first time
### Info needed
| Service  | parameter                   | default value       |
|----------|-----------------------------|---------------------|
| postgres | main port                   | 5432                |
|          | technical port              | 5434                |
|          | data dir                    | /var/lib/postgresql |
|          | root password               | mysecretpassword    |
| repmgr   | master flag                 | true                |
|          | config directory            | /etc/repmgr.d/      |
|          | node name                   | node1               |
|          | list of slave addresses     | <empty>             |
|          | password for main user      | iamrepmgr           |
|          | password for streaming user | iamstreamingrepmgr  |
| barman   | password for main user      | iambarman           |
|          | password for streaming user | iamstreamingbarman  |
|          | backup dir                  | /home/barman/data   |
|          | retention                   | 5 days              |
|          | config directory            | /etc/barman.d/      |
|----------|-----------------------------|---------------------|

### Sequence
 
1. Check flag for master
3. Start PG on technical port with specified directory
3. Add replication/backup properties to postgresql.conf and restart PG
4. Initialize users: repmgr, repmgr\_streaming, barman, barman\_streaming
5. Create DB repmgr
6. Register current node as a master of the cluster with repmgr
7. Create barman config
8. Touch cron for barman
9. Start postgresql on consumer port 5432
10. Start crond

## Slave is started for the first time
### Info needed
| Service  | parameter                   | default value       |
|----------|-----------------------------|---------------------|
| postgres | main port                   | 5432                |
|          | technical port              | 5434                |
|          | data dir                    | /var/lib/postgresql |
|          | root password               | mysecretpassword    |
| repmgr   | master flag                 | true                |
|          | config directory            | /etc/repmgr.d/      |
|          | node name                   | node1               |
|          | list of slave addresses     | <empty>             |
|          | password for main user      | iamrepmgr           |
|          | password for streaming user | iamstreamingrepmgr  |
| barman   | password for main user      | iambarman           |
|          | password for streaming user | iamstreamingbarman  |
|          | backup dir                  | /home/barman/data   |
|          | retention                   | 5 days              |
|          | config directory            | /etc/barman.d/      |
|----------|-----------------------------|---------------------|

1. Check flag for master
2. Wait until consumer port is open on the master
3. Create pg instance as a clone of master
## Master is restarted but still is the master
## Slave is restarted but still is slave
## Automatic failover 
## Restore DB from backup
 

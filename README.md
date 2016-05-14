# kkpm 
[![Build Status](https://travis-ci.org/drkaka/kkpm.svg)](https://travis-ci.org/drkaka/kkpm)
[![Coverage Status](https://codecov.io/github/drkaka/kkpm/coverage.svg?branch=master)](https://codecov.io/github/drkaka/kkpm?branch=master)  

The private message module for golang project. (***IMPORTANT:*** The API and data structure may change!)

## Database
It is using PostgreSQL as the database and will create a table:

```sql  
CREATE TABLE IF NOT EXISTS private_msg (
    id uuid primary key,
    from_userid integer,
    to_userid integer,
    message text,
    at integer
);
CREATE INDEX IF NOT EXISTS index_private_msg_to_userid ON private_msg (to_userid);
CREATE INDEX IF NOT EXISTS index_private_msg_from_userid ON private_msg (from_userid);
CREATE INDEX IF NOT EXISTS index_private_msg_at ON private_msg (at);
```

## Dependence

```Go
go get github.com/jackc/pgx
go get github.com/satori/go.uuid
```

## Usage 

####First need to use the module with the pgx pool passed in:
```Go
err := kkpm.Use(pool)
```

####Get the unread messages count:
```Go
count, err := GetUnreadCount(2);
```

####Mark all messages from the user id as read:
```Go
err := ReadFrom(2, 3);
```

####Get the messages sent:
```Go
result, err := kkpm.GetSentMessages(3, 0);
```
The second parameter is unixtime, the messages sent later than that will be got.

####Get the messages received:
```Go
result, err := kkpm.GetReveivedMessages(3, 0);
```
The second parameter is unixtime, the messages received later than that will be got.

####Get the messages between two people:
```Go
result, err := kkpm.GetPeerChat(3, 2, 0);
```
The third parameter is unixtime, the messages later than that will be got.

## TODO

#### Delete

#### Manage

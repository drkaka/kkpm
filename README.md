# kkpm 
[![Build Status](https://travis-ci.org/drkaka/kkpm.svg)](https://travis-ci.org/drkaka/kkpm)
[![Coverage Status](https://codecov.io/github/drkaka/kkpm/coverage.svg?branch=master)](https://codecov.io/github/drkaka/kkpm?branch=master)  

The private message module for golang project. (***IMPORTANT:*** The API and data structure may change!)

Current database created by:

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
```
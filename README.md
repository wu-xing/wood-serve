# Wood 
> (backend repository)
a notes/wiki write system.


## configure
Wood configure used [Viper](https://github.com/spf13/viper).

All configure options here. The prefie is `WOOD_`

| name            | Description                 | Default     |
| -----------     | -----------                 | ---------   |
| IMAGE_DIR       | store uploaded images place | `./upload`  |
| DISABLE_SIGN_UP  | disable sign up flag        | true        |
| DB_HOST         | database host               | `localhost` |

## development

### dependency
- go get github.com/lib/pq


## deploy 
### log

## Know Issue

- `pq: function uuid_generate_v4() does not exist` 

you need grand pg user as SUPERUSER, or use SUPERUSER execute sql `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`


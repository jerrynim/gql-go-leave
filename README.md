to merge '.graphql' in '/schema'

> $ gqlmerge --indent=2s schema schema.graphql

to generate gqlgen with 'schema.graphql'

> $ gqlgen generate

to start & reloading server

> nodemon --exec go run server.go --signal SIGTERM

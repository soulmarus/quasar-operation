# quasar-operation

curl -x POST http://localhost:8080/topsecret

curl -X POST "http://localhost:8080/topsecret" -H "accept: */*" -H "Content-Type: application/json" -d "{ "satellites":[ { "name":"kenobi", "distance":"100.0", "message":[ "este", "", "", "mensaje", "" ] }, { "name":"skywalker", "distance":"115.5", "message":[ "", "es", "", "", "secreto" ] }, { "name":"sato", "distance":"142.7", "message":[ "este", "", "un", "", "" ] } ] }"


curl -X POST "http://localhost:8080/topsecret" -H "accept: */*" -H "Content-Type: application/json" -d "{ "satellites":[ { "name":"kenobi", "distance":"100.0", "message":[ "este", "", "", "mensaje", "" ] }, { "name":"skywalker", "distance":"115.5", "message":[ "", "es", "", "", "secreto" ] }, { "name":"sato", "distance":"142.7", "message":[ "este", "", "un", "", "" ] } ] }"

{ "satellites":[ { "name":"kenobi", "distance":"100.0", "message":[ "este", "", "", "mensaje", "" ] }, { "name":"skywalker", "distance":"115.5", "message":[ "", "es", "", "", "secreto" ] }, { "name":"sato", "distance":"142.7", "message":[ "este", "", "un", "", "" ] } ] }

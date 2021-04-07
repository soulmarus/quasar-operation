# Operação Quasar

Esse projeto consiste na resolução da operação Quasar.

O projeto está estruturado utilizando uma estrutura de DDD e Arquitetura Hexagonal, segue uma explicação dos pacotes:

    cmd
        topsecret-server -> contem o entry point da aplicação
    pkg
        http
            rest -> adaptar para entrada rest
        storage
            memory -> adaptador para operações
        tracking -> camada de algoritmos e models

Para subir o projeto rodar:
    go run cmd/topsecret-server/main.go 

Para executar os testes unitarios rodar:
    go test ./... -v

Testes funcionais estão na collection no diretorio postman e podem ser executados apontando para a URL do serviço.
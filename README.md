# microservice: Golang

  - go get -u gopkg.in/mail.v2
  - go get -u github.com/confluentinc/confluent-kafka-go/kafka

## kafka:
  
  - docker-compose exec kafka bash
  - kafka-console-producer --bootstrap-server=localhost:9092 --topic=emails
  - {"emails": ["rapha.pse@gmail.com"], "subject": "Teste micro service", "body": "Acesse o link" }

``` consumer -> chan message -> send ```


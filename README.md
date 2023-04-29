# Go gotenberg sandbox

(Example pdf file)[]

## API

`create-action` service send the action to create PDF. It can be not only http, but amqp protocol by RabbitMQ. It forse the gotenberg service to make PDF from given url wich served by `url-generator`. Note: `url-generator` create static html with images wich served by nginx service.


## TODO
- [ ] - datasource wich can generate some data for charts
- [ ] - image generator wich will save image in folder wich will served by `nginx` service.
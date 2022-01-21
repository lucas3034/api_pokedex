What is Rest ? -
Api rest is an API architecture that uses the http protocol(from version 1.1) for communication, they are more standardized and advanced than
the others, it is widely used for application development.
.
why do we use REST ? -
REST itself are used to transfer information, returning the state of the information in question at the time of the request.
Rest also allows you to do actions such as inserting or deleting a record.
It is used as a standard for exchanging information, and it is also widely used for application development,
such as: web services and APIs.
Using it, facilitates the design and use by other devs, since the pattern is the same.
.
about HTTP verbs -
are responsible for communicating the rest.
GET - is the most used and is responsible for the query.
PUT- is used to update the information
DELETE - is used to remove information
POST- it is used to send information
PATCH - is used to apply partial modifications
OPTIONS - is used to describe communication options
.
What is the REST maturity model ? -
This model is used to measure the efficiency of an API, based on its use of URI, HTTP methods and HATEOAS.
The more an API uses these factors, the more "mature" it is. This model divides the APIs into 4 parts:

lvl 0 services - Does not use URI, HTTP Methods and HATEOAS resources.
It usually only uses a URI and a single HTTP method (usually POST or GET)

lvl 1 services - Efficiently uses URIs, and has resources mapped, however, despite using verbs
better than at lvl 0, still don't use them efficiently. (most APIs fit this lvl)

lvl 2 services - At this level the use of URIs and HTTP verbs is efficient, at this level the APIs start to get sophisticated
and more efficient.

lvl 3 services - At this level, the three factors (URIs, HTTP and HATEOAS) are used efficiently, this is the level
maximum and most mature an API can reach, the goal of every API is to reach that level of maturity.
.
What is HATEOAS ? -
Hateoas is a way to allow an API to perform navigation between its resources through links and URLs, even without
have the dense knowledge of these URLs and links.
Using HATEOAS, the client discovers the URLs as it navigates, without needing to know them previously.
its functioning very much resembles a hypertext.
.
.
.
.-----------.
.
.
.
O que é o Rest ? - 
Uma api rest é uma arquitetura de API q usa o protocolo http(a partir da versão 1.1) para a comunicação, elas são mais padronizadas e e avançadas do que
as demais, ele é mt utilizado para desenvolvimento de aplicações.
.
por que usamos o REST ? -
REST em si são usadas pata transferir informações, retornando o estado das informa em questão, no momento da requisição.
O Rest, tbm permite que faça ações como inserir ou excluir um registro.
Ele é usado como um padrão para a troca das informações, e ele também é muito usado para o desenvolvimento de aplicações,
como: web services e APIs.
Utiliza-lo, facilita o desenpenho e a utilização por outros devs, já que o padrão é o mesmo.
.
sobre os verbos HTTP -
são os responsáveis pela comunicação da rest.
GET - é o mais usado e é responsável pela consulta.
PUT- é usado para atualizar as informações
DELETE - é utilziado para remover informações
POST- ele é usado para enviar informações
PATCH - é utilizado para aplicar modificações parciais
OPTIONS - é usado para descrever as opções de comunicação
.
O que é o modelo de maturidade REST ? -
Esse modelo é utilizado para medir a eficiência de uma API, com base no seu uso de URI, métodos HTTP eo HATEOAS.
Qnt mais um API usa esses fatores, mais "maduro" ele é. Esse modelo divide as APIs em 4 partes:

lvl 0 services - Não usa recursos de URI, HTTP Methods e HATEOAS.
Geralmente só usa um URI e um único method HTTP(geralmente o POST ou o GET)

lvl 1 services - Utiliza de maneira eficiente as URIs, e tem os recursos mapeados, porém, apesar de utilizar os verbos
melhor do que n o lvl 0, ainda não os utilizam de maneira eficiente. (a maioria das APIs se encaixam nesse lvl)

lvl 2 services - Nesse nivel, o uso de URIs e verbos HTTP é eficiente, nesse nivel as APIs começam a ficar sofisticadas
e mais eficientes.

lvl 3 services - Nesse nível, os três fatores (URIs, HTTP e HATEOAS) são utilizados de forma eficiente, esse é o nivel
máximo e mais maduro que uma API pode chegar, o objetivo de toda API é chegar nesse nivel de maturidade.
.
O que é HATEOAS ? -
Hateoas é uma maneira de de permitir q uma API realize uma navegação entre seus recursos pelos links e URLs, mesmo sem
ter o conhecimento denso dessas URLs e links.
Utilizado o HATEOAS, o cliente vai descobrindo as URLs, a medida que vai anvegando, sem precisar conhece-las previamente.
o seu funcionamente lembra muito um hypertext.


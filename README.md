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
-What is Rest ? -
Api rest is an API architecture that uses the http protocol(from version 1.1) for communication, they are more standardized and advanced than
the others, it is widely used for application development.

.


-Why do we use REST ? -
REST itself are used to transfer information, returning the state of the information in question at the time of the request.
Rest also allows you to do actions such as inserting or deleting a record.
It is used as a standard for exchanging information, and it is also widely used for application development,
such as: web services and APIs.
Using it, facilitates the design and use by other devs, since the pattern is the same.

.


-About HTTP verbs -
are responsible for communicating the rest.
GET - is the most used and is responsible for the query.
PUT- is used to update the information
DELETE - is used to remove information
POST- it is used to send information
PATCH - is used to apply partial modifications
OPTIONS - is used to describe communication options

.


-What is the REST maturity model ? -
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


-What is HATEOAS ? -
Hateoas(Hypermedia As the Engine Of Application State) is a way to allow an API to perform navigation between its resources through links and URLs, even without
have the dense knowledge of these URLs and links.
Using HATEOAS, the client discovers the URLs as it navigates, without needing to know them previously.
its functioning very much resembles a hypertext.

.


-What is the advantage of having HATEOAS?
Hateoas is useful in helping customers to use the service without having to know the API well. It is a facilitating tool. making using the API simpler and more popular.

.


.-----------.

.


-O que ?? o Rest ? - 
Uma api rest ?? uma arquitetura de API q usa o protocolo http(a partir da vers??o 1.1) para a comunica????o, elas s??o mais padronizadas e e avan??adas do que
as demais, ele ?? mt utilizado para desenvolvimento de aplica????es.

.


-Por que usamos o REST ? -
REST em si s??o usadas pata transferir informa????es, retornando o estado das informa em quest??o, no momento da requisi????o.
O Rest, tbm permite que fa??a a????es como inserir ou excluir um registro.
Ele ?? usado como um padr??o para a troca das informa????es, e ele tamb??m ?? muito usado para o desenvolvimento de aplica????es,
como: web services e APIs.
Utiliza-lo, facilita o desenpenho e a utiliza????o por outros devs, j?? que o padr??o ?? o mesmo.

.


-Sobre os verbos HTTP -
s??o os respons??veis pela comunica????o da rest.
GET - ?? o mais usado e ?? respons??vel pela consulta.
PUT- ?? usado para atualizar as informa????es
DELETE - ?? utilziado para remover informa????es
POST- ele ?? usado para enviar informa????es
PATCH - ?? utilizado para aplicar modifica????es parciais
OPTIONS - ?? usado para descrever as op????es de comunica????o

.


-O que ?? o modelo de maturidade REST ? -
Esse modelo ?? utilizado para medir a efici??ncia de uma API, com base no seu uso de URI, m??todos HTTP eo HATEOAS.
Qnt mais um API usa esses fatores, mais "maduro" ele ??. Esse modelo divide as APIs em 4 partes:

lvl 0 services - N??o usa recursos de URI, HTTP Methods e HATEOAS.
Geralmente s?? usa um URI e um ??nico method HTTP(geralmente o POST ou o GET)

lvl 1 services - Utiliza de maneira eficiente as URIs, e tem os recursos mapeados, por??m, apesar de utilizar os verbos
melhor do que n o lvl 0, ainda n??o os utilizam de maneira eficiente. (a maioria das APIs se encaixam nesse lvl)

lvl 2 services - Nesse nivel, o uso de URIs e verbos HTTP ?? eficiente, nesse nivel as APIs come??am a ficar sofisticadas
e mais eficientes.

lvl 3 services - Nesse n??vel, os tr??s fatores (URIs, HTTP e HATEOAS) s??o utilizados de forma eficiente, esse ?? o nivel
m??ximo e mais maduro que uma API pode chegar, o objetivo de toda API ?? chegar nesse nivel de maturidade.

.

-O que ?? HATEOAS ? -
Hateoas(Hypermedia As the Engine Of Application State) ?? uma maneira de de permitir q uma API realize uma navega????o entre seus recursos pelos links e URLs, mesmo sem
ter o conhecimento denso dessas URLs e links.
Utilizado o HATEOAS, o cliente vai descobrindo as URLs, a medida que vai anvegando, sem precisar conhece-las previamente.
o seu funcionamente lembra muito um hypertext.

.


-Qual ?? a vantagem de ter HATEOAS ?
O Hateoas ?? ??til para ajudar os clientes a utilizarem o servi??o sem precisar conhecer bem a API. Ela ?? uma ferramenta facilitadora. tornando o uso da API algo mais simples e popular. 
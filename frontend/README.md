# frontend challenge

<img src="https://raw.githubusercontent.com/hubbeco/challenge/master/frontend/assets/images/Lobster-Telephone-Edward-James-And-Salvador-Dali-Collaboration.jpg" width="188">

Desafios de código (challenges) tem como objetivo de testar suas habilidades na resolução de problemas do mundo real.
Em frontend-challenge você poderá ajudar milhares de empresas a crescer auxiliando seus vendedores em uma tarefa muito 
comum e repetitiva: encontrar o endereço e telefone de contato de potenciais clientes.

## Apresentação do problema

No desafio, você irá ajudar a desenvolver parte de um sistema para o cliente fictício. Chamado **Publicante**, nosso cliente
é um portal de notícias muito conhecido e além de entregar notícias para toda sua audiência digital, vende espaço
publicitário (anúncios) para que outras empresas possam divulgar sua marca. 

Os vendedores da Publicante gastam uma grande parte do seu dia procurando informações de contato sobre outras empresas.
Essas informações são utilizadas para identificar possíveis clientes e agendar visitas. 

Os melhores vendedores conseguem conversar 200 novas empresas todos os meses e para isso recebem listas de
empresas que costumam comprar anúncios em todo o Brasil. As listas apresentam o nome, email ou site das empresas,
mas não informam o endereço ou o telefone de contato.

Nesse desafio, sua missão é ajudar essas empresas e diminuir o tempo gasto reunindo informações de seus clientes.

## Como fazer

Você irá construir um sistema que captura informações de endereço e telefone na internet. O vendedores informará o site
da empresa alvo e o sistema deverá retornar telefone e endereço, se houver.

Dividir e conquistar!

Conheça cada etapa abaixo, elabore uma estratégia para resolvê-las e implemente.

1) O usuário do sistema informa o site do potencial cliente, onde a busca será realizada
2) O sistema consulta um serviço externo que retorna o conteúdo do site
3) O sistema varre o html/conteúdo em busca de telefones e endereços
4) O sistema armazena os dados encontrados no frontend
5) O sistema retorna os dados de telefone e endereço encontrados, ordenados por critério de relevância

#### Importante:
 
Você não precisa desenvolver um servidor/backend. Para te ajudar nesse desafio, disponibilizamos um serviço web (bot)
que vai até o site informado e retorna o HTML bruto.

* URL do serviço: http://htmler.sandbox.hubbe.co/
* Exemplo de uso: http://htmler.sandbox.hubbe.co/crawl?url=https://hubbe.co

```javascript
const targetUrl = 'https://hubbe.co';
const serviceUrl = 'http://htmler.sandbox.hubbe.co/crawl?url='

fetch(serviceUrl + targetUrl)
```

#### Design:
 
Incluimos algumas referências de Design na pasta [assets/layout](assets/layout).
Para tipografia, siga as instruções abaixo:

```html
<link href="https://fonts.googleapis.com/css?family=Source+Code+Pro|Source+Sans+Pro&display=swap" rel="stylesheet">
...
<style>
    body, html {
        font-family: 'Source Sans Pro', sans-serif;
    }
    .hint-font {
        font-family: 'Source Code Pro', monospace;
    }
</style>
```

## Resultado (o que será avaliado)

* Faça o máximo que conseguir dentro do seu nível de conhecimento
* Utilize bons padrões (patterns) da linguagem
* Pode fazer uso de frameworks (React, Vue, etc), mas talvez não precise
* Utilize o Github para compartilhar seu código
  - Comece fazendo um Fork deste projeto
  - Implemente a solução localmente
  - Ao concluir, submeta um Pull Request com o código final 

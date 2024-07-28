# data-science challenge

<img src="https://raw.githubusercontent.com/hubbeco/challenge/master/data-science/assets/images/under-construction-669123b5e6c3d0c7.png" width="188">

### TODO



### Descrição do desafio

> **_NOTE:_** Este é um desafio aberto, não há uma solução predefinida, nem objetivos específicos.
> Então, divirta-se 😃

Disponilizamos 2 datasets, cada um deles em 2 formatos. Eles variam em grau de dificuldade e em tipos de habilidades testadas.

* Todos os datasets estão disponíveis em AWS S3 e em Google Drive.
* Todos os datasets são de domínio público, porem representam atividades que você vai encontrar usualmente na Zedia e nas empresas do grupo.
* Todos os datasets permitem diferentes análises; o seu senso critico de decidir o que analisar faz parte do desafio.
* Todos os datasets podems ser analisados sem a necessidade de hardware ou software proprietário ou pago.

Escolha apenas 1 dentre os desafios - aquele que você julgar que se enquadra melhor no seu perfil
(seus conhecimentos, seus interesses, e aquilo que você quer demonstrar para nós que você domina).

Analize o dataset com as ferramentas que você preferir: python, matlab, serviços cloud, softwares low-code, qualquer ferramenta que você julgar adequada.
Ao final do desafio, você deve nos enviar o link para um repositório (github, gitlab ou similar) contendo a implementação de sua análise
(qualquer código, script, ou similar utilizado para realizar a anĺise) e um arquivo pdf chamado "relatorio.pdf".
O relatório deve ter entre 1 e 5 páginas, em qualquer formatação que você preferir, contendo qualquer quantidade de imagens/diagramas/tabelas/etc que você preferir.

Deve conter ao menos as seguintes seções:

* Sumário: um ou dois paragrafos descrevendo o que foi analisado e quais resultados foram obtidos. Seja breve, não inclua imagens nem referências.
* Implementação: dê uma visão geral de como você implementou, ferramentas utilizadas, etc. Não é preciso ser uma extensa documentação, 
nos vamos olhar o seu código de qualquer maneira. Código limpo e comentado vale mais do que mil páginas de documentação 😁
* Resultados: o que você descobriu sobre os dados, métricas, etc. (por exemplo, acuráciacaso se você treinou algum modelo).

#### Foco no problema/solução

Não há necessidade de capa, linguajar super rebustado, ou qualquer outra formalidade.
Mantenha simples e direto.

Não há necessidade de "atirar para todos os lados" para demonstrar o seu conhecimento, seja focado e claro:
defina aquilo que você quer fazer/analisar, implemente o que for necessário, apresente os seus resultados. Ponto final.

Idealmente, você deve definir seus objetivos na forma de uma única pergunta, seguir um processo metodológico
para responder/resolver aquilo que você se propôs, e no final obter resultados que respondem a pergunta.

Você é livre para incluir qualquer trecho de código não produzido por você
(incluindo stack overflow, chatgpt, tutoriais, artigos científicos, repositórios abertos, etc),
porém referencie. Inclua links do stack overflow no código sempre que pegar um trecho de uma solução de terceiros.
Mencione tutoriais sempre que tomar forte inspiração deles.

Com tudo isso em mente. Escolha abaixo um dos desafios e boa sorte!
----------------------

### Desafio 1

Principais habilidades: processamento de dados tabulares, séries temporais, telemetria.
Esse dataset é uma tabela contendo registros de acessos a uma aplicação web. Cada linha
representa um cliente fazendo uma requisição para um backend, indicando que o usuário ainda
está conectado à aplicação ou indicando que o usuário realizou alguma atividade. O dataset
possui as seguintes colunas:

* timestamp (unix timestamp de quando a requisição foi recebida)
* ...
* ...

O dataset está disponível em dois formatos:

a) **Delta.** O dataset possui X linhas. Link do S3: https://abc. Link do google drive:
> https://abc.

b) **CSV.** O dataset foi reduzido para X linhas. Link do S3: https://abc. Link do google drive:
> https://abc.

Sugestões do que pode ser analisado:

* Treinar um modelo de classificação para a coluna X
* Agregar os dados realizando o count de acessos vs cidade, e treinar um modelo de regressão para o count
* Time series forecasting sobre o count ao longo do tempo
* Enriquescer os dados com outras fontes públicas (por exemplo, IBGE)

### Desafio 2

Principais habilidades: manipulação de media, visão computacional, processamento de linguagem
natural, processamento de audio.
Este dataset é sobre um episódio do cartoon "Popeye". O episódio foi dividido em trechos de 10
segundos. O idioma é inglês.


O dataset está disponível em dois formatos:

a) **Delta.** Cada linha corresponde a um trecho de 10s. As seguintes colunas estão disponíveis:
* video (binário, contendo o conteúdo do arquivo mp4 desse trecho de video)
* time_start (float, tempo em segundos do começo do trecho)
* time_end (float)
* transcription (string, contendo o texto falado no trecho).
> Link do S3: https://abc.
> Link do google drive: https://abc.

b) **Arquivos mp4.** Diversos arquivos mp4 estão em uma pasta, cada um correspondente um
trecho de 10s.
> Link do S3: https://abc.
> Link do google drive: https://abc.

Sugestões do que pode ser analisado:

* Realizar a transcripção novamente (melhorada de alguma forma)
* Analisar objetos presentes no episódio
* Agrupar os trechos por tópicos
* Visualizar uma núvem de palavras


#### Boa sorte!

# data-science challenge


![date-everywhere-data](https://github.com/user-attachments/assets/cbad2a94-ad8c-442a-a82b-18d843046812)

> **_NOTE:_** Este é um desafio aberto, não há uma solução predefinida, nem objetivos específicos.
> Então, divirta-se 😃

Disponibilizamos 2 datasets, cada um deles em 2 formatos. Eles variam em grau de dificuldade e em tipos de habilidades testadas.
Características comuns a todos os datasets:
* Estão disponíveis em AWS S3 e em Google Drive.
* São de domínio público, porem representam atividades que você vai encontrar usualmente na Zedia e nas empresas do grupo.
* Permitem diferentes análises; o seu senso critico de decidir o que analisar faz parte do desafio.
* Podem ser analisados sem a necessidade de hardware ou software proprietário ou pago.

**Escolha apenas 1 dentre os desafios**, aquele que você julgar que se enquadra melhor no seu perfil (seus conhecimentos, seus interesses, e aquilo que você quer demonstrar para nós que você domina).
Analize o dataset com as ferramentas que você preferir: python, matlab, serviços cloud, softwares low-code, qualquer ferramenta que você julgar adequada.

Você é livre para incluir qualquer trecho de código não produzido por você (incluindo stack overflow, chatgpt, tutoriais, artigos científicos, repositórios abertos, etc), porém **referencie**. Inclua links do stack overflow no código sempre que pegar um trecho de uma solução de terceiros.
Mencione tutoriais sempre que tomar forte inspiração deles.

## Entrega dos resultados
Ao final do desafio, você deve enviar um email para [vagas@zedia.com.br](mailto:vagas@zedia.com.br) com o título "<nome do candidato> - Entrega do desafio de Data Science" contendo o link para um repositório (github, gitlab ou similar). O repositório deve conter a implementação de sua análise (qualquer código, script, ou similar utilizado) e um arquivo pdf chamado "relatorio.pdf".

O **relatório** deve ter entre 1 e 5 páginas, em qualquer formatação que você preferir, contendo qualquer quantidade de imagens/diagramas/tabelas/etc que você preferir. Deve conter ao menos as seguintes seções:

* **Sumário:** um ou dois paragrafos descrevendo o que foi analisado e quais resultados foram obtidos. Seja breve, não inclua imagens nem referências.
* **Implementação:** dê uma visão geral de como você implementou, ferramentas utilizadas, etc. Não é preciso ser uma extensa documentação, 
nos vamos olhar o seu código de qualquer maneira. Código limpo e comentado vale mais do que mil páginas de documentação 😁
* **Resultados:** o que você descobriu sobre os dados, métricas, etc. (por exemplo, acuráciacaso se você treinou algum modelo).

## Foco no problema/solução

Não há necessidade de capa, linguajar super rebustado, ou qualquer outra formalidade.
Mantenha simples e direto.

Não há necessidade de "atirar para todos os lados" para demonstrar o seu conhecimento, seja focado e claro:
defina aquilo que você quer fazer/analisar, implemente o que for necessário, apresente os seus resultados. Ponto final.

Tente definir seus objetivos na forma de uma única pergunta, seguir um processo metodológico para responder/resolver aquilo que você se propôs, e no final obter resultados que respondem a pergunta.

## Desafio 1

Principais habilidades: processamento de dados tabulares, séries temporais, telemetria.

Esse dataset é uma tabela contendo registros de **acessos a uma aplicação web**. Cada linha representa um cliente fazendo uma requisição para um backend, indicando que o usuário ainda está conectado à aplicação ou indicando que o usuário realizou alguma atividade. Você perceberá que os acessor estão distribuidos em "ativações": durante alguns minutos há muitos acessos, e depois não há nenhum acesso.

O dataset possui as seguintes colunas: `created_at` (unix timestamp de quando a requisição foi recebida), `publisher` (string, uuid identificando o local onde o usuário estava navegando), `event_type` (string, identifica o tipo requisição), `ip` (string, semi anonimizado), `user_id` (string, uuid identificando o usuário), `geo_lat` (double, latitude onde o usuário estava fisicamente), `geo_lon` (double, longiture onde o usuário estava fisicamente), `geo_country` (string, país da localização latitude-longitude), `geo_region` (string, estado da localização latitude-longitude), `geo_city` (string, cidade da localização latitude-longitude).


O dataset está disponível em dois formatos:

1. **Delta:** o dataset possui 1047577 linhas. [Link do S3](https://abc.TODO). [Link do Google Drive](https://drive.google.com/file/d/1gNgnI_lLqfSrNmrH6ag8d9Bxd6zl-6hR/view?usp=sharing).

2. **CSV:** o dataset foi reduzido para 100000 linhas. [Link do S3](https://abc.TODO). [Link do Google Drive](https://drive.google.com/file/d/193NSK-3UzUpDpfCLRT6Yhup84oB3Qgdk/view?usp=sharing).

Sugestões do que pode ser analisado:
* Treinar um modelo de classificação para a coluna event_type.
* Agregar os dados realizando a contagem de acessos por cidade, e treinar um modelo de regressão para a contagem.
* Time series forecasting sobre a contagem ao longo do tempo (de toda a série temporal ou apenas dentro de uma ativação).
* Enriquescer os dados com outras fontes públicas (por exemplo, IBGE).
* Comparação de estatisticas entre cidades, entre ativações, etc.

## Desafio 2

Principais habilidades: manipulação de media, visão computacional, processamento de linguagem natural, processamento de audio.

Este dataset é sobre um **episódio do cartoon "Popeye"**. O episódio foi dividido em trechos de 10 segundos. O idioma é inglês.

O dataset está disponível em dois formatos:
1. **Delta:** cada linha corresponde a um trecho de 10s. As seguintes colunas estão disponíveis: `chunk` (binário, contendo o conteúdo do arquivo mp4 desse trecho de video), `time_start` (int, tempo em segundos do começo do trecho), `time_end` (int), `transcript` (string, contendo o texto falado no trecho). [Link do S3](https://abc.TODO). [Link do Google Drive](https://drive.google.com/file/d/1qTu3LYnGcS88SWd6QZWYr0jzCK7rbTO4/view?usp=sharing).
2. **Arquivos mp4:** diversos arquivos mp4 estão em uma pasta, cada um correspondente um trecho de 10s. [Link do S3](https://abc.TODO). [Link do Google Drive](https://drive.google.com/file/d/11b1ARDT1787MacXdAHEKGK0M99wE9Vie/view?usp=sharing).

Sugestões do que pode ser analisado:
* Realizar a transcripção novamente (melhorada de alguma forma).
* Analisar objetos presentes no episódio.
* Agrupar os trechos por tópicos.
* Visualizar uma núvem de palavras.

## **Com tudo isso em mente, escolha um dos desafios e boa sorte!** 🙋‍♂️

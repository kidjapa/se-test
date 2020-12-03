# Serasa - Teste para analista desenvolvedor
Olá, obrigado pelo interesse em fazer parte da nossa equipe.  
O objetivo deste teste é verificar (até certo ponto) suas habilidades de codificação e arquitetura. Para isso você receberá um problema simples onde poderá mostrar suas técnicas de desenvolvimento.
Nós encorajamos você a exagerar um pouco na solução para mostrar do que você é capaz.
Considere um cenário em que você esteja construindo uma aplicação pronta para produção, onde outros desenvolvedores precisarão trabalhar e manter essa aplicação ao longo do tempo.  
Você **PODE** e **DEVE** usar bibliotecas de terceiros, usando ou não um framework, você decide. Lembre-se, um desenvolvedor eficaz sabe o que construir e o que reutilizar.
Na entrevista de "code review", esteja preparado para responder algumas perguntas sobre arquitetura, bibliotecas e, caso utilize, sobre o framework. Como e por que você as escolheu e com quais outras alternativas você está familiarizado, serão algumas dessas perguntas.
Como este é um processo de "code review", evite adicionar código gerado ao projeto.
***Obs***: Para realizar esse teste, não crie um repositório público! Esse desafio é compartilhado apenas com pessoas que estamos entrevistando e gostaríamos que permanecesse assim.  
Aqui no Serasa, nós utilizamos o [Docker](https://www.docker.com/products/docker) para executar as aplicações, por isso, pedimos que você faça o mesmo neste teste. Isso garante que tenhamos um resultado idêntico ao seu quando testarmos sua aplicação.

## Descrição do problema:
A empresa XPTO tem uma aplicação legada (mainframe) de negativações que não está suportando a demanda atual, esta aplicação não pode ser alterada, portanto, é necessário construir um serviço que atenda a demanda, como Façade.
Esta aplicação deve consumir os dados do legado e persistir numa base intermediária, devendo também expor uma API para acesso ao dados das negativações recebendo como parametro o CPF do cliente. 

## Instruções
- Para simular a aplicação legada a ser consumida, você deverá subir um servidor com o arquivo negativacoes.json (sugestão: json-server)
- A API **DEVE**:
  - possuir no minimo dois endpoints: um para consulta por cpf das negativações e outro para executar a integração
  - seguir os padrões REST
  - ser autenticada
- Banco de dados ao seu critério
- Camada de cache é opcional
- Testes unitários
- Documentação de setup e do funcionamento da API (um Makefile cai muito bem!)

## Desenvolvendo
- crie um repositório no seu github (ou sistema de sua preferencia) este repositório
- Crie uma nova branch chamada `develop`
- Desenvolva a aplicação
- Crie uma "pull request" da branch `develop` para a "branch" `master`. Essa PR deve conter as instruções para executarmos as suas aplicações, as tecnologias que você decidiu usar, por que decidiu utilizá-las e também as decisões que você teve quanto ao design do seu código

## Critérios de avaliação
Dê uma atenção especial aos seguintes aspectos:
- Você **DEVE** usar bibliotecas de terceiros, e pode escolher usar um framework, utilizar não vai ser uma penalidade, mas você vai precisar justificar a sua escolha.
- Sua aplicação **DEVE** executar em container Docker.
- Sua aplicação **DEVE** retornar um JSON válido e **DEVEM** conter os recursos citados anteriormente.
- Você **DEVE** escrever um código testável e demonstrar isso escrevendo testes unitários.
- Você **DEVE** prestar atenção nas melhores práticas para segurança de APIs.
- Você **NÃO** precisa desenvolver um "frontend" (telas) para esse teste.

## Pontos que consideramos um bônus:
- Fazer uso de uma criptografia reversível de dados sensíveis do usuário antes de persisti-los no banco de dados
- Suas respostas durante o code review
- Sua descrição do que foi feito na sua "pull request"
- Setup da aplicação em apenas um comando ou um script que facilite esse setup
- Outros tipos de testes, como: testes funcionais e de integração
- Histórico do seus commits, com mensagens descritivas do que está sendo desenvolvido.
---

Boa sorte!
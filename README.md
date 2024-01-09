# Desafio time Metaplane 

<div align="center">
  <img src="img/lojas-americanas.png" alt="Lojas Americanas" width="40%">
</div>

---

## :handshake: Realizado por:

<img width="100" alt="Foto de Perfil da Dayane" src="img/perfil_dayane.jpg"> <br>  `Dayane Cristina Santos` <br> Santos - SP <br><br> [![Github Badge](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/dayane-cristina-santos/)

# Briefing

## Objetivo
Este projeto consiste em mesclar duas listas ordenadas de números e também oferecer uma API simples para salvar essas listas e realizar a operação de mesclagem.

<br>

## 🔍 Como Rodar

1. Certifique-se de ter instalado em sua máquina o `git`
  
1. Clone o repositório

    `git clone https://github.com/DayaneCristina/merge-lists`

2. Navegue até o diretório do projeto

    `cd merge-lists`

3. Caso tenha o `docker` / `docker compose` instalados na sua máquina, execute o comando abaixo para subir a API: 
  
    `docker-compose up --build`

    Se não possuir o `docker`, certifique-se de ter o `go` instalado e rode o comando abaixo:

    `go run main.go `

## API
#### [POST] /saveLists
Endpoint para salvar duas listas ordenadas. Deve receber em seu `body` os dois campos `list1` e `list2` contendo um *array* de números inteiros.

| cURL
```
curl --location 'localhost:8080/saveLists' \
--header 'Content-Type: application/json' \
--data '{
    "list1": [1,3,2,4],
    "list2": [2,5,6,7]
}'
```

#### [GET] /merge
Endpoint para mesclar as listas previamente salvas.

| cURL
```
curl --location 'localhost:8080/merge'
```

### Testes unitários:
Os testes podem ser encontrados no diretório *tests* e podem ser executados usando o comando abaixo após subir a aplicação via `docker` (como demonstrado no passo 4.)

  `docker exec -it americanas.teste.api go test ./tests`

Ou, caso não possua o docker, rodar (já com o `go` devidamente instalado):

`go test ./tests`

## **Tecnologias Utilizadas:**

<div style="display: inline_block">
  <img align="center" alt="icone-GIT" height="60" src="https://github.com/devicons/devicon/blob/master/icons/git/git-original.svg">
  &nbsp;&nbsp;
  <img align="center" alt="icone-Go" height="60" src="https://github.com/devicons/devicon/blob/master/icons/github/github-original.svg">
  &nbsp;&nbsp;
  <img align="center" alt="icone-VS-CODE" height="60" src="https://github.com/devicons/devicon/blob/master/icons/vscode/vscode-original.svg">
  &nbsp;&nbsp;
  <img align="center" alt="icone-Go" height="60" src="https://github.com/devicons/devicon/blob/master/icons/go/go-original.svg">
  &nbsp;&nbsp;
  <img align="center" alt="icone-Docker" height="60" src="https://github.com/devicons/devicon/blob/master/icons/docker/docker-original.svg">
  &nbsp;&nbsp;
</div>

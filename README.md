# estudo-golang

códigos referente ao estudo de golang


#obtendo pacotes via git

go get -u url_do_projeto_git


#relatorios de teste

go test -coverprofile=arquivo.out

go tool cover -func=arquivo.out

go tool cover -html=arquivo.out

library(RCurl)
library(XML)
library(magrittr)
library(tibble)

link <- ("https://www.fundamentus.com.br/resultado.php")
User <- "Mozilla/5.0 (Windows NT 6.1; WOW64)"

html <- getURL("https://www.fundamentus.com.br/resultado.php",
               httpheader = c("User-Agent"=User))

DADOS <- readHTMLTable(html)
TABELA <- DADOS$resultado 
TABELA$Liq.2meses <-  gsub("\\." , "" ,TABELA$Liq.2meses) %>%  gsub("\\," , "." ,.) %>% as.integer(., length  = 20)
TABELA$Liq.2meses <- ifelse(TABELA$Liq.2meses <= 	100000,NA,TABELA$Liq.2meses) 
TABELA <- na.omit(TABELA)
TABELA <- TABELA[,1:3]

colnames(TABELA) <- c("Papel","Preço","P_L")
TABELA$P_L <- gsub("\\." , "" ,TABELA$P_L) %>%  gsub("\\," , "." ,.) %>% as.numeric()
TABELA$Preço <- gsub("\\." , "" ,TABELA$Preço) %>%  gsub("\\," , "." ,.) %>% as.numeric()
TABELA$Preço <- ifelse(TABELA$Preço == 0,NA,TABELA$Preço)
TABELA$P_L <- ifelse(TABELA$P_L == 0,NA,TABELA$P_L)

TABELA <- na.omit(TABELA)
TABELA$E_Y <- 1/TABELA$P_L
TABELA <- TABELA[order(TABELA$E_Y, decreasing = T),]

view(TABELA)

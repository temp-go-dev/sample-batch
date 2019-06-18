# sample-batch


## 実行MEMO

sample-api-ginで使用してるDBが必要


* 正常終了  
`set ENV_DBMS=mysql&set ENV_USER=user&set ENV_PASS=password&set ENV_PROTOCOL=tcp(localhost:3306)&set ENV_DBNAME=sampledb& go run main.go 123`  


* exitCd=1で終了するルート  
第一引数を`9999`  
`set ENV_DBMS=mysql&set ENV_USER=user&set ENV_PASS=password&set ENV_PROTOCOL=tcp(localhost:3306)&set ENV_DBNAME=sampledb& go run main.go 9999`  

  * 実行結果
    ```
    >set ENV_DBMS=mysql&set ENV_USER=user&set ENV_PASS=password&set ENV_PROTOCOL=tcp(localhost:3306)&set ENV_DBNAME=sampledb& go run main.go
    running
    errorHandle!! ExitErrorType
    messageID:0000
    err:引数0件エラー
    exitCode:1
    exit status 1
    ```
gRPC cli

github ktr0731/evans  > https://github.com/ktr0731/evans

download the file from release -> https://github.com/ktr0731/evans/releases/tag/v0.10.9 -> or find latest one

copy to ~/.local/bin/evans

evans --host=localhost --port=5051 --reflection repl
in evens cli 
>> show package   : list packages

>> package greet :: select package to work on

    >> show messages  --> list all messages

    >> show service 


    >> service GreetService  --> select service to work on

        >> call Greet --> to call gRPC endpoint

        >> for client streaming it will keep asking for inputs --> use ctrl+D to finish inputs stream and see result
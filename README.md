A simple command line tool for generating bcrypted passwords.

Download the precompiled tool for your operating system from the [releases page](https://github.com/ericchiang/bcrypt-cli/releases).

```
$ wget https://github.com/ericchiang/bcrypt-cli/releases/download/v1.0.0/bcrypt-cli-v1.0.0-linux-amd64.tar.gz
$ sha512sum bcrypt-cli-v1.0.0-linux-amd64.tar.gz
# verify results
$ tar -zxvf bcrypt-cli-v1.0.0-linux-amd64.tar.gz
./
./bcrypt-cli/
./bcrypt-cli/bcrypt-cli
```

Use the tool to create a bcrypted password.

```
$ ./bcrypt-cli/bcrypt-cli 
Enter password: 
Re-enter password: 
$2a$10$C9DliP580ooWFpr6XTQwOuFd3znBNDMVSytJmm9viL4XLZuy.PnMm
```

Increase the cost using the `-cost` flag (defaults to 10).

```
$ ./bcrypt-cli/bcrypt-cli -cost=14
Enter password: 
Re-enter password: 
$2a$14$830vR/vhwIFKXaLtHaP41OInok.XsUBAgrG.EC3JqyQUMrbj091qG
```

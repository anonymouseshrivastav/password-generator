## Password generator

Usecase:- Suppose you know the correct username and you try some passwords related to that. In this case you can use this to generate passwords related to the username. Another best usecase of this the wordpress login where you can easly enumurate the username and for password you can use this tool.
### Supported in:
Linux <br/>
Termux <br/>
Mac Os <br/>
Windows <br/>


### Installation:
(Linux)
```bash
sudo apt install golang git
```

(Termux)
```bash
pkg install golang git
```

```bash
git clone https://github.com/sudityashrivastav/password-generator
```

```bash
cd password-generator
```

```bash
go build
```

### Usage:
```bash
passgen <username>
```

A new file ```pass_list.txt``` will be created in current directory. 

### Access the program from any directory.

(Linux)
```bash
cp passgen /usr/bin/passgen
```

(Termux)
```bash
cp passgen /data/data/com.termux/files/usr/bin/passgen
```


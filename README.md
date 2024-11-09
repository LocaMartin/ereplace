# ereplace
### This tool is tested on ```go 1.23.2```
#### replace URL parameter value with ease with ereplace
This project is inspired by <a herf='https://github.com/tomnomnom/qsreplace'>qsreplece</a> project .

#### Installation

```
go install github.com/LocaMartin/ereplace@latest
```
##### Or you can download prebuild binary using `wget` & `curl` command and it move to `~/go/bin` directory

```
wget -O ereplace https://github.com/LocaMartin/ereplace/pkg/ereplace
```
```
curl -o ereplace https://github.com/LocaMartin/ereplace/pkg/ereplace
```
```
mv ereplace ~/go/bin
```


#### Flags

- `-h`: Show help message
- `-u <url>`: Single URL to modify
- `-uf <file>`: File containing multiple URLs to modify (one per line)
- `-p`: Single payload
- `-pf <file>`: Single payload file to modify
(can specify multiple files)
- `-s <output_file>`: Save modified payloads to a file
- `-version`,`-v`: Print version information

#### Usage

##### single url single payload:
```
ereplace -u "httpx://example.com/file=help" -p "/../etc/passwd"
```
##### sigle url payload file:

```
ereplace -u "httpx://example.com/file=help" -pf payload.txt
```
##### single url multiple payload file:
```
ereplace -u "httpx://example.com/file=help" -pf payload1.txt payload2.txt payload3.txt
```
##### save results:

```
ereplace -u "httpx://example.com/file=help" -p "/../etc/passwd" -s results.txt
```

#### Intigration with other tools & command

##### echo:
```
echo "httpx://example.com/file=help" | ereplace -p "/../etc/passwd" -s results.txt
```
##### cat:
```
cat url.txt | ereplace -p "/../etc/passwd" -s results.txt
```
##### httpx:

```
echo "httpx://example.com/file=help" | ereplace -p "/../etc/passwd" | httpx -silent -mr "root:x"
```
#### Support My Work

If you find this tool helpful, please consider supporting my work:

<p align="center"><a href="https://buymeacoffee.com/locabomartin"><img  src="https://img.shields.io/badge/Buy%20Me%20a%20Coffee-ffdd00?style=for-the-badge&logo=buy-me-a-coffee&logoColor=black"/></a></p>

Your support helps me continue developing and maintaining open-source tools like ereplace.
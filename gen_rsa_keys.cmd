@echo off

echo Download necessary components
mkdir tmp
powershell -Command "& {$client = new-object System.Net.WebClient; $client.DownloadFile('https://mirror.firedaemon.com/OpenSSL/openssl-3.0.0.zip','openssl.zip');}"
powershell -Command "& {Add-Type -AssemblyName System.IO.Compression.FileSystem; [System.IO.Compression.ZipFile]::ExtractToDirectory('openssl.zip', 'tmp');}"
move tmp\openssl-3\x64\bin tmp\openssl >NUL

echo Creating secrets directory and keys
rd /S /Q secrets >NUL
mkdir secrets >NUL
tmp\openssl\openssl genpkey -algorithm RSA -out secrets\token.private -pkeyopt rsa_keygen_bits:4096 >NUL
tmp\openssl\openssl rsa -pubout -in secrets\token.private -out secrets\token.public >NUL

echo Clean no more necessary components
rd /S /Q tmp >NUL
del /Q /F openssl.zip >NUL

# Develop JVM by go

## init

```shell
go mod init github.com/taoistwar/go-jvm
```

## java demo

```shell
cd java
javac com/github/taoistwar/java/HelloWorld.java
```

## run

```shell
./go-jvm -cp demo com.github.taoistwar.java.HelloWorld
```

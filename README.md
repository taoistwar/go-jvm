# Develop JVM by go

## init

```shell
go mod init github.com/taoistwar/go-jvm
```

## java demo

```shell
cd demo
javac com/github/taoistwar/java/HelloWorld.java
javac com/github/taoistwar/java/ClassFileTest.java
```

## run

```shell
./go-jvm -cp demo com.github.taoistwar.java.HelloWorld
```

## 字节码

不同版本的 JDK，对字节码格式有各自的扩展。

```shell
cd demo
javap -v com.github.taoistwar.java.ClassFileTest
```

结果：

```class

```

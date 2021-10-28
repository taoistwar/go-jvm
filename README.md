# Develop JVM by go

TODO 适配 JDK17

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
javap -l -c -v com.github.taoistwar.java.ClassFileTest
```

### local vars

```shell
javap -l -c -v com.github.taoistwar.java.LocalVarDemo
```

## 随想

- 表态类型，内存中只存储数据。
- 动态类型，内存中存储类型和数据。

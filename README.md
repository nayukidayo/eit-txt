## EIT TXT TO JSON

### 运行程序

`eit-txt.exe -mode="模式" -txt="文件路径"`

转换成功输出【完成】，并且在输入文件同级目录下生成同名 JSON 文件，其他输出代表转换失败

### 显示帮助

`./eit-txt.exe -h`

```
-mode string
      模式 uell400 uell uref cirs cirs812
-txt string
      文件路径
```

### uell400：400 帧数据

`./eit-txt.exe -mode="uell400" -txt="./test/uell400.txt"`

### uell：400 帧数据随机选取 1 帧

`./eit-txt.exe -mode="uell" -txt="./test/uell.txt"`

### uref：空场矩阵 208 x 1

`./eit-txt.exe -mode="uref" -txt="./test/uref.txt"`

### cirs：灵敏度矩阵 208 x 1024

`./eit-txt.exe -mode="cirs" -txt="./test/cirs1024.txt"`

### cirs812：灵敏度矩阵 208 x 812

`./eit-txt.exe -mode="cirs812" -txt="./test/cirs812.txt"`

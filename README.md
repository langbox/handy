# handy
helpful tools

## install

```
go get github.com/langbox/handy
```

## 功能说明

### 获取随机id

```
id, err := handy.GenShortID()
fmt.Println(id, err)
```

### 清除字符串两边空格

```
str := " 1111    "
s := handy.Trim(str)
fmt.Println(len(s))

或者 使用 strings包
str := " 1111    "
s := strings.Trim(str," ")
fmt.Println(len(s))
```
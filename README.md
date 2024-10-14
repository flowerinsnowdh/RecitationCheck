# RecitationCheck
“抽查背诵”，抽查拼读

# 背景
在学苏州话，想抽查一下自己会不会某字词的发音

# 用法
是向导模式

```shell
请选择词库
1. 称呼（你、我、Ta...）
2. 打招呼（你好）
```

此时输入数字即可开始抽查

```console
我
```

此时输入`ok`进行下一个

此时输入`tip`显示提示词

此时输入`?`查询词典

不保证顺序

# 配置内容
- `dict` string - 词典链接
- `rows` [][RowObject](#rowobject) - 词库

## RowObject
- `desc` string - 显示在列表上的名称
- `words` map[string]string - 练习词语，键值对，键是要练习的词汇，值是提示词

# 工具
临时个人工具，不会进行维护

# 版权信息
- [go](https://github.com/golang/go) Copyright 2009 The Go Authors. ([BSD 3-Clause "New" or "Revised" License](https://github.com/golang/go/blob/master/LICENSE))
- [go-toml v2](https://github.com/pelletier/go-toml) Copyright (c) 2021 - 2023 Thomas Pelletier [MIT License](https://github.com/pelletier/go-toml/blob/v2/LICENSE)

[config.toml](config.toml)文件中使用了苏白学堂的教学内容
- 学堂版拼音摘录自 [av11828488](https://www.bilibili.com/video/av11828488)
- 入门日常语摘录自 [av24054870](https://www.bilibili.com/video/av24054870)
- 侵权请联系我，我会第一时间删除

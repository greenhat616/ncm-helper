# 网易云助手

这是一个实验性（还在活跃开发中）的项目，因此缺乏特性说明。  
当项目进入 **稳定**(Stable) 阶段，我们将补充配置，特性等说明。

<div align="center">

![Go](https://github.com/greenhat616/ncm-helper/workflows/Go/badge.svg)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/greenhat616/ncm-helper)
[![Maintainability](https://api.codeclimate.com/v1/badges/3eedabb10c8fa983538d/maintainability)](https://codeclimate.com/github/greenhat616/ncm-helper/maintainability)
[![codecov](https://codecov.io/gh/greenhat616/ncm-helper/branch/master/graph/badge.svg)](https://codecov.io/gh/greenhat616/ncm-helper)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/greenhat616/ncm-helper?sort=semver)

</div>

## 状态：开发
* 新特性会不断添加
* 机制存在缺陷，甚至不能运行
* 可能会丢失数据
* 运行时可能会出现不合预期的问题
* **不能用于生产环境**

## 鸣谢

网易云助手离不开 **萌创团队** 以及 **一言项目组** 的支持，更离不开  [JetBrains](https://www.jetbrains.com/?from=hitokoto-osc) 为开源项目免费提供具有强生产力的 IDE 等相关授权。
[<img src=".github/jetbrains-variant-3.png" width="200"/>](https://www.jetbrains.com/?from=hitokoto-osc)

## 许可证
项目原协议版权归 **网易公司**，项目其他代码遵守：  
**GNU General Public License v3.0**。  

此外，项目只用于学习目的，不会通过任何途径 **签发** 或 **授权** 商用行为（commercial use）。


## 开发
现在，让我们简单阐述下怎样参与咱们的开发。

### 依赖
* 项目协议来自 [Binaryify/NeteaseCloudMusicApi](https://github.com/Binaryify/NeteaseCloudMusicApi) 
* 框架（重要的外部依赖）
  * 配置：viper
  * 日志： logrus
  * Web： gin
  * flag 解析：pflag
  * CI/CD：Github Action（后期前端内容的继承也将通过此服务）

### 编译
```bash
$ make build
```  

### 测试
```bash
$ make test
```

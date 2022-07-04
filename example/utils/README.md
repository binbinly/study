## 特性

- 👏 全面、高效、可复用
- 💪 180+常用go工具函数，支持string、slice、datetime、net、crypt...
- 💅 只依赖go标准库
- 🌍 所有导出函数单元测试覆盖率100%

## 安装

```go
go get github.com/duke-git/lancet
```

## 用法

lancet是以包的结构组织代码的，使用时需要导入相应的包名。例如：如果使用字符串相关函数，需要导入strutil包:

```go
import "github.com/duke-git/lancet/strutil"
```

## 例子

此处以字符串工具函数ReverseStr（逆序字符串）为例，需要导入strutil包:

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    s := "hello"
    rs := strutil.ReverseStr(s)
    fmt.Println(rs) //olleh
}
```

## API文档
### convertor转换器包支持一些常见的数据类型转换。

```go
import "github.com/duke-git/lancet/convertor"
```
#### 函数列表:
- [ColorHexToRGB](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ColorHexToRGB)
- [ColorRGBToHex](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ColorRGBToHex)
- [ToBool](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToBool)
- [ToBytes](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToBytes)
- [ToChar](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToChar)
- [ToInt](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToInt)
- [ToJson](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToJson)
- [ToString](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToString)
- [StructToMap](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#StructToMap)
  
### cryptor加密包支持数据加密和解密，获取md5，hash值。支持base64, md5, hmac, aes, des, rsa。

```go
import "github.com/duke-git/lancet/cryptor"
```

#### 函数列表:
- [AesEcbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesEcbEncrypt)
- [AesEcbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesEcbDecrypt)
- [AesCbcEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesCbcEncrypt)
- [AesCbcDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesCbcDecrypt)
- [AesCtrCrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesCtrCrypt)
- [AesCfbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesCfbEncrypt)
- [AesCfbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesCfbDecrypt)
- [AesOfbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesOfbEncrypt)
- [AesOfbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesOfbDecrypt)
- [Base64StdEncode](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Base64StdEncode)
- [Base64StdDecode](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Base64StdDecode)
- [DesEcbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesEcbEncrypt)
- [DesEcbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesEcbDecrypt)
- [DesCbcEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesCbcEncrypt)
- [DesCbcDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesCbcDecrypt)
- [DesCtrCrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesCtrCrypt)
- [DesCfbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesCfbEncrypt)
- [DesCfbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesCfbDecrypt)
- [DesOfbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesOfbEncrypt)
- [DesOfbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesOfbDecrypt)
- [HmacMd5](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#HmacMd5)
- [HmacSha1](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#HmacSha1)
- [HmacSha256](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#HmacSha256)
- [HmacSha512](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#HmacSha512)
- [Md5String](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Md5String)
- [Md5File](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Md5File)
- [Sha1](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Sha1)
- [Sha256](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Sha256)
- [Sha512](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Sha512)
- [GenerateRsaKey](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#GenerateRsaKey)
- [RsaEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#RsaEncrypt)
- [RsaDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#RsaDecrypt)

### datetime日期时间处理包，格式化日期，比较日期。


```go
import "github.com/duke-git/lancet/datetime"
```
#### 函数列表:
- [AddDay](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#AddDay)
- [AddHour](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#AddHour)
- [AddMinute](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#AddMinute)
- [BeginOfMinute](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfMinute)
- [BeginOfHour](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfHour)
- [BeginOfDay](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfDay)
- [BeginOfWeek](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfWeek)
- [BeginOfMonth](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfMonth)
- [BeginOfYear](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfYear)
- [EndOfMinute](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfMinute)
- [EndOfHour](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfHour)
- [EndOfDay](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfDay)
- [EndOfWeek](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfWeek)
- [EndOfMonth](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfMonth)
- [EndOfYear](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfYear)
- [GetNowDate](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#GetNowDate)
- [GetNowTime](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#GetNowTime)
- [GetNowDateTime](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#GetNowDateTime)
- [GetZeroHourTimestamp](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#GetZeroHourTimestamp)
- [GetNightTimestamp](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#GetNightTimestamp)
- [FormatTimeToStr](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#FormatTimeToStr)
- [FormatStrToTime](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#FormatStrToTime)

### fileutil包支持文件基本操作。

```go
import "github.com/duke-git/lancet/fileutil"
```

#### 函数列表：

- [ClearFile](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#ClearFile)
- [CreateFile](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#CreateFile)
- [CopyFile](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#CopyFile)
- [FileMode](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#FileMode)
- [MiMeType](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#MiMeType)
- [IsExist](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#IsExist)
- [IsLink](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#IsLink)
- [IsDir](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#IsDir)
- [ListFileNames](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#ListFileNames)
- [RemoveFile](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#RemoveFile)
- [ReadFileToString](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#ReadFileToString)
- [ReadFileByLine](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#ReadFileByLine)
- [Zip](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#Zip)
- [UnZip](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#UnZip)

### formatter格式化器包含一些数据格式化处理方法。

```go
import "github.com/duke-git/lancet/formatter"
```
#### 函数列表:
- [Comma](https://github.com/duke-git/lancet/blob/main/docs/formatter_zh-CN.md#Comma)


### function函数包控制函数执行流程，包含部分函数式编程。

```go
import "github.com/duke-git/lancet/function"
```

#### 函数列表:
- [After](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#After)
- [Before](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Before)
- [Curry](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Curry)
- [Compose](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Compose)
- [Debounced](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Debounced)
- [Delay](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Delay)
- [Watcher](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Watcher)

### mathutil包实现了一些数学计算的函数。

```go
import "github.com/duke-git/lancet/mathutil"
```

#### Function list:
- [Exponent](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Exponent)
- [Fibonacci](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Fibonacci)
- [Factorial](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Factorial)
- [Percent](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Percent)
- [RoundToFloat](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#RoundToFloat)
- [RoundToString](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#RoundToString)
- [TruncRound](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#TruncRound)

### netutil网络包支持获取ip地址，发送http请求。

```go
import "github.com/duke-git/lancet/netutil"
```

#### 函数列表:
- [ConvertMapToQueryString](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#ConvertMapToQueryString)
- [GetInternalIp](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#GetInternalIp)
- [GetIps](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetIps)
- [GetMacAddrs](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetMacAddrs)
- [GetPublicIpInfo](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#GetPublicIpInfo)
- [IsPublicIP](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#IsPublicIP)
- [HttpGet](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpGet)
- [HttpDelete](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpDelete)
- [HttpPost](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpPost)
- [HttpPut](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpPut)
- [HttpPatch](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpPatch)
- [ParseHttpResponse](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#ParseHttpResponse)

### random随机数生成器包，可以生成随机[]bytes, int, string。

```go
import "github.com/duke-git/lancet/random"
```

#### 函数列表:
- [RandBytes](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#RandBytes)
- [RandInt](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#RandInt)
- [RandString](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#RandString)
- [UUIdV4](https://github.com/duke-git/lancet/blob/main/docs/random.md#UUIdV4)
### retry重试执行函数直到函数运行成功或被context cancel。

```go
import "github.com/duke-git/lancet/retry"
```

#### 函数列表:
- [Context](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#Context)
- [Retry](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#Retry)
- [RetryFunc](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#RetryFunc)
- [RetryDuration](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#RetryDuration)
- [RetryTimes](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#RetryTimes)


### slice包包含操作切片的方法集合。

```go
import "github.com/duke-git/lancet/slice"
```

#### 函数列表:
- [Contain](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Contain)
- [ContainSubSlice](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#ContainSubSlice)
- [Chunk](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Chunk)
- [Compact](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Compact)
- [Concat](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Concat)
- [Count](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Count)
- [Difference](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Difference)
- [DifferenceBy](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#DifferenceBy)
- [DeleteByIndex](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#DeleteByIndex)
- [Drop](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Drop)
- [Every](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Every)
- [Filter](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Filter)
- [Find](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Find)
- [FindLast](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#FindLast)
- [FlattenDeep](#FlattenDeep)
- [ForEach](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#ForEach)
- [GroupBy](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#GroupBy)
- [IntSlice](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#IntSlice)
- [InterfaceSlice](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#InterfaceSlice)
- [Intersection](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Intersection)
- [InsertByIndex](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#InsertByIndex)
- [Map](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Map)
- [ReverseSlice](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#ReverseSlice)
- [Reduce](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Reduce)
- [Shuffle](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Shuffle)
- [SortByField](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#SortByField)
- [Some](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Some)
- [StringSlice](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#StringSlice)
- [Unique](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Unique)
- [Union](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Union)
- [UpdateByIndex](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#UpdateByIndex)
- [Without](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Without)


### strutil包含处理字符串的相关函数。

```go
import "github.com/duke-git/lancet/strutil"
```

#### 函数列表:

- [After](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#After)
- [AfterLast](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#AfterLast)
- [Before](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#Before)
- [BeforeLast](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#BeforeLast)
- [CamelCase](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#CamelCase)
- [Capitalize](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#Capitalize)
- [IsString](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#IsString)
- [KebabCase](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#KebabCase)
- [LowerFirst](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#LowerFirst)
- [UpperFirst](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#UpperFirst)
- [PadEnd](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#PadEnd)
- [PadStart](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#PadStart)
- [ReverseStr](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#ReverseStr)
- [SnakeCase](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#SnakeCase)
- [Wrap](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#Wrap)
- [Unwrap](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#Unwrap)
  

### system包含os, runtime, shell command相关函数。

```go
import "github.com/duke-git/lancet/system"
```

#### 函数列表:
- [IsWindows](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN.md#IsWindows)
- [IsLinux](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN.md#IsLinux)
- [IsMac](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN.md#IsMac)
- [GetOsEnv](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN.md#GetOsEnv)
- [SetOsEnv](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN.md#SetOsEnv)
- [RemoveOsEnv](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN.md#RemoveOsEnv)
- [CompareOsEnv](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN.md#CompareOsEnv)
- [ExecCommand](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN.md#ExecCommand)

### validator验证器包，包含常用字符串格式验证函数。

```go
import "github.com/duke-git/lancet/validator"
```
#### 函数列表:

- [ContainChinese](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#ContainChinese)
- [ContainLetter](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#ContainLetter)
- [ContainLower](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#ContainLower)
- [ContainUpper](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#ContainUpper)
- [IsAlpha](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsAlpha)
- [IsAllUpper](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsAllUpper)
- [IsAllLower](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsAllLower)
- [IsBase64](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsBase64)
- [IsChineseMobile](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsChineseMobile)
- [IsChineseIdNum](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsChineseIdNum)
- [IsChinesePhone](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsChinesePhone)
- [IsCreditCard](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsCreditCard)
- [IsDns](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsDns)
- [IsEmail](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsEmail)
- [IsEmptyString](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsEmptyString)
- [IsFloatStr](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsFloatStr)
- [IsNumberStr](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsNumberStr)
- [IsJSON](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsJSON)
- [IsRegexMatch](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsRegexMatch)
- [IsIntStr](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsIntStr)
- [IsIp](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsIp)
- [IsIpV4](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsIpV4)
- [IsIpV6](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsIpV6)
- [IsStrongPassword](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsStrongPassword)
- [IsUrl](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsUrl)
- [IsWeakPassword](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsWeakPassword)


## 如何贡献代码

非常感激任何的代码提交以使lancet的功能越来越强大。创建pull request时请遵守以下规则。

1. Fork lancet仓库。
2. 创建自己的特性分支。
3. 提交变更。
4. Push分支。
5. 创建新的pull request。
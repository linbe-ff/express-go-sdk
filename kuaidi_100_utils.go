package express

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func (k *kuaiDi100) DoRequest(t string, param string, postUrl string) (string, error) {
	// 计算签名
	signStr := param + t + k.Key + k.Secret
	hash := md5.New()
	hash.Write([]byte(signStr))
	sign := hex.EncodeToString(hash.Sum(nil))
	sign = strings.ToUpper(sign)

	// 构造form表单数据
	formData := url.Values{}
	formData.Add("key", k.Key)
	formData.Add("t", t)
	formData.Add("sign", sign)
	formData.Add("param", param)

	return k.execute(postUrl, formData)
}

/*
*
多一个Method入参
*/
func (k *kuaiDi100) DoMethodRequest(method string, t string, param string, postUrl string) (string, error) {
	// 计算签名
	signStr := param + t + k.Key + k.Secret
	hash := md5.New()
	hash.Write([]byte(signStr))
	sign := hex.EncodeToString(hash.Sum(nil))
	sign = strings.ToUpper(sign)

	// 构造form表单数据
	formData := url.Values{}
	formData.Add("key", k.Key)
	formData.Add("method", method)
	formData.Add("t", t)
	formData.Add("sign", sign)
	formData.Add("param", param)

	return k.execute(postUrl, formData)
}

/*
*
使用customer鉴权
*/
func (k *kuaiDi100) CustomerRequest(param string, postUrl string) (string, error) {
	// 计算签名
	signStr := param + k.Key + k.Customer
	hash := md5.New()
	hash.Write([]byte(signStr))
	sign := hex.EncodeToString(hash.Sum(nil))
	sign = strings.ToUpper(sign)

	// 构造form表单数据
	formData := url.Values{}
	formData.Add("customer", k.Customer)
	formData.Add("sign", sign)
	formData.Add("param", param)

	return k.execute(postUrl, formData)
}

/*
*
根据map传入form数据
*/
func (k *kuaiDi100) DoMapRequest(m map[string]string, postUrl string) (string, error) {
	formData := url.Values{}
	// 由map生成form表单数据
	for key, value := range m {
		formData.Add(key, value)
	}
	return k.execute(postUrl, formData)
}

func (k *kuaiDi100) DoFileRequest(m map[string]string, file *os.File, postUrl string) (string, error) {
	// 创建一个缓冲区来存储表单数据
	buf := bytes.NewBuffer(nil)

	// 创建一个新的 multipart/writer
	writer := multipart.NewWriter(buf)

	// 添加 map 中的数据到表单
	for key, value := range m {
		if err := writer.WriteField(key, value); err != nil {
			return "", err
		}
	}

	// 添加文件到表单
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		return "", err
	}
	if _, err = io.Copy(part, file); err != nil {
		return "", err
	}

	// 关闭 writer，否则请求体可能会不完整
	writer.Close()

	// 创建HTTP客户端
	client := &http.Client{}

	// 创建一个新的 http 请求
	req, err := http.NewRequest("POST", postUrl, buf)
	if err != nil {
		return "", err
	}

	// 设置请求头中的 Content-Type
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送HTTP请求
	fmt.Printf("请求信息: %v\n", req)
	fmt.Printf("请求参数: %v\n", fmt.Sprintf("%v", buf))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return "请求失败", err
	}
	defer resp.Body.Close()

	// 读取响应内容
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return string(respBody), err
	}

	// 打印响应内容
	fmt.Println("响应内容:", string(respBody))
	return string(respBody), err
}

/**
*执行HTTP请求
 */
func (k *kuaiDi100) execute(postUrl string, formData url.Values) (string, error) {
	// 创建HTTP客户端
	client := &http.Client{}

	// 创建HTTP请求
	req, err := http.NewRequest("POST", postUrl, strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return "创建请求失败", err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送HTTP请求
	fmt.Printf("请求信息: %v\n", req)
	fmt.Printf("请求参数: %v\n", formData)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return "请求失败", err
	}
	defer resp.Body.Close()

	// 读取响应内容
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return string(respBody), err
	}

	// 打印响应内容
	fmt.Println("响应内容:", string(respBody))
	return string(respBody), err
}

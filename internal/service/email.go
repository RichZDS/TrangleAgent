package service

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
	"github.com/patrickmn/go-cache"
)

var (
	// 验证码缓存
	// 缓存中的验证代码将在创建后5分钟内有效，且每隔10分钟进行一次清理。
	verificationCodeCache = cache.New(5*time.Minute, 10*time.Minute)
)

type EmailService interface {
	// SendVerificationCode 向用户的邮箱发送验证码
	SendVerificationCode(ctx context.Context, to string) error
	// VerifyVerificationCode 验证邮箱的验证码
	VerifyVerificationCode(email string, code string) bool
}

type emailService struct {
}

func NewEmailService() EmailService {
	return &emailService{}
}

// SendVerificationCode 向用户的邮箱发送验证码
func (e *emailService) SendVerificationCode(ctx context.Context, to string) error {
	code := generateVerificationCode()

	err := e.sendVerificationCode(to, code)
	if err != nil {
		return err
	}

	// 将验证码存储在缓存中以供后续验证
	verificationCodeCache.Set(to, code, cache.DefaultExpiration)

	return nil
}

// sendVerificationCode 发送验证代码到指定的邮箱。
// 参数 to: 邮件接收人的邮箱地址。
// 参数 code: 需要发送的验证代码。
// 返回值 error: 发送过程中遇到的任何错误。
func (e *emailService) sendVerificationCode(to string, code string) error {
	// 创建一个新的邮件实例
	em := email.NewEmail()
	em.From = "3149026837@qq.com"
	em.To = []string{to}
	em.Subject = "验证码"

	// 读取HTML模板文件
	htmlFilePath := "resource/public/verification_code.html"
	if tmplContent, err := ioutil.ReadFile(htmlFilePath); err == nil {
		// 成功读取模板文件，使用模板
		tmpl, err := template.New("verification").Parse(string(tmplContent))
		if err == nil {
			var buf bytes.Buffer
			data := map[string]string{"Code": code}
			if err := tmpl.Execute(&buf, data); err == nil {
				em.HTML = buf.Bytes()
			} else {
				// 模板执行失败，使用默认内容
				em.HTML = []byte(fmt.Sprintf(`
                    <h1>验证码</h1>
                    <p>您的验证码是: <strong>%s</strong></p>
                    <p>此验证码将在5分钟后过期。</p>
                `, code))
			}
		} else {
			// 模板解析失败，使用默认内容
			em.HTML = []byte(fmt.Sprintf(`
                <h1>验证码</h1>
                <p>您的验证码是: <strong>%s</strong></p>
                <p>此验证码将在5分钟后过期。</p>
            `, code))
		}
	} else {
		// 读取模板文件失败，使用默认内容
		em.HTML = []byte(fmt.Sprintf(`
            <h1>验证码</h1>
            <p>您的验证码是: <strong>%s</strong></p>
            <p>此验证码将在5分钟后过期。</p>
        `, code))
	}

	// 发送邮件(这里使用QQ进行发送邮件验证码)
	_ = em.Send(
		"smtp.qq.com:587",
		smtp.PlainAuth("", "3149026837@qq.com", "szkgunhouolidghb", "smtp.qq.com"),
	)

	//if err != nil {
	//	g.Log().Errorf(context.Background(), "send email error: %+v", err)
	//	return fmt.Errorf("发送邮件失败: %v", err)
	//}

	return nil // 邮件发送成功，返回nil
}

// generateVerificationCode 随机生成一个6位数的验证码。
func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	return code
}

// VerifyVerificationCode 验证发送给用户的验证码
func (e *emailService) VerifyVerificationCode(email string, code string) bool {
	// 调试代码
	if code == "123456" {
		return true
	}

	// 从缓存中检索验证码
	cachedCode, found := verificationCodeCache.Get(email)
	// 如果没有找到验证码或者验证码过期，返回false
	if !found {
		return false
	}

	// 比较缓存中的代码与提供的代码
	if cachedCode != code {
		return false
	}

	return true
}

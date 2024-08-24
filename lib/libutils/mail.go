package libutils

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"io"
	"mime"
	"os"
)

type Mail struct {
	user   string
	passwd string
	host   string
	port   int
}

// user:pass@host:port
func NewMail(user string, password string, host string, port int) *Mail {
	return &Mail{
		user:   user,
		passwd: password,
		host:   host,
		port:   port,
	}
}

func (mail *Mail) SendMsg(msg *gomail.Message) error {
	diaConn := gomail.NewDialer(mail.host, mail.port, mail.user, mail.passwd)
	return diaConn.DialAndSend(msg)
}

// SendContent 发送邮件
func (mail *Mail) SendContent(title, content string, toList, ccList []string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", mail.user)
	msg.SetHeader("Cc", ccList...)
	msg.SetHeader("To", toList...)
	msg.SetHeader("Subject", title)
	msg.SetBody("text/plain", fmt.Sprintf("%v", content))

	return mail.SendMsg(msg)
}

// SendAttach 通过文件路径添加附件
func (mail *Mail) SendAttach(filepath string, name string, title, content string, toList, ccList []string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	msg := gomail.NewMessage()
	msg.Attach(mime.QEncoding.Encode("UTF-8", name),
		gomail.SetCopyFunc(func(writer io.Writer) error {
			_, err = writer.Write(data)
			return err
		}))

	msg.SetHeader("From", mail.user)
	msg.SetHeader("Cc", ccList...)
	msg.SetHeader("To", toList...)
	msg.SetHeader("Subject", title)
	msg.SetBody("text/plain", fmt.Sprintf("%v", content))

	return mail.SendMsg(msg)
}

// SendHtml 发送邮件
func (mail *Mail) SendHtml(subject, content string, toList, ccList []string) error {
	msg := gomail.NewMessage(
		//发送文本时设置编码，防止乱码。 如果txt文本设置了之后还是乱码，那可以将原txt文本在保存时
		//就选择utf-8格式保存
		gomail.SetEncoding(gomail.Base64),
	)
	msg.SetHeader("From", mail.user)  // 添加别名
	msg.SetHeader("To", toList...)    // 发送给用户(可以多个)
	msg.SetHeader("Cc", ccList...)    // 抄送给用户(可以多个)
	msg.SetHeader("Subject", subject) // 设置邮件主题
	msg.SetBody("text/html", content) // 设置邮件正文

	return mail.SendMsg(msg)
}

func (mail *Mail) SendContentToOne(title, content string, target string) error {
	return mail.SendContent(title, content, []string{target}, nil)
}

func (mail *Mail) GetAccountInfo() string {
	return fmt.Sprintf("%s:%s@%s:%v", mail.user, mail.passwd, mail.host, mail.port)
}

package controllers

import (
	"bejigo/models"
	"bejigo/syserrors"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"io"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

const SESSION_USER_KEY = "SESSION_USER_KEY" // 定义一个常量，用于保存用户登录session

type BaseController struct {
	beego.Controller
	User    models.User
	IsLogin bool // 用户是否登录
}

type M map[string]interface{}

// 返回json数据
func (p *BaseController) ReturnJson(msg, url string) {
	p.Data["json"] = map[string]interface{}{
		"code": 0,
		"msg":  msg,
		"url":  url,
	}
	p.ServeJSON()
}

// 返回json提示信息
func (p *BaseController) ReturnJsonCode(msg string) {
	p.Data["json"] = map[string]interface{}{
		"code": 0,
		"msg":  msg,
	}
	p.ServeJSON()
}

// 返回json数据但没有跳转
func (p *BaseController) ReturnJsonNoUrl(msg string, data M) {
	data["code"] = 0
	data["msg"] = msg
	data["json"] = data
	p.ServeJSON()
}

func (p *BaseController) Prepare() {
	u, ok := p.GetSession("SESSION_USER_KEY").(models.User) // 类型判断
	p.IsLogin = false                                       // 把登录状态设置为false
	if ok {
		p.User = u
		p.IsLogin = true
		p.Data["ID"] = p.User.ID
		p.Data["UserName"] = p.User.UserName
		p.Data["Email"] = p.User.Email
		p.Data["Avatar"] = p.User.Avatar
		p.Data["Role"] = p.User.Role
	}
}

// 权限控制
func (p *BaseController) permission() {
	u, ok := p.GetSession("SESSION_USER_KEY").(models.User)
	if ok {
		if u.Role != 1 {
			p.Abort("暂无权限！")
		}
	}
}

// 错误提示
func (p *BaseController) About500(err error) {
	p.Data["error"] = err
	p.Abort("500")
}

// 字符串不能为空
func (p *BaseController) GetMustString(key, msg string) string {
	key = p.GetString(key)
	if utf8.RuneCountInString(key) == 0 {
		p.About500(errors.New(msg))
	}
	return key
}

// 字符串长度限制
func (p *BaseController) GetAllowMaxString(key, msg string, length int) string {
	key = p.GetString(key)
	if utf8.RuneCountInString(key) > length {
		p.About500(errors.New(msg))
	}
	return key
}

// 不能为空和长度限制并返回验证过后的值
func (p *BaseController) GetMustAndInlen(key, nullErrNotice, overLenErrNotice string, maxLength int) string {
	key = p.GetString(key)
	if utf8.RuneCountInString(key) == 0 {
		p.About500(errors.New(nullErrNotice))
	}
	if utf8.RuneCountInString(key) > maxLength {
		p.About500(errors.New(overLenErrNotice))
	}
	return key
}

// 必须登录
func (p *BaseController) MustLogin() {
	if !p.IsLogin { // 如果没登录，提示错误
		p.Ctx.Redirect(301, "/admin/login/index")
		return
	}
}

// 获取所有分类
func (p *BaseController) GetAllCategory() (cate []models.Category) {
	var category []models.Category
	err := models.Category{}.GetAllCategory(&category)
	if err != nil {
		p.About500(syserrors.New("暂无分类！", err))
	}
	return category
}

// 获取所有的标签
func (p *BaseController) GetAllTag() (t []models.Tag) {
	// 取出所有的标签
	var tag []models.Tag
	err1 := models.GetAllTag(&tag)
	if err1 != nil {
		p.About500(syserrors.New("暂无标签！", err1))
	}
	return tag
}

//字串截取
func (p *BaseController) SubString(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

// 初始化搜索数据，用于前台搜索
func (p *BaseController) InitSearchData() {
	var searchData []models.Search
	models.Article{}.GetArticleDataForSearch(&searchData)
	//content, err := json.MarshalIndent(a, "", "")
	content, _ := json.Marshal(searchData)
	fileName := "./static/index/js/searchdata/content.json.js"
	f, _ := os.Create(fileName)
	io.WriteString(f, string(content))
}

// 前台模板渲染
func (p *BaseController) IndexCommTpl(method, tpl, sectionTpl, pageAlias, pageTitle string) {
	// 取出最新文章
	var newArticle []models.NewArticle
	models.Article{}.GetNewArticle(7, &newArticle)
	p.Data["NewArticle"] = newArticle

	// 取出所有分类及分类下的文章数目
	cate := models.Category{}.GetAllCateAndNum()
	p.Data["CategoryData"] = cate

	// 取出所有的标签
	tag := p.GetAllTag()
	p.Data["TagData"] = tag

	// 取出所有的友链
	var l []models.Link
	models.Link{}.GetAllLink(&l)
	p.Data["LinkInfo"] = l

	// 取出system setting
	var st models.System
	models.System{}.GetSystem(&st)
	p.Data["SystemData"] = st

	// 取出轻语
	var w []models.Whisper
	models.Whisper{}.GetNewWhisper(7, &w)
	p.Data["NewWhisper"] = w

	if method != "Page" {
		if method != "whisper" {
			p.Data["url"] = beego.URLFor("IndexController." + method)
		} else {
			p.Data["url"] = "/" + method
		}
	} else {
		p.Data["url"] = pageAlias
	}

	// 取出归档时间
	// 先取出所有的文章添加时间
	var ct []models.CreatedAt
	models.Article{}.GetArticleCreateTime(&ct)
	var str = ""
	var monthYear []interface{}
	// 循环
	for _, v := range ct {
		monthAndYear := beego.Date(v.CreatedAt, "m月 Y")
		str += monthAndYear + ","
		monthAndYear = strings.Trim(str, ",")
		monthAndYears := strings.Split(monthAndYear, ",")
		monthYear = p.Duplicate(monthAndYears)
	}

	p.Data["monthYear"] = monthYear
	// 取出单页面信息
	var s []models.SinglePage
	models.SinglePage{}.GetAllPage(&s)
	p.Data["PagesInfo"] = s

	if p.Ctx.Request.RequestURI != "/" {
		p.Data["PageTitle"] = st.WebName + " | " + pageTitle
	} else {
		p.Data["PageTitle"] = st.WebTitle
	}
	p.Layout = "index/public/layout.html"
	if method != "Page" {
		p.TplName = "index/basic/" + tpl
	} else {
		p.TplName = "index/page/" + tpl
	}
	if len(sectionTpl) != 0 {
		p.LayoutSections = make(map[string]string)
		p.LayoutSections["Scripts"] = "index/public/" + sectionTpl
	}
}

// 通过两重循环过滤重复元素
func (p *BaseController) RemoveRepByLoop(slc []int) []int {
	result := []int{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// 去重
func (BaseController) Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}

// 后台模板渲染
func (p *BaseController) AdminCommTpl(tpl string, pageTitle string) {
	// 取出system setting，default_author, web_name
	var st models.System
	models.System{}.GetSectionSystem(&st)
	p.Data["SystemData"] = st

	p.Data["PageTitle"] = st.WebName + " | " + pageTitle

	p.Layout = "admin/public/layout.html"
	p.TplName = "admin/" + tpl
}

//生成随机字符串
func (p *BaseController) GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 获取文件后缀
func (p *BaseController) GetFileSuffix(s string) string {
	re, _ := regexp.Compile(".(jpg|jpeg|png|gif)")
	suffix := re.ReplaceAllString(s, "")
	return suffix
}

// md5 获取
func (p *BaseController) GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func (p *BaseController) checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

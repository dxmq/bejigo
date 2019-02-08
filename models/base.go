package models

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io"
	"os"
)

var (
	db *gorm.DB
)

var count int

func init() {
	var err error
	if err = os.MkdirAll("data", 0777); err != nil { // 创建数据库目录
		panic("failed" + err.Error())
	}
	db, err = gorm.Open("sqlite3", "data/data.db")
	if err != nil {
		panic("连接数据库失败")
	}
	db.AutoMigrate(&User{}, &Tag{}, &Category{}, &ArticleTag{}, &Article{}, &Link{}, &SinglePage{}, &System{}, &Whisper{})
	// 如果user数据表里边没有数据，新增一条user记录
	if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
		password := GetMd5String("123456")
		db.Create(&User{
			UserName: "admin",
			PassWord: password,
			Email:    "admin888@qq.com",
			Avatar:   "/static/index/img/face.png",
			Role:     1,
		})
	}
	// 如果system数据表里没有数据，新增一条user记录
	if err := db.Model(&System{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&System{
			ID:              1,
			RecordNumber:    "滇ICP备10000000号",
			CopyRight:       "2018-2019 xxx.com",
			WebName:         "The Lost Eden",
			WebTitle:        "Lost | 愿你岁月无波澜，敬我余生不悲欢",
			WebKeywords:     "The Lost Eden 博客",
			DefaultAuthor:   "zuo",
			WebDescription:  "Lost | 愿你岁月无波澜，敬我余生不悲欢",
			WebSlogan:       "愿你岁月无波澜，敬我余生不悲欢。",
			IndexShowNumber: 5,
		})
	}

	// 初始化category
	if err := db.Model(&Category{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&Category{
			CategoryName: "PHP",
		})
	}
	// 初始化tag
	if err := db.Model(&Tag{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&Tag{
			Name: "PHP",
		})
	}

	// 初始化articleTag
	if err := db.Model(&ArticleTag{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&ArticleTag{
			ArticleId: 1,
			TagId:     "1",
		})
	}

	// 初始化Link
	if err := db.Model(&Link{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&Link{
			LinkName: "h1ml博客",
			LinkUrl:  "http://h1ml.com",
		})
	}

	// 初始化whisper
	if err := db.Model(&Whisper{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&Whisper{
			UserId:  1,
			Whisper: "人各有命，上天注定",
		})
	}
	InitArticle()
	InitSearchData() // 初始化搜索数据
	InitSinglePage()
}

// 初始化Article
func InitArticle() {
	if err := db.Model(&Article{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&Article{
			Title:      "写在前面",
			Author:     "zuo",
			CategoryId: 1,
			IsShow:     1,
			IsTop:      1,
			Summary: `本博客使用golang和beego框架开发，数据库采用更为轻量的sqlite3，有几个优点：

- 简洁优雅，小而美
- 性能优秀，并发好
- 采用editormd作为文章编辑器，专注于写作`,
			Content: `> 本博客使用golang和beego框架开发，数据库采用更为轻量的sqlite3，有几个优点：

- 简洁优雅，小而美
- 性能优秀，并发好
- 采用editormd作为文章编辑器，专注于写作

**安装**
1. 安装所需库
go get github.com/astaxie/beego
go get github.com/beego/bee
go get github.com/mattn/go-sqlite3

2. 克隆项目到GOPATH/src目录
cd $GOPATH/src/
git clone https://gitee.com/h1ml/bejigo

3. 运行项目
bee run

4. 如需nginx代理服务，请参考https://beego.me/docs/deploy/nginx.md>

>注：我是初学者，水平有限，如有任何好的建议和意见或遇到任何bug，请联系我。邮箱：1697859639@qq.com`,
		})
	}

}

func InitSinglePage() {
	// 初始化singlePage
	if err := db.Model(&SinglePage{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&SinglePage{
			PageName:      "关于",
			Sort:          10,
			PageAlias:     "about",
			PageIconClass: "fa fa-user",
			IsShow:        1,
			Content: `<div class="layout-l">
     <div class="content_container">
      <div class="post">
       <h1 class="post-title">关于</h1>
       <div class="post-content">
        <div class="author-page">
         <a href="javascript:;" class="photo"><img src="/static/index/img/face.png" /></a>
        <div class="author">
          <p><i class="fa fa-user">name：</i>zuo</p>
          <p><i class="fa fa-email">email：</i>xxx@qq.com</p>
          <p><i class="fa fa-github">github：</i><a href="https://github.com/dxmq">dxmq</a></p>
         </div>
        </div>
        <h6 id="为什么会有这个？"><a href="#为什么会有这个？" class="headerlink" title="为什么会有这个？"></a>为什么会有这个？</h6>
        <p>不想让生活博客被代码充斥而变了味，谈生活，不谈技术，谈技术，就不谈生活，过的纯粹些。</p> 
        <h6 id="要怎么做？"><a href="#要怎么做？" class="headerlink" title="要怎么做？"></a>要怎么做？</h6>
        <p>每天好好学习之后总结点东西发上来罢了。</p> 
        <h6 id="我是谁？"><a href="#我是谁？" class="headerlink" title="我是谁？"></a>我是谁？</h6>
        <p>1) 不喜热闹罢了;<br />2) 喜欢简洁纯粹的设计;<br />3) 喜阅读，然于书海中徜徉多年，所获不及万分之一，深感书海无涯;<br /></p> 
       </div>
       <div id="comments">
        <div id="lv-container" data-id="city" data-uid="MTAyMC80MjMxOC8xODg2NQ=="></div>
       </div>
      </div>
     </div>
    </div>
<script src="/static/admin/js/jquery.min.js"></script>
<script>window._bd_share_config={"common":{"bdSnsKey":{},"bdText":"","bdMini":"2","bdMiniList":["mshare","weixin","tsina","qzone","linkedin","fbook","twi","print","renren","sqq","evernotecn","bdysc","tqq","tqf","bdxc","kaixin001","tieba","douban","bdhome","thx","ibaidu","meilishuo","mogujie","diandian","huaban","duitang","hx","fx","youdao","sdo","qingbiji","people","xinhua","mail","isohu","yaolan","wealink","ty","iguba","h163","copy"],"bdPic":"","bdStyle":"1","bdSize":"16"},"share":{},"image":{"viewList":["tsina","qzone","weixin","fbook","twi","linkedin","youdao","evernotecn","mshare"],"viewText":"分享到：","viewSize":"16"},"selectShare":{"bdContainerClass":null,"bdSelectMiniList":["tsina","qzone","weixin","fbook","twi","linkedin","youdao","evernotecn","mshare"]}};with(document)0[(getElementsByTagName('head')[0]||head).appendChild(createElement('script')).src='http://bdimg.share.baidu.com/static/api/js/share.js?v=89860593.js?cdnversion='+~(-new Date()/36e5)];
</script>
<script>(function(d, s) {
    var j, e = d.getElementsByTagName('body')[0];
    if (typeof LivereTower === 'function') { return; }
    j = d.createElement(s);
    j.src = '/static/index/js/embed.dist.js'/*tpa=https://cdn-city.livere.com/js/embed.dist.js*/;
    j.async = true;
    e.appendChild(j);
})(document, 'script');
</script>`,
		})
	}
}

// md5 获取
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// 初始化搜索数据，用于前台搜索
func InitSearchData() {
	var searchData []Search
	Article{}.GetArticleDataForSearch(&searchData)
	//content, err := json.MarshalIndent(a, "", "")
	content, _ := json.Marshal(searchData)
	fileName := "./static/index/js/searchdata/content.json.js"
	f, _ := os.Create(fileName)
	io.WriteString(f, string(content))
}
